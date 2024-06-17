// package controllers

// import (
//     "encoding/json"
//     "net/http"

//     "diet-plan-generator/services"
// )

// type RequestBody struct {
//     Prompt string `json:"prompt"`
// }

// type ResponseBody struct {
//     DietPlan string `json:"dietPlan"`
// }

// func GenerateDietPlan(w http.ResponseWriter, r *http.Request) {
//     var reqBody RequestBody
//     if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     dietPlan, err := services.GenerateDietPlan(reqBody.Prompt)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

//     responseBody := ResponseBody{DietPlan: dietPlan}
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(responseBody)
// }

package controllers

import (
    "encoding/json"
    "net/http"

    "diet-plan-generator/models"
    "diet-plan-generator/services"
)

func GenerateDietPlan(w http.ResponseWriter, r *http.Request) {
    var reqBody models.RequestBody
    if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response, err := services.GenerateDietPlan(reqBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
