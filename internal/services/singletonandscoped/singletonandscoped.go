package singletonandscoped

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"

	contracts_singletonandscoped "github.com/fluffy-bunny/sarulabsdi/internal/contracts/singletonandscoped"
)

type (
	service struct {
		name string
	}
)

func (s *service) GetName() string {
	return s.name
}

// AddScopedISingletonAndScoped helper
func AddScopedISingletonAndScoped(builder *di.Builder) {
	log.Info().Msg("DI: AddScopedISingletonAndScoped")
	contracts_singletonandscoped.AddScopedISingletonAndScopedByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "scoped",
			}, nil
		})
}

// AddSingletonISingletonAndScoped helper
func AddSingletonISingletonAndScoped(builder *di.Builder) {
	log.Info().Msg("DI: AddSingletonISingletonAndScoped")
	contracts_singletonandscoped.AddSingletonISingletonAndScopedByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "singleton",
			}, nil
		})
}
