package controllers

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "diet-plan-generator/database"
    "diet-plan-generator/models"
    "diet-plan-generator/services"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string      `json:"token"`
    User  models.User `json:"user"`
}

type SignupRequest struct {
    Email       string `json:"email"`
    Password    string `json:"password"`
    FullName    string `json:"fullName"`
    Username    string `json:"username"`
    PhoneNumber string `json:"phoneNumber"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    var req SignupRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    user := models.User{
        ID:          primitive.NewObjectID(),
        Email:       req.Email,
        Password:    string(hashedPassword),
        FullName:    req.FullName,
        Username:    req.Username,
        PhoneNumber: req.PhoneNumber,
    }

    collection := database.DB.Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var existingUser models.User
    if err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser); err == nil {
        http.Error(w, "User already exists", http.StatusBadRequest)
        return
    } else if err != mongo.ErrNoDocuments {
        log.Printf("Error checking for existing user: %v", err)
        http.Error(w, "Error checking for existing user", http.StatusInternalServerError)
        return
    }

    _, err = collection.InsertOne(ctx, user)
    if err != nil {
        log.Printf("Error inserting user: %v", err)
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var loginReq LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := services.GetUserByEmail(loginReq.Email)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    token, err := services.GenerateToken(user.ID.Hex())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := LoginResponse{
        Token: token,
        User:  *user, 
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
