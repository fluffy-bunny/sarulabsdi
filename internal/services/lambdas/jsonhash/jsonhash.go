package jsonhash

import (
	"encoding/json"

	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_hasher "github.com/fluffy-bunny/sarulabsdi/internal/contracts/lambdas/hasher"

	"crypto/sha1"
	"fmt"
)

func sha1Hash(obj interface{}) (name string, hash string, err error) {
	name = "json sha1"
	jsonObj, err := json.Marshal(obj)
	if err == nil {
		h := sha1.New()
		h.Write(jsonObj)
		bs := h.Sum(nil)
		hash = fmt.Sprintf("%x", bs)
	}
	return
}

func AddHashFunc(builder *di.Builder) {
	contracts_hasher.AddHashFunc(builder, sha1Hash)
}
