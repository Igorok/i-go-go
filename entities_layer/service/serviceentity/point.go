// GeoJSON:
// {"type": "Point", "coordinates": [lon, lat]}

package serviceentity

// Loc is the entity of geographic point
type Loc struct {
	Type        string
	Coordinates [2]int
}

// Point is the entity of point from delivery
type Point struct {
	ID         string
	IDCl       string
	IDEst      string
	IDDlv      string
	IDUsr      string
	DateCreate string
	Loc        Loc
}
