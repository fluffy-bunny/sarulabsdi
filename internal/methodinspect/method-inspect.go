package methodinspect

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/fluffy-bunny/sarulabsdi/internal/methodset"
)

type (
	MethodInspect struct {
		Type         reflect.Type
		MethodOffset int
		Kind         reflect.Kind
	}
)

func NewMethodInspect(rt reflect.Type) (*MethodInspect, error) {
	kind := rt.Kind()
	switch kind {
	case reflect.Ptr:
		return &MethodInspect{
			Type:         rt,
			Kind:         kind,
			MethodOffset: 1,
		}, nil

	case reflect.Interface:
		return &MethodInspect{
			Type:         rt,
			Kind:         kind,
			MethodOffset: 0,
		}, nil

	}
	return nil, errors.New("only pointer or interface type is supported")
}

func getMethods(rt reflect.Type) []reflect.Method {
	var methods []reflect.Method
	for i := 0; i < rt.NumMethod(); i++ {
		method := rt.Method(i)
		if method.PkgPath != "" {
			continue
		}
		methods = append(methods, method)
	}
	return methods
}
func (s *MethodInspect) Implements(interfaceType reflect.Type) (bool, methodset.MethodSet, error) {
	switch s.Kind {
	case reflect.Ptr:
		interfaceMethodInspect, err := NewMethodInspect(interfaceType)
		if err != nil {
			return false, nil, err
		}
		if interfaceMethodInspect.Kind != reflect.Interface {
			return false, nil, fmt.Errorf("interfaceType must be an interface type")
		}
		if s.Type.Implements(interfaceType) {
			return true, nil, nil
		}
		methodsInterface := interfaceMethodInspect.GetMethods()
		methodsObj := s.GetMethods()

		result := methodsInterface.Copy()
		for method := range methodsInterface {
			if methodsObj.Contains(method) {
				result.Remove(method)
			}
		}

		return false, result, nil
	default:
		return false, nil, errors.New("only pointer can be expected")
	}

}

func (s *MethodInspect) GetMethods() methodset.MethodSet {
	methods := getMethods(s.Type)

	mSet := make(methodset.MethodSet)
	for _, method := range methods {
		builderFull := strings.Builder{}
		sPtr := ""
		switch s.Type.Kind() {
		case reflect.Ptr:
			sPtr = "*"
		}

		builderFull.WriteString(fmt.Sprintf("func (s %s%s) ", sPtr, s.Type.Name()))
		builder := strings.Builder{}
		builder.WriteString(fmt.Sprintf("%s(", method.Name))

		for i := s.MethodOffset; i < method.Type.NumIn(); i++ {
			builder.WriteString(fmt.Sprint(method.Type.In(i)))

			if i < method.Type.NumIn()-1 {
				builder.WriteString(", ")

			}
		}
		builder.WriteString(")")
		numOut := method.Type.NumOut()
		if numOut > 0 {
			if numOut > 1 {
				builder.WriteString(" (")
			} else {
				builder.WriteString(" ")
			}
			for i := 0; i < numOut; i++ {
				builder.WriteString(fmt.Sprint(method.Type.Out(i)))
				if i < numOut-1 {
					builder.WriteString(", ")
				}
			}
		}
		if numOut > 1 {
			builder.WriteString(")")
		}
		sFuncN := builder.String()
		builderFull.WriteString(sFuncN)
		mSet[builder.String()] = builderFull.String()

	}
	return mSet
}
