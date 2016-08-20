package gokani

import "fmt"

// VocabularyListContainer wraps instances of Items
type VocabularyListContainer struct {
	VocabularyList Items `json:"requested_information"`
}

/*
VocabularyList returns a full list of vocabulary items with user's stats ordered
by ascending levels.. An optional argument of declaring a single or
comma-delimited list of levels is available. An example of a comma-delimited list
of levels is 1,2,5,9. The levels can be a minimum of 1 and a maximum of 60
(Maximum is always increasing as we continue to add more content). If the levels
argument is not specified, then all levels up to the user's level will be
returned.
*/
func (c Client) VocabularyList(levels []string) (Items, error) {
	v := VocabularyListContainer{}
	url := APIURL + "/user/" + c.Key + "/vocabulary/"
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

	return v.VocabularyList, nil
}
