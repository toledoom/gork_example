package usecases

import (
	"github.com/toledoom/gork/pkg/gork"
	"github.com/toledoom/gork_example/internal/app/query"
)

type GetRankInput struct {
	PlayerID string
}

type GetRankOutput struct {
	Rank uint64
}

func GetRank(cr *gork.CommandRegistry, qr *gork.QueryRegistry) gork.UseCase[GetRankInput, GetRankOutput] {
	return func(gri GetRankInput) (GetRankOutput, error) {
		getRankQuery := query.GetRank{
			PlayerID: gri.PlayerID,
		}

		queryResponse, err := gork.HandleQuery[*query.GetRank, *query.GetRankResponse](qr, &getRankQuery)
		if err != nil {
			return GetRankOutput{}, err
		}

		return GetRankOutput{
			Rank: queryResponse.Rank,
		}, nil

	}
}
