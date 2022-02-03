package timefuncs

import (
	"fmt"
	"testing"
	"time"

	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_timefuncs "github.com/fluffy-bunny/sarulabsdi/internal/contracts/timefuncs"
	"github.com/stretchr/testify/require"
)

func TestTimeFuncNow(t *testing.T) {
	builder, _ := di.NewBuilder()
	AddTimeNow(builder)
	AddTimeNowFunc(builder, NewMockITimeYearMonthDate(2022, time.January))
	app := builder.Build()
	require.NotNil(t, app)
	now, err := app.SafeGetByType(contracts_timefuncs.RT_Now)
	require.NoError(t, err)
	require.NotNil(t, now)
	d := now.(func() time.Time)()
	require.NotNil(t, d)
	fmt.Println(d)

	nows, err := app.SafeGetManyByType(contracts_timefuncs.RT_Now)
	require.NoError(t, err)
	require.NotNil(t, nows)
	for _, now := range nows {
		d := now.(func() time.Time)()
		require.NotNil(t, d)
		fmt.Println(d)
	}
}
