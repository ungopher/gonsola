package gonsola_test

import (
	"gonsola"
	"testing"
)

func TestBasicFunctions(t *testing.T) {
	g := gonsola.CreateConsola()

	g.Info("Using Gonsola v0.0.1")
	g.Start("Building project")
	g.Warn("A new version of Gonsola 0.0.2")
	g.Success("Project built")
	g.Error("Nothing is wrong. All fine!")

	g.Box("This is cramped box", gonsola.BoxOptions{})

	g.Box("probably a better box.\n\n\np.s. wow i can this", gonsola.BoxOptions{Fill: true, Padding: 8, Style: "double"})
}
