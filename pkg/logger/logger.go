package logger

import "github.com/fatih/color"

func Error(s string) {
	color.Red(s)
}
