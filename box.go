package gonsola

import (
	"fmt"
	"strings"
)

type boxStyle struct {
	tl string
	tr string
	bl string
	br string
	h  string
	v  string
}

var boxStyles = map[string]boxStyle{
	"solid": {
		tl: "┌",
		tr: "┐",
		bl: "└",
		br: "┘",
		h:  "─",
		v:  "│",
	},
	"double": {
		tl: "╔",
		tr: "╗",
		bl: "╚",
		br: "╝",
		h:  "═",
		v:  "║",
	},
	"doubleSingle": {
		tl: "╓",
		tr: "╖",
		bl: "╙",
		br: "╜",
		h:  "─",
		v:  "║",
	},
	"doubleSingleRounded": {
		tl: "╭",
		tr: "╮",
		bl: "╰",
		br: "╯",
		h:  "─",
		v:  "║",
	},
	"singleThick": {
		tl: "┏",
		tr: "┓",
		bl: "┗",
		br: "┛",
		h:  "━",
		v:  "┃",
	},
	"singleDouble": {
		tl: "╒",
		tr: "╕",
		bl: "╘",
		br: "╛",
		h:  "═",
		v:  "│",
	},
	"singleDoubleRounded": {
		tl: "╭",
		tr: "╮",
		bl: "╰",
		br: "╯",
		h:  "═",
		v:  "│",
	},
	"rounded": {
		tl: "╭",
		tr: "╮",
		bl: "╰",
		br: "╯",
		h:  "─",
		v:  "│",
	},
}

type boxing struct {
	options    BoxOptions
	message    string
	term_width int
}

func (b *boxing) Render() {
	width := 0
	messages := []string{}
	style := boxStyles[b.options.Style]

	if style == (boxStyle{}) {
		style = boxStyles["solid"]
	}

	if b.options.Fill {
		width = b.term_width - 2
	}

	for _, v := range strings.Split(b.message, "\n") {
		if b.options.Center {
			v = strings.Trim(v, "\t ")
		}

		v = strings.ReplaceAll(v, "\t", "    ")

		if len(v) > width {
			width = len(v)
		}

		messages = append(messages, v)

	}

	b.RenderTop(width, style)

	for i := 0; i < b.options.Padding/4; i++ {
		b.RenderYPad(width, style)
	}

	for _, v := range messages {
		leftPad := width - len(v) + b.options.Padding
		rightPad := width - len(v) + b.options.Padding

		fmt.Print(style.v, strings.Repeat(" ", leftPad), v, strings.Repeat(" ", rightPad), style.v, "\n")
	}

	for i := 0; i < b.options.Padding/4; i++ {
		b.RenderYPad(width, style)
	}

	b.RenderBottom(width, style)

}

func (b *boxing) RenderTop(width int, style boxStyle) {
	fmt.Print(style.tl, strings.Repeat(style.h, width), style.tr, "\n")
}

func (b *boxing) RenderYPad(width int, style boxStyle) {
	fmt.Print(style.v, strings.Repeat(" ", width), style.v, "\n")
}

func (b *boxing) RenderBottom(width int, style boxStyle) {
	fmt.Print(style.bl, strings.Repeat(style.h, width), style.br, "\n\n")
}
