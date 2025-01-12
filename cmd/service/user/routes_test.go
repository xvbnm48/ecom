package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/xvbnm48/ecom/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.UserRegisterPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)
		fmt.Println(rr.Body.String())
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

	})
	t.Run("should is user correctly", func(t *testing.T) {
		payload := types.UserRegisterPayload{
			FirstName: "vini",
			LastName:  "vini",
			Email:     "vini1@gmail.com",
			Password:  "123456",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)
		fmt.Println(rr.Body.String())
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	// t.Run("should return 400 if email already exists", func(t *testing.T) {
	// 	payload := types.UserRegisterPayload{
	// 		FirstName: "John",
	// 		LastName:  "Doe",
	// 		Email:     "",
	// 		Password:  "password",
	// 	}
	// 	marshalled, _ := json.Marshal(payload)
	// 	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	rr := httptest.NewRecorder()
	// 	router := mux.NewRouter()
	// 	router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
	// 	router.ServeHTTP(rr, req)

	// 	if rr.Code != http.StatusBadRequest {
	// 		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	// 	}
	// })
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
