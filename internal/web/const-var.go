package web

import (
	"embed"

	"github.com/aceberg/ExerciseDiary/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf

	//
	exData models.AllExData
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS
