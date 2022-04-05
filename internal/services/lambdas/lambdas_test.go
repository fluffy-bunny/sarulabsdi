package timefuncs

import (
	"fmt"
	"reflect"
	"testing"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/fluffy-bunny/sarulabsdi/internal"
	contracts_lambdas_hasher "github.com/fluffy-bunny/sarulabsdi/internal/contracts/lambdas/hasher"
	contracts_lambdas_jsonhash "github.com/fluffy-bunny/sarulabsdi/internal/services/lambdas/jsonhash"
	contracts_lambdas_recursivedeephash "github.com/fluffy-bunny/sarulabsdi/internal/services/lambdas/recursivedeephash"
	"github.com/stretchr/testify/require"
)

type service struct {
	Hashers []contracts_lambdas_hasher.Hash `inject:""`
}

func TestManyHashLambdas(t *testing.T) {
	builder, _ := di.NewBuilder()
	di.AddSingleton(builder, reflect.TypeOf(&service{}))
	contracts_lambdas_jsonhash.AddHashFunc(builder)
	contracts_lambdas_recursivedeephash.AddHashFunc(builder)
	app := builder.Build()
	di.Dump(app)

	require.NotNil(t, app)
	hashFunc, err := contracts_lambdas_hasher.SafeGetHashFromContainer(app)
	require.NoError(t, err)
	require.NotNil(t, hashFunc)
	name, hashResult, err := hashFunc("test")
	require.NoError(t, err)
	require.Equal(t, "recursive_deep_hash", name)
	require.True(t, len(hashResult) > 0)
	fmt.Println(internal.PrettyJSON(struct {
		Name  string
		Hash  string
		Error error
	}{
		Name:  name,
		Hash:  hashResult,
		Error: err,
	}))
	hashFuncs, err := contracts_lambdas_hasher.SafeGetManyHashFromContainer(app)
	require.NoError(t, err)
	require.NotNil(t, hashFuncs)
	require.True(t, len(hashFuncs) == 2)
	for _, hashFunc := range hashFuncs {
		name, hashResult, err := hashFunc("test")
		require.NoError(t, err)
		require.True(t, len(hashResult) > 0)
		fmt.Println(internal.PrettyJSON(struct {
			Name  string
			Hash  string
			Error error
		}{
			Name:  name,
			Hash:  hashResult,
			Error: err,
		}))
	}

	obj, err := app.SafeGetByType(reflect.TypeOf(&service{}))
	require.NoError(t, err)
	require.NotNil(t, obj)
	hashersContainer := obj.(*service)
	require.True(t, len(hashersContainer.Hashers) == 2)
	for _, hasher := range hashersContainer.Hashers {
		name, hashResult, err := hasher("test")
		require.NoError(t, err)
		require.True(t, len(hashResult) > 0)
		fmt.Println(internal.PrettyJSON(struct {
			Name  string
			Hash  string
			Error error
		}{
			Name:  name,
			Hash:  hashResult,
			Error: err,
		}))
	}

}
