package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kot-zakhar/gobook/internal/dbconnector"
	"github.com/kot-zakhar/gobook/internal/routers/notesrouter"
	"log"
	"net/http"
)

const connectionStringEnvName = "MONGO_CONNECTION_STRING"
const hostAddress = "localhost:8080"

func getMainRouter() *mux.Router {
	router := mux.NewRouter()
	notesrouter.BindEndpointsToRouter(router.PathPrefix("/notes").Subrouter().StrictSlash(true))
	return router
}

func main() {
	connectionString := os.GetEnv(connectionStringEnvName)
	err := dbconnector.ConnectToDb(connectionString)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer dbconnector.DisconnectFromDb()

	log.Println("Server is listened on " + hostAddress)
	log.Fatal(http.ListenAndServe(hostAddress, getMainRouter()))
}
