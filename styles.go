package task

import (
	. "github.com/charmbracelet/lipgloss"
)

// Colors.
var (
	fuschia = AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
	green   = Color("#04B575")

	yellowGreen = AdaptiveColor{Light: "#04B575", Dark: "#ECFD65"}
	darkGray    = AdaptiveColor{Light: "#DDDADA", Dark: "#3C3C3C"}
	cream       = AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
	red         = AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
)

var (
	errorTitleStyle = NewStyle().
			Foreground(cream).
			Background(red).
			Padding(0, 1)

	subtleStyle = NewStyle().
			Foreground(AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"})
)
