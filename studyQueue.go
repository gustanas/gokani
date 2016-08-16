package gokani

type StudyQueueContainer struct {
	StudyQueue int `json:"requested_information"`
}

type StudyQueue struct {
	LessonsAvailable         int `json:"lessons_available"`
	ReviewsAvailable         int `json:"reviews_available"`
	NextReviewDate           int `json:"next_review_date"`
	ReviewsAvailableNextHour int `json:"reviews_available_next_hour"`
	ReviewsAvailableNextDay  int `json:"reviews_available_next_day"`
}

func (c Client) StudyQueue() (*StudyQueue, error) {
	v := StudyQueue{}
	url := APIURL + "/user/" + c.Key + "/study-queue"
	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}
	return &v, nil
}
