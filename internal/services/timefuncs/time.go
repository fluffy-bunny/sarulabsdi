package timefuncs

import (
	"reflect"
	"time"

	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_timefuncs "github.com/fluffy-bunny/sarulabsdi/internal/contracts/timefuncs"
)

type (
	timeHost struct {
		NowFunc contracts_timefuncs.TimeNow `inject:""`
	}
)

func (s *timeHost) Now() time.Time {
	return s.NowFunc()
}
func BuildBreak() contracts_timefuncs.ITime {
	return &timeHost{}
}
func AddSingletonITime(builder *di.Builder) {
	contracts_timefuncs.AddSingletonITime(builder, reflect.TypeOf(&timeHost{}))
}

var (
	// Months ...
	Months = []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
)

// NewMockITimeYearMonthDate ...
func NewMockITimeYearMonthDate(year int, month time.Month) contracts_timefuncs.TimeNow {
	return NewMockITimeDate(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// NewMockITimeYearMonthDayDate ...
func NewMockITimeYearMonthDayDate(year int, month time.Month, day int) contracts_timefuncs.TimeNow {
	return NewMockITimeDate(year, month, day, 0, 0, 0, 0, time.UTC)
}

// NewMockITimeYearMonthDayHourDate ...
func NewMockITimeYearMonthDayHourDate(year int, month time.Month, day int, hour int) contracts_timefuncs.TimeNow {
	return NewMockITimeDate(year, month, day, hour, 0, 0, 0, time.UTC)
}

// NewMockITimeYearMonthDayHourMinDate ...
func NewMockITimeYearMonthDayHourMinDate(year int, month time.Month, day int, hour int, min int) contracts_timefuncs.TimeNow {
	return NewMockITimeDate(year, month, day, hour, min, 0, 0, time.UTC)
}

// NewMockITimeDate ...
func NewMockITimeDate(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) contracts_timefuncs.TimeNow {
	mockTimeNow := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return func() time.Time {
		return mockTimeNow
	}
}

// AddTimeNow adds a singleton of Now to the container
func AddTimeNow(builder *di.Builder) {
	contracts_timefuncs.AddTimeNowFunc(builder, time.Now)
}
