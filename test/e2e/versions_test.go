package e2e_test

import (
	"testing"

	"github.com/celestiaorg/celestia-app/test/e2e"
	"github.com/stretchr/testify/require"
)

func TestVersionParsing(t *testing.T) {
	versionStr := "v1.3.0 v1.1.0 v1.2.0-rc0"
	versions := e2e.ParseVersions(versionStr)
	require.Len(t, versions, 3)
	require.Len(t, versions.FilterOutReleaseCandidates(), 2)
	require.Equal(t, versions.GetLatest(), e2e.Version{1, 3, 0, false, 0})
}

// Test case with multiple major versions and filtering out a single major version
func TestFilterMajorVersions(t *testing.T) {
	versionStr := "v2.0.0 v1.1.0 v2.1.0-rc0 v1.2.0 v2.2.0 v1.3.0"
	versions := e2e.ParseVersions(versionStr)
	require.Len(t, versions, 6)
	require.Len(t, versions.FilterMajor(1), 3)
}

// Test case to check the Order function
func TestOrder(t *testing.T) {
	versionStr := "v1.3.0 v1.1.0 v1.2.0-rc0 v1.4.0 v1.2.1 v2.0.0"
	versions := e2e.ParseVersions(versionStr)
	versions.Order()
	require.Equal(t, versions[0], e2e.Version{1, 1, 0, false, 0})
	require.Equal(t, versions[1], e2e.Version{1, 2, 0, true, 0})
	require.Equal(t, versions[2], e2e.Version{1, 2, 1, false, 0})
	require.Equal(t, versions[3], e2e.Version{1, 3, 0, false, 0})
	require.Equal(t, versions[4], e2e.Version{1, 4, 0, false, 0})
	require.Equal(t, versions[5], e2e.Version{2, 0, 0, false, 0})
	for i := len(versions) - 1; i > 0; i-- {
		require.True(t, versions[i].IsGreater(versions[i-1]))
	}
}

func TestOrderOfReleaseCandidates(t *testing.T) {
	versionsStr := "v1.0.0 v1.0.0-rc0 v1.0.0-rc1"
	versions := e2e.ParseVersions(versionsStr)
	versions.Order()
	require.Equal(t, versions[0], e2e.Version{1, 0, 0, true, 0})
	require.Equal(t, versions[1], e2e.Version{1, 0, 0, true, 1})
	require.Equal(t, versions[2], e2e.Version{1, 0, 0, false, 0})
}
