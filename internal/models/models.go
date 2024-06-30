package models

type Request struct {
	Language      string       `json:"language"`
	Preload       string       `json:"preload"`
	Code          string       `json:"code"`
	EnableNetwork bool         `json:"enable_network"`
	Dependencies  []Dependency `json:"dependencies"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RunData struct {
	Stdout string `json:"stdout"`
	Error  string `json:"error"`
}

type DependenciesData struct {
	Dependencies []Dependency `json:"dependencies"`
}

type Dependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
