package main

import (
	"time"

	"github.com/nsf/termbox-go"
	music "service.music/utils"
	"service.music/utils/animation"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	quit := make(chan struct{})
	go handleQuit(quit)

	characterIndex := 0
	waveStep := 0

	for {
		select {
		case <-quit:
			return
		default:
			audioPlaying := music.CheckSystemAudioPlaying()

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			width, height := termbox.Size()
			centerX, centerY := width/2, height/2

			if audioPlaying {
				animation.DrawCharacter(animation.Characters[characterIndex], centerX-1, centerY-1)
				characterIndex = (characterIndex + 1) % len(animation.Characters)
			} else {
				animation.DrawWave(animation.Wave, waveStep, width, height)
				waveStep = (waveStep + 1) % len(animation.Wave)
			}

			termbox.Flush()
			time.Sleep(200 * time.Millisecond) // Скорость анимации
		}
	}
}

// handleQuit обрабатывает событие выхода из программы при нажатии клавиши `q` или `Esc`.
func handleQuit(quit chan struct{}) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Ch == 'q' {
				close(quit)
				return
			}
		}
	}
}
