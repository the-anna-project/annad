package common

var (
	// LevelColors represent the log levels and their related colors the logger
	// implementations should use.
	LevelColors = map[string]string{
		"D": "cyan",
		"E": "red",
		"F": "magenta",
		"I": "white",
		"W": "yellow",
	}
)
