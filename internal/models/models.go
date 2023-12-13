package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	Icon     string
	ConfPath string
	YamlPath string
	NodePath string
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Themes  []string
	Example string
}
