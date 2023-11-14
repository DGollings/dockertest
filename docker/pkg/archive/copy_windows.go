// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package archive // import "github.com/dgollings/dockertest/v3/docker/pkg/archive"

import (
	"path/filepath"
)

func normalizePath(path string) string {
	return filepath.FromSlash(path)
}
