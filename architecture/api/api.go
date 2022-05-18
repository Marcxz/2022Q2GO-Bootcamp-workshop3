package api

// geo - interface that makes the contract to connects to an external api to do the geocoding process
type Geo interface {
	GetGeocoding(address string) (float64, float64, error)
}

type geocoder struct {
}

func NewGeo() Geo {
	return &geocoder{}
}

func (g *geocoder) GetGeocoding(address string) (float64, float64, error) {
	return -100.0, -100.0, nil
}
