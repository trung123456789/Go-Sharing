package urls

import (
	"apis/api"
	"config"
	"structdemo"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Config type
type Config structdemo.Config

// GeneralURL function
func GeneralURL() {
	cfg := config.GetEnv()
	port := fmt.Sprintf(":%s", cfg.Server.Port)

	r := mux.NewRouter()
	router := r.PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/userlist", api.PostGetUser).Methods(http.MethodPost)
	router.HandleFunc("/addUser", api.PostAddUser).Methods(http.MethodPost)
	router.HandleFunc("/updateUser", api.PutUpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/deleteUser/{userID}", api.DeleteUser).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(port, r))
}
