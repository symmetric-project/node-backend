package graph

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/symmetric-project/node-backend/errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/joho/godotenv"
)

// This file will not be regenerated automatically.
//
// It serves for dependency injection

var DATABASE_URL string
var SQ sq.StatementBuilderType
var DB *pgxpool.Pool

func init() {
	SQ = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	godotenv.Load()
	DATABASE_URL = os.Getenv("DATABASE_URL")

	var err error
	DB, err = pgxpool.Connect(context.Background(), DATABASE_URL)

	if err != nil {
		errors.Stacktrace(err)
		os.Exit(1)
	}
}

type Resolver struct{}
