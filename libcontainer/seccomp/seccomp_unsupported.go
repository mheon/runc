// +build !linux !cgo !seccomp

package seccomp

import (
	"errors"

	"github.com/opencontainers/runc/libcontainer/configs"
)

// Seccomp not supported, do nothing
func InitSeccomp(config *configs.Seccomp) error {
	return nil
}
