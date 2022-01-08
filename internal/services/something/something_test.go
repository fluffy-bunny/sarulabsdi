package something

import (
	"testing"

	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_something "github.com/fluffy-bunny/sarulabsdi/internal/contracts/something"
	"github.com/stretchr/testify/require"
)

func TestSameTypeAsSingletonTransientScoped(t *testing.T) {
	var err error
	b, _ := di.NewBuilder()
	// order maters for Singleton and Transient, they are both app scoped and the last one wins
	AddSingletonISomething(b)
	AddTransientISomething(b) // winner
	AddScopedISomething(b)    // scoped last
	app := b.Build()

	request, err := app.SubContainer()
	require.Nil(t, err)
	subrequest, err := app.SubContainer()
	require.Nil(t, err)

	meSomething := contracts_something.GetISomethingFromContainer(app)
	require.NotNil(t, meSomething)
	require.Equal(t, "transient", meSomething.GetName())

	meSomething = contracts_something.GetISomethingFromContainer(request)
	require.NotNil(t, meSomething)
	require.Equal(t, "scoped", meSomething.GetName())

	meSomething = contracts_something.GetISomethingFromContainer(subrequest)
	require.NotNil(t, meSomething)
	require.Equal(t, "scoped", meSomething.GetName())
}
func TestSameTypeAsScopedSingletonTransient(t *testing.T) {
	var err error
	b, _ := di.NewBuilder()
	// order maters for Singleton and Transient, they are both app scoped and the last one wins
	AddScopedISomething(b) // scoped first
	AddSingletonISomething(b)
	AddTransientISomething(b) // winner
	app := b.Build()

	request, err := app.SubContainer()
	require.Nil(t, err)
	subrequest, err := app.SubContainer()
	require.Nil(t, err)

	meSomething := contracts_something.GetISomethingFromContainer(app)
	require.NotNil(t, meSomething)
	require.Equal(t, "transient", meSomething.GetName())

	meSomething = contracts_something.GetISomethingFromContainer(request)
	require.NotNil(t, meSomething)
	require.Equal(t, "scoped", meSomething.GetName())

	meSomething = contracts_something.GetISomethingFromContainer(subrequest)
	require.NotNil(t, meSomething)
	require.Equal(t, "scoped", meSomething.GetName())
}
func TestSameTypeAsTransientSingletonScoped(t *testing.T) {
	var err error
	b, _ := di.NewBuilder()
	// order maters for Singleton and Transient, they are both app scoped and the last one wins
	AddTransientISomething(b)
	AddSingletonISomething(b) // winner
	AddScopedISomething(b)    // scoped last
	app := b.Build()

	request, err := app.SubContainer()
	require.Nil(t, err)
	subrequest, err := app.SubContainer()
	require.Nil(t, err)

	meSomething := contracts_something.GetISomethingFromContainer(app)
	require.NotNil(t, meSomething)
	require.Equal(t, "singleton", meSomething.GetName())

	meSomething = contracts_something.GetISomethingFromContainer(request)
	require.NotNil(t, meSomething)
	require.Equal(t, "scoped", meSomething.GetName())

	meSomething = contracts_something.GetISomethingFromContainer(subrequest)
	require.NotNil(t, meSomething)
	require.Equal(t, "scoped", meSomething.GetName())

}

func TestSameTypeAsScopedTransientSingleton(t *testing.T) {
	var err error
	b, _ := di.NewBuilder()
	// order maters for Singleton and Transient, they are both app scoped and the last one wins
	AddScopedISomething(b) // scoped first
	AddTransientISomething(b)
	AddSingletonISomething(b) // winner
	app := b.Build()

	request, err := app.SubContainer()
	require.Nil(t, err)
	subrequest, err := app.SubContainer()
	require.Nil(t, err)

	meSomething := contracts_something.GetISomethingFromContainer(app)
	require.NotNil(t, meSomething)
	require.Equal(t, "singleton", meSomething.GetName())

	meSomething = contracts_something.GetISomethingFromContainer(request)
	require.NotNil(t, meSomething)
	require.Equal(t, "scoped", meSomething.GetName())

	meSomething = contracts_something.GetISomethingFromContainer(subrequest)
	require.NotNil(t, meSomething)
	require.Equal(t, "scoped", meSomething.GetName())

}
