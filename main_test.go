package gonsola_test

import (
	"gonsola"
	"testing"
	"time"
)

func TestBasicFunctions(t *testing.T) {
	g := gonsola.Gonsola{}

	g.Info("Using Gonsola v0.0.1")
	g.Start("Building project", 0)
	time.Sleep(time.Second * 10)

	g.Warn("A new version of Gonsola 0.0.2")
}
