package log

import(
	"fmt"
	"time"
)

type Insert struct {
	ID rune;
	Value string;
}

func FormatMessage(format string, inserts []Insert) (string) {
	var output string;
	for i := 0; i < len(output_format); i++ {
		if (output_format[i] == '%') {
			i++;
			var found bool = false;
			for p := 0; p < len(inserts); p++ {
				if (inserts[p].ID == rune(output_format[i])) {
					output += inserts[p].Value;
					found = true;
					break;
				}
			}

			if (!found) {
				output += "%";
				output += string(output_format[i]);
			}
		} else {
			output += string(output_format[i])
		}
	}
	return output;
}

func GetDate() string {
	now := time.Now()
	y, m, d := now.Date()
	y = y % 100;
	month := "00";
	switch m.String() {
		case "January":   month = "01";
		case "February":  month = "02";
		case "March":     month = "03";
		case "April":     month = "04";
		case "May":       month = "05";
		case "June":      month = "06";
		case "July":      month = "07";
		case "August":    month = "08";
		case "September": month = "09";
		case "October":   month = "10";
		case "November":  month = "11";
		case "December":  month = "12";
	}

	return fmt.Sprintf("%s/%d/%d", month, d, y);
}