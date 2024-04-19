package models

type Route struct {
	ID      int64  `db:"Route_id" json:"id"`
	Name    string `db:"RouteOwner" json:"name"`
	Weekday string `db:"Weekday" json:"weekday"`
}
