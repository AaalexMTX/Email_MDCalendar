package message

import (
	"fmt"
	"math"
)

// NewCalendar 无参构建日历结构
func NewCalendar() *Calendar {
	return &Calendar{Date: *NewDate()}
}

// NewCalendarWithParams 含参参构建日历结构
func NewCalendarWithParams(year, month int) *Calendar {
	return &Calendar{Date: *NewDateWithParams(year, month)}
}

// PrintCalendar 格式化打印日历
func (c Calendar) PrintCalendar() {
	day := 1                                               // 当前日期
	monthDays := c.Date.dayOfMonth()                       // 获取当月天数
	week := c.Date.WeekOfTheFirstDay()                     // 获取当月第一天是星期几
	rowNow := math.Ceil((float64)(week-1+monthDays) / 7.0) // 计算需要打印的行数
	fmt.Print("总日期 -- 起始星期 -- 行数\n")
	fmt.Printf("%d -- %d -- %d\n", monthDays, week, int(rowNow))
	fmt.Print("| 日 | 一 | 二 | 三 | 四 | 五 | 六 |\n")
	for i := range 7 {
		if i == 6 {
			fmt.Print("|:--:|\n")
		} else {
			fmt.Print("|:--:")
		}
	}
	for i := 0; i < int(rowNow); i++ {
		fmt.Print("|")
		for j := 0; j < 7; j++ {
			if (i == 0 && j+1 > week || i > 0) && day <= monthDays {
				fmt.Printf("%2d<br><br>|", day)
				day++
			} else {
				fmt.Print("    |")
			}
		}
		if i != 6 {
			fmt.Printf("\n")
		}
	}
	for i := 0; i < 7; i++ {
		fmt.Print("| ")
		for j := 0; j < 18; j++ {
			fmt.Print("-")
		}
		if i == 6 {
			fmt.Print(" |\n")
		}
	}
}
