package server

import (
	"github.com/postredsocial/internal/infrastructure/config"
	"github.com/postredsocial/internal/infrastructure/types"
	"go.uber.org/fx"
)

type ProviderServerStorage struct {
	Providers []fx.Option
}

func (ps *ProviderServerStorage) Init() {
	ps.Providers = []fx.Option{
		fx.Provide(types.NewHandlersStore),
		fx.Provide(config.AppSettingsUnmarshalnFn),
	}
}

func (ps *ProviderServerStorage) AddModule(p []fx.Option) {

	ps.Providers = append(ps.Providers, p...)
}

func (ps *ProviderServerStorage) Up(lp ...[]fx.Option) {
	ps.Providers = append(ps.Providers, fx.Invoke(NewFiberServer))
	fx.New(ps.Providers...).Run()
}
