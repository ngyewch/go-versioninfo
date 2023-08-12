package model

type VersionInfo struct {
	Version         string
	Commit          string
	CommitTimestamp int64 // UNIX epoch seconds
}
