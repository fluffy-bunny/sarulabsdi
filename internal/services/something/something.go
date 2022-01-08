package something

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"

	contracts_something "github.com/fluffy-bunny/sarulabsdi/internal/contracts/something"
)

type (
	service struct {
		name string
	}
)

func (s *service) GetName() string {
	return s.name
}

// AddSingletonISomething helper
func AddSingletonISomething(builder *di.Builder) {
	log.Info().Msg("DI: AddSingletonISomething")
	contracts_something.AddSingletonISomethingByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "singleton",
			}, nil
		})
}

// AddTransientISomething helper
func AddTransientISomething(builder *di.Builder) {
	log.Info().Msg("DI: AddTransientISomething")
	contracts_something.AddTransientISomethingByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "transient",
			}, nil
		})
}

// AddScopedISomething helper
func AddScopedISomething(builder *di.Builder) {
	log.Info().Msg("DI: AddScopedISomething")
	contracts_something.AddScopedISomethingByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "scoped",
			}, nil
		})
}
