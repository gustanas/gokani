package gokani

// LevelProgressionContainer wraps an instance of LevelProgression
type LevelProgressionContainer struct {
	LevelProgression int `json:"requested_information"`
}

// LevelProgression holds User's current level radical and kanji progression.
type LevelProgression struct {
	RadicalsProgress int `json:"radicals_progress"`
	RadicalsTotal    int `json:"radicals_total"`
	KanjiProgress    int `json:"kanji_progress"`
	KanjiTotal       int `json:"kanji_total"`
}

// LevelProgression returns a LevelProgression struct
func (c Client) LevelProgression() (*LevelProgression, error) {
	v := LevelProgression{}
	url := APIURL + "/user/" + c.Key + "/level-progression"
	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}
	return &v, nil
}
