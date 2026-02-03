package di

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type ITestService interface {
	GetMessage() string
}

type testService struct {
	message string
}

func (s *testService) GetMessage() string {
	return s.message
}

var rTypeTestService = reflect.TypeOf(&testService{})

type TestLambda func(string) string

func TestAddSingletonT(t *testing.T) {
	b, _ := NewBuilder()

	AddSingletonT[*testService](b)

	ctn := b.Build()

	// Should be able to get by concrete type
	obj, err := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.NotNil(t, obj)
}

func TestAddScopedT(t *testing.T) {
	b, _ := NewBuilder()

	AddScopedT[*testService](b)

	ctn := b.Build()
	reqCtn, _ := ctn.SubContainer()

	// Should be able to get by concrete type in scoped container
	obj, err := SafeGetFromContainerT[*testService](reqCtn)
	require.Nil(t, err)
	require.NotNil(t, obj)
}

func TestAddTransientT(t *testing.T) {
	b, _ := NewBuilder()

	AddTransientT[*testService](b)

	ctn := b.Build()

	// Should be able to get multiple different instances
	obj1, err := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.NotNil(t, obj1)

	obj2, err := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.NotNil(t, obj2)

	// Different instances for transient
	require.NotSame(t, obj1, obj2)
}

func TestAddSingletonWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"key": "value",
	}

	AddSingletonWithMetadataT[*testService](b, metadata, rTypeTestService)

	ctn := b.Build()

	def := GetSetterDefinitionT[*testService](ctn)
	require.NotNil(t, def)
	require.Equal(t, "value", def.MetaData["key"])
}

func TestAddTransientWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"key": "transient-value",
	}

	AddTransientWithMetadataT[*testService](b, metadata)

	ctn := b.Build()

	def := GetSetterDefinitionT[*testService](ctn)
	require.NotNil(t, def)
	require.Equal(t, "transient-value", def.MetaData["key"])
	require.True(t, def.Unshared)
}

func TestAddScopedWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"key": "scoped-value",
	}

	AddScopedWithMetadataT[*testService](b, metadata)

	ctn := b.Build()
	reqCtn, _ := ctn.SubContainer()

	def := GetSetterDefinitionT[*testService](reqCtn)
	require.NotNil(t, def)
	require.Equal(t, "scoped-value", def.MetaData["key"])
	require.Equal(t, Request, def.Scope)
}

func TestAddSingletonWithImplementedTypesByFuncT(t *testing.T) {
	b, _ := NewBuilder()

	buildFunc := func(ctn Container) (interface{}, error) {
		return &testService{message: "custom"}, nil
	}

	AddSingletonWithImplementedTypesByFuncT[*testService](b, buildFunc)

	ctn := b.Build()

	obj, err := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.NotNil(t, obj)
	require.Equal(t, "custom", obj.GetMessage())
}

func TestAddTransientWithImplementedTypesByFuncT(t *testing.T) {
	b, _ := NewBuilder()

	counter := 0
	buildFunc := func(ctn Container) (interface{}, error) {
		counter++
		return &testService{message: "instance"}, nil
	}

	AddTransientWithImplementedTypesByFuncT[*testService](b, buildFunc)

	ctn := b.Build()

	obj1, _ := SafeGetFromContainerT[*testService](ctn)
	obj2, _ := SafeGetFromContainerT[*testService](ctn)

	require.NotNil(t, obj1)
	require.NotNil(t, obj2)
	require.Equal(t, 2, counter) // Build function called twice for transient
	require.NotSame(t, obj1, obj2)
}

func TestAddScopedWithImplementedTypesByFuncT(t *testing.T) {
	b, _ := NewBuilder()

	buildFunc := func(ctn Container) (interface{}, error) {
		return &testService{message: "scoped"}, nil
	}

	AddScopedWithImplementedTypesByFuncT[*testService](b, buildFunc)

	ctn := b.Build()
	reqCtn, _ := ctn.SubContainer()

	obj, err := SafeGetFromContainerT[*testService](reqCtn)
	require.Nil(t, err)
	require.NotNil(t, obj)
	require.Equal(t, "scoped", obj.GetMessage())
}

