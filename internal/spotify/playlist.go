package spotify

import (
	"net/url"
)

func (c *Client) GetPlaylistTracks(token string, playlistID string) (PaginatedType[PlaylistTrack], error) {
	var tracks PaginatedType[PlaylistTrack]
	query := url.Values{"limit": {"50"}}
	err := c.fetch("GET", "/playlists/"+playlistID+"/tracks", token, query, nil, &tracks)

	return tracks, err
}
