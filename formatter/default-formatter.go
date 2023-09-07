package formatter

import (
	"fmt"
	"github.com/ngyewch/go-versioninfo/model"
)

type DefaultFormatter struct {
	config Config
}

func NewDefaultFormatter(config Config) *DefaultFormatter {
	return &DefaultFormatter{
		config: config,
	}
}

func (converter *DefaultFormatter) Format(info *SimpleDescribeInfo) *model.VersionInfo {
	versionInfo := &model.VersionInfo{
		Commit:          info.Commit,
		CommitTimestamp: info.CommitTimestamp,
	}

	version := converter.config.ProcessTag(info.Tag)
	if info.AdditionalCommits > 0 {
		version += fmt.Sprintf("-%d-g%s", info.AdditionalCommits, info.Commit[:7])
	}
	if info.Dirty {
		version += "-dirty"
	}
	versionInfo.Version = version

	return versionInfo
}
