package checktime

// NewDateWithParams 含参创建结构
func NewDateWithParams(year, month int) *Date {
	return &Date{Year: year, Month: month}
}

// NewDate 无参创建结构
func NewDate() *Date {
	return &Date{}
}

func (d Date) isLeap() bool {
	return d.Year%400 == 0 || d.Year%4 == 0 && d.Year%100 != 0
}

// isValid 结构内容设置合理
func (d Date) isValid() bool {
	return d.Month >= 1 && d.Month <= 12 && d.Year > 1582
}

// dayOfMonth 返回该月的天数
func (d Date) dayOfMonth() int {
	days := [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if d.Month == 2 {
		// 没有三目运算符的痛苦
		return days[d.Month] + map[bool]int{false: 0, true: 1}[d.isLeap()]
	}
	return days[d.Month]
}

// WeekOfTheFirstDay week 返回本月第一天的日期是星期几
func (d Date) WeekOfTheFirstDay() int {
	/*
		公式： w=y + [y/4] + [c/4] - 2c + [26(m + 1)/10] + d - 1

		W 是星期数。
		c 是世纪数减一，也就是年份前两位。
		y 是年份后两位。
		m 是月份。m 的取值范围是 3 至 14，因为某年的 1、2 月要看作上一年的 13、14月，
		比如 2019 年的 1 月 1 日要看作 2018 年的 13 月 1 日来计算。
		d 是日数。
		[] 是取整运算。
		mod 是求余运算。
	*/
	y := d.Year
	m := d.Month
	day := 1
	if d.Month <= 2 {
		y -= 1
		m += 12
	}
	y = y % 100
	c := y / 100
	return (y + y/4 + c/4 - 2*c + 13*(m+1)/5 + day - 1) % 7
}
