package web

import (
	"embed"

	"github.com/aceberg/ExerciseDiary/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf

	// Exercise data
	exData models.AllExData
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS

// pubFS - public folder
//
//go:embed public/*
var pubFS embed.FS
