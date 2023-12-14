package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	Icon     string
	DirPath  string
	ConfPath string
	NodePath string
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Themes  []string
	Example string
}
