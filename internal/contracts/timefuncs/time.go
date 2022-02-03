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

func AddTimeNow(builder *di.Builder, fnc interface{}) {
	if f, ok := fnc.(func() time.Time); ok {
		di.AddFunc(builder, f)
	} else {
		panic("timefuncs.AddTimeNow: fnc must be a func() time.Time")
	}
}
