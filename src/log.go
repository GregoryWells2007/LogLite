package log

import (
	"fmt"
	"time"
)

var contains_console_out bool = false
var console_index int = 0

var current_target_id int = 0

func AddOutputTarget(output_type int, a ...interface{}) *OutputTarget {
	if contains_console_out && output_type == Console {
		Write(Warning, "Attempting to add multiple console outputs [LogLite]")
		return &OutputTargets[console_index]
	}

	var new_target OutputTarget = NewOutputTarget()

	if output_type == List {
		listOutput := &ListOutput{}
		new_target.OutputStream = listOutput
	} else if output_type == File {
		for i := 0; i < len(OutputTargets); i++ {
			if OutputTargets[i].OutputStream.GetOutputType() == File {
				if OutputTargets[i].OutputStream.(*FileOutput).OutputFileName == a[0].(string) {
					Write(Warning, "Attempting to add a file as an output target again, returning old target [LogLite]")
					return &OutputTargets[i]
				}
			}
		}

		fileOutput := &FileOutput{}
		new_target.OutputStream = fileOutput
	} else if output_type == Console {
		consoleOutput := &ConsoleOutput{}
		new_target.OutputStream = consoleOutput
		contains_console_out = true
		console_index = len(OutputTargets)
	} else {
		Write(Error, "Unknown Output target [LogLite]")
		return nil
	}

	OutputTargets = append(OutputTargets, new_target)

	OutputTargets[len(OutputTargets)-1].OutputStream.Init(a...)
	current_target_id++
	return &OutputTargets[len(OutputTargets)-1]
}
func RemoveOutputTarget(target *OutputTarget) {
	for i := 0; i < len(OutputTargets); i++ {
		if OutputTargets[i].GetTargetID() == target.GetTargetID() {
			OutputTargets = append(OutputTargets[:i], OutputTargets[i+1:]...)
			return
		}
	}
	WriteFormatted(Warning, "cannot removed output target as it is not a current output target [LogLite]")
}

/*
%d = date
%t = time
%l = level
%m = message
*/

func Write(level LogLevel, message string) {
	if len(OutputTargets) == 0 {
		AddOutputTarget(Console)
		Write(Warning, "No output target specified adding console [LogLite]")
	}

	for i := 0; i < len(OutputTargets); i++ {
		var output string = FormatMessage(OutputTargets[i].ouputPattern, []Insert{
			{"date", GetDate()},
			{"time", time.Now().Format("15:04:05")},
			{"level", level.name},
			{"message", message}})
		output += "\n"

		if OutputTargets[i].OutputFilter.filter[level.name] {
			OutputTargets[i].OutputStream.Output(output)
		}
	}
}

func WriteFormatted(level LogLevel, format string, a ...interface{}) {
	var formatted_string string = fmt.Sprintf(format, a...)
	Write(level, formatted_string)
}
