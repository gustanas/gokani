package gokani

// RadicalsListContainer wraps instances of Items
type RadicalsListContainer struct {
	RadicalsList Items `json:"requested_information"`
}

/*
RadicalsList returns a full list of radical items with user's stats ordered
by ascending levels. An optional argument of declaring a single or
comma-delimited list of levels is available. An example of a comma-delimited
list of levels is 1,2,5,9. The levels can be a minimum of 1 and a maximum of 60
(Maximum is always increasing as we continue to add more content). If the levels
argument is not specified, then all levels up to the user's level will be
returned.
*/
func (c Client) RadicalsList(percentage []string) (Items, error) {
	v := RadicalsListContainer{}
	url := APIURL + "/user/" + c.Key + "/radicals/"
	for i, s := range percentage {
		url += s
		if i < len(percentage)-1 {
			url += ","
		}
	}

	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}

	return v.RadicalsList, nil
}
