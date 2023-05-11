package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool                // global var to check if things are or should be in prod or development.
	Session       *scs.SessionManager // points at the session globally to be changed anywhere is referred to
}
