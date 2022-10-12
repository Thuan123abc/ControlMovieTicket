package cinemabusine

import (
	cinemamodel "ControlMovieTicket/Model/Cinema/model"
	"context"
	"errors"
)

type CreateCinemaStore interface {
	CreateCinema(context context.Context, data *cinemamodel.Cinema) error
}
type CreateCinemaBiz struct {
	store CreateCinemaStore
}

func NewCinemaBiz(store CreateCinemaStore) *CreateCinemaBiz {
	return &CreateCinemaBiz{store: store}
}
func (biz *CreateCinemaBiz) CreateCinema(context context.Context, data *cinemamodel.Cinema) error {
	if data.Name == "" {
		return errors.New("Name Ciname cannt empty")
	}
	if err := biz.CreateCinema(context, data); err != nil {
		return err
	}
	return nil
}
