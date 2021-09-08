package graph

import (
	"os"

	"github.com/symmetric-project/node-backend/errors"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
)

// This file will not be regenerated automatically.
//
// It serves for dependency injection

var DATABASE_URL string
var SQ sq.StatementBuilderType
var DB *sqlx.DB

func init() {
	SQ = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	godotenv.Load()
	DATABASE_URL = os.Getenv("DATABASE_URL")

	var err error
	DB, err = sqlx.Connect("pgx", DATABASE_URL)

	if err != nil {
		errors.Stacktrace(err)
		os.Exit(1)
	}
}

type Resolver struct{}
