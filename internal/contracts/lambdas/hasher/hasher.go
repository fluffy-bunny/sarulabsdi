package hasher

//go:generate genny   -pkg $GOPACKAGE     -in=../../../../genny/func-types.go -out=gen-func-$GOFILE gen "FuncType=Hash"

type (
	Hash func(interface{}) (name string, hash string, err error)
)
