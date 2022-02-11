package timefuncs

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITime"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE github.com/fluffy-bunny/sarulabsdi/internal/contracts/$GOPACKAGE ITime

import (
	"reflect"
	"time"

	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	ITime interface {
		Now() time.Time
	}
	TimeNow func() time.Time
)

var (
	RT_Now = reflect.TypeOf(TimeNow(nil))
)

func AddTimeNowFunc(builder *di.Builder, fnc TimeNow) {
	di.AddFunc(builder, fnc)
}

func GetTimeNowFromContainer(ctn di.Container) TimeNow {
	obj := ctn.GetByType(RT_Now)
	if f, ok := obj.(TimeNow); ok {
		return f
	} else {
		panic("timefuncs.GetTimeNowFromContainer: obj must be a TimeNow")
	}
}

func GetManyTimeNowFromContainer(ctn di.Container) []TimeNow {
	objs := ctn.GetManyByType(RT_Now)
	var results []TimeNow
	for _, obj := range objs {
		results = append(results, obj.(TimeNow))
	}
	return results
}
func SafeGetTimeNowFromContainer(ctn di.Container) (TimeNow, error) {
	obj, err := ctn.SafeGetByType(RT_Now)
	if err != nil {
		return nil, err
	}
	return obj.(TimeNow), nil
}

func SafeGetManyTimeNowFromContainer(ctn di.Container) ([]TimeNow, error) {
	objs, err := ctn.SafeGetManyByType(RT_Now)
	if err != nil {
		return nil, err
	}
	var results []TimeNow
	for _, obj := range objs {
		results = append(results, obj.(TimeNow))
	}
	return results, nil
}
