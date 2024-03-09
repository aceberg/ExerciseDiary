package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/ExerciseDiary/internal/auth"
	"github.com/aceberg/ExerciseDiary/internal/check"
	"github.com/aceberg/ExerciseDiary/internal/models"
)

// Get - read config from file or env
func Get(path string) (config models.Conf, authConf auth.Conf) {

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8851")
	viper.SetDefault("THEME", "grass")
	viper.SetDefault("COLOR", "light")
	viper.SetDefault("HEATCOLOR", "#03a70c")
	viper.SetDefault("PAGESTEP", 10)

	viper.SetDefault("AUTH_USER", "")
	viper.SetDefault("AUTH_PASSWORD", "")
	viper.SetDefault("AUTH_EXPIRE", "7d")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.Color, _ = viper.Get("COLOR").(string)
	config.HeatColor, _ = viper.Get("HEATCOLOR").(string)
	config.PageStep = viper.GetInt("PAGESTEP")

	authConf.Auth = viper.GetBool("AUTH")
	authConf.User, _ = viper.Get("AUTH_USER").(string)
	authConf.Password, _ = viper.Get("AUTH_PASSWORD").(string)
	authConf.ExpStr, _ = viper.Get("AUTH_EXPIRE").(string)

	authConf.Expire = auth.ToTime(authConf.ExpStr)
	config.Auth = authConf.Auth

	return config, authConf
}

// Write - write config to file
func Write(config models.Conf, authConf auth.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("color", config.Color)
	viper.Set("heatcolor", config.HeatColor)
	viper.Set("pagestep", config.PageStep)

	viper.Set("auth", authConf.Auth)
	viper.Set("auth_user", authConf.User)
	viper.Set("auth_password", authConf.Password)
	viper.Set("auth_expire", authConf.ExpStr)

	err := viper.WriteConfig()
	check.IfError(err)
}
