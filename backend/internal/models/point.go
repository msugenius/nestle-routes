package models

type Point struct {
	Latitude  float32 `db:"Latitude" json:"latitude"`
	Longitude float32 `db:"Longitude" json:"longitude"`
	Address   string  `db:"OLDeliveryAddress" json:"address"`
	Name      string  `db:"OlName" json:"name"`
	ID        string  `db:"OL_id" json:"id"`
	Weekday   int     `db:"Weekday" json:"weekday"`
}
