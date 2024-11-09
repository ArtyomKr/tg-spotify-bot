package spotify

import "net/url"

func (c *Client) GetCurrentTrack(token string) (PlaybackStatus, error) {
	var playbackStatus PlaybackStatus
	err := c.fetch("GET", "/me/player/currently-playing", token, nil, nil, &playbackStatus)

	return playbackStatus, err
}

func (c *Client) AddTrackToQueue(token string, URI string) error {
	query := url.Values{}
	query.Add("uri", URI)
	err := c.fetch("POST", "/me/player/queue", token, query, nil, nil)

	return err
}
