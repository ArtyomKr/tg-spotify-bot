package spotify

import "net/http"

type Client struct {
	clientID     string
	clientSecret string
	redirectURI  string
	baseUrl      string
	httpClient   *http.Client
}

type TokenResBody struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type Artist struct {
	ExternalUrls ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalUrls         ExternalURLs `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         Restrictions `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Artists              []Artist     `json:"artists"`
}

type Track struct {
	Album            Album        `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIDs      ExternalIDs  `json:"external_ids"`
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       interface{}  `json:"linked_from"`
	Restrictions     Restrictions `json:"restrictions"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
}

type ExternalIDs struct {
	ISRC string `json:"isrc"`
	EAN  string `json:"ean"`
	UPC  string `json:"upc"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Restrictions struct {
	Reason string `json:"reason"`
}

type Device struct {
	ID               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
	SupportsVolume   bool   `json:"supports_volume"`
}

type PlaybackStatus struct {
	Device       `json:"device"`
	RepeatState  string `json:"repeat_state"`
	ShuffleState bool   `json:"shuffle_state"`
	Context      struct {
		Type         string       `json:"type"`
		Href         string       `json:"href"`
		ExternalUrls ExternalURLs `json:"external_urls"`
		URI          string       `json:"uri"`
	} `json:"context"`
	Timestamp  int  `json:"timestamp"`
	ProgressMs int  `json:"progress_ms"`
	IsPlaying  bool `json:"is_playing"`
	Item       struct {
		Album            `json:"album"`
		Artists          []Artist     `json:"artists"`
		AvailableMarkets []string     `json:"available_markets"`
		DiscNumber       int          `json:"disc_number"`
		DurationMs       int          `json:"duration_ms"`
		Explicit         bool         `json:"explicit"`
		ExternalIds      ExternalIDs  `json:"external_ids"`
		ExternalUrls     ExternalURLs `json:"external_urls"`
		Href             string       `json:"href"`
		ID               string       `json:"id"`
		IsPlayable       bool         `json:"is_playable"`
		LinkedFrom       struct {
		} `json:"linked_from"`
		Restrictions Restrictions `json:"restrictions"`
		Name         string       `json:"name"`
		Popularity   int          `json:"popularity"`
		PreviewURL   string       `json:"preview_url"`
		TrackNumber  int          `json:"track_number"`
		Type         string       `json:"type"`
		URI          string       `json:"uri"`
		IsLocal      bool         `json:"is_local"`
	} `json:"item"`
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions              struct {
		InterruptingPlayback  bool `json:"interrupting_playback"`
		Pausing               bool `json:"pausing"`
		Resuming              bool `json:"resuming"`
		Seeking               bool `json:"seeking"`
		SkippingNext          bool `json:"skipping_next"`
		SkippingPrev          bool `json:"skipping_prev"`
		TogglingRepeatContext bool `json:"toggling_repeat_context"`
		TogglingShuffle       bool `json:"toggling_shuffle"`
		TogglingRepeatTrack   bool `json:"toggling_repeat_track"`
		TransferringPlayback  bool `json:"transferring_playback"`
	} `json:"actions"`
}

type QueueRes struct {
	CurrentlyPlaying Track   `json:"currently_playing"`
	Queue            []Track `json:"queue"`
}
