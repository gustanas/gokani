package gokani

// SRSDistributionContainer wraps an instance of SRSDistribution
type SRSDistributionContainer struct {
	SRSDistribution int `json:"requested_information"`
}

// SRSDistribution holds information about the Spaced repetition system
// distribution count by SRS level, then by item type
type SRSDistribution struct {
	Apprentice SRSDistributionData `json:"apprentice"`
	Guru       SRSDistributionData `json:"guru"`
	Master     SRSDistributionData `json:"master"`
	Enlighten  SRSDistributionData `json:"enlighten"`
	Burned     SRSDistributionData `json:"burned"`
}

// SRSDistributionData shows the distribution for a given level showing radical,
// kanji, vocabulary and total
type SRSDistributionData struct {
	Radical    int `json:"radical"`
	Kanji      int `json:"kanji"`
	Vocabulary int `json:"vocabulary"`
	Total      int `json:"total"`
}

// SRSDistribution returns a SRSDistribution struct
func (c Client) SRSDistribution() (*SRSDistribution, error) {
	v := SRSDistribution{}
	url := APIURL + "/user/" + c.Key + "/srs-distribution"
	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}
	return &v, nil
}
