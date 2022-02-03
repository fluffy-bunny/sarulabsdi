package timefuncs

import (
	"reflect"
	"time"

	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	Now func() time.Time
)

var (
	RT_Now = reflect.TypeOf((func() time.Time)(nil))
)

func AddTimeNow(builder *di.Builder, fnc interface{}) {
	if f, ok := fnc.(func() time.Time); ok {
		di.AddFunc(builder, f)
	} else {
		panic("timefuncs.AddTimeNow: fnc must be a func() time.Time")
	}
}
