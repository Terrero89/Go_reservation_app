package config

import (
	"html/template"
	"log"
)

// AppConfig holds configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
