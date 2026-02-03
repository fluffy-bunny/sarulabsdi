package di

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBuilder(t *testing.T) {
	var b *Builder
	var err error

	_, err = NewBuilder("app", "")
	require.NotNil(t, err, "should not be able to create a Builder with an empty scope")

	_, err = NewBuilder("app", "request", "app", "subrequest")
	require.NotNil(t, err, "should not be able to create a Builder with two identical scopes")

	b, err = NewBuilder("a", "b", "c")
	require.Nil(t, err)
	require.Equal(t, ScopeList{"a", "b", "c"}, b.Scopes())

	b, err = NewBuilder()
	require.Nil(t, err)
	require.Equal(t, ScopeList{App, Request, SubRequest}, b.Scopes())
}

func TestDefinitions(t *testing.T) {
	b, _ := NewBuilder()

	b.Add([]Def{
		{
			Name:  "o1",
			Build: func(ctn Container) (interface{}, error) { return nil, nil },
		},
		{
			Name:  "o2",
			Build: func(ctn Container) (interface{}, error) { return nil, nil },
		},
	}...)

	defs := b.Definitions()

	require.Len(t, defs, 2)
	require.Equal(t, "o1", defs["o1"].Name)
	require.Equal(t, "o2", defs["o2"].Name)
}

func TestIsDefined(t *testing.T) {
	b, _ := NewBuilder()

	b.Add(Def{
		Name:  "name",
		Scope: App,
		Build: func(ctn Container) (interface{}, error) { return nil, nil },
	})

	require.True(t, b.IsDefined("name"))
	require.False(t, b.IsDefined("undefined"))
}

func TestServiceOverride(t *testing.T) {
	b, _ := NewBuilder()

	var err error

	err = b.Add(Def{
		Name: "name",
		Build: func(ctn Container) (interface{}, error) {
			return "first", nil
		},
	})
	require.Nil(t, err)

	err = b.Add(Def{
		Name: "name",
		Build: func(ctn Container) (interface{}, error) {
			return "second", nil
		},
	})
	require.Nil(t, err)

	require.Equal(t, "second", b.Build().Get("name").(string))
}

func TestAddErrors(t *testing.T) {
	b, _ := NewBuilder()

	var err error

	buildFunc := func(ctn Container) (interface{}, error) { return nil, nil }

	err = b.Add(Def{Name: "name", Scope: App, Build: buildFunc})
	require.Nil(t, err)

	err = b.Add(Def{Name: "object", Scope: "undefined", Build: buildFunc})
	require.NotNil(t, err, "should not be able to add a Def in an undefined scope")

	err = b.Add(Def{Name: "", Scope: App, Build: buildFunc})
	require.NotNil(t, err, "should not be able to add a Def if the name is empty")

	err = b.Add(Def{Name: "object", Scope: App})
	require.NotNil(t, err, "should not be able to add a Def if Build is empty")

	err = b.Add(Def{Name: "object", Scope: App, Build: buildFunc})
	require.Nil(t, err)
}

func TestSet(t *testing.T) {
	b, _ := NewBuilder()

	var err error

	err = b.Set("", "error")
	require.NotNil(t, err, "should not be able to set an object without a name")

	err = b.Set("key", "value")
	require.Nil(t, err)

	ctn := b.Build()
	require.Equal(t, "value", ctn.Get("key").(string))
}

type TestStruct struct {
	Name string
}

func (s *TestStruct) GetName() string {
	return s.Name
}

type ISomething interface {
	GetName() string
}

var rTypeTestStruct = reflect.TypeOf(&TestStruct{})

