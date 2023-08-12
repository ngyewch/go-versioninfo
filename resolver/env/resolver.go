package env

import (
	"fmt"
	"github.com/ngyewch/go-versioninfo/model"
	"os"
	"strconv"
)

type Resolver struct {
	prefix string
}

func New(prefix string) *Resolver {
	return &Resolver{
		prefix: prefix,
	}
}

func (resolver *Resolver) Resolve() (*model.VersionInfo, error) {
	version := os.Getenv(fmt.Sprintf("%sVERSION", resolver.prefix))
	if version == "" {
		return nil, nil
	}
	info := model.VersionInfo{
		Version: version,
		Commit:  os.Getenv(fmt.Sprintf("%sCOMMIT", resolver.prefix)),
	}
	commitTimestampString := os.Getenv(fmt.Sprintf("%sCOMMIT_TIMESTAMP", resolver.prefix))
	if commitTimestampString != "" {
		timestamp, err := strconv.ParseInt(commitTimestampString, 10, 64)
		if err == nil {
			info.CommitTimestamp = timestamp
		}
	}
	return &info, nil
}
