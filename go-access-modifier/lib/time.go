package lib

const secondsInMinute int = 60
const minuteInHour int = 60
const HourInDay int = 24

func daysToHour(day int) int {
	return day * HourInDay
}

func DaysToMinutes(day int) int {
	return day * HourInDay * minuteInHour
}