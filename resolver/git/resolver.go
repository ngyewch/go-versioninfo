package git

import (
	"github.com/ngyewch/go-versioninfo/model"
)

type Resolver struct {
	config    Config
	convertor VersionInfoConverter
	helper    *Helper
}

type Config struct {
	TagPrefix     string
	DropTagPrefix bool
	FallbackTag   string
	CheckDirty    bool
}

func New(config Config, converter VersionInfoConverter) (*Resolver, error) {
	repo, err := FindRepository(".")
	if err != nil {
		return nil, err
	}
	return &Resolver{
		config:    config,
		convertor: converter,
		helper:    NewHelper(repo),
	}, nil
}

func (resolver *Resolver) Resolve() (*model.VersionInfo, error) {
	describeInfo, err := resolver.helper.Describe(resolver.config.TagPrefix, resolver.config.CheckDirty)
	if err != nil {
		return nil, err
	}
	return resolver.convertor.Convert(describeInfo), nil
}
