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
	AddSingletonITime(builder)
	contracts_timefuncs.AddTimeNowFunc(builder, NewMockITimeYearMonthDate(2022, time.January))

	app := builder.Build()
	require.NotNil(t, app)
	now, err := contracts_timefuncs.SafeGetTimeNowFromContainer(app)
	require.NoError(t, err)
	require.NotNil(t, now)
	d := now()
	require.NotNil(t, d)
	fmt.Println(d)

	nows, err := contracts_timefuncs.SafeGetManyTimeNowFromContainer(app)
	require.NoError(t, err)
	require.NotNil(t, nows)
	for _, now := range nows {
		d := now()
		require.NotNil(t, d)
		fmt.Println(d)
	}

	timeNowObj := contracts_timefuncs.GetITimeFromContainer(app)
	require.NotNil(t, timeNowObj)
	currentTime := timeNowObj.Now()
	actualTime := time.Now()
	require.Equal(t, currentTime.Year(), actualTime.Year())
}
