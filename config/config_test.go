package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstans(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", apiGithubAccessToken)
}

func TestGetGithubAccessToken(t *testing.T) {
	assert.EqualValues(t, "", GetGithubAccessToken())
}
