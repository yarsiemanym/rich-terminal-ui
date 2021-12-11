package main

import (
	"time"

	"github.com/yarsiemanym/rich-terminal-ui/widgets"
)

func main() {
	statuses := []string{
		"Reticulating Splines ...",
		"Adding Hidden Agendas ...",
		"Calculating Llama Expectoration Trajectory ...",
		"Resolving GUID Conflict ...",
		"Splatting Transforms ...",
		"Routing Neural Network Infrastructure ...",
		"Retracting Phong Shader ...",
		"Obfuscating Quigley Matrix ...",
	}
	bar := widgets.ProgressBar{}
	bar.New(0, 100, 20)

	bar.Render()

	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Advance(1)
		status := statuses[len(statuses)*i/100]
		bar.Render(status)
	}

	bar.Complete()
	bar.Render("Done!")
}
