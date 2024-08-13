package message

import "testing"

func TestCalendar_PrintCalendar(t *testing.T) {
	c := NewCalendarWithParams(2024, 8)
	c.PrintCalendar()
}
