package music

import (
	"os/exec"
	"strings"
)

// checkMusicPlaying проверяет состояние воспроизведения музыки в указанном приложении.
func checkMusicPlaying(appName string) bool {
	var script string
	if appName == "Music" {
		script = `osascript -e 'tell application "Music" to player state as string'`
	} else if appName == "Spotify" {
		script = `osascript -e 'tell application "Spotify" to player state as string'`
	}

	cmd := exec.Command("sh", "-c", script)
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(output)) == "playing"
}

// CheckSystemAudioPlaying проверяет, воспроизводится ли музыка в системных приложениях.
func CheckSystemAudioPlaying() bool {
	// Проверяем Apple Music
	if checkMusicPlaying("Music") {
		return true
	}
	// Проверяем Spotify
	if checkMusicPlaying("Spotify") {
		return true
	}

	return false
}
