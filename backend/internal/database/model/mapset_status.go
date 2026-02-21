package model

type MapsetStatus string

const (
	Graveyard MapsetStatus = "graveyard"
	Wip       MapsetStatus = "wip"
	Pending   MapsetStatus = "pending"
	Ranked    MapsetStatus = "ranked"
	Approved  MapsetStatus = "approved"
	Qualified MapsetStatus = "qualified"
	Loved     MapsetStatus = "loved"
)
