package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	FistName  string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Email     string     `json:"email"`
}
var DB *gorm.DB
var err error 

const DNS = "root:12345678@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=UTC"


func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to MySQl database")
	}
	DB.AutoMigrate(&User{})
}



func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Check if the user already exists based on an identifying attribute (e.g., email)
    var existingUser User
    result := DB.Where("email = ?", user.Email).First(&existingUser)
    if result.Error == nil {
        // User with similar email already exists, return a message or handle the case as needed
        http.Error(w, "User with this email already exists", http.StatusConflict)
        return
    }

    // No existing user found, proceed to create the new user
    DB.Create(&user)
    json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "applications/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)

}


func GetUser(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "applications/json")
	params := mux.Vars(r)
	var user User 
	DB.First(&user,params["id"])
	json.NewEncoder(w).Encode(user)
		
}

func UpdateUser(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "applications/json")
	params := mux.Vars(r)
	var user User 
	DB.First(&user,params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "applications/json")
	params := mux.Vars(r)
	var user User 
	DB.Delete(&user,params["id"])
	json.NewEncoder(w).Encode("The user deleted successfully....")
	
}