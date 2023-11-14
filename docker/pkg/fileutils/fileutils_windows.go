// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package fileutils // import "github.com/dgollings/dockertest/docker/pkg/fileutils"

// GetTotalUsedFds Returns the number of used File Descriptors. Not supported
// on Windows.
func GetTotalUsedFds() int {
	return -1
}
