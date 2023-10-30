package gonsola_test

import (
	"testing"

	"github.com/ungopher/gonsola"
)

func TestBasicFunctions(t *testing.T) {
	g := gonsola.CreateConsola()

	g.Info("Using Gonsola v0.0.1")
	g.Start("Building project")
	g.Warn("A new version of Gonsola 0.0.2")
	g.Success("Project built")
	g.Error("Nothing is wrong. All fine!")

	g.Box("This is a cramped box", gonsola.BoxOptions{})

	g.Box("probably a better box.\n\n\np.s. wow i can this", gonsola.BoxOptions{Fill: true, Padding: 8, Center: true, Style: "double"})
	g.Box("another one of your mom\n\nwhat? with no center", gonsola.BoxOptions{Padding: 8, Style: "rounded"})
}
