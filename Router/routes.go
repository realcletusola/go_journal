package routes

 
import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/cletushunsu/go_journal/Handler"
)


// router function 
func NewRouter() http.Handler {

	// define chi router 
	apiRouter := chi.NewRouter()
	
	// set middlewares
	apiRouter.Use(middleware.Logger)
	apiRouter.Use(middleware.AllowContentType("application/json"))
	apiRouter.Use(middleware.CleanPath)
	apiRouter.Use(middleware.AllowContentEncoding("deflate","gzip"))

	// create routes 
	apiRouter.Get("/", handler.GetAllJournals) // get all journals route 
	apiRouter.Get("/{id}", handler.GetJournal) // get single journal by id
	apiRouter.Post("/", handler.CreateJournal) // create new journal 
	apiRouter.Put("/{id}", handler.UpdateJournal) // update journal 
	apiRouter.Delete("/{id}", handler.DeleteJournal) // delete Journal

	return apiRouter
}