package di

import (
	"errors"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSafeGet(t *testing.T) {
	b, _ := NewBuilder()

	b.Add([]Def{
		{
			Name:  "object",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) {
				return &mockObject{}, nil
			},
		},
		{
			Name:  "unmakable",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) {
				return nil, errors.New("error")
			},
		},
	}...)

	app := b.Build()
	request, _ := app.SubContainer()
	subrequest, _ := request.SubContainer()

	var obj, objBis interface{}
	var err error

	_, err = app.SafeGet("object")
	require.NotNil(t, err, "should not be able to create the object from the app scope")

	_, err = request.SafeGet("undefined")
	require.NotNil(t, err, "should not be able to create an undefined object")

	_, err = request.SafeGet("unmakable")
	require.NotNil(t, err, "should not be able to create an object if there is an error in the Build function")

	// should be able to create the object from the request scope
	obj, err = request.SafeGet("object")
	require.Nil(t, err)
	require.Equal(t, &mockObject{}, obj.(*mockObject))

	// should retrieve the same object every time
	objBis, err = request.SafeGet("object")
	require.Nil(t, err)
	require.Equal(t, &mockObject{}, objBis.(*mockObject))
	require.True(t, obj == objBis)

	// should be able to create an object from a sub-container
	obj, err = subrequest.SafeGet("object")
	require.Nil(t, err)
	require.Equal(t, &mockObject{}, obj.(*mockObject))
	require.True(t, obj == objBis)
}

func TestGetByType_FromSubContainer(t *testing.T) {
	b, _ := NewBuilder()

	// claim that the added type implements an interface it doesn't support
	types := NewTypeSet()
	rt := GetInterfaceReflectType((*IGetterSetter)(nil))
	types.Add(rt)

	defer func() {
		require.Nil(t, recover(), "add and build should not panic")
	}()

	// add mockObject
	err := b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Scope:            Request,
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 1,
			}, nil
		},
		Unshared: true,
	})
	assert.NoError(t, err)
	err = b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Scope:            Request,
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 2,
			}, nil
		},
		Unshared: true,
	})
	assert.NoError(t, err)
	var app = b.Build()

	_, err = app.SafeGetByType(rt)
	require.Error(t, err)

	_, err = app.SafeGetManyByType(rt)
	require.Error(t, err)

	sub, err := app.SubContainer()
	require.Nil(t, err)
	obj, err := sub.SafeGetByType(rt)
	require.Nil(t, err)
	require.Equal(t, 2, obj.(*mockObject2).Value)

	retObjs, err := sub.SafeGetManyByType(rt)
	require.Nil(t, err)
	require.Equal(t, 2, len(retObjs))
	// make sure order is still preserved
	require.Equal(t, 2, retObjs[0].(*mockObject2).Value)
	require.Equal(t, 1, retObjs[1].(*mockObject2).Value)

}

func TestAddByType_ImplementedTypes_MustPanic(t *testing.T) {

	b, _ := NewBuilder()

	// claim that the added type implements an interface it doesn't support
	types := NewTypeSet()
	inter := GetInterfaceReflectType((*IGetterSetter)(nil))
	types.Add(inter)

	assert.Panics(t, func() {
		b.Add(Def{
			Type:             reflect.TypeOf(&mockObject{}),
			ImplementedTypes: types,
			Build: func(ctn Container) (interface{}, error) {
				return &mockObject{}, nil
			},
			Unshared: true,
		})
	})
}

