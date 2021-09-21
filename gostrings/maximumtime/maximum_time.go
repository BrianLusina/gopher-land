package maximumtime

import "strings"

// MaximumTime returns the maximum time from a given time string of the format hh:mm
func MaximumTime(time string) string {
	hoursMinutes := strings.Split(time, "")

	if (hoursMinutes[1] == "?" || hoursMinutes[1] <= "3") && hoursMinutes[0] == "?" {
		hoursMinutes[0] = "2"
	} else if hoursMinutes[0] == "?" {
		hoursMinutes[0] = "1"
	}

	if hoursMinutes[0] == "2" && hoursMinutes[1] == "?" {
		hoursMinutes[1] = "3"
	} else if hoursMinutes[1] == "?" {
		hoursMinutes[1] = "9"
	}

	if hoursMinutes[3] == "?" {
		hoursMinutes[3] = "5"
	}

	if hoursMinutes[4] == "?" {
		hoursMinutes[4] = "9"
	}

	return strings.Join(hoursMinutes, "")
}
