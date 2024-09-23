package log

import (
	"fmt"
	"time"
)

const (
	Log      = "Log"
	Trage    = "Trace"
	Info     = "Info"
	Message  = "Message"
	Warning  = "Warning"
	Error    = "Error"
	Critical = "Critical"
)

var contains_console_out bool = false;
var console_index int = 0;
func AddOutput(output_type int, a ...interface{}) *OutputTarget {
	if (contains_console_out && output_type == Console) {
		Write(Warning, "Attempting to add multiple console outputs [LogLite]");
		return &OutputTargets[console_index];
	}

	var new_target OutputTarget = NewOutputTarget();

	if output_type == List {
		listOutput := &ListOutput{};
		new_target.OutputStream = listOutput;
	} else if output_type == File {
		for i := 0; i < len(OutputTargets); i++ {
			if (OutputTargets[i].OutputStream.GetOutputType() == File) {
				if (OutputTargets[i].OutputStream.(*FileOutput).OutputFileName == a[0].(string)) {
					Write(Warning, "Attempting to add a file as an output target again, returning old target [LogLite]");
					return &OutputTargets[i];
				}
			}
		}

		fileOutput := &FileOutput{};
		new_target.OutputStream = fileOutput;
	} else if output_type == Console {
		consoleOutput := &ConsoleOutput{};
		new_target.OutputStream = consoleOutput;
		contains_console_out = true;
		console_index = len(OutputTargets);
	} else {
		Write(Error, "Unknown Output target [LogLite]");
		return nil;
	}

	OutputTargets = append(OutputTargets, new_target);

	OutputTargets[len(OutputTargets) - 1].OutputStream.Init(a...);
	return &OutputTargets[len(OutputTargets) - 1];
}

/*
%d = date
%t = time
%l = level
%m = message
*/

func Write(level string, message string) {
	if (len(OutputTargets) == 0) {
		AddOutput(Console);
		Write(Warning, "No output target specified adding console [LogLite]");
	}

	for i := 0; i < len(OutputTargets); i++ {
		var output string = FormatMessage(OutputTargets[i].OuputPattern, []Insert{ Insert{'d', GetDate()}, Insert{'t', time.Now().Format("15:04:05")}, Insert{'l', level}, Insert{'m', message}  });
		output += "\n";
		OutputTargets[i].OutputStream.Output(output);
	}
}

func WriteFormatted(level string, format string, a ...interface{}) {
	var formatted_string string = fmt.Sprintf(format, a...);
	Write(level, formatted_string);	
}