func TestAddSingletonWithImplementedTypesByFuncWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"version": "1.0",
	}

	buildFunc := func(ctn Container) (interface{}, error) {
		return &testService{message: "test"}, nil
	}

	AddSingletonWithImplementedTypesByFuncWithMetadataT[*testService](b, buildFunc, metadata)

	ctn := b.Build()

	def := GetSetterDefinitionT[*testService](ctn)
	require.NotNil(t, def)
	require.Equal(t, "1.0", def.MetaData["version"])

	obj, err := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.NotNil(t, obj)
}

func TestAddTransientWithImplementedTypesByFuncWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"type": "transient",
	}

	buildFunc := func(ctn Container) (interface{}, error) {
		return &testService{message: "test"}, nil
	}

	AddTransientWithImplementedTypesByFuncWithMetadataT[*testService](b, buildFunc, metadata)

	ctn := b.Build()

	def := GetSetterDefinitionT[*testService](ctn)
	require.NotNil(t, def)
	require.Equal(t, "transient", def.MetaData["type"])
	require.True(t, def.Unshared)
}

func TestAddScopedWithImplementedTypesByFuncWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"scope": "request",
	}

	buildFunc := func(ctn Container) (interface{}, error) {
		return &testService{message: "test"}, nil
	}

	AddScopedWithImplementedTypesByFuncWithMetadataT[*testService](b, buildFunc, metadata)

	ctn := b.Build()
	reqCtn, _ := ctn.SubContainer()

	def := GetSetterDefinitionT[*testService](reqCtn)
	require.NotNil(t, def)
	require.Equal(t, "request", def.MetaData["scope"])
	require.Equal(t, Request, def.Scope)
}

func TestAddSingletonWithImplementedTypesByObjT(t *testing.T) {
	b, _ := NewBuilder()

	obj := &testService{message: "prebuilt"}

	AddSingletonWithImplementedTypesByObjT[*testService](b, obj)

	ctn := b.Build()

	retrieved, err := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.NotNil(t, retrieved)
	require.Equal(t, "prebuilt", retrieved.GetMessage())
	require.Same(t, obj, retrieved) // Should be the exact same instance
}

func TestAddSingletonWithImplementedTypesByObjWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	obj := &testService{message: "prebuilt"}
	metadata := map[string]interface{}{
		"prebuilt": true,
	}

	AddSingletonWithImplementedTypesByObjWithMetadataT[*testService](b, obj, metadata)

	ctn := b.Build()

	def := GetSetterDefinitionT[*testService](ctn)
	require.NotNil(t, def)
	require.Equal(t, true, def.MetaData["prebuilt"])

	retrieved, err := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.Same(t, obj, retrieved)
}

func TestGetFromContainerT(t *testing.T) {
	b, _ := NewBuilder()

	AddSingletonT[*testService](b)

	ctn := b.Build()

	obj := GetFromContainer[*testService](ctn)
	require.NotNil(t, obj)
}

func TestSafeGetFromContainerT(t *testing.T) {
	b, _ := NewBuilder()

	AddSingletonT[*testService](b)

	ctn := b.Build()

	obj, err := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.NotNil(t, obj)

	// Test with missing type
	b2, _ := NewBuilder()
	ctn2 := b2.Build()

	obj2, err2 := SafeGetFromContainerT[*testService](ctn2)
	require.NotNil(t, err2)
	require.Nil(t, obj2)
}

func TestGetManyFromContainerT(t *testing.T) {
	b, _ := NewBuilder()

	// Add multiple instances
	AddTransientT[*testService](b)
	AddTransientT[*testService](b)
	AddTransientT[*testService](b)

	ctn := b.Build()

	objs := GetManyFromContainerT[*testService](ctn)
	require.Len(t, objs, 3)
}

