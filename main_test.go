package gonsola_test

import (
	"gonsola"
	"testing"
)

func TestBasicFunctions(t *testing.T) {
	g := gonsola.Gonsola{}

	g.Info("Using Gonsola v0.0.1")
	g.Start("Building project")
	g.Warn("A new version of Gonsola 0.0.2")
	g.Error("Nothing is wrong. All fine!")
}
