package docker

import "github.com/srepio/sdk/types"

type Container struct {
	Id         string
	Name       string
	Image      string
	Ports      []types.Port
	Volumes    []types.Volume
	Privileged bool
}
