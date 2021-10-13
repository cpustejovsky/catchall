package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/cpustejovsky/catchall/logger"
	"github.com/cpustejovsky/catchall/routes"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Addr  string
	Uri   string
	Pprof string
}

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Flag and Config Setup
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":5000", "HTTP network address")
	flag.StringVar(&cfg.Uri, "uri", "mongodb://localhost:27017/catch_all", "MongoDB URI")
	flag.StringVar(&cfg.Pprof, "pprof", ":4000", "Pprof host and port")
	flag.Parse()

	// Environemntal Variables
	mongoUriFromEnv := os.Getenv("MONGO_URI")
	if mongoUriFromEnv != "" {
		cfg.Uri = mongoUriFromEnv
	}

	//Logger Setup
	logger := logger.NewLogger()

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
	logger.InfoLog.Println("Successfully connected to database!")

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: routes.Routes(logger, client),
	}
	logger.InfoLog.Printf("Starting server on %s", cfg.Addr)

	go func() {
		log.Println(http.ListenAndServe(cfg.Pprof, nil))
	}()

	// Server Start
	err = srv.ListenAndServe()
	logger.ErrorLog.Fatal(err)
}
