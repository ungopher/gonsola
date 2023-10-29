package gonsola

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gookit/color"
	"golang.org/x/term"
)

type gonsola struct {
	unicode_supported bool
	history           []history
}

type BoxOptions struct {
	Center  bool
	Fill    bool
	Padding int
	Style   string
	style   boxStyle
}

func CreateConsola() gonsola {
	return gonsola{
		unicode_supported: IsUnicodeSupported(),
	}
}

func (g *gonsola) Info(msg string) {
	fmt.Println(color.Green.Render(UnicodeS("ℹ", "i")), color.White.Render(msg))

	g.history = append(g.history, createHistory("info", 1))
}

func (g *gonsola) Start(msg string) {
	fmt.Println(color.Magenta.Render(UnicodeS("◐", "o")), color.White.Render(msg))

	g.history = append(g.history, createHistory("start", 1))
}

func (g *gonsola) Success(msg string) {
	fmt.Println(color.Green.Render(UnicodeS("✔", "√")), color.White.Render(msg))

	g.history = append(g.history, createHistory("success", 1))
}

func (g *gonsola) Warn(msg string) {
	fmt.Printf("\n%s %s\n\n", color.BgYellow.Render(" WARN "), color.White.Render(msg))

	g.history = append(g.history, createHistory("warn", 3))
}

func (g *gonsola) Error(msg string) {
	fmt.Printf("\n%s %s\n", color.BgRed.Render(" ERROR "), color.White.Render(msg))

	stack := strings.Split(string(debug.Stack()), "\n")

	for i := 2; i < len(stack)/2; i++ {
		trimmed := strings.Trim(string(stack[i*2]), "\t ")
		last_semi := strings.LastIndex(trimmed, ":")
		p := trimmed[0:last_semi]
		line := strings.Split(trimmed[last_semi+1:], " ")[0]

		fmt.Printf("\n    %s %s (%s)", color.Gray.Render("at"), p, color.Green.Render(string(stack[i*2+1]), ":", line))

	}

	print("\n\n")

	g.history = append(g.history, createHistory("error", len(stack)/2+3))
}

func (g *gonsola) Box(msg string, options BoxOptions) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	b := boxing{
		options:    options,
		message:    msg,
		term_width: width,
	}

	b.Render()

	g.history = append(g.history, createHistory("box", 1))
}
