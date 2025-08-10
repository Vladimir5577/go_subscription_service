package main

import (
	"net/http"
	"subscriptions_service/internal/config"
	"subscriptions_service/internal/handler"
	"subscriptions_service/internal/repository"
	"subscriptions_service/internal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	_ "subscriptions_service/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Swagger Subscription Service
// @version 1.0
// @description This is a server for user subscriptions.
// @host localhost:8080
func main() {

	config.InitLogrus()

	envConfigs, err := config.NewEnvConfig()
	if err != nil {
		logrus.Fatalf("Can't read .env file, %v\n", err)
	}
	db, err := config.ConnectDB(envConfigs)
	if err != nil {
		logrus.Fatalf("Can't connect to database, %v\n", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	subscriptionRepository := repository.NewSubscriptionRepository(db)
	subscriptionService := service.NewSubscriptionService(subscriptionRepository)
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionService)

	r.Post("/subscription", subscriptionHandler.Create())
	r.Get("/subscriptions/total_summ", subscriptionHandler.GetTotalSummByFilter())
	r.Get("/subscription/{id}", subscriptionHandler.GetById())
	r.Put("/subscription", subscriptionHandler.Update())
	r.Delete("/subscription/{id}", subscriptionHandler.Delete())
	r.Get("/subscriptions", subscriptionHandler.GetAll())

	// Home page.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Subscriptions service."))
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://"+envConfigs.ServiceHost+":"+envConfigs.ServicePort+"/swagger/doc.json"), //The url pointing to API definition
	))

	logrus.Info("Server up and running on the port", envConfigs.ServicePort)
	http.ListenAndServe(":"+envConfigs.ServicePort, r)
}
