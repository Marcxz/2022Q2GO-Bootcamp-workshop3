package controller

import (
	api "architecture/api"
	database "architecture/db"
)

type Controller struct {
	geo api.Geo
	db  database.DB
}

func NewController(geo api.Geo, db database.DB) *Controller {
	return &Controller{geo, db}
}

func (c *Controller) GetInfoUserById(id string) (string, error) {
	info, err := c.db.GetInfoFromDB(id)

	if err != nil {
		return "", err
	}

	return info, nil
}

func (c *Controller) GetGeocoding(address string) (float64, float64, error) {
	lat, lon, err := c.geo.GetGeocoding(address)

	if err != nil {
		return lat, lon, err
	}

	return lat, lon, nil
}
