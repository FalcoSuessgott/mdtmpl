package commit

import (
	"errors"

	"github.com/Masterminds/semver/v3"
	"github.com/leodido/go-conventionalcommits"
	"github.com/leodido/go-conventionalcommits/parser"
)

type SemVerFunc func(*semver.Version) string

var (
	IncMajor SemVerFunc = func(v *semver.Version) string { return v.IncMajor().String() }
	IncMinor SemVerFunc = func(v *semver.Version) string { return v.IncMinor().String() }
	IncPatch SemVerFunc = func(v *semver.Version) string { return v.IncPatch().String() }
)

func ParseConventionalCommit(commit []byte) (SemVerFunc, error) {
	cc, err := parser.NewMachine(
		conventionalcommits.WithTypes(conventionalcommits.TypesConventional),
		conventionalcommits.WithBestEffort()).Parse(commit)
	if err != nil {
		return nil, err
	}

	if cc.IsBreakingChange() {
		return IncMajor, nil
	}

	if cc.IsFeat() {
		return IncMinor, nil
	}

	if cc.IsFix() {
		return IncPatch, nil
	}

	return nil, errors.New("commit is not a conventional commit")
}
