package tennis

type event interface {
	name() string
}

type MatchStarted struct {
	Players [2]string
	Sets    int
}

func (MatchStarted) name() string {
	return "MATCH_STARTED"
}

type SetScored struct {
	Scorer     string
	MatchScore Points
}

func (SetScored) name() string {
	return "SET_SCORED"
}

type GameScored struct {
	Scorer   string
	SetScore Points
}

func (GameScored) name() string {
	return "GAME_SCORED"
}

type PointScored struct {
	Scorer    string
	GameScore Points
}

func (PointScored) name() string {
	return "POINT_SCORED"
}

type MatchWon struct {
	Winner string
}

func (MatchWon) name() string {
	return "MATCH_WON"
}
