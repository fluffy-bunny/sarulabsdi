package timefuncs

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/func-types.go -out=gen-func-$GOFILE gen "FuncType=TimeNow"

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITime"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE github.com/fluffy-bunny/sarulabsdi/internal/contracts/$GOPACKAGE ITime

import (
	"time"
)

type (
	ITime interface {
		Now() time.Time
	}
	TimeNow func() time.Time
)
