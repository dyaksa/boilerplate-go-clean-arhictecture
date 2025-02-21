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
	panic("unimplemented")
}

func (d *Database[T]) QueryRow(ctx context.Context, tx *sql.Tx, query string, entity Model[T], args ...any) (*T, error) {
	panic("unimplemented")
}

func NewDatabase[T any](db *sql.DB) *Database[T] {
	return &Database[T]{wrapper: NewWrapper(db)}
}
