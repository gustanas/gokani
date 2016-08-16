package gokani

type UserInformationContainer struct {
	UserInformation UserInformation `json:"user_information"`
}

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

func (c Client) UserInformation() (*UserInformation, error) {
	v := UserInformationContainer{}
	url := APIURL + "/user/" + c.Key + "/user-information"
	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}
	return &v.UserInformation, nil
}
