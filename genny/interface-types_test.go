package genny

import (
	"fmt"
	"reflect"
	"testing"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type serviceInterfaceType struct {
	Name string
}
type IJunkInterfaceType interface {
}

var (
	// ReflectTypeServiceInterfaceType expored so that projects can refernce it to for go mod vendor to pull it
	ReflectTypeServiceInterfaceType = reflect.TypeOf(&serviceInterfaceType{})
	rtIJunkInterfaceType            = di.GetInterfaceReflectType((*IJunkInterfaceType)(nil))
)

func assert_singleton_InterfaceType_SafeGet_Get(t *testing.T, singletonContainer di.Container) {
	obj, err := SafeGetInterfaceTypeFromContainer(singletonContainer)
	assert.NoError(t, err)
	assert.NotNil(t, obj)

	assert.NotPanics(t, assert.PanicTestFunc(func() {
		obj = GetInterfaceTypeFromContainer(singletonContainer)

	}))
	assert.NotNil(t, obj)

	assert.Panics(t, assert.PanicTestFunc(func() {
		singletonContainer.GetByType(rtIJunkInterfaceType)
	}))

	junk, err := singletonContainer.SafeGetByType(rtIJunkInterfaceType)
	assert.Error(t, err)
	assert.Nil(t, junk)
}
func Test_singleton_InterfaceType_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	AddSingletonInterfaceType(builder, ReflectTypeServiceInterfaceType)
	singletonContainer := builder.Build()
	assert_singleton_InterfaceType_SafeGet_Get(t, singletonContainer)
}
func Test_singleton_InterfaceType_SafeGet_Get_with_RemoveAll(t *testing.T) {
	builder, _ := di.NewBuilder()
	AddSingletonInterfaceType(builder, ReflectTypeServiceInterfaceType)
	RemoveAllInterfaceType(builder)
	singletonContainer := builder.Build()
	assert.Panics(t, assert.PanicTestFunc(func() {
		GetInterfaceTypeFromContainer(singletonContainer)
	}))

	obj, err := SafeGetInterfaceTypeFromContainer(singletonContainer)
	assert.Error(t, err)
	assert.Nil(t, obj)

}
func Test_singleton_InterfaceType_ByFunc_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	obj := &serviceInterfaceType{}
	AddSingletonInterfaceTypeByFunc(builder,
		ReflectTypeServiceInterfaceType, func(ctn di.Container) (interface{}, error) {
			return obj, nil
		})
	singletonContainer := builder.Build()

	assert_singleton_InterfaceType_SafeGet_Get(t, singletonContainer)
}
func Test_singleton_InterfaceType_ByObj_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	obj := &serviceInterfaceType{}
	AddSingletonInterfaceTypeByObj(builder, obj)
	singletonContainer := builder.Build()
	assert_singleton_InterfaceType_SafeGet_Get(t, singletonContainer)
}

func assert_scoped_InterfaceType_SafeGet_Get(t *testing.T,
	singletonContainer di.Container,
	scopedContainer di.Container) {
	// cannot asked for scoped object from singleton container
	obj, err := SafeGetInterfaceTypeFromContainer(singletonContainer)
	assert.Error(t, err)
	assert.Nil(t, obj)

	assert.Panics(t, assert.PanicTestFunc(func() {
		singletonContainer.GetByType(ReflectTypeInterfaceType)
	}))

	// we can ask for it from a scped container
	obj, err = SafeGetInterfaceTypeFromContainer(scopedContainer)
	assert.NoError(t, err)
	assert.NotNil(t, obj)

	// should not panic either
	assert.NotPanics(t, assert.PanicTestFunc(func() {
		obj = GetInterfaceTypeFromContainer(scopedContainer)
	}))
	assert.NotNil(t, obj)

	// should panic if we ask for something from the scoped container that doesn't exist
	assert.Panics(t, assert.PanicTestFunc(func() {
		scopedContainer.GetByType(rtIJunkInterfaceType)
	}))

	junk, err := scopedContainer.SafeGetByType(rtIJunkInterfaceType)
	assert.Error(t, err)
	assert.Nil(t, junk)
}
func Test_scoped_InterfaceType_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	AddScopedInterfaceType(builder, ReflectTypeServiceInterfaceType)
	singletonContainer := builder.Build()
	scopedContainer, _ := singletonContainer.SubContainer()
	assert_scoped_InterfaceType_SafeGet_Get(t, singletonContainer, scopedContainer)
}
func Test_scoped_InterfaceType_ByFunc_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	obj := &serviceInterfaceType{}
	AddScopedInterfaceTypeByFunc(builder,
		ReflectTypeServiceInterfaceType, func(ctn di.Container) (interface{}, error) {
			return obj, nil
		})
	singletonContainer := builder.Build()
	scopedContainer, _ := singletonContainer.SubContainer()
	assert_scoped_InterfaceType_SafeGet_Get(t, singletonContainer, scopedContainer)
}

