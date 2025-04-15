package leaderboard

import (
	"errors"

	"github.com/toledoom/gork/pkg/gork"
	"github.com/toledoom/gork_example/internal/domain/player"
)

type PlayerScoreUpdatedEventHandler struct {
	r Ranking
}

func NewPlayerScoreUpdatedEventHandler(r Ranking) *PlayerScoreUpdatedEventHandler {
	return &PlayerScoreUpdatedEventHandler{
		r: r,
	}
}

func (eh *PlayerScoreUpdatedEventHandler) Handle(evt gork.Event) error {
	pse, ok := evt.(*player.ScoreUpdatedEvent)
	if !ok {
		return errors.New("invalid event. Want PlayerScoreUpdated")
	}
	err := eh.r.UpdateScore(pse.PlayerID, pse.NewScore)
	return err
}
