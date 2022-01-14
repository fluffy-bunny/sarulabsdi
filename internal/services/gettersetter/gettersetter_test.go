package gettersetter

import (
	"reflect"
	"testing"

	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_gettersetter "github.com/fluffy-bunny/sarulabsdi/internal/contracts/gettersetter"
	mocks_gettersetter "github.com/fluffy-bunny/sarulabsdi/internal/mocks/gettersetter"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestTypedObjects_ReflectBuilder_ManyAdded_OneRetrieved_ByFunc(t *testing.T) {
	b, _ := di.NewBuilder()

	// Add 2 of the same type
	contracts_gettersetter.AddTransientIGetterSetterByFunc(b, reflect.TypeOf(&getterSetterService{}), func(ctn di.Container) (interface{}, error) {
		return &getterSetterService{
			Value: 1,
		}, nil
	})
	contracts_gettersetter.AddTransientIGetterSetterByFunc(b, reflect.TypeOf(&getterSetterService{}), func(ctn di.Container) (interface{}, error) {

		return &getterSetterService{
			Value: 2,
		}, nil
	})
	di.AddTransient(b, reflect.TypeOf(&getterSetterContainer{}))

	// The last object added

	var app = b.Build()

	// get the type of the object we want to retrieve
	rt := reflect.TypeOf(&getterSetterContainer{})

	obj1, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	obj2, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	// should retrieve different object every time
	require.False(t, obj1 == obj2)

	// value must be of the last one added
	exected := 2
	require.Equal(t, exected, obj1.(*getterSetterContainer).GetterSetter.GetValue())
	require.Equal(t, exected, obj2.(*getterSetterContainer).GetterSetter.GetValue())

	require.Equal(t, 2, obj1.(*getterSetterContainer).GetterSetters[0].GetValue())
	require.Equal(t, 1, obj1.(*getterSetterContainer).GetterSetters[1].GetValue())

	manyGetterSetters := contracts_gettersetter.GetManyIGetterSetterFromContainer(app)
	require.NotNil(t, manyGetterSetters)
	require.NotEmpty(t, manyGetterSetters)
	require.Equal(t, 2, manyGetterSetters[0].GetValue())
	require.Equal(t, 1, manyGetterSetters[1].GetValue())

	manyGetterSetters, err = contracts_gettersetter.SafeGetManyIGetterSetterFromContainer(app)
	require.NotNil(t, manyGetterSetters)
	require.NoError(t, err)
	require.NotEmpty(t, manyGetterSetters)
	require.Equal(t, 2, manyGetterSetters[0].GetValue())
	require.Equal(t, 1, manyGetterSetters[1].GetValue())

}
func TestTypedObjects_ReflectBuilder_ManyAdded_OneRetrieved_ByObj(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGetterSetter := mocks_gettersetter.NewMockIGetterSetter(ctrl)
	mockGetterSetter.EXPECT().GetValue().Return(2).AnyTimes()
	b, _ := di.NewBuilder()

	// Add 2 of the same type
	contracts_gettersetter.AddSingletonIGetterSetterByFunc(b, reflect.TypeOf(&getterSetterService{}), func(ctn di.Container) (interface{}, error) {
		return &getterSetterService{
			Value: 1,
		}, nil
	})
	contracts_gettersetter.AddSingletonIGetterSetterByObj(b, mockGetterSetter)
	di.AddSingleton(b, reflect.TypeOf(&getterSetterContainer{}))

	// The last object added

	var app = b.Build()

	// get the type of the object we want to retrieve
	rt := reflect.TypeOf(&getterSetterContainer{})

	obj1, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	// value must be of the last one added
	exected := 2
	require.Equal(t, exected, obj1.(*getterSetterContainer).GetterSetter.GetValue())

	require.Equal(t, 2, obj1.(*getterSetterContainer).GetterSetters[0].GetValue())
	require.Equal(t, 1, obj1.(*getterSetterContainer).GetterSetters[1].GetValue())

	manyGetterSetters := contracts_gettersetter.GetManyIGetterSetterFromContainer(app)
	require.NotNil(t, manyGetterSetters)
	require.NotEmpty(t, manyGetterSetters)
	require.Equal(t, 2, manyGetterSetters[0].GetValue())
	require.Equal(t, 1, manyGetterSetters[1].GetValue())

	manyGetterSetters, err = contracts_gettersetter.SafeGetManyIGetterSetterFromContainer(app)
	require.NotNil(t, manyGetterSetters)
	require.NoError(t, err)
	require.NotEmpty(t, manyGetterSetters)
	require.Equal(t, 2, manyGetterSetters[0].GetValue())
	require.Equal(t, 1, manyGetterSetters[1].GetValue())

}
