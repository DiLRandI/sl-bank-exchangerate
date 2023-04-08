package config

type Config struct {
	Plugins []Plugin `json:"plugins"`
}

type Plugin struct {
	Name string `json:"name"`
	File string `json:"file"`
}