func TestSafeGetByType_Empty_container(t *testing.T) {

	b, _ := NewBuilder()

	var app = b.Build()

	// try to get mockObject
	rt := reflect.TypeOf(&mockObject{}).Elem()
	_, err := app.SafeGetByType(rt)
	require.NotNil(t, err)

}
func TestAddByType_Unshared_ImplementedTypes_NoPanic(t *testing.T) {
	b, _ := NewBuilder()

	// claim that the added type implements an interface it doesn't support
	types := NewTypeSet()
	rt := GetInterfaceReflectType((*IGetterSetter)(nil))
	types.Add(rt)

	defer func() {
		require.Nil(t, recover(), "add and build should not panic")
	}()

	// add mockObject
	b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{}, nil
		},
		Unshared: true,
	})
	var app = b.Build()

	_, err := app.SafeGetByType(rt)
	require.Nil(t, err)

}
func TestTypedObject_Unshared_OneAdded_FailedRetrieve(t *testing.T) {
	b, _ := NewBuilder()

	// add mockObject
	b.Add(Def{
		Type: reflect.TypeOf(&mockObject{}),
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject{}, nil
		},
		Unshared: true,
	})

	var app = b.Build()

	// try to get mockObject2
	rt := reflect.TypeOf(&mockObject2{}).Elem()
	_, err := app.SafeGetByType(rt)
	require.NotNil(t, err)

	// try to get IGetterSetter
	rt = GetInterfaceReflectType((*IGetterSetter)(nil))
	_, err = app.SafeGetByType(rt)
	require.NotNil(t, err)

}
func TestTypedObject_Unshared_OneAdded_OneRetrieved(t *testing.T) {
	b, _ := NewBuilder()

	b.Add(Def{
		Type: reflect.TypeOf(&mockObject{}),
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject{}, nil
		},
		Unshared: true,
	})

	var app = b.Build()
	rt := reflect.TypeOf(&mockObject{}).Elem()
	obj1, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	obj2, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	// should retrieve different object every time
	require.False(t, obj1 == obj2)
}

func TestTypedObject_Unshared_ImplementedTypes_OneAdded_OneRetrieved(t *testing.T) {
	b, _ := NewBuilder()

	// claim that the added type implements an interface it does support
	types := NewTypeSet()
	rt := GetInterfaceReflectType((*IGetterSetter)(nil))
	types.Add(rt)

	err := b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 1234,
			}, nil
		},
		Unshared: true,
	})
	require.Nil(t, err)

	var app = b.Build()

	// get the object by its interface
	obj1, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	obj2, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	// should retrieve different object every time
	require.False(t, obj1 == obj2)

	require.Equal(t, 1234, obj1.(IGetterSetter).GetValue())
	require.Equal(t, 1234, obj2.(IGetterSetter).GetValue())

	obj1.(IGetterSetter).SetValue(5555)
	require.Equal(t, 5555, obj1.(IGetterSetter).GetValue())
	require.Equal(t, 1234, obj2.(IGetterSetter).GetValue())

}

func TestTypedObject_Unshared_ImplementedTypes_ManyAdded_OneRetrieved(t *testing.T) {
	b, _ := NewBuilder()

	// claim that the added type implements an interface it does support
	types := NewTypeSet()
	rt := GetInterfaceReflectType((*IGetterSetter)(nil))
	types.Add(rt)

	err := b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 1234,
			}, nil
		},
		Unshared: true,
	})
	require.Nil(t, err)

	err = b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 9999,
			}, nil
		},
		Unshared: true,
	})
	require.Nil(t, err)

	var app = b.Build()

	// get the object by its interface
	obj1, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	obj2, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	// should retrieve different object every time
	require.False(t, obj1 == obj2)

	require.Equal(t, 9999, obj1.(IGetterSetter).GetValue())
	require.Equal(t, 9999, obj2.(IGetterSetter).GetValue())

	obj1.(IGetterSetter).SetValue(5555)
	require.Equal(t, 5555, obj1.(IGetterSetter).GetValue())
	require.Equal(t, 9999, obj2.(IGetterSetter).GetValue())

}

func TestTypedObjects_Unshared_ImplementedTypes_Failed_ManyRetrieved(t *testing.T) {
	b, _ := NewBuilder()

	rt := GetInterfaceReflectType((*IGetterSetter)(nil))
	// Add 2 of the same type
	b.Add(Def{
		Type: reflect.TypeOf(&mockObject2{}),
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{}, nil
		},
		Unshared: true,
	})
	b.Add(Def{
		Type: reflect.TypeOf(&mockObject2{}),
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{}, nil
		},
		Unshared: true,
	})

	// The last object added

	var app = b.Build()

	_, err := app.SafeGetManyByType(rt)
	require.NotNil(t, err)

}

