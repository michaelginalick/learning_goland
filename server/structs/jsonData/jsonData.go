package jsonData

import(
	"time"
)

type JsonData struct {
	SiteID       int       `json:"siteId"`
	ClientID     string    `json:"clientID"`
	ModifiedDate time.Time `json:"modifiedDate"`
}