package main

import (
	"context"
	"database/sql"
	"go-graphql-apollo-federation/users/graph"
	"go-graphql-apollo-federation/users/graph/model"
	"go-graphql-apollo-federation/users/internal/db"
	"go-graphql-apollo-federation/users/internal/middleware"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"

	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const defaultPort = "8081"
const (
	websocketsKeepAlivePingInterval = 10 * time.Second
	dataLoaderWait                  = 250 * time.Microsecond
	dataLoaderMaxBatch              = 1000
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/md_petshop")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	c := graph.Config{Resolvers: &graph.Resolver{
		UserDB: queries,
	}}

	c.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver, arg *model.UserType) (res interface{}, err error) {
		userCtx := ctx.Value("user")
		if userCtx == nil {
			return &model.LoginRequiredError{
				Message: "User is not authenticated",
			}, nil
		}

		user := userCtx.(*model.User)
		if user.Type != *arg {
			return &model.UnauthorizedError{
				Message: "Permission insufficient",
			}, nil
		}
		return next(ctx)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	/*srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: websocketsKeepAlivePingInterval,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.AddTransport(transport.SSE{})
	srv.AddTransport(transport.MultipartForm{})*/
	router := chi.NewRouter()
	router.Use(middleware.AuthMiddleware)

	//handlerAuthMiddleware := middleware.AuthMiddleware(srv)
	//http.Handle("/", handlerAuthMiddleware)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
