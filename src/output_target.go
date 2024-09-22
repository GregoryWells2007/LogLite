package log

import (
	"fmt"
	"os"
)

const (
	List    = 0
	File    = 1
	Console = 2
)

// base output target class
type OutputTarget interface {
	Init(arguments ...any)
	Output(message string);
	Close();
}

// console output
type ConsoleOutput struct {
	OutputTarget
}
func (consoleOutput *ConsoleOutput) Init(arguments ...any) {}
func (consoleOutput *ConsoleOutput) Output(message string) { fmt.Print(message); }
func (consoleOutput *ConsoleOutput) Close() { }

// file output
type FileOutput struct {
	OutputTarget
	OutputFile *os.File;
	Name string;
}
func (fileOutput *FileOutput) Init(arguments ...any) {
	var err error;
	fileOutput.OutputFile, err = os.Create(arguments[0].(string))
	if err != nil {
		panic(err);
	}
}
func (fileOutput *FileOutput) Output(message string) {
	_, err := fileOutput.OutputFile.WriteString(message);
	if err != nil {
		panic(err);
	}	
}
func (fileOutput *FileOutput) Close() { fileOutput.OutputFile.Close(); }

// list output
type ListOutput struct {
	OutputTarget
	OutputList *[]string;
}
func (listOutput *ListOutput) Init(arguments ...any) { listOutput.OutputList = arguments[0].(*[]string); }
func (listOutput *ListOutput) Output(message string) { *listOutput.OutputList = append(*listOutput.OutputList, message) }
func (listOutput *ListOutput) Close() { }

var OutputTargets []OutputTarget; 