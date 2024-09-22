package log

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