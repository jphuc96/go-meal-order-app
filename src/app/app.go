package app

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"git.d.foundation/datcom/backend/src/handler"
	"git.d.foundation/datcom/backend/src/service"
)

// App struct for router and db
type App struct {
	Router  *gin.Engine
	Service *service.Service
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

var SessionStore = sessions.NewCookieStore([]byte("secret"))
var allowedHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token, email, state, access_token"

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

	a.Service = service.NewService(db)
	a.Handler = handler.NewCoreHandler(a.Service, db)

	a.Router = gin.Default()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://datcom.netlify.com", "http://datcom.netlify.com", "http://localhost:3000"},
		AllowedHeaders:   []string{"Origin", "client_id", "id_token", "access_token"},
		AllowCredentials: true,
		Debug:            true,
	})
	a.Router.Use(cors)
	a.Router.Use(sessions.Sessions("default", SessionStore))
	a.SetRouters()

	return a, nil
}

func (a *App) SetRouters() {
	a.Router.POST("/auth/google/logout", a.GoogleLogout)
	a.Router.GET("/auth/google/callback", a.GoogleOauthCallback)
	a.Router.GET("/menus", a.GetLatestMenu)
	a.Router.POST("/menus", a.CreateMenu)
	a.Router.POST("/menus/:MenuID/time", a.ModifyMenuTime)
	a.Router.POST("/menus/:MenuID/items", a.AddItemToMenu)
	a.Router.DELETE("/items/:ItemID", a.DeleteItemFromMenu)
	a.Router.GET("/menus/:MenuID/users/:UserID/orders", a.GetOrdersOfUser)
	a.Router.POST("/menus/:MenuID/users/:UserID/orders", a.CreateOrModifyOrder)
	a.Router.DELETE("/menus/:MenuID/users/:UserID/orders", a.CancelAllOrderOfUser)
	a.Router.GET("/menus/:MenuID/people-in-charge", a.GetPeopleInCharge)

	a.Router.OPTIONS("/auth/google/logout", a.preflight)
	a.Router.OPTIONS("/auth/google/callback", a.preflight)
	a.Router.OPTIONS("/menus", a.preflight)
	a.Router.OPTIONS("/menus/:MenuID/time", a.preflight)
	a.Router.OPTIONS("/menus/:MenuID/items", a.preflight)
	a.Router.OPTIONS("/items/:ItemID", a.preflight)
	a.Router.OPTIONS("/menus/:MenuID/users/:UserID/orders", a.preflight)
	a.Router.OPTIONS("/menus/:MenuID/people-in-charge", a.preflight)
}

func (a *App) preflight(g *gin.Context) {
	g.Header("Access-Control-Allow-Origin", "*")
	g.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	g.JSON(http.StatusOK, struct{}{})
}

func (a *App) GoogleLogout(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.GoogleLogout(g)
}

func (a *App) GoogleOauthCallback(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	a.Handler.GoogleOauthCallback(g)
}

func (a *App) GetLatestMenu(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.GetLatestMenu(g)
}

func (a *App) CreateMenu(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.CreateMenu(g)
}

func (a *App) ModifyMenuTime(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.ModifyMenuTime(g)
}

func (a *App) AddItemToMenu(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.AddItemToMenu(g)
}

func (a *App) DeleteItemFromMenu(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.DeleteItemFromMenu(g)
}

func (a *App) GetOrdersOfUser(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.GetOrdersOfUser(g)
}

func (a *App) CreateOrModifyOrder(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.CreateOrModifyOrder(g)
}

func (a *App) CancelAllOrderOfUser(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.CancelAllOrderOfUser(g)
}

func (a *App) GetPeopleInCharge(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "application/json")
	err := a.Service.AuthCheck(g.Request)
	if err != nil {
		a.Handler.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	a.Handler.GetPeopleInCharge(g)
}

func (a *App) RunServer(host string) {
	log.Println("server is running at " + host)
	log.Fatal(a.Router.RunTLS(host, "cert.pem", "privkey.pem"))
	// log.Fatal(a.Router.Run(host))
}
