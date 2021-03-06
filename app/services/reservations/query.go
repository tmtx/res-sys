package reservations

import (
	"context"

	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/aggregates"
	"github.com/tmtx/res-sys/pkg/event"
)

func (s *Service) GetAll() (result []app.Aggregate, err error) {
	// TODO: this is not entirely correct. Should restore aggregates from target events,
	// but without making query for each reservation
	events, err := s.EventRepository.FindAllWithFilter(
		context.Background(),
		event.Filter{
			"key": app.ReservationCreated,
		},
	)
	if err != nil {
		return result, err
	}

	for _, e := range events {
		ag := &aggregates.Reservation{}
		ag.HydrateFromParams(e.Params)
		result = append(result, ag)
	}

	return result, err
}
