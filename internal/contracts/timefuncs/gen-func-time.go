// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package timefuncs

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeTimeNow used when your service claims to implement TimeNow
var ReflectTypeTimeNow = reflect.TypeOf(TimeNow(nil))

// AddSingletonTimeNowFunc adds a func to the DI
func AddTimeNowFunc(builder *di.Builder, fnc TimeNow) {
	di.AddFunc(builder, fnc)
}

// RemoveAllTimeNowFunc removes all TimeNow functions from the DI
func RemoveAllTimeNowFunc(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeTimeNow)
}

// GetTimeNowFromContainer alternative to SafeGetTimeNowFromContainer but panics of object is not present
func GetTimeNowFromContainer(ctn di.Container) TimeNow {
	return ctn.GetByType(ReflectTypeTimeNow).(TimeNow)
}

// GetManyTimeNowFromContainer alternative to SafeGetManyTimeNowFromContainer but panics of object is not present
func GetManyTimeNowFromContainer(ctn di.Container) []TimeNow {
	objs := ctn.GetManyByType(ReflectTypeTimeNow)
	var results []TimeNow
	for _, obj := range objs {
		results = append(results, obj.(TimeNow))
	}
	return results
}

// SafeGetTimeNowFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetTimeNowFromContainer(ctn di.Container) (TimeNow, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeTimeNow)
	if err != nil {
		return nil, err
	}
	return obj.(TimeNow), nil
}

// SafeGetManyTimeNowFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyTimeNowFromContainer(ctn di.Container) ([]TimeNow, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeTimeNow)
	if err != nil {
		return nil, err
	}
	var results []TimeNow
	for _, obj := range objs {
		results = append(results, obj.(TimeNow))
	}
	return results, nil
}
