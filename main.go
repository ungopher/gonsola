package gonsola

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

// type FormatOptions struct {
//   columns int
//   date bool
//   colors bool
//   compact int
// }

// type InputLogObject struct {
//   level int
//   tag string
//   log_type string
//   message string
//   additional []string
//   args []any
//   date Date
// }

// type LogObject struct {
//   level int
//   log_type string
//   tag string
//   args []any
//   date Date
// }

// type ReporterContext struct {
// 	options GonsolaOptions
// }

// type GonsolaReporter func(logObj LogObject, ctx ReporterContext) void

// type MockFunc func(log_type string, defaults InputLogObject) void

// type GonsolaOptions struct {
//   reporters []ConsolaReporter
//   types string
//   level int
//   defaults InputLogObject
//   throttle int
//   throttleMin int
//   stdout NodeJS.WriteStream
//   stderr *NodeJS.WriteStream
//   mockFn *MockFunc
//   prompt typeof import("./prompt").prompt | undefined
//   formatOptions FormatOptions
// }

type GonsolaStart struct {
	stdout     *os.File
	stderr     *os.File
	null       *os.File
	characters []string
}

type Gonsola struct {
}

func (g *Gonsola) New() Gonsola {
	return Gonsola{}
}

func (g *Gonsola) Info(msg string) {
	fmt.Println(color.GreenString("ℹ"), color.WhiteString(msg))
}

func (g *Gonsola) Start(msg string, character_set int) GonsolaStart {
	null, _ := os.Open(os.DevNull)
	stdout := os.Stdout
	stderr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)

	go func() {
		icons := []string{"◐", "◓", "◑", "◒"}
		i := 0

		for {
			fmt.Fprint(stdout, "\033[K", color.MagentaString(icons[i]), " ", color.WhiteString(msg))

			time.Sleep(time.Second)

			if i < 3 {
				i++
			} else {
				i = 0
			}
		}
	}()

	return GonsolaStart{
		stdout:     stdout,
		stderr:     stderr,
		null:       null,
		characters: CHARACTER_SETS[character_set],
	}
}

func (gs *GonsolaStart) Msg(msg string) {
}

func (gs *GonsolaStart) Stop(msg string) {
	defer gs.null.Close()
	os.Stdout = gs.stdout
	os.Stderr = gs.stderr
	log.SetOutput(os.Stderr)
}

func (g *Gonsola) Warn(msg string) {
	warn := color.New(color.BgYellow, color.FgBlack).SprintFunc()(" WARN ")

	fmt.Println(warn, color.WhiteString(msg))
}
