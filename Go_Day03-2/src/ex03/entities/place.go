package entities

type Place struct {
	ID       uint   `csv:"ID" json:"id"`
	Name     string `csv:"Name" json:"name"`
	Address  string `csv:"Address" json:"address"`
	Phone    string `csv:"Phone" json:"phone"`
	Location struct {
		Longitude float64 `csv:"Longitude" json:"lon"`
		Latitude  float64 `csv:"Latitude" json:"lat"`
	} `csv:"location" json:"location"`
}
