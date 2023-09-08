package resolver

import (
	"github.com/ngyewch/go-versioninfo/model"
)

type Resolver interface {
	Resolve() (*model.VersionInfo, error)
}

type Resolvers []Resolver

func (resolvers Resolvers) Resolve() (*model.VersionInfo, error) {
	for _, r := range resolvers {
		info, err := r.Resolve()
		if err != nil {
			continue
		}
		if info != nil {
			return info, nil
		}
	}
	return nil, nil
}
