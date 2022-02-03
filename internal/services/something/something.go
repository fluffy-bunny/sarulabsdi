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

// BuildBreak is here to stop the compile becuase getterSetterService doesn't implement IGetterSetter
// a NewGetterSetter would work, but that is not how these objects are to be instantiated
func BuildBreak() contracts_something.ISomething {
	return &service{}
}
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

// AddSingletonISomething2 helper
func AddSingletonISomething2(builder *di.Builder) {
	log.Info().Msg("DI: AddSingletonISomething2")
	contracts_something.AddSingletonISomething2ByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "singleton2",
			}, nil
		})
} // AddSingletonISomething3 helper
func AddSingletonISomething3(builder *di.Builder) {
	log.Info().Msg("DI: AddSingletonISomething3")
	contracts_something.AddSingletonISomething3ByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "singleton3",
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

// AddTransientISomething helper
func AddTransientISomething2(builder *di.Builder) {
	log.Info().Msg("DI: AddTransientISomething2")
	contracts_something.AddTransientISomething2ByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "transient2",
			}, nil
		})
}

// AddTransientISomething helper
func AddTransientISomething3(builder *di.Builder) {
	log.Info().Msg("DI: AddTransientISomething3")
	contracts_something.AddTransientISomething3ByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "transient3",
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

// AddScopedISomething2 helper
func AddScopedISomething2(builder *di.Builder) {
	log.Info().Msg("DI: AddScopedISomething2")
	contracts_something.AddScopedISomething2ByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "scoped2",
			}, nil
		})
}

// AddScopedISomething3 helper
func AddScopedISomething3(builder *di.Builder) {
	log.Info().Msg("DI: AddScopedISomething3")
	contracts_something.AddScopedISomething3ByFunc(builder, reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				name: "scoped3",
			}, nil
		})
}
