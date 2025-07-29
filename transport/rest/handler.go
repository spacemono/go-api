package rest

import (
	"github.com/gorilla/mux"
	"github.com/spacemono/go-api/generated/openapi"
	"github.com/spacemono/go-api/service"
)

func NewHandler(userService *service.User) *mux.Router {
	router := mux.NewRouter()
	openapi.HandlerFromMux(NewUserHandlers(userService), router)

	return router
}
