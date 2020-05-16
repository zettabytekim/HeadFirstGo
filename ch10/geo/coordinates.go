package geo

import "errors"

// Coordinates struct...
type Coordinates struct {
	latitude  float64
	longitude float64
}

// Latitude ...
func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

// Logitude ...
func (c *Coordinates) Longitude() float64 {
	return c.longitude
}

// SetLatitude ...
func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}

// SetLongitude ...
func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("invalid longitude")
	}
	c.longitude = longitude
	return nil
}
