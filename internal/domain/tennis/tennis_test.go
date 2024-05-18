package tennis

import (
	"testing"
)

var (
	player1 string = "Nick"
	player2 string = "Rupe"
	sets    int    = 3
)

func TestStartsTennis(t *testing.T) {
	tennis := New(player1, player2, sets)

	assertEvents([]event{
		MatchStarted{
			Players: [2]string{player1, player2},
			Sets:    sets,
		},
	}, tennis.events, t)
}

func TestScoreing(t *testing.T) {
	tennis := New(player1, player2, sets)
	tennis.events = []event{}

	tennis.Score(player2)
	tennis.Score(player2)
	tennis.Score(player1)

	assertEvents([]event{
		PointScored{
			Scorer:    player2,
			GameScore: Points{0, 1},
		},
		PointScored{
			Scorer:    player2,
			GameScore: Points{0, 2},
		},
		PointScored{
			Scorer:    player1,
			GameScore: Points{1, 2},
		},
	}, tennis.events, t)
}

func setUp() *Tennis {
	tennis := New(player1, player2, sets)
	tennis.events = []event{}

	return tennis
}

func TestGame(t *testing.T) {
	t.Run("clean game", func(t *testing.T) {
		tennis := setUp()

		tennis.Score(player1)
		tennis.Score(player1)
		tennis.Score(player1)
		tennis.Score(player1)

		assertEvents([]event{
			PointScored{
				Scorer:    player1,
				GameScore: Points{3, 0},
			},
			GameScored{
				Scorer:   player1,
				SetScore: Points{1, 0},
			},
		},
			tennis.events[len(tennis.events)-2:],
			t,
		)
	})

	t.Run("win by two points", func(t *testing.T) {
		tennis := setUp()

		tennis.Score(player2)
		tennis.Score(player2)
		tennis.Score(player2)
		// 0 3 (0 45)
		tennis.Score(player1)
		tennis.Score(player1)
		tennis.Score(player1)
		tennis.Score(player1)
		// 4 3 (A 45)
		tennis.Score(player2)
		// 4 4 (45 45)
		tennis.Score(player2)
		// 4 5 (A 45)
		tennis.Score(player2)
		// Player 2 wins!

		assertEvents([]event{
			PointScored{
				Scorer:    player2,
				GameScore: Points{4, 5},
			},
			GameScored{
				Scorer:   player2,
				SetScore: Points{0, 1},
			},
		},
			tennis.events[len(tennis.events)-2:],
			t,
		)
	})
}

func TestSet(t *testing.T) {
	t.Run("clean set", func(t *testing.T) {
		tennis := setUp()

		cleanGame(tennis, player1)
		cleanGame(tennis, player1)
		cleanGame(tennis, player1)
		cleanGame(tennis, player1)
		cleanGame(tennis, player1)
		cleanGame(tennis, player1)

		assertEvents([]event{
			PointScored{
				Scorer:    player1,
				GameScore: Points{3, 0},
			},
			SetScored{
				Scorer:     player1,
				MatchScore: Points{1, 0},
			},
		}, tennis.events[len(tennis.events)-2:], t)
	})

	t.Run("Two Games Clear", func(t *testing.T) {
		tennis := setUp()

		cleanGame(tennis, player2)
		cleanGame(tennis, player2)
		cleanGame(tennis, player2)
		cleanGame(tennis, player2)
		cleanGame(tennis, player2)

		cleanGame(tennis, player1)
		cleanGame(tennis, player1)
		cleanGame(tennis, player1)
		cleanGame(tennis, player1)
		cleanGame(tennis, player1)
		cleanGame(tennis, player1)

		cleanGame(tennis, player2)
		cleanGame(tennis, player2)
		cleanGame(tennis, player2)

		assertEvents([]event{
			SetScored{
				Scorer:     player2,
				MatchScore: Points{0, 1},
			},
		}, tennis.events[len(tennis.events)-1:], t)
	})
}

func TestMatch(t *testing.T) {
	t.Run("Clean Sets 3", func(t *testing.T) {
		tennis := setUp()

		cleanSet(tennis, player1)
		cleanSet(tennis, player1)

		assertEvents([]event{
			PointScored{
				Scorer:    player1,
				GameScore: Points{3, 0},
			},
			MatchWon{
				Winner: player1,
			},
		}, tennis.events[len(tennis.events)-2:], t)
	})

	t.Run("2-1 Win", func(t *testing.T) {
		tennis := setUp()

		cleanSet(tennis, player1)
		cleanSet(tennis, player2)
		cleanSet(tennis, player2)

		assertEvents([]event{
			PointScored{
				Scorer:    player2,
				GameScore: Points{0, 3},
			},
			MatchWon{
				Winner: player2,
			},
		}, tennis.events[len(tennis.events)-2:], t)
	})

	t.Run("No play after win", func(t *testing.T) {
		tennis := setUp()

		cleanSet(tennis, player1)
		cleanSet(tennis, player1)

		err := tennis.Score(player2)

		if err == nil {
			t.Error("Play continued")
		}
	})
}

func TestScorePlayerError(t *testing.T) {
	tennis := setUp()

	err := tennis.Score("not a player")

	if err == nil {
		t.Error("Score accepted unknown player")
	}
}

func cleanGame(tennis *Tennis, player string) {
	tennis.Score(player)
	tennis.Score(player)
	tennis.Score(player)
	tennis.Score(player)
}

func cleanSet(tennis *Tennis, player string) {
	cleanGame(tennis, player)
	cleanGame(tennis, player)
	cleanGame(tennis, player)
	cleanGame(tennis, player)
	cleanGame(tennis, player)
	cleanGame(tennis, player)
}

func assertEvents(expected []event, actual []event, t *testing.T) {
	if len(expected) != len(actual) {
		t.Error("Lenth missmatch. Expected", len(expected), "got", len(actual))
		t.Error(expected, actual)
		return
	}

	for i, expectedEvent := range expected {
		if expectedEvent.name() != actual[i].name() {
			t.Error("type name missmatch. expected", expectedEvent.name(), "got", actual[i].name())
		}

		if expectedEvent != actual[i] {
			t.Error("data missmatch at position", i, "expected", expectedEvent, "got", actual[i])
		}
	}
}
