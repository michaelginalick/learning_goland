package menu

type Menu struct {
	ID       int `json:"id"`
	Schedule struct {
		ID int `json:"id"`
	} `json:"schedule"`
}