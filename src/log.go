package log

import (
	"fmt"
	"time"
)

const (
	Message = "Message"
	Warning = "Warning"
	Error   = "Error"
)

var contains_console_out bool = false;
func AddOutput(output_type int, a ...interface{}) any {
	if (contains_console_out && output_type == Console) {
		Write(Warning, "Attempting to add multiple console outputs [LogLite]");
		return nil;
	}

	if output_type == List {
		listOutput := &ListOutput{};
		OutputTargets = append(OutputTargets, listOutput);
	} else if output_type == File {
		for i := 0; i < len(OutputTargets); i++ {
			if (OutputTargets[i].GetOutputType() == File) {
				
				if (OutputTargets[i].(*FileOutput).OutputFileName == a[0].(string)) {
					Write(Warning, "Attempting to add a file as an output target again, returning old target [LogLite]");
					return OutputTargets[i];
				}
			}
		}

		fileOutput := &FileOutput{};
		OutputTargets = append(OutputTargets, fileOutput);
	} else if output_type == Console {
		consoleOutput := &ConsoleOutput{};
		OutputTargets = append(OutputTargets, consoleOutput);
		contains_console_out = true;
	} else {
		Write(Error, "Unknown Output target [LogLite]");
		return nil;
	}

	OutputTargets[len(OutputTargets) - 1].Init(a...);
	return &OutputTargets[len(OutputTargets) - 1];
}

/*
%t = time
%l = level
%m = message
*/
var output_format string = "[%t] %l: %m"

func SetFormat(format string) {
	output_format = format
}

func Write(level string, message string) {
	var output string

	for i := 0; i < len(output_format); i++ {
		if string(output_format[i]) == "%" {
			i++
			format := output_format[i]
			if format == 't' {
				output += time.Now().Format("15:04:05")
			} else if format == 'l' {
				output += level
			} else if format == 'm' {
				output += message
			} else {
				// Write(Error, "invalid log code\n");
				return
			}
		} else {
			output += string(output_format[i])
		}
	}
	output += "\n";

	for i := 0; i < len(OutputTargets); i++ {
		OutputTargets[i].Output(output);
	}
}

func WriteFormatted(level string, format string, a ...interface{}) {
	var formatted_string string = fmt.Sprintf(format, a...);
	Write(level, formatted_string);	
}