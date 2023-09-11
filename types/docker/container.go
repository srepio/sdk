package docker

type Container struct {
	Id         string
	Name       string
	Image      string
	Ports      []Port
	Volumes    []Volume
	Privileged bool
}
