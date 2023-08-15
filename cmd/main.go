package main

import (
	"github.com/osvaldoabel/user-api/configs"
	webServer "github.com/osvaldoabel/user-api/internal/adapters/http/server"
)

//	@title			User API Docs
//	@version		1.0
//	@description	User API with auhtentication
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Osvaldo Abel
//	@contact.url	osvaldoabel.dev
//	@contact.email	osvalldo.abel@gmail.com

//	@license.name	MIT
//	@license.url	http://osvaldoabel.dev

// @host						localhost:8800
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	config, _ := configs.LoadConfig(".")

	ws := webServer.NewWebServer(*config)
	ws.Start()
}
