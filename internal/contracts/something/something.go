package something

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ISomething,ISomething2,ISomething3"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE github.com/fluffy-bunny/sarulabsdi/internal/contracts/$GOPACKAGE ISomething,ISomething2,ISomething3

type (
	// ISomething helper
	ISomething interface {
		GetName() string
	}
	ISomething2 interface {
		ISomething
	}
	ISomething3 interface {
		ISomething
	}
)
