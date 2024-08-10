package checktime

import (
	"fmt"
	"testing"
)

func TestDateWeekOfTheFirstDay(t *testing.T) {
	today := NewDateWithParams(2024, 7)
	fmt.Println(today.WeekOfTheFirstDay())
}
