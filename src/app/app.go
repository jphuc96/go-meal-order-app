package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"git.d.foundation/datcom/backend/src/handler"
	"git.d.foundation/datcom/backend/src/service"
)

// App struct for router and db
type App struct {
	Router  *mux.Router
	Service *service.Service
	Handler *handler.CoreHandler
}

func (a *App) NewApp() (*App, error) {
	connectionString := "user=postgres dbname=datcom sslmode=disable password=datcom host=localhost port=5432"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.New("could not connect to database")
	}
	a.Service = service.NewService(db)
	a.Handler = handler.NewCoreHandler(a.Service, db)
	a.Router = mux.NewRouter()
	a.SetRouters()

	return a, nil
}

func (a *App) SetRouters() {
	a.Router.HandleFunc("/auth/google/login-url", a.GetGoogleLoginURL).Methods("GET")
	a.Router.HandleFunc("/auth/login/google", a.VerifyGoogleUserLogin).Methods("POST")
	a.Router.HandleFunc("/menus", a.GetLatestMenu).Methods("GET")
	a.Router.HandleFunc("/menus", a.CreateMenu).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}/time", a.ModifyMenuTime).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}/items", a.AddItemToMenu).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}/items/{ItemID}", a.DeleteItemFromMenu).Methods("DELETE")
	a.Router.HandleFunc("/menus/{MenuID}/summary", a.GetMenuSummary).Methods("GET")
	a.Router.HandleFunc("/menus/{MenuID}/users/{UserID}/orders", a.GetOrdersOfUser).Methods("GET")
	a.Router.HandleFunc("/menus/{MenuID}/users/{UserID}/orders", a.CreateOrModifyOrder).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}/users/{UserID}/orders", a.CancelAllOrderOfUser).Methods("DELETE")
	a.Router.HandleFunc("/menus/{MenuID}/pic", a.GetPeopleInCharge).Methods("GET")
	a.Router.HandleFunc("/orders/{UserID}", a.AddPeopleInCharge).Methods("POST")
}

func (a *App) GetGoogleLoginURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) VerifyGoogleUserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetLatestMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) CreateMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ModifyMenuTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) AddItemToMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) DeleteItemFromMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetMenuSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetOrdersOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a.Handler.GetOrdersOfUser(w, r)
}

func (a *App) CreateOrModifyOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a.Handler.CreateOrModifyOrder(w, r)
}

func (a *App) CancelAllOrderOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a.Handler.CancelAllOrderOfUser(w, r)
}

func (a *App) GetPeopleInCharge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (a *App) AddPeopleInCharge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *App) RunServer(host string) {
	fmt.Printf("server in running at %s\n", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
