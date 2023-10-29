package gonsola

import (
	"runtime"
	"syscall"
)

func IsUnicodeSupported() bool {
	term, f_term := syscall.Getenv("TERM")
	_, f_ci := syscall.Getenv("CI")
	_, f_wt := syscall.Getenv("WT_SESSION")
	_, f_terminus := syscall.Getenv("TERMINUS_SUBLIME")
	cet, f_cet := syscall.Getenv("ConEmuTask")
	term_pro, f_term_pro := syscall.Getenv("TERM_PROGRAM")
	terminal_emu, f_terminal_emu := syscall.Getenv("TERMINAL_EMULATOR")

	if runtime.GOOS == "windows" {
		if !f_term {
			return true
		}

		return term != "LINUX"
	}

	return f_ci || f_wt || f_terminus ||
		(f_cet && cet == "{cmd::Cmder}") ||
		(f_term_pro && (term_pro == "Terminus-Sublime" || term_pro == "vscode")) ||
		(f_term && (term == "xterm-256color" || term == "alacritty")) ||
		(f_terminal_emu && terminal_emu == "JetBrains-JediTerm")
}

func UnicodeS(c string, fallback string) string {
	if IsUnicodeSupported() {
		return c
	}

	return fallback
}