func TestTypedObjects_Unshared_ImplementedTypes_ManyAdded_ManyRetrieved(t *testing.T) {
	b, _ := NewBuilder()

	// claim that the added type implements an interface it does support
	types := NewTypeSet()
	rt := GetInterfaceReflectType((*IGetterSetter)(nil))
	types.Add(rt)

	expected1 := 1234
	expected2 := 9999
	expected3 := 5555

	// Add 2 of the same type
	b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: expected1,
			}, nil
		},
		Unshared: true,
	})
	b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: expected2,
			}, nil
		},
		Unshared: true,
	})

	// The last object added

	var app = b.Build()

	many1, err := app.SafeGetManyByType(rt)
	require.Nil(t, err)
	require.Equal(t, 2, len(many1))

	many2, err := app.SafeGetManyByType(rt)
	require.Nil(t, err)
	require.Equal(t, 2, len(many2))

	// should retrieve different object every time
	require.False(t, many1[0] == many1[1])
	require.False(t, many2[0] == many2[1])

	require.False(t, many1[0] == many2[0])
	require.False(t, many1[1] == many2[1])

	// last one added must be first in the list returned

	require.Equal(t, expected2, many1[0].(IGetterSetter).GetValue())
	require.Equal(t, expected1, many1[1].(IGetterSetter).GetValue())

	many1[0].(IGetterSetter).SetValue(expected3)
	require.Equal(t, expected3, many1[0].(IGetterSetter).GetValue())
	require.Equal(t, expected1, many1[1].(IGetterSetter).GetValue())
}

func TestTypedObjects_Unshared_ManyAdded_OneRetrieved(t *testing.T) {
	b, _ := NewBuilder()

	// Add 2 of the same type
	b.Add(Def{
		Type: reflect.TypeOf(&mockObject2{}),
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 2,
			}, nil
		},
		Unshared: true,
	})
	b.Add(Def{
		Type: reflect.TypeOf(&mockObject2{}),
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 1,
			}, nil
		},
		Unshared: true,
	})

	// The last object added

	var app = b.Build()

	// get the type of the object we want to retrieve
	rt := reflect.TypeOf(&mockObject2{}).Elem()

	obj1, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	obj2, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	// should retrieve different object every time
	require.False(t, obj1 == obj2)

	// value must be of the last one added
	exected := 1
	require.Equal(t, exected, obj1.(*mockObject2).Value)
	require.Equal(t, exected, obj2.(*mockObject2).Value)

}
func TestTypedObjects_Unshared_ManyAdded_ManyRetrieved(t *testing.T) {
	b, _ := NewBuilder()

	// Add 2 of the same type
	b.Add(Def{
		Type: reflect.TypeOf(&mockObject2{}),
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 2,
			}, nil
		},
		Unshared: true,
	})
	b.Add(Def{
		Type: reflect.TypeOf(&mockObject2{}),
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 1,
			}, nil
		},
		Unshared: true,
	})

	// The last object added

	var app = b.Build()

	// get the type of the object we want to retrieve
	rt := reflect.TypeOf(&mockObject2{}).Elem()

	many1, err := app.SafeGetManyByType(rt)
	require.Nil(t, err)

	many2, err := app.SafeGetManyByType(rt)
	require.Nil(t, err)

	// should retrieve different object every time
	require.False(t, many1[0] == many1[1])
	require.False(t, many2[0] == many2[1])

	require.False(t, many1[0] == many2[0])
	require.False(t, many1[1] == many2[1])

	// last one added must be first in the list returned
	require.Equal(t, 1, many1[0].(*mockObject2).Value)
	require.Equal(t, 2, many1[1].(*mockObject2).Value)

}

func TestUnsharedObjects(t *testing.T) {
	b, _ := NewBuilder()

	b.Add(Def{
		Name: "unshared",
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject{}, nil
		},
		Unshared: true,
	})

	var app = b.Build()

	obj1, err := app.SafeGet("unshared")
	require.Nil(t, err)

	obj2, err := app.SafeGet("unshared")
	require.Nil(t, err)

	// should retrieve different object every time
	require.False(t, obj1 == obj2)
}

func TestBuildPanic(t *testing.T) {
	b, _ := NewBuilder()

	b.Add(Def{
		Name:  "object",
		Scope: App,
		Build: func(ctn Container) (interface{}, error) {
			panic("panic in Build function")
		},
	})

	app := b.Build()

	defer func() {
		require.Nil(t, recover(), "SafeGet should not panic")
	}()

	_, err := app.SafeGet("object")
	require.NotNil(t, err, "should not panic but not be able to create the object either")
}

