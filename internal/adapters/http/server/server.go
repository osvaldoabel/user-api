package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/osvaldoabel/user-api/configs"
	_ "github.com/osvaldoabel/user-api/docs"
	authHandler "github.com/osvaldoabel/user-api/internal/adapters/http/handlers/auth"
	userHandler "github.com/osvaldoabel/user-api/internal/adapters/http/handlers/user"
	"github.com/osvaldoabel/user-api/internal/container"
	userService "github.com/osvaldoabel/user-api/internal/services/user"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	ROUTE_GET_USERS   = "/users"
	ROUTE_CREATE_USER = "/users"
	ROUTE_GET_USER    = "/users/{id}"
	ROUTE_UPDATE_USER = "/users/{id}"
	ROUTE_DELETE_USER = "/users/{id}"

	ROUTE_GENERATE_TOKEN = "/users/generate_token"
	ROUTE_DOCS           = "/docs/*"
	ROUTE_SWAGGER_DOCS   = "http://localhost:8800/docs/doc.json"
)

type WebServer struct {
	Port         int
	Router       *chi.Mux
	Configs      configs.Conf
	Dependencies container.DependencyContainer
	UserHandler  *userHandler.UserHandler
	TokenHandler *authHandler.TokenHandler
}

func NewWebServer(conf configs.Conf) WebServer {
	r := chi.NewRouter()
	deps := container.NewDependenciesContainer(conf)
	userService := userService.NewUserService(deps)

	return WebServer{
		Port:         conf.WebServerPort,
		Router:       r,
		Configs:      conf,
		Dependencies: deps,
		UserHandler:  userHandler.NewUserHandler(userService),
		TokenHandler: authHandler.NewTokenHandler(userService),
	}
}

func (s *WebServer) GetMiddlewares() []func(http.Handler) http.Handler {
	var _middlewares []func(http.Handler) http.Handler
	_middlewares = append(_middlewares, middleware.Logger)
	return _middlewares
}

func (s *WebServer) SetMiddlewares(middlewares []func(http.Handler) http.Handler) {
	for _, mid := range middlewares {
		s.Router.Use(mid)
	}
}

func (s *WebServer) Start() {
	s.LoadRoutes()
	port := fmt.Sprintf(":%v", s.Port)
	log.Default().Println("running server in port: " + port)
	log.Fatal(http.ListenAndServe(port, s.Router))

}

func (s *WebServer) LoadRoutes() {
	middlewares := s.GetMiddlewares()
	s.SetMiddlewares(middlewares)

	s.Router.Use(middleware.WithValue("jwt", s.Configs.TokenAuth))
	s.Router.Use(middleware.WithValue("JwtExperesIn", s.Configs.JWTExperesIn))

	s.Router.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(s.Configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Put(ROUTE_UPDATE_USER, s.UserHandler.UpdateUser)
		r.Delete(ROUTE_DELETE_USER, s.UserHandler.DeleteUser)
	})

	s.Router.Get(ROUTE_GET_USERS, s.UserHandler.ListUsers)
	s.Router.Get(ROUTE_GET_USER, s.UserHandler.GetUser)
	s.Router.Post(ROUTE_CREATE_USER, s.UserHandler.CreateUser)
	s.Router.Post(ROUTE_GENERATE_TOKEN, s.TokenHandler.GetToken)
	s.Router.Get(ROUTE_DOCS, httpSwagger.Handler(httpSwagger.URL(ROUTE_SWAGGER_DOCS)))
}
