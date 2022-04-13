package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/boooscaaa/clean-go/adapter/http/graphql/schema"
	"github.com/boooscaaa/clean-go/adapter/postgres"
	"github.com/boooscaaa/clean-go/di"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/boooscaaa/clean-go/adapter/http/rest/docs"
	"github.com/boooscaaa/clean-go/adapter/http/rest/middleware"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Clean GO API Docs
// @version 1.0.0
// @contact.name Vinícius Boscardin
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:port
// @BasePath /
func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	productService := di.ConfigProductDI(conn)
	productGraphQLService := di.ConfigProductGraphQLDI(conn)

	router := mux.NewRouter()
	graphQLRouter := schema.Config(productGraphQLService)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	jsonApiRouter := router.PathPrefix("/").Subrouter()
	jsonApiRouter.Use(middleware.Cors)

	jsonApiRouter.Handle("/product", http.HandlerFunc(productService.Create)).Methods("POST", "OPTIONS")
	jsonApiRouter.Handle("/product", http.HandlerFunc(productService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET", "OPTIONS")

	router.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := schema.ExecuteQuery(r.URL.Query().Get("query"), graphQLRouter)
		json.NewEncoder(w).Encode(result)
	})

	port := viper.GetString("server.port")

	if port == "" {
		port = os.Getenv("PORT")
	}
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
