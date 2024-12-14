package bot

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var spotifyTrackPattern = regexp.MustCompile(`(?:https?://)?(?:open\.)?spotify\.com/track/([a-zA-Z0-9]+)`)

func formatDuration(ms int) string {
	d := time.Duration(ms) * time.Millisecond
	return fmt.Sprintf("%d:%02d", int(d.Minutes()), int(d.Seconds())%60)
}

func formatPlaybackProgress(progress, total int) string {
	const barLength = 20
	percentage := float64(progress) / float64(total)
	completed := int(percentage * float64(barLength))

	bar := strings.Repeat("▰", completed) + strings.Repeat("▱", barLength-completed)

	// Add emoji indicators based on progress
	var indicator string
	switch {
	case percentage < 0.33:
		indicator = "⏳"
	case percentage < 0.66:
		indicator = "🎵"
	default:
		indicator = "🎶"
	}

	return fmt.Sprintf("%s %s <code>%s</code> %s",
		indicator,
		formatDuration(progress),
		bar,
		formatDuration(total),
	)
}

func getSpotifyURI(url string) []string {
	matches := spotifyTrackPattern.FindAllStringSubmatch(url, -1)
	URIs := make([]string, 0, len(matches))

	for _, match := range matches {
		if len(match) >= 2 {
			URI := "spotify:track:" + match[1]
			URIs = append(URIs, URI)
		}
	}

	return URIs
}
