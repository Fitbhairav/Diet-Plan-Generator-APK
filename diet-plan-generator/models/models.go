// package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

// type ChatGPTRequest struct {
//     Messages  []ChatGPTMessage `json:"messages"`
//     MaxTokens int              `json:"max_tokens"`
//     Model     string           `json:"model"`
// }

// type ChatGPTMessage struct {
//     Role    string `json:"role"`
//     Content string `json:"content"`
// }

// type ChatGPTResponse struct {
//     Choices []struct {
//         Message ChatGPTMessage `json:"message"`
//     } `json:"choices"`
// }

// type User struct {
//     ID          primitive.ObjectID `bson:"_id,omitempty"`
//     Email       string             `bson:"email"`
//     Password    string             `bson:"password"`
//     FullName    string             `bson:"fullName"`
//     Username    string             `bson:"username"`
//     PhoneNumber string             `bson:"phoneNumber"`
// }
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RequestBody struct {
    DietryPrefs            string `json:"dietry_prefs"`
    Goal                   string `json:"goal"`
    CaloricIntake          int    `json:"caloric_intake"`
    Proteins               int    `json:"proteins"`
    Fats                   int    `json:"fats"`
    Carbohydrates          int    `json:"carbohydrates"`
    AllergiesOrRestrictions string `json:"allergies_or_restrictions"`
    PlanType               string `json:"plan_type"`
    MealsFrequency         int    `json:"meals_frequency"`
}

type User struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Email       string             `bson:"email"`
    Password    string             `bson:"password"`
    FullName    string             `bson:"fullName"`
    Username    string             `bson:"username"`
    PhoneNumber string             `bson:"phoneNumber"`
}
