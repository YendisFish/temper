package temper

import (
	"fmt"
	"time"
)

func Error(msg string, info ...interface{}) {
	tm := ""
	inf := ""

	for _, dat := range info {
		switch d := dat.(type) {
		case LogTime:
			tm = " " + Colorize(d.Color, time.Now().Format(d.Format))
		case []string:
			if len(d) == 2 {
				inf = inf + "    " + Bold(Colorize(Red, d[0]+": ")) + d[1] + "\n"
			}
		}
	}

	error := Bold(Colorize(Red, "Error")) + tm + Bold(Colorize(Red, ": ")) + msg + "\n" + inf
	fmt.Print(error)
}

func Info(msg string, info ...interface{}) {
	tm := ""
	inf := ""

	for _, dat := range info {
		switch d := dat.(type) {
		case LogTime:
			tm = " " + Colorize(d.Color, time.Now().Format(d.Format))
		case []string:
			if len(d) == 2 {
				inf = inf + "    " + Bold(Colorize(Green, d[0]+": ")) + d[1] + "\n"
			}
		}
	}

	error := Bold(Colorize(Green, "Info")) + tm + Bold(Colorize(Green, ": ")) + msg + "\n" + inf
	fmt.Print(error)
}

func Warn(msg string, info ...interface{}) {
	tm := ""
	inf := ""

	for _, dat := range info {
		switch d := dat.(type) {
		case LogTime:
			tm = " " + Colorize(d.Color, time.Now().Format(d.Format))
		case []string:
			if len(d) == 2 {
				inf = inf + "    " + Bold(Colorize(Yellow, d[0]+": ")) + d[1] + "\n"
			}
		}
	}

	error := Bold(Colorize(Yellow, "Warning")) + tm + Bold(Colorize(Yellow, ": ")) + msg + "\n" + inf
	fmt.Print(error)
}

func Custom(msg []string, color int, info ...interface{}) {
	tm := ""
	inf := ""

	for _, dat := range info {
		switch d := dat.(type) {
		case LogTime:
			tm = " " + Colorize(d.Color, time.Now().Format(d.Format))
		case []string:
			if len(d) == 2 {
				inf = inf + "    " + Bold(Colorize(color, d[0]+": ")) + d[1] + "\n"
			}
		}
	}

	if len(msg) != 2 {
		Error("Not enough parameters in msg!", []string{"Info", "Length must be 2"})
	}

	error := Bold(Colorize(color, msg[0])) + tm + Bold(Colorize(color, ": ")) + msg[1] + "\n" + inf
	fmt.Print(error)
}
