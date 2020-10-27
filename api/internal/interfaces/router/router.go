package router

import (
	"api/internal/common/validation"
	"api/internal/registry"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

const (
	basepath        = "/api"
	healthcheckpath = basepath + "/healthcheck"
	userpath        = basepath + "/user"
	wordlistpath    = basepath + "/wordlist"
)

// Router route
type Router struct {
	Router *echo.Echo
}

// NewRouter router構造体の生成
func NewRouter() *Router {
	return &Router{
		Router: echo.New(),
	}
}

// Init Routerを初期化
func (r *Router) Init(rg *registry.Registry) {
	r.Router.Use(middleware.Logger())
	r.Router.Use(middleware.Recover())
	r.Router.Validator = &validation.CustomValidator{Validator: validator.New()}

	r.Router.GET(healthcheckpath, rg.HealthCheckH.HealthCheck())
	r.Router.POST(userpath+"/signup", rg.UserH.Signup())
	r.Router.POST(userpath+"/login", rg.UserH.Login())
	r.Router.GET(wordlistpath, rg.WordListH.GETWordList())
}

// StartServer サーバーの立ち上げ
func (r *Router) StartServer(port string) {
	fmt.Printf("Port[%s]でサーバーをスタートします。", port)
	r.Router.Start(port)
}
