package user

import (
	"bytes"
	"encoding/json"
	"enfermeria_go/types"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}

func TestUserServiceHandles(t *testing.T) {
	userStore := &mockUserStore{}
	handlre := NewHandler(userStore)

	t.Run("should fail if user payload is invalid", func(t *testing.T) {

		// test logic
		type Payload struct {
			FirstName int    `json:"first_name"`
			LastName  int    `json:"last_name"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		}

		payload := Payload{
			FirstName: 1,
			LastName:  1,
			Email:     "",
			Password:  "123456",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()

		router := mux.NewRouter()

		router.HandleFunc("/register", handlre.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", rec.Code)
		}
	})
}