func assert_transient_InterfaceType_SafeGet_Get(t *testing.T,
	singletonContainer di.Container,
	scopedContainer di.Container,
	subScopedContainer di.Container,
) {
	// can get transient from anywhere

	// Singleton
	obj, err := SafeGetInterfaceTypeFromContainer(singletonContainer)
	assert.NoError(t, err)
	assert.NotNil(t, obj)

	// should not panic either
	assert.NotPanics(t, assert.PanicTestFunc(func() {
		obj = GetInterfaceTypeFromContainer(singletonContainer)
	}))

	// Scoped
	obj, err = SafeGetInterfaceTypeFromContainer(scopedContainer)
	assert.NoError(t, err)
	assert.NotNil(t, obj)

	// should not panic either
	assert.NotPanics(t, assert.PanicTestFunc(func() {
		obj = GetInterfaceTypeFromContainer(scopedContainer)
	}))

	// Sub Scoped
	obj, err = SafeGetInterfaceTypeFromContainer(subScopedContainer)
	assert.NoError(t, err)
	assert.NotNil(t, obj)

	// should not panic either
	assert.NotPanics(t, assert.PanicTestFunc(func() {
		obj = GetInterfaceTypeFromContainer(subScopedContainer)
	}))
}
func Test_transient_InterfaceType_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	AddTransientInterfaceType(builder, ReflectTypeServiceInterfaceType)
	singletonContainer := builder.Build()
	scopedContainer, _ := singletonContainer.SubContainer()
	subScopedContainer, _ := scopedContainer.SubContainer()
	assert_transient_InterfaceType_SafeGet_Get(t, singletonContainer, scopedContainer, subScopedContainer)
}
func Test_transient_InterfaceType_ByFunc_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()

	AddTransientInterfaceTypeByFunc(builder, ReflectTypeServiceInterfaceType, func(ctn di.Container) (interface{}, error) {
		return &serviceInterfaceType{}, nil
	})
	singletonContainer := builder.Build()
	scopedContainer, _ := singletonContainer.SubContainer()
	subScopedContainer, _ := scopedContainer.SubContainer()
	assert_transient_InterfaceType_SafeGet_Get(t, singletonContainer, scopedContainer, subScopedContainer)
}
func Test_singleton_InterfaceType_WithMetadata_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	metaData := map[string]interface{}{"key": "value"}
	for i := 0; i < 2; i++ {
		AddSingletonInterfaceTypeWithMetadata(builder, ReflectTypeServiceInterfaceType, metaData)
	}
	singletonContainer := builder.Build()
	scopedContainer, _ := singletonContainer.SubContainer()
	subScopedContainer, _ := scopedContainer.SubContainer()
	containers := []di.Container{singletonContainer, scopedContainer, subScopedContainer}
	for _, container := range containers {
		def := container.GetDefinitionByType(ReflectTypeServiceInterfaceType)
		require.NotNil(t, def)
		val, ok := def.MetaData["key"]
		require.True(t, ok)
		require.Equal(t, val, "value")
		require.NotPanics(t, assert.PanicTestFunc(func() {
			obj := container.Get(def.Name).(*serviceInterfaceType)
			require.NotNil(t, obj)
		}))
		defs := container.GetDefinitionsByType(ReflectTypeServiceInterfaceType)
		require.Len(t, defs, 2)
		for _, def := range defs {
			require.NotNil(t, def)
			require.NotPanics(t, assert.PanicTestFunc(func() {
				obj := container.Get(def.Name).(*serviceInterfaceType)
				require.NotNil(t, obj)
			}))
		}
		objs, err := container.SafeGetManyByType(ReflectTypeServiceInterfaceType)
		require.NoError(t, err)
		require.Len(t, objs, 2)
		require.NotPanics(t, assert.PanicTestFunc(func() {
			objs := container.GetManyByType(ReflectTypeServiceInterfaceType)
			require.Len(t, objs, 2)
		}))
	}
}
func Test_transient_InterfaceType_WithMetadata_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	metaData := map[string]interface{}{"key": "value"}
	for i := 0; i < 2; i++ {
		AddTransientInterfaceTypeWithMetadata(builder, ReflectTypeServiceInterfaceType, metaData)
	}
	singletonContainer := builder.Build()
	scopedContainer, _ := singletonContainer.SubContainer()
	subScopedContainer, _ := scopedContainer.SubContainer()
	containers := []di.Container{singletonContainer, scopedContainer, subScopedContainer}
	for _, container := range containers {
		def := container.GetDefinitionByType(ReflectTypeServiceInterfaceType)
		require.NotNil(t, def)
		val, ok := def.MetaData["key"]
		require.True(t, ok)
		require.Equal(t, val, "value")
		require.NotPanics(t, assert.PanicTestFunc(func() {
			obj := container.Get(def.Name).(*serviceInterfaceType)
			require.NotNil(t, obj)
		}))
		defs := container.GetDefinitionsByType(ReflectTypeServiceInterfaceType)
		require.Len(t, defs, 2)
		for _, def := range defs {
			require.NotNil(t, def)
			require.NotPanics(t, assert.PanicTestFunc(func() {
				obj := container.Get(def.Name).(*serviceInterfaceType)
				require.NotNil(t, obj)
			}))
		}
		objs, err := container.SafeGetManyByType(ReflectTypeServiceInterfaceType)
		require.NoError(t, err)
		require.Len(t, objs, 2)
		require.NotPanics(t, assert.PanicTestFunc(func() {
			objs := container.GetManyByType(ReflectTypeServiceInterfaceType)
			require.Len(t, objs, 2)
		}))
	}
}
func Test_scoped_InterfaceType_WithMetadata_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	metaData := map[string]interface{}{"key": "value"}
	for i := 0; i < 2; i++ {
		AddScopedInterfaceTypeWithMetadata(builder, ReflectTypeServiceInterfaceType, metaData)
	}
	singletonContainer := builder.Build()
	def := singletonContainer.GetDefinitionByType(ReflectTypeServiceInterfaceType)
	require.Nil(t, def)

	scopedContainer, _ := singletonContainer.SubContainer()
	subScopedContainer, _ := scopedContainer.SubContainer()
	containers := []di.Container{scopedContainer, subScopedContainer}
	for _, container := range containers {
		def := container.GetDefinitionByType(ReflectTypeServiceInterfaceType)
		require.NotNil(t, def)
		val, ok := def.MetaData["key"]
		require.True(t, ok)
		require.Equal(t, val, "value")
		require.NotPanics(t, assert.PanicTestFunc(func() {
			obj := container.Get(def.Name).(*serviceInterfaceType)
			require.NotNil(t, obj)
		}))
		defs := container.GetDefinitionsByType(ReflectTypeServiceInterfaceType)
		require.Len(t, defs, 2)
		for _, def := range defs {
			require.NotNil(t, def)
			require.NotPanics(t, assert.PanicTestFunc(func() {
				obj := container.Get(def.Name).(*serviceInterfaceType)
				require.NotNil(t, obj)
			}))
		}
		objs, err := container.SafeGetManyByType(ReflectTypeServiceInterfaceType)
		require.NoError(t, err)
		require.Len(t, objs, 2)
		require.NotPanics(t, assert.PanicTestFunc(func() {
			objs := container.GetManyByType(ReflectTypeServiceInterfaceType)
			require.Len(t, objs, 2)
		}))
	}
}
func Test_InterfaceType_ByObj_WithMetadata_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	metaData := map[string]interface{}{"key": "value"}
	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("%d", i)
		obj := &serviceInterfaceType{
			Name: name,
		}
		AddSingletonInterfaceTypeByObjWithMetadata(builder, obj, metaData)
	}
	singletonContainer := builder.Build()
	scopedContainer, _ := singletonContainer.SubContainer()
	subScopedContainer, _ := scopedContainer.SubContainer()
	containers := []di.Container{singletonContainer, scopedContainer, subScopedContainer}
	for _, container := range containers {
		def := container.GetDefinitionByType(ReflectTypeServiceInterfaceType)
		require.NotNil(t, def)
		val, ok := def.MetaData["key"]
		require.True(t, ok)
		require.Equal(t, val, "value")
		require.NotPanics(t, assert.PanicTestFunc(func() {
			obj := container.Get(def.Name).(*serviceInterfaceType)
			require.NotNil(t, obj)
			require.Equal(t, obj.Name, "1")
		}))
		defs := container.GetDefinitionsByType(ReflectTypeServiceInterfaceType)
		require.Len(t, defs, 2)
		i := 1
		for _, def := range defs {
			name := fmt.Sprintf("%d", i)
			i--
			require.NotNil(t, def)
			require.NotPanics(t, assert.PanicTestFunc(func() {
				obj := container.Get(def.Name).(*serviceInterfaceType)
				require.NotNil(t, obj)
				require.Equal(t, obj.Name, name)
			}))
		}
		objs, err := container.SafeGetManyByType(ReflectTypeServiceInterfaceType)
		require.NoError(t, err)
		require.Len(t, objs, 2)
		require.NotPanics(t, assert.PanicTestFunc(func() {
			objs := container.GetManyByType(ReflectTypeServiceInterfaceType)
			require.Len(t, objs, 2)
		}))
	}
}
func Test_transient_InterfaceType_ByFunc_WithMetadata_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	metaData := map[string]interface{}{"key": "value"}
	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("%d", i)
		AddTransientInterfaceTypeByFuncWithMetadata(builder, ReflectTypeServiceInterfaceType, func(ctn di.Container) (interface{}, error) {
			return &serviceInterfaceType{
				Name: name,
			}, nil
		}, metaData)
	}
	singletonContainer := builder.Build()
	scopedContainer, _ := singletonContainer.SubContainer()
	subScopedContainer, _ := scopedContainer.SubContainer()
	containers := []di.Container{singletonContainer, scopedContainer, subScopedContainer}
	for _, container := range containers {
		def := container.GetDefinitionByType(ReflectTypeServiceInterfaceType)
		require.NotNil(t, def)
		val, ok := def.MetaData["key"]
		require.True(t, ok)
		require.Equal(t, val, "value")
		require.NotPanics(t, assert.PanicTestFunc(func() {
			obj := container.Get(def.Name).(*serviceInterfaceType)
			require.NotNil(t, obj)
			require.Equal(t, obj.Name, "1")
		}))
		defs := container.GetDefinitionsByType(ReflectTypeServiceInterfaceType)
		require.Len(t, defs, 2)
		i := 1
		for _, def := range defs {
			name := fmt.Sprintf("%d", i)
			i--
			require.NotNil(t, def)
			require.NotPanics(t, assert.PanicTestFunc(func() {
				obj := container.Get(def.Name).(*serviceInterfaceType)
				require.NotNil(t, obj)
				require.Equal(t, obj.Name, name)
			}))
		}
		objs, err := container.SafeGetManyByType(ReflectTypeServiceInterfaceType)
		require.NoError(t, err)
		require.Len(t, objs, 2)
		require.NotPanics(t, assert.PanicTestFunc(func() {
			objs := container.GetManyByType(ReflectTypeServiceInterfaceType)
			require.Len(t, objs, 2)
		}))
	}
}
func Test_singleton_InterfaceType_ByFunc_WithMetadata_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	metaData := map[string]interface{}{"key": "value"}
	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("%d", i)
		AddSingletonInterfaceTypeByFuncWithMetadata(builder, ReflectTypeServiceInterfaceType, func(ctn di.Container) (interface{}, error) {
			return &serviceInterfaceType{
				Name: name,
			}, nil
		}, metaData)
	}
	singletonContainer := builder.Build()
	scopedContainer, _ := singletonContainer.SubContainer()
	subScopedContainer, _ := scopedContainer.SubContainer()
	containers := []di.Container{singletonContainer, scopedContainer, subScopedContainer}
	for _, container := range containers {
		def := container.GetDefinitionByType(ReflectTypeServiceInterfaceType)
		require.NotNil(t, def)
		val, ok := def.MetaData["key"]
		require.True(t, ok)
		require.Equal(t, val, "value")
		require.NotPanics(t, assert.PanicTestFunc(func() {
			obj := container.Get(def.Name).(*serviceInterfaceType)
			require.NotNil(t, obj)
			require.Equal(t, obj.Name, "1")
		}))
		defs := container.GetDefinitionsByType(ReflectTypeServiceInterfaceType)
		require.Len(t, defs, 2)
		i := 1
		for _, def := range defs {
			name := fmt.Sprintf("%d", i)
			i--
			require.NotNil(t, def)
			require.NotPanics(t, assert.PanicTestFunc(func() {
				obj := container.Get(def.Name).(*serviceInterfaceType)
				require.NotNil(t, obj)
				require.Equal(t, obj.Name, name)
			}))
		}
		objs, err := container.SafeGetManyByType(ReflectTypeServiceInterfaceType)
		require.NoError(t, err)
		require.Len(t, objs, 2)
		require.NotPanics(t, assert.PanicTestFunc(func() {
			objs := container.GetManyByType(ReflectTypeServiceInterfaceType)
			require.Len(t, objs, 2)
		}))
	}
}
func Test_scoped_InterfaceType_ByFunc_WithMetadata_SafeGet_Get(t *testing.T) {
	builder, _ := di.NewBuilder()
	metaData := map[string]interface{}{"key": "value"}
	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("%d", i)
		AddScopedInterfaceTypeByFuncWithMetadata(builder, ReflectTypeServiceInterfaceType, func(ctn di.Container) (interface{}, error) {
			return &serviceInterfaceType{
				Name: name,
			}, nil
		}, metaData)
	}
	singletonContainer := builder.Build()
	def := singletonContainer.GetDefinitionByType(ReflectTypeServiceInterfaceType)
	require.Nil(t, def)

	scopedContainer, _ := singletonContainer.SubContainer()
	subScopedContainer, _ := scopedContainer.SubContainer()
	containers := []di.Container{scopedContainer, subScopedContainer}

	for _, container := range containers {
		def := container.GetDefinitionByType(ReflectTypeServiceInterfaceType)
		require.NotNil(t, def)
		val, ok := def.MetaData["key"]
		require.True(t, ok)
		require.Equal(t, val, "value")
		require.NotPanics(t, assert.PanicTestFunc(func() {
			obj := container.Get(def.Name).(*serviceInterfaceType)
			require.NotNil(t, obj)
			require.Equal(t, obj.Name, "1")
		}))
		defs := container.GetDefinitionsByType(ReflectTypeServiceInterfaceType)
		require.Len(t, defs, 2)
		i := 1
		for _, def := range defs {
			name := fmt.Sprintf("%d", i)
			i--
			require.NotNil(t, def)
			require.NotPanics(t, assert.PanicTestFunc(func() {
				obj := container.Get(def.Name).(*serviceInterfaceType)
				require.NotNil(t, obj)
				require.Equal(t, obj.Name, name)
			}))
		}
		objs, err := container.SafeGetManyByType(ReflectTypeServiceInterfaceType)
		require.NoError(t, err)
		require.Len(t, objs, 2)
		require.NotPanics(t, assert.PanicTestFunc(func() {
			objs := container.GetManyByType(ReflectTypeServiceInterfaceType)
			require.Len(t, objs, 2)
		}))
	}
}
