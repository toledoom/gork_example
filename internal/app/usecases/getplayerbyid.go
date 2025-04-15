package usecases

import (
	"github.com/toledoom/gork/pkg/gork"
	"github.com/toledoom/gork_example/internal/app/query"
	"github.com/toledoom/gork_example/internal/domain/player"
)

type GetPlayerByIDInput struct {
	PlayerID string
}

type GetPlayerByIDOutput struct {
	Player *player.Player
}

func GetPlayerByID(cr *gork.CommandRegistry, qr *gork.QueryRegistry) gork.UseCase[GetPlayerByIDInput, GetPlayerByIDOutput] {
	return func(gpbid GetPlayerByIDInput) (GetPlayerByIDOutput, error) {
		q := query.GetPlayerByID{
			PlayerID: gpbid.PlayerID,
		}
		response, err := gork.HandleQuery[*query.GetPlayerByID, *query.GetPlayerByIDResponse](qr, &q)
		if err != nil {
			return GetPlayerByIDOutput{}, err
		}

		return GetPlayerByIDOutput{
			Player: response.Player,
		}, nil
	}
}
