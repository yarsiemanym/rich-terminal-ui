package widgets

import (
	"fmt"
	"strings"

	"github.com/yarsiemanym/rich-terminal-ui/vt100"
)

type ProgressBar struct {
	start           int
	end             int
	current         int
	percentComplete int
	width           int
}

func (bar *ProgressBar) New(start int, end int, width int) {
	bar.start = start
	bar.end = end
	bar.width = width
	bar.current = start
	bar.percentComplete = start * 100 / end
}

func (bar *ProgressBar) Advance(delta int) {
	bar.current += delta
	bar.percentComplete = bar.current * 100 / bar.end
}

func (bar *ProgressBar) Set(current int) {
	bar.current = current
	bar.percentComplete = bar.current * 100 / bar.end
}

func (bar *ProgressBar) Complete() {
	bar.current = bar.end
	bar.percentComplete = bar.current * 100 / bar.end
}

func (bar *ProgressBar) Reset() {
	bar.current = bar.start
	bar.percentComplete = bar.current * 100 / bar.end
}

func (bar *ProgressBar) Render(parameters ...string) {
	filled := bar.width * bar.percentComplete / 100
	unfilled := bar.width - filled

	vt100.SaveCursor()
	vt100.EraseCurrentLine()
	vt100.Print("Progress: ")
	percentageText := fmt.Sprintf("%3d%% ", bar.percentComplete)
	vt100.Print(percentageText, vt100.Yellow)
	barText := fmt.Sprintf("[%s%s] ", strings.Repeat("=", filled), strings.Repeat("-", unfilled))
	vt100.Print(barText)

	if len(parameters) > 0 {
		status := parameters[0]
		vt100.Print(status, vt100.Cyan)
	}

	vt100.RestoreCursor()

}
