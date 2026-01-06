package main

import (
	"github.com/joho/godotenv"
	"github.com/postredsocial/internal/infrastructure/server"
	modules "github.com/postredsocial/internal/module"
)

func main() {

	godotenv.Overload()

	app := server.ProviderServerStorage{}
	app.Init()
	app.AddModule(modules.ModulePublicationsProvider())
	app.Up()

}
