package constants

import "time"

var (
	Loc, _ = time.LoadLocation("Asia/Ho_Chi_Minh")
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
)
