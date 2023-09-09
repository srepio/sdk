package docker

type Port struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}

type Volume struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}
