package config

type ExtractionStrategy string

const (
	SceneExtractionStrategy ExtractionStrategy = "scene"
	TimerExtractionStrategy ExtractionStrategy = "timer"
)

type RotationStrategy string

const (
	SequentialRotationStrategy RotationStrategy = "sequential"
	ShuffleRotationStrategy    RotationStrategy = "shuffle"
)

type LogLevel string

const (
	DebugLogLevel LogLevel = "debug"
	InfoLogLevel  LogLevel = "info"
	WarnLogLevel  LogLevel = "warn"
	ErrorLogLevel LogLevel = "error"
)
