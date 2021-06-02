package bouncerule

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// TODO: look into Chi for router and gorilla mux differences (not compatible with standard)

// Create database connection and wire up routes.
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbname)

	log.Printf("Connecting to %s...", connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.DB = db
	a.Router = mux.NewRouter()
	a.Router.Use(prometheusMiddleware)
	a.initializeRoutes()
}

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "bouncerulemanager_http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}

func (a *App) initializeRoutes() {
	a.Router.Path("/metrics").Handler(promhttp.Handler())
	a.Router.HandleFunc("/bounce_rules", a.getBounceRules).Methods("GET")
	a.Router.HandleFunc("/bounce_rules", a.createBounceRule).Methods("POST")
	a.Router.HandleFunc("/bounce_rules/{id:[0-9]+}", a.getBounceRule).Methods("GET")
	a.Router.HandleFunc("/bounce_rules/{id:[0-9]+}", a.updateBounceRule).Methods("PUT")
	a.Router.HandleFunc("/bounce_rules/{id:[0-9]+}", a.deleteBounceRule).Methods("DELETE")
	a.Router.HandleFunc("/bounce_rule_changes", a.getBounceRuleChanges).Methods("GET")
	a.Router.HandleFunc("/bounce_rule_changes/{id:[0-9]+}", a.getBounceRuleChangesForBounceRule).Methods("GET")
}

// Start up the application.
func (a *App) Run(addr string) {
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

func (a *App) getBounceRules(w http.ResponseWriter, r *http.Request) {
	bounceRules, err := getBounceRules(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, bounceRules)
}

func (a *App) createBounceRule(w http.ResponseWriter, r *http.Request) {
	var br BounceRule
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&br); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid bounce rule request payload")
		return
	}
	defer r.Body.Close()

	if err := br.createBounceRule(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, br)
}

func (a *App) getBounceRule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid bounce rule ID")
		return
	}

	br := BounceRule{ID: id}
	err = br.getBounceRule(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Bounce rule not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, br)
}

func (a *App) updateBounceRule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid bounce rule ID")
		return
	}

	var br BounceRule
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&br); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid bounce rule request payload")
		return
	}
	defer r.Body.Close()

	br.ID = id
	if err := br.updateBounceRule(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, br)
}

func (a *App) deleteBounceRule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid bounce rule ID")
		return
	}

	br := BounceRule{ID: id}
	if err := br.deleteBounceRule(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) getBounceRuleChanges(w http.ResponseWriter, r *http.Request) {
	bounceRuleChanges, err := getBounceRuleChanges(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, bounceRuleChanges)
}

func (a *App) getBounceRuleChangesForBounceRule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid bounce rule ID")
		return
	}

	bounceRuleChanges, err := getBounceRuleChangesForBounceRule(a.DB, id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Bounce rule changes not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, bounceRuleChanges)
}
