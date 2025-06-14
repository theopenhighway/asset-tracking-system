package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	database "asset-tracking-system/db"
	// setup "asset-tracking-system/setup"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	db, err := database.ConnectDB()

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer db.Close()

	h := setup.SetupHandlers(db)
	g := setup.SetupRoutes(h)

	srv := http.Server{
		Addr:    os.Getenv("BASE_URL"),
		Handler: g,
	}

	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatalf("error server listen and serve: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)

	<-quit
	log.Println("shutdown server..")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = srv.Shutdown(ctx)

	if err != nil {
		log.Fatal("Server shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}

	log.Println("server exiting")

}
