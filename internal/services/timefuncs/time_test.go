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
	AddSingletonITimeHost(builder)

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

	host, err := contracts_timefuncs.SafeGetITimeHostFromContainer(app)
	require.NoError(t, err)
	require.NotNil(t, host)
	tNow := host.Now()
	fmt.Println(tNow)

	require.Equal(t, "2022-01-01 00:00:00 +0000 UTC", tNow.Format("2006-01-02 15:04:05 -0700 MST"))

}
