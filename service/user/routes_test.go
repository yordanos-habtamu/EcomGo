package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/yordanos-habtamu/EcomGo.git/types"
)

func TestUserServiceHandlers(t *testing.T){
    userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T){
		dateStr := "2000-01-01"
	layout := "2006-01-02" // The layout for parsing (Go's reference date is "2006-01-02 15:04:05")
	dob, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	fmt.Println("Parsed DoB:", dob)
		payload := types.RegisterUserPayload{
			FirstName: "user",
		    LastName: "123",
			Email:"hh@gmail.com",
			Password: "asd",
			Age:21,
			DoB: dob,
			Sex : "Male",
		}
		marshalled,_ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost,"/register",bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
        router.HandleFunc("/register",handler.handleRegister)
		router.ServeHTTP(rr,req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d",http.StatusBadRequest,rr.Code)
		}
	
})
} 

type mockUserStore struct {

}

func (m *mockUserStore) GetUserByEmail (email string) (*types.User,error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById (id int) (*types.User,error) {
	return nil, nil
}
func (m *mockUserStore) CreateUser (types.User) error {
	return nil
}