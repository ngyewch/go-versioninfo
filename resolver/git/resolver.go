package git

import (
	"github.com/ngyewch/go-versioninfo/formatter"
	"github.com/ngyewch/go-versioninfo/model"
)

type Resolver struct {
	config    Config
	formatter formatter.Formatter
}

type Config struct {
	TagPrefix  string
	CheckDirty bool
}

func New(config Config, formatter formatter.Formatter) *Resolver {
	return &Resolver{
		config:    config,
		formatter: formatter,
	}
}

func (resolver *Resolver) Resolve() (*model.VersionInfo, error) {
	repo, err := FindRepository(".")
	if err != nil {
		return nil, err
	}
	helper := NewHelper(repo)

	describeInfo, err := helper.Describe(resolver.config.TagPrefix, resolver.config.CheckDirty)
	if err != nil {
		return nil, err
	}

	simpleDescribeInfo := &formatter.SimpleDescribeInfo{
		AdditionalCommits: describeInfo.AdditionalCommits,
		Dirty:             describeInfo.Dirty,
	}
	if describeInfo.Tag != nil {
		simpleDescribeInfo.Tag = describeInfo.Tag.Name
	}
	if describeInfo.Commit != nil {
		simpleDescribeInfo.Commit = describeInfo.Commit.Hash.String()
		simpleDescribeInfo.CommitTimestamp = describeInfo.Commit.Committer.When.Unix()
	}
	return resolver.formatter.Format(simpleDescribeInfo), nil
}
