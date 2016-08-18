package gokani

// UserInformationContainer wraps an instance of UserInformation
type UserInformationContainer struct {
	UserInformation UserInformation `json:"user_information"`
}

// UserInformation holds User-specific information
type UserInformation struct {
	Username     string `json:"username"`
	Gravatar     string `json:"gravatar"`
	Level        int    `json:"level"`
	Title        string `json:"title"`
	About        string `json:"about"`
	Website      string `json:"website"`
	Twitter      string `json:"twitter"`
	TopicsCount  int    `json:"topics_count"`
	PostsCount   int    `json:"posts_count"`
	CreationDate int    `json:"creation_date"`
	VacationDate int    `json:"vacation_date"`
}

// UserInformation returns a UserInformation struct
func (c Client) UserInformation() (*UserInformation, error) {
	v := UserInformationContainer{}
	url := APIURL + "/user/" + c.Key + "/user-information"
	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}
	return &v.UserInformation, nil
}
