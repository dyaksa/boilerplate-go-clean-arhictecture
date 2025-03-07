package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/infrastructure/pqsql"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/log"
)

func NewPostgres(env *Env, l log.Logger) pqsql.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName
	dbSSL := env.DBSSL

	client, err := pqsql.NewClient(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, dbSSL))
	if err != nil {
		l.Warn("failed to connect to postgres", log.Error("error", err))
	}

	if err = client.PingContext(ctx); err != nil {
		l.Warn("failed to ping postgres", log.Error("error", err))
	}

	return client
}

func CloseConnection(client pqsql.Client, l log.Logger) {
	if client == nil {
		return
	}

	if err := client.Close(); err != nil {
		l.Warn("failed to close postgres connection", log.Error("error", err))
	}

	l.Info("postgres connection closed")
}
