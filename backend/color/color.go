package color

import (
	"fmt"
)

const (
	FgBlack = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgPurple
	FgCyan
	FgWhite
)

const (
	BgBlack = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgPurple
	BgCyan
	BgWhite
)

const (
	FgHiBlack = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

func Color(fg int, bg int, str string) string {
	return fmt.Sprintf("\033[0;%d;%dm%s\033[0m", fg, bg, str)
}
