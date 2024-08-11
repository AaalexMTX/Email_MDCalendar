package message

// NewCalendar 无参构建日历结构
func NewCalendar() *Calendar {
	return &Calendar{Date: *NewDate()}
}

// NewCalendarWithParams 含参参构建日历结构
func NewCalendarWithParams(year, month int) *Calendar {
	return &Calendar{Date: *NewDateWithParams(year, month)}
}
