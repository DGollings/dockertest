// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package system // import "github.com/dgollings/dockertest/docker/pkg/system"

// Umask is not supported on the windows platform.
func Umask(newmask int) (oldmask int, err error) {
	// should not be called on cli code path
	return 0, ErrNotSupportedPlatform
}