func TestSafeGetManyFromContainerT(t *testing.T) {
	b, _ := NewBuilder()

	// Add multiple instances
	AddSingletonT[*testService](b)
	AddSingletonT[*testService](b)

	ctn := b.Build()

	objs, err := SafeGetManyFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.Len(t, objs, 2)

	// Test with missing type
	b2, _ := NewBuilder()
	ctn2 := b2.Build()

	objs2, err2 := SafeGetManyFromContainerT[*testService](ctn2)
	require.NotNil(t, err2)
	require.Nil(t, objs2)
}

func TestGetSetterDefinitionT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"test": "value",
	}

	AddSingletonWithMetadataT[*testService](b, metadata)

	ctn := b.Build()

	def := GetSetterDefinitionT[*testService](ctn)
	require.NotNil(t, def)
	require.Equal(t, "value", def.MetaData["test"])
	require.Equal(t, App, def.Scope)
}

func TestAddSingletonWithImplementedTypesT(t *testing.T) {
	b, _ := NewBuilder()

	AddSingletonWithImplementedTypesT[*testService](b)

	ctn := b.Build()

	obj, err := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err)
	require.NotNil(t, obj)
}

func TestAddTransientWithImplementedTypesWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"transient": true,
	}

	AddTransientWithImplementedTypesWithMetadataT[*testService](b, metadata)

	ctn := b.Build()

	def := GetSetterDefinitionT[*testService](ctn)
	require.NotNil(t, def)
	require.True(t, def.MetaData["transient"].(bool))
	require.True(t, def.Unshared)
}

func TestAddScopedWithImplementedTypesT(t *testing.T) {
	b, _ := NewBuilder()

	AddScopedWithImplementedTypesT[*testService](b)

	ctn := b.Build()
	reqCtn, _ := ctn.SubContainer()

	obj, err := SafeGetFromContainerT[*testService](reqCtn)
	require.Nil(t, err)
	require.NotNil(t, obj)

	// Same instance within same scope
	obj2, _ := SafeGetFromContainerT[*testService](reqCtn)
	require.Same(t, obj, obj2)
}

func TestAddScopedWithImplementedTypesWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"scoped": true,
	}

	AddScopedWithImplementedTypesWithMetadataT[*testService](b, metadata)

	ctn := b.Build()
	reqCtn, _ := ctn.SubContainer()

	def := GetSetterDefinitionT[*testService](reqCtn)
	require.NotNil(t, def)
	require.True(t, def.MetaData["scoped"].(bool))
	require.Equal(t, Request, def.Scope)
}

func TestAddSingletonWithImplementedTypesWithMetadataT(t *testing.T) {
	b, _ := NewBuilder()

	metadata := map[string]interface{}{
		"singleton": true,
		"version":   2,
	}

	AddSingletonWithImplementedTypesWithMetadataT[*testService](b, metadata)

	ctn := b.Build()

	def := GetSetterDefinitionT[*testService](ctn)
	require.NotNil(t, def)
	require.True(t, def.MetaData["singleton"].(bool))
	require.Equal(t, 2, def.MetaData["version"])
	require.Equal(t, App, def.Scope)
}

func TestGenericFunctions_WithMultipleInterfaces(t *testing.T) {
	// Test that we can register the same concrete type multiple times
	b, _ := NewBuilder()

	// Add the same concrete type multiple times
	AddSingletonT[*testService](b)
	AddSingletonT[*testService](b)

	ctn := b.Build()

	obj1, err1 := SafeGetFromContainerT[*testService](ctn)
	require.Nil(t, err1)
	require.NotNil(t, obj1)

	// Get many should return multiple registrations
	objs, err2 := SafeGetManyFromContainerT[*testService](ctn)
	require.Nil(t, err2)
	require.Len(t, objs, 2)
}
