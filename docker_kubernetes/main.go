package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func apiHandler(w http.ResponseWriter, r  *http.Request){
	w.Header().Set("Content-Type" , "appliction/json")
	fmt.Fprintf(w, `{"message" : "Hello From API", "status" : "success"}`)
}

func healthHandler(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "<h1>Welcome to Go HTTP Server! </h1>")
	fmt.Fprintf(w, "<p>Server is running successfully! </p>")
	
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>About Page</h1>")
	fmt.Fprintf(w, "Simple HTTP Server")
}

//https://pkg.go.dev/cloud.google.com/go/cloudsqlconn
func connect(){
	// Register the Cloud SQL Go Connector's driver
	cleanup , err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN()) //Optional IAM AUTH
	if err != nil {
		log.Fatalf("pgxv4.RegisterDriver: %v" ,err)
	}
	defer cleanup() //important for stopping background processes

	db , err := sql.Open("cloudsql-postgres", "host=go-database:us-central1:public-instance-1 user=go-postgres-service-account@gcloud-learn-483710.iam password=Postgresql@123 dbname=learn-cloud-sql-go sslmode=disable")

	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil{
		log.Fatalf("db.Ping: %v", err)
	}
	fmt.Println("Successfully connected to Cloud SQL")

}
func main(){
	//database connection 
	connect()
	//Routes
	http.HandleFunc("/", apiHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/about", aboutHandler)

	//server configuration
	port := ":8080"

	err := http.ListenAndServe(port, nil)
	if err != nil {
	log.Fatal("Server is not Started: ", err)
	}	
}
