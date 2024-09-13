package main

import (
	"database/sql"
	"go-graphql-apollo-federation/carts/graph"
	"go-graphql-apollo-federation/carts/internal/db"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8085"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/md_petshop")
	if err != nil {
		panic(err)
	}

	queries := db.New(dbConn)
	//db.New()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CartDB: queries,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
