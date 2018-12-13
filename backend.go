// This file is a part of github.com/thekondor/confita-prefixed-env

package prefixed_env

import (
	"context"
	"errors"
	confita_backend "github.com/heetch/confita/backend"
	confita_backend_env "github.com/heetch/confita/backend/env"
	"strings"
)

var (
	ErrInvalidPrefix = errors.New("Invalid or empty prefix")
)

// NewDefaultBackend() returns Confita Backend configured with underscore ('_') prefix delimiter
func NewDefaultBackend(prefix string) confita_backend.Backend {
	return NewBackend(prefix, "_")
}

// NewDefault() returns Confita Backend configured with a prefixDelimiter provided
func NewBackend(prefix string, prefixDelimeter string) confita_backend.Backend {
	impl := confita_backend_env.NewBackend()

	return confita_backend.Func("prefixed_env", func(ctx context.Context, key string) ([]byte, error) {
		if 0 == len(prefix) {
			return nil, ErrInvalidPrefix
		}

		// NOTE: no validations of `fullKey' conformity against system's environment variable names are made intentionally
		fullKey := strings.Join([]string{prefix, key}, prefixDelimeter)

		return impl.Get(ctx, fullKey)
	})
}
