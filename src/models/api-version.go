package models

import "fmt"

var (
	// - Deprecating hashing access tokens
	VERSION_2_0_0 = APIVersion{2, 0, 0}
	// - Get data endpoints
	// - Adding metrics prefix to endpoints
	VERSION_1_1_0 = APIVersion{1, 1, 0}

	// - Initial release
	VERSION_1_0_0 = APIVersion{1, 0, 0}
)

type APIVersion struct {
	Major int
	Minor int
	Patch int
}

func (version *APIVersion) IsHigherThan(comparisionVersion APIVersion) bool {
	return version.Major > comparisionVersion.Major ||
		version.Minor > comparisionVersion.Minor ||
		version.Patch > comparisionVersion.Patch
}

func (version *APIVersion) IsHigherOrEqualTo(comparisionVersion APIVersion) bool {
	return (version.Major > comparisionVersion.Major ||
		version.Minor > comparisionVersion.Minor ||
		version.Patch > comparisionVersion.Patch) || (version.Major == comparisionVersion.Major &&
		version.Minor == comparisionVersion.Minor &&
		version.Patch == comparisionVersion.Patch)
}

func (version *APIVersion) ToString() string {
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
}
