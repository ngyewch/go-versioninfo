package formatter

import (
	"github.com/ngyewch/go-versioninfo/model"
	"strings"
)

type SimpleDescribeInfo struct {
	Tag               string
	Commit            string
	CommitTimestamp   int64
	AdditionalCommits int
	Dirty             bool
}

type Config struct {
	FallbackTag string
	TagPrefix   string
}

type Formatter interface {
	Format(info *SimpleDescribeInfo) *model.VersionInfo
}

func (config *Config) ProcessTag(tag string) string {
	if tag == "" {
		tag = config.FallbackTag
	}
	if (config.TagPrefix != "") && strings.HasPrefix(tag, config.TagPrefix) {
		tag = tag[len(config.TagPrefix):]
	}
	return tag
}
