package mycalendar

type MyCalendar struct {
	calendar [][2]int
}

func NewMyCalendar() MyCalendar {
	return MyCalendar{
		calendar: [][2]int{},
	}
}

func (cal *MyCalendar) Book(start int, end int) bool {
	for _, pair := range cal.calendar {
		if pair[0] >= end || pair[1] <= start {
			continue
		}

		return false
	}

	cal.calendar = append(cal.calendar, [2]int{start, end})
	return true
}
