package bot

import (
	"fmt"
	"strings"
	"telegram-bot/internal/spotify"
	"time"
)

func formatDuration(ms int) string {
	d := time.Duration(ms) * time.Millisecond
	return fmt.Sprintf("%d:%02d", int(d.Minutes()), int(d.Seconds())%60)
}

func formatPlaybackProgress(progress, total int) string {
	const barLength = 18
	percentage := float64(progress) / float64(total)
	completed := int(percentage * float64(barLength))

	bar := strings.Repeat("▰", completed) + strings.Repeat("▱", barLength-completed)

	// Add emoji indicators based on progress
	var indicator string
	switch {
	case percentage < 0.33:
		indicator = "⏳"
	case percentage < 0.66:
		indicator = "⌛️"
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

func formatQueueItems(tracks *[]spotify.Track) string {
	if len(*tracks) == 0 {
		return "Queue is empty"
	}

	var messageBuilder strings.Builder

	messageBuilder.WriteString("<b>📊 Queue:</b>\n")
	for i, track := range *tracks {
		messageBuilder.WriteString(fmt.Sprintf("%d. <a href=\"%s\"><b>%s</b></a>\n", i+1, track.ExternalURLs.Spotify, track.Name))
		messageBuilder.WriteString(fmt.Sprintf("   by <i>%s</i>\n", track.Artists[0].Name))
		messageBuilder.WriteString("\n")
	}

	return messageBuilder.String()
}

func formatCurrentlyPlaying(status spotify.PlaybackStatus) string {
	var messageBuilder strings.Builder

	messageBuilder.WriteString("<b>🎵 Now Playing:</b>\n\n")
	messageBuilder.WriteString(fmt.Sprintf("<a href=\"%s\"><b>%s</b></a>\n", status.Item.ExternalUrls.Spotify, status.Item.Name))
	messageBuilder.WriteString(fmt.Sprintf("by <i>%s</i>\n", status.Item.Artists[0].Name))
	messageBuilder.WriteString(fmt.Sprintf("from <a href=\"%s\"><i>%s</i></a>\n\n", status.Item.Album.ExternalUrls.Spotify, status.Item.Album.Name))

	messageBuilder.WriteString(formatPlaybackProgress(status.ProgressMs, status.Item.DurationMs))

	return messageBuilder.String()
}
