package gettersetter

import (
	contracts_gettersetter "github.com/fluffy-bunny/sarulabsdi/internal/contracts/gettersetter"
)

type (
	getterSetterService struct {
		Value int
	}
	getterSetterContainer struct {
		GetterSetter  contracts_gettersetter.IGetterSetter   `inject:""`
		GetterSetters []contracts_gettersetter.IGetterSetter `inject:""`
	}
)

func (s *getterSetterService) GetValue() int {
	return s.Value
}
func (s *getterSetterService) SetValue(value int) {
	s.Value = value
}
