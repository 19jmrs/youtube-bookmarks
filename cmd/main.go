package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/19jmrs/youtube-bookmarks/db"
	"github.com/19jmrs/youtube-bookmarks/handlers"
	"github.com/19jmrs/youtube-bookmarks/views"
	"github.com/a-h/templ"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("hello from main")

	client, err := mongo.Connect(context.TODO(),options.Client().ApplyURI(os.Getenv("DBURI")))
	if err != nil{
		log.Fatal(err)
	}

	//userHandler := handlers.NewUserHandler(db.NewMongoUserStore(client))
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fileServer))

	r.Get("/", templ.Handler(views.Landing()).ServeHTTP)
	r.Get("/register", templ.Handler(views.Register()).ServeHTTP)
	

	http.ListenAndServe(":8000", r)

}
