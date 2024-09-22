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

func AddOutput(output_type int, a ...interface{}) any {
	if output_type == List {
		listOutput := &ListOutput{};
		OutputTargets = append(OutputTargets, listOutput);
	} else if output_type == File {
		fileOutput := &FileOutput{};
		OutputTargets = append(OutputTargets, fileOutput);
	} else if output_type == Console {
		consoleOutput := &ConsoleOutput{};
		OutputTargets = append(OutputTargets, consoleOutput);
	} else {
		Write(Error, "Unknown Output target");
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