func TestDependencies(t *testing.T) {
	b, _ := NewBuilder()

	appObject := &mockObject{}

	b.Add([]Def{
		{
			Name:  "appObject",
			Scope: App,
			Build: func(ctn Container) (interface{}, error) {
				return appObject, nil
			},
		},
		{
			Name:  "objWithDependency",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) {
				return &mockObjectWithDependency{
					Object: ctn.Get("appObject").(*mockObject),
				}, nil
			},
		},
	}...)

	app := b.Build()
	request, _ := app.SubContainer()

	objWithDependency := request.Get("objWithDependency").(*mockObjectWithDependency)
	require.True(t, appObject == objWithDependency.Object)
}

func TestDependenciesError(t *testing.T) {
	b, _ := NewBuilder()

	b.Add([]Def{
		{
			Name:  "reqObject",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) {
				return &mockObject{}, nil
			},
		},
		{
			Name:  "objWithDependency",
			Scope: App,
			Build: func(ctn Container) (interface{}, error) {
				return &mockObjectWithDependency{
					Object: ctn.Get("reqObject").(*mockObject),
				}, nil
			},
		},
	}...)

	app := b.Build()
	request, _ := app.SubContainer()

	_, err := request.SafeGet("objWithDependency")
	require.NotNil(t, err, "an App object should not depends on a Request object")
}

func TestGet(t *testing.T) {
	b, _ := NewBuilder()

	b.Add(Def{
		Name:  "object",
		Scope: Request,
		Build: func(ctn Container) (interface{}, error) {
			return 10, nil
		},
	})

	app := b.Build()
	request, _ := app.SubContainer()

	object := request.Get("object").(int)
	require.Equal(t, 10, object)
}

func TestGetPanic(t *testing.T) {
	b, _ := NewBuilder()

	b.Add(Def{
		Name: "object",
		Build: func(ctn Container) (interface{}, error) {
			return 10, errors.New("build error")
		},
	})

	app := b.Build()

	require.Panics(t, func() {
		app.Get("object")
	})
}

func TestFill(t *testing.T) {
	b, _ := NewBuilder()

	b.Add(Def{
		Name:  "object",
		Scope: App,
		Build: func(ctn Container) (interface{}, error) {
			return 10, nil
		},
	})

	app := b.Build()

	var err error
	var object int
	var wrongType string

	err = app.Fill("unknown", &wrongType)
	require.NotNil(t, err)

	err = app.Fill("object", &wrongType)
	require.NotNil(t, err, "should have failed to fill an object with the wrong type")

	err = app.Fill("object", &object)
	require.Nil(t, err)
	require.Equal(t, 10, object)
}

func TestDeleteDuringBuild(t *testing.T) {
	built := false
	closed := false

	b, _ := NewBuilder()

	b.Add(Def{
		Name: "object",
		Build: func(ctn Container) (interface{}, error) {
			ctn.Delete()
			built = true
			return 10, nil
		},
		Close: func(obj interface{}) error {
			closed = true
			return nil
		},
	})

	app := b.Build()

	_, err := app.SafeGet("object")
	require.NotNil(t, err)
	require.True(t, app.IsClosed())
	require.True(t, built)
	require.True(t, closed)
}

func TestDeleteDuringBuildWithCloseError(t *testing.T) {
	built := false
	closed := false

	b, _ := NewBuilder()

	b.Add(Def{
		Name: "object",
		Build: func(ctn Container) (interface{}, error) {
			ctn.Delete()
			built = true
			return 10, nil
		},
		Close: func(obj interface{}) error {
			closed = true
			return errors.New("could not close object")
		},
	})

	app := b.Build()

	_, err := app.SafeGet("object")
	require.NotNil(t, err)
	require.True(t, app.IsClosed())
	require.True(t, built)
	require.True(t, closed)
}

