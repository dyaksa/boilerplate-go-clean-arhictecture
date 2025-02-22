package pqsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var _ Collection[any] = &Database[any]{}

type Model[T any] interface {
	ScanDestinations() []any
	To() *T
}

type Collection[T any] interface {
	QueryContext(ctx context.Context, tx *sql.Tx, query string, entity Model[T], args ...any) ([]*T, error)
	QueryRowContext(ctx context.Context, tx *sql.Tx, query string, entity Model[T], args ...any) (*T, error)
	ExecContext(ctx context.Context, tx *sql.Tx, query string, args ...any) error
	Query(ctx context.Context, tx *sql.Tx, query string, entity Model[T], args ...any) ([]*T, error)
	QueryRow(ctx context.Context, tx *sql.Tx, query string, entity Model[T], args ...any) (*T, error)
	Exec(ctx context.Context, tx *sql.Tx, query string, args ...any) error
}

type Database[T any] struct {
	wrapper *WrapperTx
}

func (d *Database[T]) ExecContext(ctx context.Context, tx *sql.Tx, query string, args ...any) error {
	_, err := tx.ExecContext(ctx, query, args...)
	return err
}

func (d *Database[T]) QueryContext(ctx context.Context, tx *sql.Tx, query string, entity Model[T], args ...any) ([]*T, error) {
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []*T
	for rows.Next() {
		err = rows.Scan(entity.ScanDestinations()...)
		if err != nil {
			return nil, err
		}

		results = append(results, entity.To())
	}

	return results, nil
}

func (d *Database[T]) QueryRowContext(ctx context.Context, tx *sql.Tx, query string, entity Model[T], args ...any) (*T, error) {
	err := tx.QueryRowContext(ctx, query, args...).
		Scan(entity.ScanDestinations()...)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, fmt.Errorf("no rows found")
	case err != nil:
		return nil, err
	default:
		return entity.To(), nil
	}
}

func (d *Database[T]) Exec(ctx context.Context, tx *sql.Tx, query string, args ...any) error {
	_, err := tx.Exec(query, args...)
	return err
}

func (d *Database[T]) Query(ctx context.Context, tx *sql.Tx, query string, entity Model[T], args ...any) ([]*T, error) {
	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []*T
	for rows.Next() {
		err := rows.Scan(entity.ScanDestinations()...)
		if err != nil {
			return nil, err
		}

		results = append(results, entity.To())
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (d *Database[T]) QueryRow(ctx context.Context, tx *sql.Tx, query string, entity Model[T], args ...any) (*T, error) {
	err := tx.QueryRow(query, args...).Scan(entity.ScanDestinations()...)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, fmt.Errorf("no rows found")
	case err != nil:
		return nil, err
	default:
		return entity.To(), nil
	}
}

type Client interface {
	Database() *Database[any]
	PingContext(ctx context.Context) error
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
	Conn(ctx context.Context) (*sql.Conn, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Ping() error
	Close() error
}

type postgresClient struct {
	db *sql.DB
}

func (c *postgresClient) Database() *Database[any] {
	return &Database[any]{wrapper: &WrapperTx{db: c.db}}
}

func (c *postgresClient) PingContext(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

func (c *postgresClient) Ping() error {
	return c.db.Ping()
}

func (c *postgresClient) Close() error {
	return c.db.Close()
}

func (c *postgresClient) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return c.db.BeginTx(ctx, opts)
}

func (c *postgresClient) Begin() (*sql.Tx, error) {
	return c.db.Begin()
}

func (c *postgresClient) Conn(ctx context.Context) (*sql.Conn, error) {
	return c.db.Conn(ctx)
}

func (c *postgresClient) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}

func (c *postgresClient) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return c.db.QueryContext(ctx, query, args...)
}

func (c *postgresClient) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return c.db.QueryRowContext(ctx, query, args...)
}

func (c *postgresClient) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return c.db.PrepareContext(ctx, query)
}

func NewClient(connection string) (Client, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	return &postgresClient{db: db}, nil
}
