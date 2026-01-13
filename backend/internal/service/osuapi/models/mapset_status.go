package osuapimodels

type MapsetStatusAPIOption string

const (
	Graveyard MapsetStatusAPIOption = "graveyard"
	Loved     MapsetStatusAPIOption = "loved"
	Pending   MapsetStatusAPIOption = "pending"
	Ranked    MapsetStatusAPIOption = "ranked"

	// Nominated we don't use this cause it shows maps that user nominated (from others) which breaks mapset FK
	Nominated MapsetStatusAPIOption = "nominated"
)
