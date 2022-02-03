package timefuncs

import (
	"reflect"
	"time"

	di "github.com/fluffy-bunny/sarulabsdi"
)

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITimeHost"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE github.com/fluffy-bunny/sarulabsdi/internal/contracts/$GOPACKAGE ITimeHost

type (
	ITimeHost interface {
		Now() time.Time
	}
)

var (
	RT_Now = reflect.TypeOf((func() time.Time)(nil))
)

func AddTimeNowFunc(builder *di.Builder, fnc interface{}) {
	if f, ok := fnc.(func() time.Time); ok {
		di.AddFunc(builder, f)
	} else {
		panic("timefuncs.AddTimeNow: fnc must be a func() time.Time")
	}
}

func GetTimeNowFromContainer(ctn di.Container) func() time.Time {
	obj := ctn.GetByType(RT_Now)
	if f, ok := obj.(func() time.Time); ok {
		return f
	} else {
		panic("timefuncs.GetTimeNowFromContainer: obj must be a func() time.Time")
	}
}

func GetManyTimeNowFromContainer(ctn di.Container) []func() time.Time {
	objs := ctn.GetManyByType(RT_Now)
	var results []func() time.Time
	for _, obj := range objs {
		results = append(results, obj.(func() time.Time))
	}
	return results
}
func SafeGetTimeNowFromContainer(ctn di.Container) (func() time.Time, error) {
	obj, err := ctn.SafeGetByType(RT_Now)
	if err != nil {
		return nil, err
	}
	return obj.(func() time.Time), nil
}

func SafeGetManyTimeNowFromContainer(ctn di.Container) ([]func() time.Time, error) {
	objs, err := ctn.SafeGetManyByType(RT_Now)
	if err != nil {
		return nil, err
	}
	var results []func() time.Time
	for _, obj := range objs {
		results = append(results, obj.(func() time.Time))
	}
	return results, nil
}
