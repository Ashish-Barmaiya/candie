package collection

type Collection struct {
	Name     string   `json:"name"`
	Path     string   `json:"path"`
	Manifest Manifest `json:"manifest"`
}
