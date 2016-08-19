package gokani

// RecentUnlocksListContainer wraps instances of Items
type RecentUnlocksListContainer struct {
	RecentUnlocksList Items `json:"requested_information"`
}

// Items represents a list of Item
type Items []Item

// Item holds information about the kanji, radical or vocabulary item
type Item struct {
	ItemType         string `json:"type"`
	Meaning          string `json:"meaning"`
	Character        string `json:"character"`
	Onyomi           string `json:"onyomi"`
	Kunyomi          string `json:"kunyomi"`
	Nanori           string `json:"nanori"`
	ImportantReading string `json:"important_reading"`
	Level            int    `json:"level"`
	UnlockedDate     int    `json:"unlocked_date"`
	Kana             string `json:"kana"`
	Image            string `json:"image"`
	Percentage       string `json:"percentage"`
}

// RecentUnlocksList returns a list of RecentUnlocksList
// The limit can be a minimum of 1 and a maximum of 100
// If the limit argument is not specified or outside the range, a default of 10
// will be used.
func (c Client) RecentUnlocksList(limit string) (Items, error) {
	v := RecentUnlocksListContainer{}
	url := APIURL + "/user/" + c.Key + "/recent-unlocks/" + limit
	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}
	return v.RecentUnlocksList, nil
}
