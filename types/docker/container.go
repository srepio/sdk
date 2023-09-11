package docker

type Container struct {
	Id         string
	Name       string
	Image      string
	Ports      []Port
	Volumes    []Volume
	Privileged bool
	PlayID     int64
}

func (c *Container) GetPlayID() int64 {
	return c.PlayID
}

func (c *Container) SetPlayID(id int64) {
	c.PlayID = id
}
