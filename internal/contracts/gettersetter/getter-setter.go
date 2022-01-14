package gettersetter

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IGetterSetter"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE github.com/fluffy-bunny/sarulabsdi/internal/contracts/$GOPACKAGE IGetterSetter

type (
	IGetterSetter interface {
		GetValue() int
		SetValue(value int)
	}
)
