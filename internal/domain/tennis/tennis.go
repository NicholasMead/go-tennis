package tennis

import (
	"errors"
	"fmt"
)

type Tennis struct {
	events []event

	sets     int
	score    Score
	players  [2]string
	gameOver bool
}

func New(player1 string, player2 string, sets int) *Tennis {
	tennis := Tennis{
		events:   make([]event, 0),
		gameOver: false,
	}

	tennis.dispatch(MatchStarted{
		Players: [2]string{player1, player2},
		Sets:    sets,
	})

	return &tennis
}

func (t *Tennis) Score(player string) error {
	num, err := t.getPlayerNum(player)

	if err != nil {
		return err
	}

	if t.gameOver {
		return errors.New("Game has ended")
	}

	score := t.score.addPoint(num)

	points := score.Game[num]
	diff := points - score.Game[1-num]

	if points < 4 || diff < 2 {
		return t.dispatch(PointScored{
			Scorer:    player,
			GameScore: score.Game,
		})
	}

	score = score.addGame(num)

	games := score.Set[num]
	diff = games - score.Set[1-num]

	if games < 6 || diff < 2 {
		return t.dispatch(
			GameScored{
				Scorer:   player,
				SetScore: score.Set,
			},
		)
	}

	score = score.addSet(num)

	if score.Match[num]/2 >= t.sets/2 {
		return t.dispatch(MatchWon{player})
	}

	return t.dispatch(SetScored{
		Scorer:     player,
		MatchScore: score.Match,
	})
}

func (t *Tennis) dispatch(event event) error {
	t.events = append(t.events, event)
	return t.on(event)
}

func (t *Tennis) on(event event) error {
	switch ev := event.(type) {
	case MatchStarted:
		t.players = ev.Players
		t.sets = ev.Sets

	case MatchWon:
		t.gameOver = true

	case SetScored:
		t.score.Match = ev.MatchScore
		t.score.Set = Points{}
		t.score.Game = Points{}

	case GameScored:
		t.score.Set = ev.SetScore
		t.score.Game = Points{}

	case PointScored:
		t.score.Game = ev.GameScore

	default:
		return fmt.Errorf("unknown event %s", event)
	}

	return nil
}

func (t Tennis) getPlayerNum(player string) (int, error) {
	for i, n := range t.players {
		if n == player {
			return i, nil
		}
	}

	return -1, errors.New("Player not found")
}
