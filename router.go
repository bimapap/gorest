package main

import (
	"log"
	"net/http"

	"github.com/bimapap/gorest/model"
	"github.com/bimapap/gorest/repository"
	"github.com/bimapap/gorest/service"
	"github.com/gorilla/mux"
)

func RestRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	customerRouter(api)
	r.Use(LoggingMiddleware)
	return r
}

func customerRouter(r *mux.Router) {
	var dbConn, err = NewDBConnection("customer.db")
	if err != nil {
		log.Fatalf("DB Connection error : %v", err)
	}
	dbConn.AutoMigrate(&model.Customer{})
	var custRepository = repository.NewCustomerRepository(dbConn)
	var custService = service.NewCustomerService(custRepository)
	var custHandler = NewCustomerHandler(custService)

	r.HandleFunc("/customers:{limi}", custHandler.GetAll).Queries("limit", "{limit}", "offset", "{offset}").Methods(http.MethodGet)
	r.HandleFunc("/customers/{id}", custHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/customers", custHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/customers/{id}", custHandler.Put).Methods(http.MethodPut)
	r.HandleFunc("/customers/{id}", custHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", custHandler.NotFound)
}
