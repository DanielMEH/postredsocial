package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/postredsocial/internal/infrastructure/config"
	"github.com/postredsocial/internal/infrastructure/constants"
	"github.com/postredsocial/internal/infrastructure/types"
	"github.com/postredsocial/internal/module/adapters/apis/handlers"
	"github.com/postredsocial/internal/module/adapters/services/database"
	"github.com/postredsocial/internal/module/application"
	"github.com/postredsocial/internal/module/domains/ports"

	"go.uber.org/fx"
)

func configureRouterUser(
	newPublication *handlers.PublicationAccountHandler,
	likePublication *handlers.LikePublicationHandler,
	getPublication *handlers.GetPublicationHandler,
	Hstore *types.HandlersStore, _ *config.AppSettings) {

	HandlerRouter := types.SliceHandlers{
		Prefix: "",
		Routes: []types.HandlerModule{
			{
				Route:   constants.API_ROUTER_STABLE + "/add_new_publication",
				Method:  fiber.MethodPost,
				Handler: newPublication.RunNewPublicationAccountHandler,
			},
			{
				Route:   constants.API_ROUTER_STABLE + "/all_publications",
				Method:  fiber.MethodGet,
				Handler: getPublication.RunGetPublicationHandler,
			},
		},
	}
	Hstore.Handlers = append(Hstore.Handlers, HandlerRouter)
}

func ModulePublicationsProvider() []fx.Option {
	return []fx.Option{

		// 1. Proveemos el Handler (el controlador)
		fx.Provide(handlers.NewPublicationAccountHandler),
		fx.Provide(handlers.NewLikePublicationHandler),
		fx.Provide(handlers.NewGetPublicationHandler),

		// 2. Dominios puertos
		fx.Provide(database.NewServicesDatabase,
			func(adapter *database.ServicesDatabaseAdapter) ports.PortsRepositoryPosts {
				return adapter
			}),

		// fx.Provide(usecases.NewUserUseCase),
		fx.Provide(application.NewPublicationAccountUseCase),
		fx.Provide(application.NewLikePublicationUseCase),
		fx.Provide(application.NewGetPublicationUseCase),

		// 3. Invocamos la configuración de rutas
		fx.Invoke(configureRouterUser),
	}
}