func TestRemoveAll(t *testing.T) {

	b, _ := NewBuilder()

	var err error
	obj := &TestStruct{
		Name: "hello",
	}
	types := NewTypeSet()
	inter := GetInterfaceReflectType((*ISomething)(nil))
	types.Add(inter)

	err = b.Add(Def{
		Scope:            App,
		Type:             rTypeTestStruct,
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	})
	require.Nil(t, err)

	b.RemoveAllByType(rTypeTestStruct)
	ctn := b.Build()

	obj2, err := ctn.SafeGetByType(rTypeTestStruct)
	require.Nil(t, obj2)
	require.NotNil(t, err)

	b, _ = NewBuilder()
	types = NewTypeSet()
	types.Add(inter)
	err = b.Add(Def{
		Scope:            App,
		Type:             rTypeTestStruct,
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	})
	b.RemoveAllByType(inter)
	ctn = b.Build()

	obj2, err = ctn.SafeGetByType(rTypeTestStruct.Elem())
	require.NotNil(t, obj2)
	require.Nil(t, err)

	obj2, err = ctn.SafeGetByType(inter)
	require.Nil(t, obj2)
	require.NotNil(t, err)
}
func TestBuild(t *testing.T) {
	ctn := (&Builder{}).Build()
	require.Nil(t, ctn, "should have at least one scope to use Build")

	b, _ := NewBuilder()

	b.Add([]Def{
		{
			Name:  "o1",
			Build: func(ctn Container) (interface{}, error) { return nil, nil },
		},
		{
			Name:  "o2",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) { return nil, nil },
		},
	}...)

	app := b.Build()
	require.Equal(t, App, app.Scope())
	require.Len(t, app.Definitions(), 2)
	require.Equal(t, App, app.Definitions()["o1"].Scope)
	require.Equal(t, Request, app.Definitions()["o2"].Scope)
}

func TestRemoveAllByTypeT_WithInterface(t *testing.T) {
	b, _ := NewBuilder()

	obj := &TestStruct{
		Name: "hello",
	}
	types := NewTypeSet()
	inter := GetInterfaceReflectType((*ISomething)(nil))
	types.Add(inter)

	err := b.Add(Def{
		Scope:            App,
		Type:             rTypeTestStruct,
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	})
	require.Nil(t, err)

	// Remove by interface type using generic function
	RemoveAllByTypeT[ISomething](b)
	ctn := b.Build()

	// Should not be able to get by interface type
	obj2, err := ctn.SafeGetByType(inter)
	require.Nil(t, obj2)
	require.NotNil(t, err)

	// Should still be able to get by concrete type since we only removed by interface
	obj3, err := ctn.SafeGetByType(rTypeTestStruct)
	require.NotNil(t, obj3)
	require.Nil(t, err)
}

func TestRemoveAllByTypeT_WithConcreteType(t *testing.T) {
	b, _ := NewBuilder()

	obj := &TestStruct{
		Name: "hello",
	}

	err := b.Add(Def{
		Scope: App,
		Type:  rTypeTestStruct,
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	})
	require.Nil(t, err)

	// Remove by concrete type using generic function
	RemoveAllByTypeT[*TestStruct](b)
	ctn := b.Build()

	obj2, err := ctn.SafeGetByType(rTypeTestStruct)
	require.Nil(t, obj2)
	require.NotNil(t, err)
}

type TestFunc func() string

func TestRemoveAllByTypeT_WithFuncType(t *testing.T) {
	b, _ := NewBuilder()

	// Add a function type
	testFunc := TestFunc(func() string {
		return "test"
	})

	AddFunc(b, testFunc)

	// Verify it was added
	ctn := b.Build()
	funcType := reflect.TypeOf(testFunc)
	obj, err := ctn.SafeGetByType(funcType)
	require.Nil(t, err)
	require.NotNil(t, obj)

	// Now remove using generic function with function type
	b2, _ := NewBuilder()
	AddFunc(b2, testFunc)

	RemoveAllByTypeT[TestFunc](b2)
	ctn2 := b2.Build()

	obj2, err2 := ctn2.SafeGetByType(funcType)
	require.Nil(t, obj2)
	require.NotNil(t, err2)
}

func TestRemoveAllByTypeT_WithMultipleInstances(t *testing.T) {
	b, _ := NewBuilder()

	obj1 := &TestStruct{Name: "first"}
	obj2 := &TestStruct{Name: "second"}

	types := NewTypeSet()
	inter := GetInterfaceReflectType((*ISomething)(nil))
	types.Add(inter)

	// Add multiple instances with same interface
	b.Add(Def{
		Scope:            App,
		Type:             rTypeTestStruct,
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return obj1, nil
		},
	})

	types2 := NewTypeSet()
	types2.Add(inter)
	b.Add(Def{
		Scope:            App,
		Type:             rTypeTestStruct,
		ImplementedTypes: types2,
		Build: func(ctn Container) (interface{}, error) {
			return obj2, nil
		},
	})

	// Remove all by interface type
	RemoveAllByTypeT[ISomething](b)
	ctn := b.Build()

	// Should not be able to get any instances by interface
	objs, err := ctn.SafeGetManyByType(inter)
	require.NotNil(t, err)
	require.Nil(t, objs)
}
