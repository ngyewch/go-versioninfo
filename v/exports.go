package v

import (
	"github.com/ngyewch/go-versioninfo/model"
	"strconv"
)

var (
	Version         string
	Commit          string
	CommitTimestamp string // UNIX epoch seconds
)

func GetVersionInfo() *model.VersionInfo {
	v := &model.VersionInfo{
		Version: Version,
		Commit:  Commit,
	}
	commitTimestamp, err := strconv.ParseInt(CommitTimestamp, 10, 64)
	if err == nil {
		v.CommitTimestamp = commitTimestamp
	}
	return v
}
