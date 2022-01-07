package singletonandscoped

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ISingletonAndScoped"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE github.com/fluffy-bunny/sarulabsdi/internal/contracts/$GOPACKAGE ISingletonAndScoped

type (
	// ISingletonAndScoped helper
	ISingletonAndScoped interface {
		GetName() string
	}
)
