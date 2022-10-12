package cinemastorage

import (
	cinemamodel "ControlMovieTicket/Model/Cinema/model"
	"context"
)

func (c *sqlStore) CreateCinema(context context.Context, data *cinemamodel.Cinema) error {
	if err := c.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
