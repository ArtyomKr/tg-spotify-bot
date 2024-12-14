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
	const barLength = 14
	percentage := float64(progress) / float64(total)
	completed := int(percentage * float64(barLength))
	bar := strings.Repeat("─", completed) + "●" + strings.Repeat("━", barLength-completed)
	return fmt.Sprintf("%s %s %s",
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
