package Pojo

type Activity struct {
	ActivityId    string `json:"activityId"`
	ActivityType  string `json:"activityType"`
	UserId        string `json:"userId"`
	DocumentId    string `json:"documentId"`
	ViewingTime   uint64 `json:"viewingTime"`
	ViewedPages   string `json:"viewedPages"`
	NoOfDownloads uint64 `json:"noOfDownloads"`
	LoginTime     uint64 `json:"loginTime"`
	ProfileUpdate bool   `json:"profileUpdate"`
}
