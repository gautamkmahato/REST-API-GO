package main

// NOTE: BSON = Binary JSON

import (
	"context" // manage multiple requests
	"fmt" // Println() function
	"os"
	"reflect" // get an object type
	"time"
	"encoding/hex" // hexadecimal encoding of BSON obj

	// import 'mongo-go-driver' package libraries
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database Schema
type ArticlesDatabase struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Title       string             `bson:"title" json:"title"`
	Subtitle    string             `bson:"Subtitle" json:"Subtitle"`
	Description string             `bson:"description" json:"description"`
}

func databaseConnection(){
	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func findArticlesById(){
	// call the collection's Find() method and return Cursor object into result
	fmt.Println(`bson.M{"_id": docID}:`, bson.M{"_id": docID})
	err = col.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)

	// Check for any errors returned by MongoDB FindOne() method call
	if err != nil {
		fmt.Println("FindOne() ObjectIDFromHex ERROR:", err)
		os.Exit(1)
	} 
	else {
		// Print out data from the document result
		fmt.Println("result AFTER:", result, "\n")

		// Struct instances are objects with MongoDB fields that can be accessed as attributes
		fmt.Println("FindOne() result:", result)
		fmt.Printf("result doc ID: %v\n", result.ID)
		fmt.Println("result.title:", result.title)
		fmt.Println("result.Subtitle:", result.Subtitle)
		fmt.Println("result.description:", result.description)
	}
}

func handleRequests() {
	// all routes
    http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/articlesById", findArticlesById)
	log.Fatal(http.ListenAndServe(":10000", nil))
	
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the artickes!")
	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(idStr)
	fmt.Println("\nID ObjectIDFromHex:", docID)
	fmt.Println("ID ObjectIDFromHex err:", err)
	fmt.Println("ID hexByte type:", reflect.TypeOf(docID))

	// Declare a struct instance of the MongoDB fields that will contain the document returned
	result := ArticlesDatabase{}
	fmt.Println("\nresult type:", reflect.TypeOf(result))
	fmt.Println("result BEFORE:", result)

    json.NewEncoder(w).Encode(Articles)
}

func main() {
	// Declare host and port options to pass to the Connect() method
	databaseConnection();

	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// Access a MongoDB collection through a database
	col := client.Database("SomeDatabase").Collection("Some Collection")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	// Raw string representation of the MongoDB doc _id
	idStr := "5d2399ef96fb765873a24bae"

	// The MongoDB BSON ObjectID is essentially a byte slice with a length of 12
	hexByte, err := hex.DecodeString(idStr)
	fmt.Println("ID hexByte len:", len(hexByte))
	fmt.Println("ID hexByte type:", reflect.TypeOf(hexByte))

	handleRequests()
}