package spotify

import "net/url"

func (c *Client) GetAlbum(token string, albumID string) (Album, error) {
	var album Album
	err := c.fetch("GET", "/albums/"+albumID, token, nil, nil, &album)

	return album, err
}

func (c *Client) GetAlbumTracks(token string, albumID string) (PaginatedType[Track], error) {
	var tracks PaginatedType[Track]
	query := url.Values{"limit": {"50"}}
	err := c.fetch("GET", "/albums/"+albumID+"/tracks", token, query, nil, &tracks)

	return tracks, err
}
