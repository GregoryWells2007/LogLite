package log

import (
	"fmt"
	"time"
	"os"
)

const (
	Message = "Message"
	Warning = "Warning"
	Error   = "Error"
)

const (
	List    = 0
	File    = 1
	Console = 2
)

var output_lists []*[]string
var output_files []*os.File;
var output_console bool = false

func AddOutput(output_type int, a ...interface{}) {
	if output_type == List {
		output_lists = append(output_lists, a[0].(*[]string))
	} else if output_type == File {
		file, err := os.Create(a[0].(string))
		if err != nil {
			panic(err);
		}
		output_files = append(output_files, file)
	} else if output_type == Console {
		output_console = true
	}
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

	for i := 0; i < len(output_lists); i++ {
		*output_lists[i] = append(*output_lists[i], output)
	}
	for i := 0; i < len(output_files); i++ {
        file := output_files[i];

		_, err := file.WriteString(output);
		if err != nil {
			panic(err);
		}
	}
	if output_console {
		fmt.Printf("%s", output)
	}
}
