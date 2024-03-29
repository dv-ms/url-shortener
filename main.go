package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/Dev-ManavSethi/url-shortener/controllers"
	"github.com/Dev-ManavSethi/url-shortener/models"
	"github.com/Dev-ManavSethi/url-shortener/utils"
)

func init() {

	models.DummyError = nil
	models.Templates, models.DummyError = template.ParseGlob("templates/*")
	utils.HandleErr(models.DummyError, "Error parsing glob from templates", "Parsed glob from templates")

	err := godotenv.Load(".env")
	utils.HandleErr(err, "Error loading env vaiables", "Loaded env variables")

	models.DummyError= nil
	models.Map = make(map[string]string)
	models.Map, models.DummyError= utils.LoadMapBackup()
	utils.HandleErr(models.DummyError, "Error loading backup map on start", "Loaded map.backup")


	models.Router = mux.NewRouter()

}

func main() {

	models.Router.HandleFunc("/", controllers.Home)
	models.Router.HandleFunc("/succedess", controllers.Success).Methods("GET")
	models.Router.HandleFunc("/search", controllers.Search).Methods("GET")

	models.Router.HandleFunc("/all", controllers.AllLinks).Methods("GET")
	models.Router.HandleFunc("/{slug:[a-zA-Z0-9]+}", controllers.Redirect).Methods("GET")

	http.Handle("/", models.Router)


	log.Println("Listening on port: "+ os.Getenv("PORT") )
	err := http.ListenAndServe(":"+os.Getenv("PORT"), models.Router)
	utils.HandleErr(err, "Error starting HTTP server", "")

}
