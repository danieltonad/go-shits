package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Response struct {
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

var userCache = make(map[int]User, 1)

var userCacheMutex sync.RWMutex

func main(){
	fmt.Println("Starting Server...")	

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/user/{id}", getUser)
	mux.HandleFunc("POST /user", createUser)
	mux.HandleFunc("/users", getUsers)

	server := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	fmt.Println("Server started at http://localhost:8080")
	server.ListenAndServe()

	

}

func handleRoot(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json_response, err := json.Marshal(Response{Message: "Hello World", Data: nil})
	if err != nil {
		http.Error(w, "Error converting user to JSON", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", json_response)
}


func getUsers(w http.ResponseWriter, r *http.Request){
	userCacheMutex.RLock()
	defer userCacheMutex.RUnlock()
	var users []User
	for _, user := range userCache {
		users = append(users, user)
	}

	json_response, err := json.Marshal(Response{Message: "All Users", Data: users})
	if err != nil {
		http.Error(w, "Error converting user to JSON", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", json_response)
}


func getUser(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)	
	}
	userCacheMutex.RLock()
	user, ok := userCache[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	userCacheMutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json_response, err := json.Marshal(Response{Message: "User found!", Data: user})
	if err != nil {
		http.Error(w, "Error converting user to JSON", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", json_response)
}


func createUser(w http.ResponseWriter, r *http.Request){
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	
	if user.Name == "" || user.Email == "" || user.Phone == "" {
		http.Error(w, "Name, Email and Phone are required", http.StatusBadRequest)
		return
	}
	
	userCacheMutex.Lock()
	userCache[len(userCache) + 1] = user
	userCacheMutex.Unlock()
	fmt.Printf("DEBUG => %v\n", len(userCache))
	w.WriteHeader(http.StatusCreated)


	w.Header().Set("Content-Type", "application/json")
	json_response, err := json.Marshal(Response{Message: "User created successfully", Data: nil})
	if err != nil {
		http.Error(w, "Error converting user to JSON", http.StatusInternalServerError)
		return
	}
	
	fmt.Fprintf(w, "%s", json_response)
	
}