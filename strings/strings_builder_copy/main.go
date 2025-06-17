package main

import "strings"

// В этом варианте new и old будут шарить один и тот же участок памяти
func firstWay() {
	old := strings.Builder{}
	// manipulating with old..
	new := old
	_ = new
}

// Более правильный способ, при котором создается новый
// strings.Builder, в который затем копируется старое
// значение строки.
func secondWay() {
	old := strings.Builder{}
	// manipulating with old..
	new := strings.Builder{}
	new.WriteString(old.String())
	_ = new
}
