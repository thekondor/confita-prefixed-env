// This file is a part of github.com/thekondor/confita-prefixed-env

package prefixed_env

import (
	"context"
	confita_backend "github.com/heetch/confita/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_Errors_OnNonExistingKey(t *testing.T) {
	sut := NewDefaultBackend("VENDOR_PREFIX")

	_, err := sut.Get(context.Background(), "NON_EXISTING_ENV_VAR")
	assert.Equal(t, confita_backend.ErrNotFound, err)
}

func Test_Errors_OnEmptyPrefix(t *testing.T) {
	sut := NewDefaultBackend("")

	_, err := sut.Get(context.Background(), "KEY")
	assert.Equal(t, ErrInvalidPrefix, err)
}

func Test_ReturnsValue_OnExistingKey_WithSameCase(t *testing.T) {
	sut := NewDefaultBackend("VENDOR_PREFIX")
	os.Setenv("VENDOR_PREFIX_ENV_VAR", "overwritten value")

	value, err := sut.Get(context.Background(), "ENV_VAR")
	require.NoError(t, err)
	assert.Equal(t, []byte("overwritten value"), value)
}

func Test_ReturnsValue_OnExistingKey_WithDifferentCase(t *testing.T) {
	sut := NewDefaultBackend("VENDOR_PREFIX")

	os.Setenv("VENDOR_PREFIX_EnV_VaR", "overwritten value")

	value, err := sut.Get(context.Background(), "ENV_VAR")
	require.NoError(t, err)
	assert.Equal(t, []byte("overwritten value"), value)
}
