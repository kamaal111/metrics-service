package models

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
