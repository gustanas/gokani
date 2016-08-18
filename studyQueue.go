package gokani

// StudyQueueContainer wraps an instance of StudyQueue
type StudyQueueContainer struct {
	StudyQueue int `json:"requested_information"`
}

// StudyQueue holds information about the user's study queue
type StudyQueue struct {
	LessonsAvailable         int `json:"lessons_available"`
	ReviewsAvailable         int `json:"reviews_available"`
	NextReviewDate           int `json:"next_review_date"`
	ReviewsAvailableNextHour int `json:"reviews_available_next_hour"`
	ReviewsAvailableNextDay  int `json:"reviews_available_next_day"`
}

// StudyQueue returns a StudyQueue
func (c Client) StudyQueue() (*StudyQueue, error) {
	v := StudyQueue{}
	url := APIURL + "/user/" + c.Key + "/study-queue"
	e := Call("GET", url, &v)
	if e != nil {
		return nil, e
	}
	return &v, nil
}
