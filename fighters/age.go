package fighters

import "time"

func AgeOn(dob, date time.Time) int {
	dob = dob.UTC()
	date = date.UTC()
	age := date.Year() - dob.Year()

	if date.Month() < dob.Month() ||
		(date.Month() == dob.Month() && date.Day() < dob.Day()) {
		age--
	}

	return age

}
