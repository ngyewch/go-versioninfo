package git

import (
	"fmt"
	"github.com/ngyewch/go-versioninfo/model"
	"strings"
)

type DefaultVersionInfoConverter struct {
	config Config
}

func NewDefaultVersionInfoConverter(config Config) *DefaultVersionInfoConverter {
	return &DefaultVersionInfoConverter{
		config: config,
	}
}

func (converter *DefaultVersionInfoConverter) Convert(describeInfo *DescribeInfo) *model.VersionInfo {
	var versionInfo model.VersionInfo

	if describeInfo.Commit != nil {
		versionInfo.Commit = describeInfo.Commit.Hash.String()
		versionInfo.CommitTimestamp = describeInfo.Commit.Committer.When.Unix()
	}

	tagName := converter.config.FallbackTag
	if describeInfo.Tag != nil {
		tagName = describeInfo.Tag.Name
	}
	versionFromTagName := tagName
	if converter.config.DropTagPrefix && strings.HasPrefix(versionFromTagName, converter.config.TagPrefix) {
		versionFromTagName = versionFromTagName[len(converter.config.TagPrefix):]
	}

	s := versionFromTagName
	if describeInfo.AdditionalCommits > 0 {
		s += fmt.Sprintf("-%d-g%s", describeInfo.AdditionalCommits, describeInfo.Commit.Hash.String()[:7])
	}
	if describeInfo.Dirty {
		s += "-dirty"
	}

	versionInfo.Version = s

	return &versionInfo
}
