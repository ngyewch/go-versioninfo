package resolver

import (
	"github.com/ngyewch/go-versioninfo/model"
)

type Resolver interface {
	Resolve() (*model.VersionInfo, error)
}
