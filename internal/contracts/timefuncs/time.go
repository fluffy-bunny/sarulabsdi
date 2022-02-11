package timefuncs

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITime,TimeNow"

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
