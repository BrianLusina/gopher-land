package gostructs

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

// NewSemanticVersion creates a new Semantic Version
func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (sv *SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

func (sv *SemanticVersion) incrementMajor() {
	sv.major += 1
}

func (sv *SemanticVersion) incrementMinor() {
	sv.minor += 1
}

func (sv *SemanticVersion) incrementPatch() {
	sv.patch += 1
}
