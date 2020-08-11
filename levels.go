package logoru

import "github.com/fatih/color"

var levels = []string{"DEBUG", "INFO", "SUCCESS", "WARNING", "ERROR", "CRITICAL"}

// Output a debugging message
func Debug(msg interface{}) error {
	err := output(levels[0], color.New(color.FgBlue, color.Bold), msg)
	return err
}

// Output an info message
func Info(msg interface{}) error {
	err := output(levels[1], color.New(color.FgWhite, color.Bold), msg)
	return err
}

// Output a success message
func Success(msg interface{}) error {
	err := output(levels[2], color.New(color.FgGreen, color.Bold), msg)
	return err
}

// Output a warning message
func Warning(msg interface{}) error {
	err := output(levels[3], color.New(color.FgYellow, color.Bold), msg)
	return err
}

// Output an error message
func Error(msg interface{}) error {
	err := output(levels[4], color.New(color.FgRed, color.Bold), msg)
	return err
}

// Output a critical message
func Critical(msg interface{}) error {
	err := output(levels[5], color.New(color.FgWhite, color.BgRed, color.Bold), msg)
	return err
}
