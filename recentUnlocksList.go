package gokani

// RecentUnlocksListContainer wraps instances of Items
type RecentUnlocksListContainer struct {
	RecentUnlocksList Items `json:"requested_information"`
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
