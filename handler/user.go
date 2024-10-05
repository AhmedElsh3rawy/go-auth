package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AhmedElsh3rawy/go-auth/database"
	"github.com/AhmedElsh3rawy/go-auth/helpers"
	"github.com/AhmedElsh3rawy/go-auth/users"
	"github.com/google/uuid"
)

type CreatedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u CreatedUser
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println("error while decoding request body")
		return
	}

	ctx := context.Background()
	_, err = database.Query.GetUser(ctx, u.Email)
	if err == nil {
		http.Error(w, "User is already registered", http.StatusBadRequest)
		return
	}

	hash, err := helpers.HashPassword([]byte(u.Password))
	if err != nil {
		fmt.Println("error hashing password")
		return
	}

	err = database.Query.CreateUser(ctx, users.CreateUserParams{
		ID:       uuid.New(),
		Username: u.Username,
		Email:    u.Email,
		Password: string(hash),
	})
	if err != nil {
		fmt.Println("error quering database")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	// get the body email and password
	// compare passwords
	// generate jwt token
	// response
}

func Logout(w http.ResponseWriter, r *http.Request) {}
