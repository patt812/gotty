package game

import (
	"fmt"
	"strings"
	"time"
)

const progressBase = 20

func RunTimer(startTime time.Time, stopTimer chan struct{}, pauseTimer chan bool) {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	running := true

	for {
		select {
		case <-stopTimer:
			return
		case pause := <-pauseTimer:
			running = !pause
		case <-ticker.C:
			if running {
				elapsed := time.Since(startTime)
				timerString := fmt.Sprintf("\033[2K\r%02d.%03d", int(elapsed.Seconds()), int(elapsed.Milliseconds()%1000))
				fmt.Print("\033[2B")
				fmt.Print(timerString)
				fmt.Print("\033[2A")
			}
		}
	}
}

func ShowProgressBar(current, total int) {
	progress := int(float64(current) / float64(total) * progressBase)
	bar := strings.Repeat("=", progress-1) + ">" + strings.Repeat("-", progressBase-progress)
	fmt.Printf("\033[2K\r%d / %d [%s]", current, total, bar)
}
