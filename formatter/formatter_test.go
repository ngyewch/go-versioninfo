package formatter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessTag(t *testing.T) {
	config := Config{
		FallbackTag: "v0.0.0",
		TagPrefix:   "v",
	}

	{
		actual := config.ProcessTag("")
		expected := "0.0.0"
		assert.Equal(t, expected, actual, "expected %s, actual %s", expected, actual)
	}

	{
		actual := config.ProcessTag("v1.2.3")
		expected := "1.2.3"
		assert.Equal(t, expected, actual, "expected %s, actual %s", expected, actual)
	}

	{
		actual := config.ProcessTag("x1.2.3")
		expected := "x1.2.3"
		assert.Equal(t, expected, actual, "expected %s, actual %s", expected, actual)
	}
}
