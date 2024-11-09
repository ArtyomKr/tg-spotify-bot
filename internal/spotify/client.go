package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func NewClient() *Client {
	return &Client{
		httpClient:   &http.Client{},
		clientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		clientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		redirectURI:  os.Getenv("SPOTIFY_REDIRECT_URI"),
		baseUrl:      SPOTIFY_BASE_URL,
	}
}

func (c *Client) fetch(method string, path string, token string, query url.Values, body io.Reader, result any) error {
	u, err := url.Parse(c.baseUrl + path)
	if err != nil {
		return fmt.Errorf("creating request url failed: %w", err)
	}
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return fmt.Errorf("creating request failed: %w", err)
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("making request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("bad status: %d", res.StatusCode)
	}

	if result != nil {
		if err := json.NewDecoder(res.Body).Decode(result); err != nil {
			return fmt.Errorf("decoding json failed: %d", res.StatusCode)
		}
	}

	return nil
}
