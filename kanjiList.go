package gokani

import "fmt"

// KanjiListContainer wraps instances of Items
type KanjiListContainer struct {
	KanjiList Items `json:"requested_information"`
}

/*
KanjiList returns a full list of kanji items with user's stats ordered by
 ascending levels.. An optional argument of declaring a single or
comma-delimited list of levels is available. An example of a comma-delimited
list of levels is 1,2,5,9. The levels can be a minimum of 1 and a maximum of 60
(Maximum is always increasing as we continue to add more content). If the levels
argument is not specified, then all levels up to the user's level will be
returned.
*/
func (c Client) KanjiList(levels []string) (Items, error) {
	v := KanjiListContainer{}
	url := APIURL + "/user/" + c.Key + "/kanji/"
	for i, s := range levels {
		url += s
		if i < len(levels)-1 {
			url += ","
		}
	}
	fmt.Println(url)
	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}

	return v.KanjiList, nil
}
