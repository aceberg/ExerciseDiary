package web

import (
	// "embed"

	"github.com/aceberg/ExerciseDiary/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf
)

// TemplHTML - html templates
//
// //go:embed templates/*
// var TemplHTML embed.FS

// // TemplPath - path to html templates
// const TemplPath = "templates/"

// TemplPath - path to html templates
const TemplPath = "../../internal/web/templates/"
