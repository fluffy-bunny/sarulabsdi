package singletonandscoped

import (
	"testing"

	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_singletonandscoped "github.com/fluffy-bunny/sarulabsdi/internal/contracts/singletonandscoped"
	"github.com/stretchr/testify/require"
)

func TestSameTypeAsSingletonAndScoped(t *testing.T) {
	var err error
	b, _ := di.NewBuilder()
	AddSingletonISingletonAndScoped(b)
	AddScopedISingletonAndScoped(b)
	app := b.Build()

	request, err := app.SubContainer()
	require.Nil(t, err)

	meSingleton := contracts_singletonandscoped.GetISingletonAndScopedFromContainer(app)
	require.NotNil(t, meSingleton)
	require.Equal(t, "singleton", meSingleton.GetName())

	meScoped := contracts_singletonandscoped.GetISingletonAndScopedFromContainer(request)
	require.NotNil(t, meScoped)
	require.Equal(t, "scoped", meScoped.GetName())
}
