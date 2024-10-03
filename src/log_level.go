package log

type LogLevel struct {
	name string
}

var registeredLevels []*LogLevel

func CreateLogLevel(value string) LogLevel {
	newLevel := LogLevel{value}
	registeredLevels = append(registeredLevels, &newLevel)
	return newLevel
}

var (
	Log      = CreateLogLevel("Log")
	Trace    = CreateLogLevel("Trace")
	Info     = CreateLogLevel("Info")
	Message  = CreateLogLevel("Message")
	Warning  = CreateLogLevel("Warning")
	Error    = CreateLogLevel("Error")
	Critical = CreateLogLevel("Critical")
)
