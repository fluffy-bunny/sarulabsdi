package changeme

import (
	"github.com/cheekybits/genny/generic"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// InterfaceType ...
type InterfaceType generic.Type

// ReflectTypeInterfaceType used when your service claims to implement InterfaceType
var ReflectTypeInterfaceType = di.GetInterfaceReflectType((*InterfaceType)(nil))
