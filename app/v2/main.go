package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cpustejovsky/catchall/internal/models/mongodb"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	errorLog *log.Logger
	infoLog  *log.Logger
)

type Config struct {
	Addr string
	Uri  string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	domains  interface {
		UpdateDelivered(string) error
		UpdateBounced(string) error
		CheckStatus(string) (string, error)
	}
}

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
	}

	// Logging
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Llongfile)
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
}

func main() {
	// Flag and Config Setup
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":5000", "HTTP network address")
	flag.StringVar(&cfg.Uri, "uri", "mongodb://localhost:27017/catch_all", "MongoDB URI")
	flag.Parse()

	// Environemntal Variables
	mongoUriFromEnv := os.Getenv("MONGO_URI")
	if mongoUriFromEnv != "" {
		cfg.Uri = mongoUriFromEnv
	}

	// DB Setup
	clientOptions := options.Client().
		ApplyURI(cfg.Uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	database := client.Database("catchall_domains_v2")
	infoLog.Println("Successfully connected to database!")

	// Application and Server Initialization
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		domains:  &mongodb.DomainModelV2{DB: database},
	}

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: app.routes(),
	}
	infoLog.Printf("Starting server on %s", cfg.Addr)

	// Server Start
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
