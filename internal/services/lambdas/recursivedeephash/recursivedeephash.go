package recursivedeephash

import (
	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_hasher "github.com/fluffy-bunny/sarulabsdi/internal/contracts/lambdas/hasher"

	recursive_deep_hash "github.com/panospet/recursive-deep-hash"
)

func recursiveDeepHash(obj interface{}) (name string, hash string, err error) {
	name = "recursive_deep_hash"
	hash, err = recursive_deep_hash.ConstructHash(obj)
	return
}
func AddHashFunc(builder *di.Builder) {
	contracts_hasher.AddHashFunc(builder, recursiveDeepHash)
}
