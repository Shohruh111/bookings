package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]string
	FloatMap  map[string]string
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
