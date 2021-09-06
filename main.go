package main

import (
	"context"
	"log"
	"net/http"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var MODE string
var HTTPClient *http.Client
var DATABASE_URL string
var JWT_SECRET string

var DBPool *pgxpool.Pool
var SQ sq.StatementBuilderType

var G *gin.Engine

var COOKIE_DOMAIN string
var COOKIE_SECURE bool

func init() {
	godotenv.Load()
	MODE = os.Getenv("MODE")
	HTTPClient = &http.Client{}
	DATABASE_URL = os.Getenv("DATABASE_URL")
	JWT_SECRET = os.Getenv("JWT_SECRET")

	switch MODE {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	default:
		err := errors.New("unable to determine running mode (dev or prod)")
		HandleError(err)
		os.Exit(1)
	}

	SQ = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	var err error
	DBPool, err = pgxpool.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		HandleError(err)
		os.Exit(1)
	}

	jwt, err := GenerateBackdoorJWT()
	if err != nil {
		HandleError(err)
		os.Exit(1)
	}
	log.Println("Backdoor JWT: " + jwt)

	switch MODE {
	case "dev":
		COOKIE_DOMAIN = os.Getenv("COOKIE_DOMAIN_DEV")
		COOKIE_SECURE = false
	case "prod":
		COOKIE_DOMAIN = os.Getenv("COOKIE_DOMAIN_PROD")
		COOKIE_SECURE = true
	default:
		err := errors.New("unable to determine running mode (dev or prod)")
		HandleError(err)
		os.Exit(1)
	}
}

func main() {
	G = gin.Default()
	G.Use(cors.New(cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"http://symmetric.localhost:3000", "https://symmetric.nodename.com"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin", "Accept", "X-Requested-With", "Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Access-Control-Allow-Origin", "Cache-Control", "Credential"},
	}))
	G.Run(":4000")
}
