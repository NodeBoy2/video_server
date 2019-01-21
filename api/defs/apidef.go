package defs

// requests

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type VideoInfo struct {
	ID           string
	AuthorID     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	Id         string
	VideoId    string
	AuthorName string
	Content    string
}

type SimpleSession struct {
	Username string // login name
	TTL      int64
}
