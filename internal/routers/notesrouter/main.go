package notesrouter

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kot-zakhar/gobook/internal/dbconnector"
	"go.mongodb.org/mongo-driver/bson"
	//	"go.mongodb.org/mongo-driver/mongo"
	//	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func BindEndpointsToRouter(notesRouter *mux.Router) {
	notesRouter.HandleFunc("/", listNotes).Methods("GET")

	//	notesRouter.HandleFunc("/{id}", getNoteById).Methods("GET")
}

func listNotes(res http.ResponseWriter, req *http.Request) {
	fmt.Println("listing notes")
	client, err := dbconnector.GetMongoClient()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
	}

	notesColl := client.Database("gobook").Collection("notes")

	cursor, err := notesColl.Find(context.TODO(), bson.D{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
	}
	fmt.Fprintf(res, "%v", results)
}
