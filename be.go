package today

import "time"

// BeCssNakedDay indicates if today is CSS Naked Day
//
// https://css-naked-day.github.io/
func BeCssNakedDay() bool {
	return IsInternationalDay(9, time.April, time.Now())
}

// BeJsNakedDay indicates if today is JS Naked Day
//
// https://js-naked-day.org/
func BeJsNakedDay() bool {
	return IsInternationalDay(24, time.April, time.Now())
}

// BeTowelDay indicates if today is Towel Day.
func BeTowelDay() bool {
	return IsInternationalDay(25, time.May, time.Now())
}
