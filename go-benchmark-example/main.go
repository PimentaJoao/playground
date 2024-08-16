package main

var dayMap map[int]string = map[int]string{
	1: "Sunday",
	2: "Monday",
	3: "Tuesday",
	4: "Wednesday",
	5: "Thursday",
	6: "Friday",
	7: "Saturday",
}

var daySlice []string = []string{
	"",
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func daySelectorIf(day int) string {
	if day == 1 {
		return "Sunday"
	} else if day == 2 {
		return "Monday"
	} else if day == 3 {
		return "Tuesday"
	} else if day == 4 {
		return "Wednesday"
	} else if day == 5 {
		return "Thursday"
	} else if day == 6 {
		return "Friday"
	} else if day == 7 {
		return "Saturday"
	} else {
		return "None"
	}
}

func daySelectorSwitch(day int) string {
	switch day {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "None"
	}
}

func daySelectorMap(day int) string {
	return dayMap[day]
}

func daySelectorSlice(day int) string {
	return daySlice[day]
}

func main() {}
