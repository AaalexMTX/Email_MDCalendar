package checktime

import (
	"email_mdCalendar/app/message"
	"fmt"
	"testing"
)

func TestDateWeekOfTheFirstDay(t *testing.T) {
	today := message.NewDateWithParams(2024, 7)
	fmt.Println(today.WeekOfTheFirstDay())
}
