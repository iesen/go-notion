package notion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient_UsesDefaultNotionApiBaseUrl(t *testing.T) {
	client := NewClient()
	assert.Equal(t, DefaultBaseUrl, client.config.BaseUrl)
}

func Test(t *testing.T) {

}
