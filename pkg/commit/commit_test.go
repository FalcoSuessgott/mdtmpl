package commit

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"
)

func TestParseConventionalCommit(t *testing.T) {
	testCases := []struct {
		name   string
		commit string
		v      string
		expV   string
		err    bool
	}{
		{
			name:   "minor",
			commit: "feat: add new feature",
			v:      "1.2.3",
			expV:   "1.3.0",
		},
		{
			name:   "patch",
			commit: "fix: add new feature",
			v:      "1.2.3",
			expV:   "1.2.4",
		},
		{
			name:   "breaking",
			commit: "chore!: add new feature",
			v:      "1.2.3",
			expV:   "2.0.0",
		},
		// {
		// 	name:   "breaking",
		// 	commit: "chore!: add new feature",
		// 	v:      "v1.2.3",
		// 	expV:   "v2.0.0",
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ParseConventionalCommit([]byte(tc.commit))

			if tc.err {
				require.Error(t, err, tc.name)

				return
			}

			require.NoError(t, err, tc.name)

			v, err := semver.NewVersion(tc.v)
			require.NoError(t, err, tc.name)
			require.Equal(t, tc.expV, res(v), tc.name)
		})
	}
}
