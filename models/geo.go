package models

type GeoSearchParams struct {
	GeoMember

	// Distance and unit when using ByRadius option.
	// Can use m, km, ft, or mi. Default is km.
	Radius     float64 `json:"radius"`
	RadiusUnit string  `json:"radius_unit"`

	// Can be ASC or DESC. Default is no sort order.
	Sort string `json:"sort"`
}

type GeoAddParams struct {
	Key     string      `json:"key"`
	Members []GeoMember `json:"members"`
}

type GeoMember struct {
	Order     int     `json:"order,omitempty"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
