package today

import (
	"fmt"
	"testing"
	"time"
)

func TestIsInternationalDay(t *testing.T) {
	for n, params := range []struct {
		offset time.Duration
		expect bool
	}{
		// >= -14 to < +36 should yield true
		{-15 * time.Hour, false},
		{-14*time.Hour - time.Second, false},
		{-14*time.Hour - time.Second - 500*time.Millisecond, false}, // check rounding
		{-14 * time.Hour, true},
		{-13 * time.Hour, true},
		{-time.Hour, true},
		{0, true},
		{35*time.Hour + 59*time.Minute + 59*time.Second, true},
		{35*time.Hour + 59*time.Minute + 59*time.Second + 500*time.Millisecond, false}, // check rounding
		{36 * time.Hour, false},
	} {
		t.Run(fmt.Sprintf("num=%d", n), func(t *testing.T) {
			now := time.Date(1968, time.April, 9, 0, 0, 0, 0, time.UTC)
			if r := IsInternationalDay(9, time.April, now.Add(params.offset)); r != params.expect {
				t.Fatalf("for %s expected %t but got %t", params.offset, params.expect, r)
			}
		})
	}
}

func TestIsDayMonth(t *testing.T) {
	for n, params := range []struct {
		when   time.Time
		day    int
		month  time.Month
		expect bool
	}{
		{
			when:   time.Now(),
			day:    time.Now().Day(),
			month:  time.Now().Month(),
			expect: true,
		},
		{
			// wrong month
			when:   time.Date(1974, time.January, 13, 0, 0, 0, 0, time.UTC),
			day:    13,
			month:  time.April,
			expect: false,
		},
		{
			// wrong day
			when:   time.Date(1974, time.April, 11, 0, 0, 0, 0, time.UTC),
			day:    13,
			month:  time.April,
			expect: false,
		},
		{
			// wrong month, wrong day
			when:   time.Date(1974, time.January, 11, 0, 0, 0, 0, time.UTC),
			day:    13,
			month:  time.April,
			expect: false,
		},

		{
			when:   time.Date(1974, time.April, 13, 0, 0, 0, 0, time.UTC),
			day:    13,
			month:  time.April,
			expect: true,
		},
	} {
		t.Run(fmt.Sprintf("%d", n), func(t *testing.T) {
			if r := IsDayMonth(params.day, params.month, params.when); r != params.expect {
				t.Fatalf("for #%d expected %t but got %t", n, params.expect, r)
			}
		})
	}
}
