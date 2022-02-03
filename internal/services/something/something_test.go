package something

import (
	"testing"

	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_something "github.com/fluffy-bunny/sarulabsdi/internal/contracts/something"
	"github.com/stretchr/testify/require"
)

func TestNoConflictWithSameUnderlyingImplementation(t *testing.T) {
	var err error
	builer, _ := di.NewBuilder()
	AddSingletonISomething(builer)
	AddTransientISomething(builer)
	AddScopedISomething(builer)

	AddSingletonISomething2(builer)
	AddTransientISomething2(builer)
	AddScopedISomething2(builer)

	AddSingletonISomething3(builer)
	AddTransientISomething3(builer)
	AddScopedISomething3(builer)

	app := builer.Build()

	request, err := app.SubContainer()
	require.Nil(t, err)

	meSomething := contracts_something.GetISomethingFromContainer(app)
	require.NotNil(t, meSomething)
	require.Equal(t, "transient", meSomething.GetName())

	meSomething2 := contracts_something.GetISomething2FromContainer(app)
	require.NotNil(t, meSomething2)
	require.Equal(t, "transient2", meSomething2.GetName())

	meSomething3 := contracts_something.GetISomething3FromContainer(app)
	require.NotNil(t, meSomething3)
	require.Equal(t, "transient3", meSomething3.GetName())

	meSomething = contracts_something.GetISomethingFromContainer(request)
	require.NotNil(t, meSomething)
	require.Equal(t, "scoped", meSomething.GetName())
}

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
