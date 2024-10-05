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

func Register(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("error while decoding request body")
		return
	}

	ctx := context.Background()
	_, err = database.Query.GetUser(ctx, body.Email)
	if err == nil {
		http.Error(w, "User is already registered", http.StatusBadRequest)
		return
	}

	hash, err := helpers.HashPassword([]byte(body.Password))
	if err != nil {
		fmt.Println("error hashing password")
		return
	}

	err = database.Query.CreateUser(ctx, users.CreateUserParams{
		ID:       uuid.New(),
		Username: body.Username,
		Email:    body.Email,
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
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("error while decoding request body")
		return
	}

	ctx := context.Background()
	u, err := database.Query.GetUser(ctx, body.Email)
	if err != nil {
		fmt.Println("error quering database")
		return
	}

	err = helpers.ComparePasswords([]byte(u.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}
	w.Write([]byte("All good!"))
	// generate jwt token
	// response
}

func Logout(w http.ResponseWriter, r *http.Request) {}
