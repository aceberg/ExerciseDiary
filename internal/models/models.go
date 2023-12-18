package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	Icon     string
	DBPath   string
	DirPath  string
	ConfPath string
	NodePath string
}

// Exercise - one exercise
type Exercise struct {
	ID     int    `db:"ID"`
	Group  string `db:"GR"`
	Place  string `db:"PLACE"`
	Name   string `db:"NAME"`
	Descr  string `db:"DESCR"`
	Image  string `db:"IMAGE"`
	Color  string `db:"COLOR"`
	Weight int    `db:"WEIGHT"`
	Reps   int    `db:"REPS"`
}

// Set - one set
type Set struct {
	ID     int    `db:"ID"`
	Date   string `db:"DATE"`
	Name   string `db:"NAME"`
	Color  string `db:"COLOR"`
	Weight int    `db:"WEIGHT"`
	Reps   int    `db:"REPS"`
}

// AllExData - all sets and exercises
type AllExData struct {
	Exs  []Exercise
	Sets []Set
}

// GuiData - web gui data
type GuiData struct {
	Config   Conf
	Themes   []string
	ExData   AllExData
	GroupMap map[string]string
	OneEx    Exercise
}
