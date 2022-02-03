package timefuncs

import (
	"reflect"
	"time"

	di "github.com/fluffy-bunny/sarulabsdi"
	contracts_timefuncs "github.com/fluffy-bunny/sarulabsdi/internal/contracts/timefuncs"
)

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
func NewMockITimeYearMonthDate(year int, month time.Month) func() time.Time {
	return NewMockITimeDate(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// NewMockITimeYearMonthDayDate ...
func NewMockITimeYearMonthDayDate(year int, month time.Month, day int) func() time.Time {
	return NewMockITimeDate(year, month, day, 0, 0, 0, 0, time.UTC)
}

// NewMockITimeYearMonthDayHourDate ...
func NewMockITimeYearMonthDayHourDate(year int, month time.Month, day int, hour int) func() time.Time {
	return NewMockITimeDate(year, month, day, hour, 0, 0, 0, time.UTC)
}

// NewMockITimeYearMonthDayHourMinDate ...
func NewMockITimeYearMonthDayHourMinDate(year int, month time.Month, day int, hour int, min int) func() time.Time {
	return NewMockITimeDate(year, month, day, hour, min, 0, 0, time.UTC)
}

// NewMockITimeDate ...
func NewMockITimeDate(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) func() time.Time {
	mockTimeNow := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return func() time.Time {
		return mockTimeNow
	}
}

func Now() time.Time {
	return time.Now()
}

// AddTimeNowFunc adds a singleton of Now to the container
func AddTimeNowFunc(builder *di.Builder, now func() time.Time) {
	di.AddFunc(builder, now)
}

// AddTimeNow adds a singleton of Now to the container
func AddTimeNow(builder *di.Builder) {
	AddTimeNowFunc(builder, Now)
}

type (
	service struct {
		NowFunc func() time.Time `inject:""`
	}
)

func (s *service) Now() time.Time {
	return s.NowFunc()
}

func AddSingletonITimeHost(builder *di.Builder) {
	contracts_timefuncs.AddSingletonITimeHost(builder, reflect.TypeOf(&service{}))
}
