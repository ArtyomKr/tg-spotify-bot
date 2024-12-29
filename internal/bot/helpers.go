package bot

import (
	"regexp"
)

var spotifyTrackPattern = regexp.MustCompile(`(?:https?://)?(?:open\.)?spotify\.com/track/([a-zA-Z0-9]+)`)
var spotifyAlbumPattern = regexp.MustCompile(`(?:https?://)?(?:open\.)?spotify\.com/album/([a-zA-Z0-9]+)`)
var spotifyPlaylistPattern = regexp.MustCompile(`(?:https?://)?(?:open\.)?spotify\.com/playlist/([a-zA-Z0-9]+)`)

func getTrackIDsFromString(url string) []string {
	matches := spotifyTrackPattern.FindAllStringSubmatch(url, -1)

	if len(matches) > 0 {
		trackIDs := make([]string, 0, len(matches))

		for _, match := range matches {
			if len(match) >= 2 {
				trackIDs = append(trackIDs, match[1])
			}
		}

		return trackIDs
	}

	return []string{}
}

func getAlbumIDFromString(url string) string {
	var ID string

	match := spotifyAlbumPattern.FindStringSubmatch(url)
	if len(match) >= 2 {
		ID = match[1]
	}

	return ID
}

func getPlaylistIDFromString(url string) string {
	var ID string

	match := spotifyPlaylistPattern.FindStringSubmatch(url)
	if len(match) >= 2 {
		ID = match[1]
	}

	return ID
}

func wrapTrackIdsIntoURIs(IDs []string) []string {
	URIs := make([]string, len(IDs))

	for i, ID := range IDs {
		URIs[i] = "spotify:track:" + ID
	}

	return URIs
}
