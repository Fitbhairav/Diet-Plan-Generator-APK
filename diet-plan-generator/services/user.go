package services

import (
    "context"
    "errors"
    "time"

    "diet-plan-generator/database"
    "diet-plan-generator/models"
    "github.com/dgrijalva/jwt-go"
    "go.mongodb.org/mongo-driver/bson"
    "os"
)

func CreateUser(user *models.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := database.DB.Collection("users").InsertOne(ctx, user)
    return err
}

func GetUserByEmail(email string) (*models.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user models.User
    err := database.DB.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
    if err != nil {
        return nil, errors.New("user not found")
    }
    return &user, nil
}

func GenerateToken(userID string) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return "", errors.New("JWT_SECRET environment variable is not set")
    }

    claims := &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        Subject:   userID,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
