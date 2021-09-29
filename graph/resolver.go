package graph

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/symmetric-project/node-backend/env"
	"github.com/symmetric-project/node-backend/utils"

	sq "github.com/Masterminds/squirrel"
)

// This file will not be regenerated automatically.
//
// It serves for dependency injection

var SQ sq.StatementBuilderType
var DB *pgxpool.Pool

func init() {
	SQ = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	var err error
	DB, err = pgxpool.Connect(context.Background(), env.CONFIG.DATABASE_URL)

	if err != nil {
		utils.StacktraceError(err)
		os.Exit(1)
	}
}

type Resolver struct{}
