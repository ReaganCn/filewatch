package models

type Path struct {
	ID           int    `json:"id"`
	Path         string `json:"path"`
	LastModified string `json:"last_modified"`
	AddedOn      string `json:"added_on"`

	/* File events to to watch */
	Events []Event `json:"changes"`
}

type Event struct {
	Type string `json:"type"` // CREATE, MODIFY, DELETE
}
