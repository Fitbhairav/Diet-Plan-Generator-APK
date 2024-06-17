package routers

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"

    "diet-plan-generator/controllers"
)

func InitializeRoutes() {
    r := mux.NewRouter()

    r.HandleFunc("/generate-diet-plan", controllers.GenerateDietPlan).Methods("POST")
    r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")
    r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", r))
}
