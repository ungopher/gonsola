package gonsola

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/gookit/color"
)

type Gonsola struct {
}

func (g *Gonsola) New() Gonsola {
	return Gonsola{}
}

func (g *Gonsola) Info(msg string) {
	fmt.Println(color.Green.Render("ℹ"), color.White.Render(msg))
}

func (g *Gonsola) Start(msg string) {
	fmt.Println(color.Magenta.Render("◐"), color.White.Render(msg))
}

func (g *Gonsola) Warn(msg string) {
	fmt.Printf("\n%s %s\n", color.BgYellow.Render(" WARN "), color.White.Render(msg))

}

func (g *Gonsola) Error(msg string) {
	fmt.Printf("\n\n%s %s\n", color.BgRed.Render(" ERROR "), color.White.Render(msg))

	stack := strings.Split(string(debug.Stack()), "\n")

	for i := 2; i < len(stack)/2; i++ {
		trimmed := strings.Trim(string(stack[i*2]), "\t ")
		last_semi := strings.LastIndex(trimmed, ":")
		p := trimmed[0:last_semi]
		line := strings.Split(trimmed[last_semi+1:], " ")[0]

		fmt.Printf("\n    %s %s (%s:%s)", color.Gray.Render("at"), p, color.Green.Render(string(stack[i*2+1])), line)

	}

	print("\n\n")
}
