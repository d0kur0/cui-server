package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/d0kur0/cui-server/auth"
	"github.com/go-chi/chi"

	"github.com/d0kur0/cui-server/config"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/d0kur0/cui-server/graph"
	"github.com/d0kur0/cui-server/graph/generated"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (server *Server) Init(appConfig *config.Config) error {
	router := chi.NewRouter()
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", appConfig.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), router)
}
