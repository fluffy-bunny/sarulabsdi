package something

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ISomething"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE github.com/fluffy-bunny/sarulabsdi/internal/contracts/$GOPACKAGE ISomething

type (
	// ISomething helper
	ISomething interface {
		GetName() string
	}
)
