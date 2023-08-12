package git

import "github.com/ngyewch/go-versioninfo/model"

type VersionInfoConverter interface {
	Convert(describeInfo *DescribeInfo) *model.VersionInfo
}
