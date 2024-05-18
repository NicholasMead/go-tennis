package tennis

type Points [2]int

type Score struct {
	Match Points
	Set   Points
	Game  Points
}

func (s Score) addPoint(player int) Score {
	s.Game[player]++
	return s
}

func (s Score) addGame(player int) Score {
	s.Set[player]++
	s.Game = Points{}

	return s
}

func (s Score) addSet(player int) Score {
	s.Match[player]++
	s.Set = Points{}
	s.Game = Points{}

	return s
}
