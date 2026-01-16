package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliction/json")
	fmt.Fprintf(w, `{"message" : "Hello From API", "status" : "success"}`)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to Go HTTP Server! </h1>")
	fmt.Fprintf(w, "<p>Server is running successfully! </p>")

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About Page</h1>")
	fmt.Fprintf(w, "Simple HTTP Server")
}

// https://pkg.go.dev/cloud.google.com/go/cloudsqlconn

func connect() *sql.DB {


	//set JSON key file
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "key.json")

	// Cloud SQL connection details
	var (
		instanceConnectionName = "gcloud-learn-483710:us-central1:go-database" // PROJECT:REGION:INSTANCE
		user                   = "go-postgres-service-account@gcloud-learn-483710.iam"
		dbName                 = "learn-cloud-sql-go"
	)

	fmt.Println("🔄 Connecting to Cloud SQL...")
	fmt.Printf("Instance: %s\n", instanceConnectionName)
	fmt.Printf("User: %s\n", user)
	fmt.Printf("Database: %s\n", dbName)

	// Create Cloud SQL dialer with IAM authentication
	d, err := cloudsqlconn.NewDialer(context.Background(), cloudsqlconn.WithIAMAuthN())
	if err != nil {
		log.Fatalf("❌ cloudsqlconn.NewDialer failed: %v", err)
	}

	// Build DSN for pgx v4
	dsn := fmt.Sprintf("user=%s database=%s", user, dbName)

	// Parse config (pgx v4 syntax)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("❌ pgx.ParseConfig failed: %v", err)
	}

	// Set custom dialer to use Cloud SQL connector
	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceConnectionName)
	}

	// Register the config with stdlib (pgx v4)
	dbURI := stdlib.RegisterConnConfig(config)

	// Open database connection
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Fatalf("❌ sql.Open failed: %v", err)
	}

	// Test connection
	fmt.Println("🔄 Testing connection with Ping...")
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ db.Ping failed: %v", err)
	}

	fmt.Println("✅ Successfully connected to Cloud SQL!")
	return db
}
func main() {
	// Database connection
	db := connect()
	defer db.Close()

	//Routes
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/about", aboutHandler)

	//server configuration
	port := ":8080"
	fmt.Printf("🚀 Server starting on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server is not Started: ", err)
	}
}
