// +build linux

package cgroups

import (
	"github.com/opencontainers/runc/libcontainer/configs"
	"github.com/pkg/errors"
)

var (
	ErrNoCgroups = errors.New("not available when cgroups are disabled")
)

type NullCgroups struct{}

func NewNullCgroupsManager() (func(config *configs.Cgroup, paths map[string]string) Manager, error) {
	return func(config *configs.Cgroup, paths map[string]string) Manager {
		return &NullCgroups{}
	}, nil
}

func (n *NullCgroups) Apply(pid int) error {
	// Don't error. Our CGroup is being managed externally.
	return nil
}

func (n *NullCgroups) GetPids() ([]int, error) {
	return nil, ErrNoCgroups
}

func (n *NullCgroups) GetAllPids() ([]int, error) {
	return nil, ErrNoCgroups
}

func (n *NullCgroups) GetStats() (*Stats, error) {
	return nil, ErrNoCgroups
}

func (n *NullCgroups) Freeze(state configs.FreezerState) error {
	return ErrNoCgroups
}

func (n *NullCgroups) Destroy() error {
	// Don't error. Our CGroup is being managed externally.
	return nil
}

func (n *NullCgroups) GetPaths() map[string]string {
	return nil
}

func (n *NullCgroups) Set(container *configs.Config) error {
	return nil
}
