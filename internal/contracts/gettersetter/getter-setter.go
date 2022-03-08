package gettersetter

//go:generate genny   -pkg $GOPACKAGE     -in=../../../genny/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IGetterSetter,IGetterSetter2"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE github.com/fluffy-bunny/sarulabsdi/internal/contracts/$GOPACKAGE IGetterSetter,IGetterSetter2

type (
	IGetterSetter interface {
		GetValue() int
		SetValue(value int)
	}
	IGetterSetter2 interface {
		GetValue() int
		SetValue(value int)
	}
)
