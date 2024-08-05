package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/nsf/termbox-go"
)

type MoodRequest struct {
	Mood string `json:"mood"`
}

type PlaylistResponse struct {
	Playlist []string `json:"playlist"`
}

type LikeRequest struct {
	TrackID string `json:"track_id"`
	Mood    string `json:"mood"`
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// go server.StartServer()

	quit := make(chan struct{})
	go handleQuit(quit)

	// characterIndex := 0
	// waveStep := 0

	for {
		select {
		case <-quit:
			return
		default:
			// mood := determineMood(heartRate)

			// playlist := getPlaylistFromMood(mood)

			// log.Printf("Current Heart Rate: %.2f, Mood: %s, Playlist: %v", heartRate, mood, playlist)

			// termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			// width, height := termbox.Size()
			// centerX, centerY := width/2, height/2

			// if heartRate > 0 {
			// 	animation.DrawCharacter(animation.Characters[characterIndex], centerX-1, centerY-1)
			// 	characterIndex = (characterIndex + 1) % len(animation.Characters)
			// } else {
			// 	animation.DrawWave(animation.Wave, waveStep, width, height)
			// 	waveStep = (waveStep + 1) % len(animation.Wave)
			// }

			termbox.Flush()
			time.Sleep(200 * time.Millisecond) // Скорость анимации
		}
	}
}

func determineMood(heartRate float64) string {
	if heartRate > 100 {
		return "angry"
	} else if heartRate > 80 {
		return "happy"
	} else if heartRate > 60 {
		return "relaxed"
	} else {
		return "sad"
	}
}

func getPlaylistFromMood(mood string) []string {
	url := "http://python-app:5000/get_playlist"
	moodRequest := MoodRequest{Mood: mood}
	requestBody, _ := json.Marshal(moodRequest)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Printf("Error sending request to Python app: %v", err)
		return []string{}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var playlistResponse PlaylistResponse
	json.Unmarshal(body, &playlistResponse)

	return playlistResponse.Playlist
}

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
