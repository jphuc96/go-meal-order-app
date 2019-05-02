package app

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"git.d.foundation/datcom/backend/src/handler"
	"git.d.foundation/datcom/backend/src/service"
)

// App struct for router and db
type App struct {
	Router  *mux.Router
	Handler *handler.CoreHandler
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (a *App) NewApp(dbConfig *DBConfig) (*App, error) {

	connString := "host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.DBName +
		" sslmode=" + dbConfig.SSLMode

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, errors.New("could not connect to database")
	}
	log.Println("database config:")
	log.Println("host: " + dbConfig.Host)
	log.Println("port: " + dbConfig.Port)
	log.Println("user: " + dbConfig.User)
	log.Println("password: " + dbConfig.Password)
	log.Println("ssl mode: " + dbConfig.SSLMode)

	svc := service.NewService(db)
	a.Handler = handler.NewCoreHandler(svc, db)
	a.Router = mux.NewRouter()
	a.SetRouters()

	return a, nil
}

func (a *App) SetRouters() {
	a.Router.HandleFunc("/auth/google/login-url", a.GetGoogleLoginURL).Methods("GET")
	a.Router.HandleFunc("/auth/google/callback", a.VerifyGoogleUserLogin).Methods("GET")
	a.Router.HandleFunc("/menus", a.GetLatestMenu).Methods("GET")
	a.Router.HandleFunc("/menus", a.CreateMenu).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}/time", a.ModifyMenuTime).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}/items", a.AddItemToMenu).Methods("POST")
	a.Router.HandleFunc("/items/{ItemID}", a.DeleteItemFromMenu).Methods("DELETE")
	a.Router.HandleFunc("/menus/{MenuID}/users/{UserID}/orders", a.GetOrdersOfUser).Methods("GET")
	a.Router.HandleFunc("/menus/{MenuID}/users/{UserID}/orders", a.CreateOrModifyOrder).Methods("POST")
	a.Router.HandleFunc("/menus/{MenuID}/users/{UserID}/orders", a.CancelAllOrderOfUser).Methods("DELETE")
	a.Router.HandleFunc("/menus/{MenuID}/people-in-charge", a.GetPeopleInCharge).Methods("GET")
}

func (a *App) GetGoogleLoginURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.GetGoogleLoginURL(w, r)
}

func (a *App) VerifyGoogleUserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.VerifyGoogleUserLogin(w, r)
}

func (a *App) GetLatestMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.GetLatestMenu(w, r)
}

func (a *App) CreateMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.CreateMenu(w, r)
}

func (a *App) ModifyMenuTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.ModifyMenuTime(w, r)
}

func (a *App) AddItemToMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.AddItemToMenu(w, r)
}

func (a *App) DeleteItemFromMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.DeleteItemFromMenu(w, r)
}

func (a *App) GetOrdersOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.GetOrdersOfUser(w, r)
}

func (a *App) CreateOrModifyOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.CreateOrModifyOrder(w, r)
}

func (a *App) CancelAllOrderOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a.Handler.CancelAllOrderOfUser(w, r)
}

func (a *App) GetPeopleInCharge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

}

func (a *App) RunServer(host string) {
	log.Println("server is running at " + host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
