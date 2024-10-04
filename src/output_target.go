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
type IOutputStream interface {
	Init(arguments ...any)
	Output(message string)
	Close()

	GetOutputType() int
}

// console output
type ConsoleOutput struct {
	IOutputStream
}

func (consoleOutput *ConsoleOutput) GetOutputType() int    { return Console }
func (consoleOutput *ConsoleOutput) Init(arguments ...any) {}
func (consoleOutput *ConsoleOutput) Output(message string) { fmt.Print(message) }
func (consoleOutput *ConsoleOutput) Close()                {}

// file output
type FileOutput struct {
	IOutputStream
	OutputFile     *os.File
	OutputFileName string
}

func (fileOutput *FileOutput) GetOutputType() int { return File }
func (fileOutput *FileOutput) Init(arguments ...any) {
	var err error
	fileOutput.OutputFile, err = os.Create(arguments[0].(string))
	if err != nil {
		panic(err)
	}
	fileOutput.OutputFileName = arguments[0].(string)
}
func (fileOutput *FileOutput) Output(message string) {
	_, err := fileOutput.OutputFile.WriteString(message)
	if err != nil {
		panic(err)
	}
}
func (fileOutput *FileOutput) Close() { fileOutput.OutputFile.Close() }

// list output
type ListOutput struct {
	IOutputStream
	OutputList *[]string
}

func (listOutput *ListOutput) GetOutputType() int { return List }
func (listOutput *ListOutput) Init(arguments ...any) {
	listOutput.OutputList = arguments[0].(*[]string)
}
func (listOutput *ListOutput) Output(message string) {
	*listOutput.OutputList = append(*listOutput.OutputList, message)
}
func (listOutput *ListOutput) Close() {}

type OutputFilter struct {
	filter map[string]bool
}

func (target *OutputFilter) internal_InitFilter() {
	target.filter = make(map[string]bool)
	target.UnfilterAll()
}

func (target *OutputFilter) FilterAll() {
	for i := 0; i < len(registeredLevels); i++ {
		target.filter[registeredLevels[i].name] = false
	}
}

func (target *OutputFilter) UnfilterAll() {
	for i := 0; i < len(registeredLevels); i++ {
		target.filter[registeredLevels[i].name] = true
	}
}

func (target *OutputFilter) Filter(level LogLevel) {
	target.filter[level.name] = false
}

func (target *OutputFilter) Unfilter(level LogLevel) {
	target.filter[level.name] = true
}

type OutputTarget struct {
	OutputStream IOutputStream
	OutputFilter OutputFilter

	ouputPattern string
	targetID     int
}

var currentTargetID int = 0

func NewOutputTarget() OutputTarget {
	newTarget := OutputTarget{nil, OutputFilter{}, "%{message}", currentTargetID}
	newTarget.OutputFilter.internal_InitFilter()
	currentTargetID++
	return newTarget
}
func (target *OutputTarget) GetTargetID() int {
	return target.targetID
}
func (target *OutputTarget) SetOutputPattern(pattern string) { target.ouputPattern = pattern }
func (target *OutputTarget) GetOutputPattern() string        { return target.ouputPattern }

var OutputTargets []OutputTarget
