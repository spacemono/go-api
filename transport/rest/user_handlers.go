package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spacemono/go-api/generated/openapi"
	"github.com/spacemono/go-api/service"
	"github.com/spacemono/go-api/service/command"
)

type UserHandlers struct {
	userService *service.User
}

func NewUserHandlers(userService *service.User) *UserHandlers {

	return &UserHandlers{userService: userService}
}

func (h *UserHandlers) GetUsersId(w http.ResponseWriter, r *http.Request, id string) {

	// Add error handling
	user := h.userService.GetUserById(id)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {

	// Add error handling
	users := h.userService.GetAll()

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandlers) PostUsers(w http.ResponseWriter, r *http.Request) {
	var body openapi.User
	_ = json.NewDecoder(r.Body).Decode(&body)

	cmd := &command.CreateUser{Username: body.Username, Name: body.Name, Password: body.Password}
	result, err := h.userService.Create(cmd)

	if err != nil {
		fmt.Println("error creating user", err)
		// Change to new error from Openapi
		// json.NewEncoder(w).Encode("error creating user")
	}

	json.NewEncoder(w).Encode(result)
}
