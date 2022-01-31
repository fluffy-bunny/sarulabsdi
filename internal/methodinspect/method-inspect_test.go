package methodinspect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type (
	ISomething interface {
		Get() string
	}
	Something    struct{}
	NotSomething struct{}
)

var RT_ISomething = reflect.TypeOf((*ISomething)(nil)).Elem()
var RT_Something = reflect.TypeOf(&Something{})
var RT_NotSomething = reflect.TypeOf(&NotSomething{})

func (s *Something) Get() string {
	return "Something"
}
func (s *NotSomething) Hello() string {
	return "hello there"
}
func TestDoesImplement(t *testing.T) {
	objRT, err := NewMethodInspect(RT_Something)
	require.NoError(t, err)
	require.NotNil(t, objRT)

	interfaceRT, err := NewMethodInspect(RT_ISomething)
	require.NoError(t, err)
	require.NotNil(t, interfaceRT)

	ok, methods, err := objRT.Implements(interfaceRT.Type)
	require.NoError(t, err)
	require.True(t, ok)
	require.Nil(t, methods)
}

func TestDoesNotImplement(t *testing.T) {
	objRT, err := NewMethodInspect(RT_NotSomething)
	require.NoError(t, err)
	require.NotNil(t, objRT)

	interfaceRT, err := NewMethodInspect(RT_ISomething)
	require.NoError(t, err)
	require.NotNil(t, interfaceRT)

	ok, methods, err := objRT.Implements(interfaceRT.Type)
	require.NoError(t, err)
	require.False(t, ok)
	require.NotNil(t, methods)

	fmt.Println(methods)
}
