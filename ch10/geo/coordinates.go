package geo

// Coordinates struct...
type Coordinates struct {
	Latutude  float64
	Longitude float64
}

// SetLatitude ...
func (c *Coordinates) SetLatitude(latitude float64) {
	c.Latutude = latitude
}

// SetLongitude ...
func (c *Coordinates) SetLongitude(longitude float64) {
	c.Longitude = longitude
}
