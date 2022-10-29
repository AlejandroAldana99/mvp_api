package models

type PackageData struct {
	Weight      float32 `json:"weight"`
	Size        string  `json:"size"`
	Service     string  `json:"service" default:"Standard"`
	Description string  `json:"description,omitempty"`
}
