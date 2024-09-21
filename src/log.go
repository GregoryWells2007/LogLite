package log
import(
	"fmt"
	"time"
)

const (
	Message = "Message"
	Warning = "Warning"
	Error   = "Error"
)

const (
	List = 0
	Console = 1
)

var output_lists []*[]string;
var output_console bool = false;
func AddOutput(output_type int, a ...interface{}) {
	if (output_type == List) {
		output_lists = append(output_lists, a[0].(*[]string));
	} else if (output_type == Console) {
		output_console = true;
	}
}

/*
%t = time
%l = level
%m = message
*/
var output_format string = "[%t] %l: %m";

func SetFormat(format string) {
	output_format = format;
}

func Write(level string, message string) {
	var output string;

	for i := 0; i < len(output_format); i++ {
		if (string(output_format[i]) == "%") {
			i++;
			format := output_format[i];
			if (format == 't') {
				output += time.Now().Format("15:04:05");
			} else if (format == 'l') {
				output += level;
			} else if (format == 'm') {
				output += message;
			} else {
				// Write(Error, "invalid log code\n");
				return;
			}
		} else {
			output += string(output_format[i]);
		}
	}	
	
	for i := 0; i < len(output_lists); i++ {
		*output_lists[i] = append(*output_lists[i], output); 
	}
	
	if (output_console) {
		fmt.Printf("%s\n", output);
	}
}