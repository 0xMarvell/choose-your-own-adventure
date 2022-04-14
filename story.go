package cyoa

import (
	"encoding/json"
	"io"

	cyoa "github.com/Marvellous-Chimaraoke/choose-your-own-adventure"
)

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}
