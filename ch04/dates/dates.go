package dates

// DaysInWeek ...
const DaysInWeek int = 7

// WeeksToDays ...
func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

// DaysToWeeks ...
func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
