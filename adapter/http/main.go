package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/boooscaaa/clean-go/adapter/postgres"
	"github.com/boooscaaa/clean-go/di"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/boooscaaa/clean-go/adapter/http/docs"
	"github.com/boooscaaa/clean-go/adapter/http/middleware"
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

	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	jsonApiRouter := router.PathPrefix("/").Subrouter()
	jsonApiRouter.Use(middleware.Cors)

	jsonApiRouter.Handle("/product", http.HandlerFunc(productService.Create)).Methods("POST")
	jsonApiRouter.Handle("/product", http.HandlerFunc(productService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	port := viper.GetString("server.port")

	if port == "" {
		port = os.Getenv("PORT")
	}
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
