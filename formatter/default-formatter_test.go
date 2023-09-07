package formatter

import (
	"github.com/ngyewch/go-versioninfo/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultFormatter(t *testing.T) {
	formatter := NewDefaultFormatter(Config{
		TagPrefix:   "v",
		FallbackTag: "v0.0.0",
	})

	{
		actual := formatter.Format(&SimpleDescribeInfo{
			Tag:               "v1.2.3",
			Commit:            "96251e73293db1b47ae17c56320a077d608aec67",
			CommitTimestamp:   1694092640,
			AdditionalCommits: 0,
			Dirty:             false,
		})
		expected := &model.VersionInfo{
			Version:         "1.2.3",
			Commit:          "96251e73293db1b47ae17c56320a077d608aec67",
			CommitTimestamp: 1694092640,
		}
		assert.Equal(t, expected, actual, "expected %+v, actual %+v", expected, actual)
	}

	{
		actual := formatter.Format(&SimpleDescribeInfo{
			Tag:               "v1.2.3",
			Commit:            "96251e73293db1b47ae17c56320a077d608aec67",
			CommitTimestamp:   1694092640,
			AdditionalCommits: 5,
			Dirty:             false,
		})
		expected := &model.VersionInfo{
			Version:         "1.2.3-5-g96251e7",
			Commit:          "96251e73293db1b47ae17c56320a077d608aec67",
			CommitTimestamp: 1694092640,
		}
		assert.Equal(t, expected, actual, "expected %+v, actual %+v", expected, actual)
	}

	{
		actual := formatter.Format(&SimpleDescribeInfo{
			Tag:               "v1.2.3",
			Commit:            "96251e73293db1b47ae17c56320a077d608aec67",
			CommitTimestamp:   1694092640,
			AdditionalCommits: 0,
			Dirty:             true,
		})
		expected := &model.VersionInfo{
			Version:         "1.2.3-dirty",
			Commit:          "96251e73293db1b47ae17c56320a077d608aec67",
			CommitTimestamp: 1694092640,
		}
		assert.Equal(t, expected, actual, "expected %+v, actual %+v", expected, actual)
	}

	{
		actual := formatter.Format(&SimpleDescribeInfo{
			Tag:               "1.2.3",
			Commit:            "96251e73293db1b47ae17c56320a077d608aec67",
			CommitTimestamp:   1694092640,
			AdditionalCommits: 10,
			Dirty:             true,
		})
		expected := &model.VersionInfo{
			Version:         "1.2.3-10-g96251e7-dirty",
			Commit:          "96251e73293db1b47ae17c56320a077d608aec67",
			CommitTimestamp: 1694092640,
		}
		assert.Equal(t, expected, actual, "expected %+v, actual %+v", expected, actual)
	}
}
