package throughputrule

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *chi.Mux
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbname)

	log.Printf("Connecting to %s...", connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.DB = db
	a.Router = chi.NewRouter()
	a.Router.Use(middleware.Logger)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.Route("/throughput_rules", func(r chi.Router) {
		r.Get("/", a.getThroughputRules)
		r.Post("/", a.createThroughputRule)
		r.Get("/{id:[0-9]+}", a.getThroughputRule)
		r.Put("/{id:[0-9]+}", a.updateThroughputRule)
		r.Delete("/{id:[0-9]+}", a.deleteThroughputRule)
	})

	a.Router.Route("/throughput_rule_changes", func(r chi.Router) {
		r.Get("/", a.getThroughputRuleChanges)
		r.Get("/{id:[0-9]+}", a.getThroughputRuleChangesForThroughputRule)
	})
}

func (a *App) Run(addr string) {
	log.Printf("Starting up server with addr %s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getThroughputRules(w http.ResponseWriter, r *http.Request) {
	throughputRules, err := getThroughputRules(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, throughputRules)
}

func (a *App) createThroughputRule(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, []string{"throughput1", "throughput2"})
}

func (a *App) getThroughputRule(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid throughput rule ID")
		return
	}
	log.Printf("Getting throughput rule with id %d", id)
	respondWithJSON(w, http.StatusOK, []string{"throughput1", "throughput2"})
}

func (a *App) updateThroughputRule(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid throughput rule ID")
		return
	}
	log.Printf("Updating throughput rule with id %d", id)
	respondWithJSON(w, http.StatusOK, []string{"throughput1", "throughput2"})
}

func (a *App) deleteThroughputRule(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid throughput rule ID")
		return
	}
	log.Printf("Deleting throughput rule with id %d", id)
	respondWithJSON(w, http.StatusOK, []string{"throughput1", "throughput2"})
}

func (a *App) getThroughputRuleChanges(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, []string{"throughputchange1", "throughputchange2"})
}

func (a *App) getThroughputRuleChangesForThroughputRule(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid throughput rule ID")
		return
	}
	log.Printf("Getting throughput rule changes for throughput rule with id %d", id)
	respondWithJSON(w, http.StatusOK, []string{"throughputchange1", "throughputchange2"})
}
