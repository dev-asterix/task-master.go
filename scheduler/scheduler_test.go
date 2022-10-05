package scheduler

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewIntervalBased(t *testing.T) {
	type args struct {
		freq     int64
		unit     Duration
		timezone *time.Location
	}
	tests := []struct {
		name string
		args args
		want *Interval
	}{
		{
			name: "new 1 second interval based scheduler",
			args: args{
				freq:     1,
				unit:     Second,
				timezone: time.Local,
			},
			want: &Interval{
				frequency: 1,
				unit:      Second,
				timezone:  time.Local,
			},
		},
		{
			name: "new 1 month interval based scheduler",
			args: args{
				freq: 1,
				unit: Month,
			},
			want: &Interval{
				frequency: 1,
				unit:      Month,
				timezone:  time.Local,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIntervalBased(tt.args.freq, tt.args.unit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntervalBased() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Frequency(t *testing.T) {
	type fields struct {
		frequency int64
		unit      Duration
		timezone  *time.Location
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1 nano second",
			fields: fields{
				frequency: 1,
				unit:      NanoSecond,
				timezone:  time.Local,
			},
			want: "1 * ns",
		},
		{
			name: "1 micro second",
			fields: fields{
				frequency: 1,
				unit:      MicroSecond,
				timezone:  time.Local,
			},
			want: "1 * µs",
		},
		{
			name: "1 milli second",
			fields: fields{
				frequency: 1,
				unit:      MilliSecond,
				timezone:  time.Local,
			},
			want: "1 * ms",
		},
		{
			name: "1 second",
			fields: fields{
				frequency: 1,
				unit:      Second,
				timezone:  time.Local,
			},
			want: "1 * sec",
		},
		{
			name: "1 month",
			fields: fields{
				frequency: 1,
				unit:      Month,
				timezone:  time.Local,
			},
			want: "1 * month",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := NewIntervalBased(tt.fields.frequency, tt.fields.unit)
			if got := i.Frequency(); got != tt.want {
				t.Errorf("Interval.Frequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Next(t *testing.T) {
	type fields struct {
		frequency int64
		unit      Duration
		timezone  *time.Location
	}

	timeNow = func(loc *time.Location) time.Time {
		return time.Date(2020, 1, 1, 0, 0, 0, 0, loc)
	}

	tests := []struct {
		name   string
		fields fields
		want   *Interval
	}{
		{
			name: "1000 nano second",
			fields: fields{
				frequency: 1000,
				unit:      NanoSecond,
				timezone:  time.Local,
			},
			want: &Interval{
				frequency: 1000,
				unit:      NanoSecond,
				nextRun:   time.Date(2020, 1, 1, 0, 0, 0, 1000, time.Local),
				timezone:  time.Local,
			},
		},
		{
			name: "500 micro second",
			fields: fields{
				frequency: 500,
				unit:      MicroSecond,
				timezone:  time.Local,
			},
			want: &Interval{
				frequency: 500,
				unit:      MicroSecond,
				nextRun:   time.Date(2020, 1, 1, 0, 0, 0, 500000, time.Local),
				timezone:  time.Local,
			},
		},
		{
			name: "200 milli second",
			fields: fields{
				frequency: 200,
				unit:      MilliSecond,
				timezone:  time.Local,
			},
			want: &Interval{
				frequency: 200,
				unit:      MilliSecond,
				nextRun:   time.Date(2020, 1, 1, 0, 0, 0, 200000000, time.Local),
				timezone:  time.Local,
			},
		},
		{
			name: "1 second",
			fields: fields{
				frequency: 1,
				unit:      Second,
				timezone:  time.Local,
			},
			want: &Interval{
				frequency: 1,
				unit:      Second,
				nextRun:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.Local),
				timezone:  time.Local,
			},
		},
		{
			name: "5 minutes",
			fields: fields{
				frequency: 5,
				unit:      Minute,
				timezone:  time.Local,
			},
			want: &Interval{
				frequency: 5,
				unit:      Minute,
				nextRun:   time.Date(2020, 1, 1, 0, 5, 0, 0, time.Local),
				timezone:  time.Local,
			},
		},
		{
			name: "6 hours",
			fields: fields{
				frequency: 6,
				unit:      Hour,
				timezone:  time.Local,
			},
			want: &Interval{
				frequency: 6,
				unit:      Hour,
				nextRun:   time.Date(2020, 1, 1, 6, 0, 0, 0, time.Local),
				timezone:  time.Local,
			},
		},
		{
			name: "15 days",
			fields: fields{
				frequency: 15,
				unit:      Day,
				timezone:  time.Local,
			},
			want: &Interval{
				frequency: 15,
				unit:      Day,
				nextRun:   time.Date(2020, 1, 16, 0, 0, 0, 0, time.Local),
				timezone:  time.Local,
			},
		},
		{
			name: "3 weeks",
			fields: fields{
				frequency: 3,
				unit:      Week,
				timezone:  time.Local,
			},
			want: &Interval{
				frequency: 3,
				unit:      Week,
				nextRun:   time.Date(2020, 1, 22, 0, 0, 0, 0, time.Local),
				timezone:  time.Local,
			},
		},
		{
			name: "2 years",
			fields: fields{
				frequency: 2,
				unit:      Year,
				timezone:  time.Local,
			},
			want: &Interval{
				frequency: 2,
				unit:      Year,
				nextRun:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				timezone:  time.Local,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := NewIntervalBased(tt.fields.frequency, tt.fields.unit)
			if got := i.Next(); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got.nextRun.String(), tt.want.nextRun.String())
				t.Errorf("Interval.Next() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
