package log

import (
	"fmt"
	"time"
)

type Insert struct {
	Text  string
	Value string
}

func FormatMessage(format string, inserts []Insert) string {
	var output string
	for i := 0; i < len(format); i++ {
		if format[i] == '%' {
			i2 := i
			var insert_text string
			for i2 < len(format) && format[i2] != '}' {
				if format[i2] == ' ' {
					break
				}

				insert_text += string(format[i2])
				i2++
			}
			insert_text = insert_text[2:]

			if i2 < len(format) && format[i2] == '}' {
				i = i2
				exists := false
				for v := 0; v < len(inserts); v++ {
					if inserts[v].Text == insert_text {
						output += inserts[v].Value
						exists = true
						break
					}
				}
				if !exists {
					output += "%{" + insert_text + "}"
				}
			} else {
				output += "%"
			}
		} else {
			output += string(format[i])
		}
	}
	return output
}

func GetDate() string {
	now := time.Now()
	y, m, d := now.Date()
	y = y % 100
	month := "00"
	switch m.String() {
	case "January":
		month = "01"
	case "February":
		month = "02"
	case "March":
		month = "03"
	case "April":
		month = "04"
	case "May":
		month = "05"
	case "June":
		month = "06"
	case "July":
		month = "07"
	case "August":
		month = "08"
	case "September":
		month = "09"
	case "October":
		month = "10"
	case "November":
		month = "11"
	case "December":
		month = "12"
	}

	return fmt.Sprintf("%s/%d/%d", month, d, y)
}
