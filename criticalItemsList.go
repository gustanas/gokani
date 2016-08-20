package gokani

import "fmt"

// CriticalItemsListContainer wraps instances of Items
type CriticalItemsListContainer struct {
	CriticalItemsList Items `json:"requested_information"`
}

/*
CriticalItemsList returns a list of CriticalItemsList ordered by ascending percentage
Optional Arguments:
Declaring a maximum percentage is available. All items with a percentage
correct less than the declared argument will return. The percentage can be a
minimum of 0 and a maximum of 100. If the percentage argument is not specified
or outside the range, a default of 75 will be used.
*/
func (c Client) CriticalItemsList(percentage string) (Items, error) {
	v := CriticalItemsListContainer{}
	url := APIURL + "/user/" + c.Key + "/critical-items/" + percentage
	fmt.Println(url)
	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}
	return v.CriticalItemsList, nil
}
