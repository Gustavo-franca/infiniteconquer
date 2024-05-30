package findgamemap

type GameRepository interface {
	Find(gmID string) (Map, error)
}

type (
	Map struct {
		ID          string      `json:"id"`
		MapName     string      `json:"map_name"`
		SizeX       int64       `json:"size_x"`
		SizeY       int64       `json:"size_y"`
		Cities      []City      `json:"cities"`
		Territories []Territory `json:"territories"`
	}

	City struct {
		CityName  string `json:"city Name"`
		PositionX int64  `json:"position_X"`
		PositionY int64  `json:"position_Y"`
	}

	Territory struct {
		Type      string `json:"type"`
		PositionX int64  `json:"position_X"`
		PositionY int64  `json:"position_Y"`
	}

	FindGameMap struct {
		gameRepository GameRepository
	}
)

func NewMapFinder(gr GameRepository) FindGameMap {
	return FindGameMap{
		gameRepository: gr,
	}
}

func (f FindGameMap) FindMap(id string) (Map, error) {
	return f.gameRepository.Find(id)
}
