package gettersetter

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_gettersetter "github.com/fluffy-bunny/sarulabsdi/internal/contracts/gettersetter"
)

type (
	getterSetterService struct {
		Value int
	}
	getterSetterContainer struct {
		Container     di.Container                           `inject:""`
		GetterSetter  contracts_gettersetter.IGetterSetter   `inject:""`
		GetterSetters []contracts_gettersetter.IGetterSetter `inject:""`
	}
)

func (s *getterSetterContainer) Ctor() {

}

// BuildBreak is here to stop the compile becuase getterSetterService doesn't implement IGetterSetter
// a NewGetterSetter would work, but that is not how these objects are to be instantiated
func BuildBreak() contracts_gettersetter.IGetterSetter {
	return &getterSetterService{}
}
func (s *getterSetterService) GetValue() int {
	return s.Value
}
func (s *getterSetterService) SetValue(value int) {
	s.Value = value
}

func AddSingletonIGetterSetter(builder *di.Builder) {
	contracts_gettersetter.AddSingletonIGetterSetter(builder, reflect.TypeOf(&getterSetterService{}))
}
func AddScopedIGetterSetter(builder *di.Builder) {
	contracts_gettersetter.AddScopedIGetterSetter(builder, reflect.TypeOf(&getterSetterService{}))
}
func AddTransientIGetterSetter(builder *di.Builder) {
	contracts_gettersetter.AddTransientIGetterSetter(builder, reflect.TypeOf(&getterSetterService{}))
}
