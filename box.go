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
	max_length := 0
	messages := []string{}
	style := boxStyles[b.options.Style]

	if style == (boxStyle{}) {
		style = boxStyles["solid"]
	}

	if b.options.Fill {
		max_length = b.term_width - 2
	}

	for _, v := range strings.Split(b.message, "\n") {
		if b.options.Center {
			v = strings.Trim(v, "\t ")
		}

		v = strings.ReplaceAll(v, "\t", "    ")

		if len(v) > max_length {
			max_length = len(v)
		}

		messages = append(messages, v)

	}

	b.RenderTop(max_length, style)

	for i := 0; i < b.options.Padding/4; i++ {
		b.RenderYPad(max_length, style)
	}

	for _, v := range messages {
		length := len(v)
		left_pad := b.options.Padding
		right_pad := b.options.Padding

		if v == "" {
			fmt.Print(style.v, strings.Repeat(" ", max_length), style.v, "\n")

			continue
		}

		if b.options.Center && length < max_length {
			left_pad = (max_length - length) / 2
		}

		if !b.options.Center && b.options.Fill {
			right_pad = max_length - length - left_pad
		}

		fmt.Print(style.v)

		for i := 0; i < left_pad; i++ {
			fmt.Print(" ")
		}

		fmt.Print(stringsPadTo(v, " ", max_length))

		for i := 0; i < right_pad; i++ {
			fmt.Print(" ")
		}

		fmt.Print(style.v, "\n")
	}

	for i := 0; i < b.options.Padding/4; i++ {
		b.RenderYPad(max_length, style)
	}

	b.RenderBottom(max_length, style)

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