func TestConcurrentBuild(t *testing.T) {
	var numBuild uint64
	var numClose uint64

	b, _ := NewBuilder()

	b.Add(Def{
		Name: "object",
		Build: func(ctn Container) (interface{}, error) {
			time.Sleep(250 * time.Millisecond)
			atomic.AddUint64(&numBuild, 1)
			return nil, nil
		},
		Close: func(obj interface{}) error {
			atomic.AddUint64(&numClose, 1)
			return nil
		},
	})

	app := b.Build()

	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			req, _ := app.SubContainer()
			req.Get("object")
			req.Delete()
			wg.Done()
		}()
	}

	wg.Wait()

	require.Equal(t, uint64(1), atomic.LoadUint64(&numBuild))
	require.Equal(t, uint64(0), atomic.LoadUint64(&numClose))

	app.Delete()

	require.Equal(t, uint64(1), atomic.LoadUint64(&numClose))
}

func TestTypedObjects_map_type(t *testing.T) {
	type MyMapType map[string]string
	rtPtr := reflect.TypeOf(&MyMapType{})
	rtElem := rtPtr.Elem()

	b, _ := NewBuilder()
	b.Add(Def{
		Type:     rtPtr,
		Unshared: true,
		Build: func(ctn Container) (interface{}, error) {
			result := MyMapType{
				"test": "dog",
			}
			return &result, nil
		},
	})
	var app = b.Build()
	obj := app.GetByType(rtElem).(*MyMapType)
	assert.NotNil(t, obj)
	dObj := *obj

	assert.Equal(t, "dog", dObj["test"])
}

func TestTypedObjects_slice_type(t *testing.T) {
	type MySliceType []string
	rtPtr := reflect.TypeOf(&MySliceType{})
	rtElem := rtPtr.Elem()

	b, _ := NewBuilder()
	b.Add(Def{
		Type:     rtPtr,
		Unshared: true,
		Build: func(ctn Container) (interface{}, error) {
			result := MySliceType{
				"dog",
			}
			return &result, nil
		},
	})
	var app = b.Build()
	obj := app.GetByType(rtElem).(*MySliceType)
	assert.NotNil(t, obj)
	dObj := *obj

	assert.Equal(t, "dog", dObj[0])
}

func TestTypedObjects_ReflectBuilder_panic(t *testing.T) {
	b, _ := NewBuilder()
	b.Add(Def{
		Type:     reflect.TypeOf(&mockObjectDependencyDoesNotExist{}),
		Unshared: true,
	})
	var app = b.Build()
	assert.Panics(t, func() {
		rt := reflect.TypeOf(&mockObjectDependencyDoesNotExist{}).Elem()
		app.GetByType(rt)
	})
}

func TestTypedObjects_ReflectBuilder_panic_must_not(t *testing.T) {
	b, _ := NewBuilder()
	b.Add(Def{
		Type:       reflect.TypeOf(&mockObjectDependencyDoesNotExist{}),
		Unshared:   true,
		SafeInject: true,
	})
	var app = b.Build()

	rt := reflect.TypeOf(&mockObjectDependencyDoesNotExist{}).Elem()
	obj := app.GetByType(rt).(*mockObjectDependencyDoesNotExist)
	assert.NotNil(t, obj)
	assert.Nil(t, obj.NotHere)
}
func TestTypedObjects_ReflectBuilder_ManyAdded_OneRetrieved(t *testing.T) {
	b, _ := NewBuilder()
	types := NewTypeSet()
	rt := GetInterfaceReflectType((*IGetterSetter)(nil))
	types.Add(rt)
	// Add 2 of the same type
	b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {

			return &mockObject2{
				Value: 2,
			}, nil
		},
		Unshared: true,
	})
	b.Add(Def{
		Type:             reflect.TypeOf(&mockObject2{}),
		ImplementedTypes: types,
		Build: func(ctn Container) (interface{}, error) {
			return &mockObject2{
				Value: 1,
			}, nil
		},
		Unshared: true,
	})

	b.Add(Def{
		Type:     reflect.TypeOf(&mockObject3{}),
		Unshared: true,
	})
	// The last object added

	var app = b.Build()

	// get the type of the object we want to retrieve
	rt = reflect.TypeOf(&mockObject3{}).Elem()

	obj1, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	obj2, err := app.SafeGetByType(rt)
	require.Nil(t, err)

	// should retrieve different object every time
	require.False(t, obj1 == obj2)

	// value must be of the last one added
	exected := 1
	require.Equal(t, exected, obj1.(*mockObject3).GetterSetter.GetValue())
	require.Equal(t, exected, obj2.(*mockObject3).GetterSetter.GetValue())

}
