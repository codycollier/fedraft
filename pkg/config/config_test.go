package acct

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
	// required config vars
	os.Setenv("FED_USER", "Foo")
	os.Setenv("FED_PASS", "SomePass")
}

func cleanup() {
	os.Unsetenv("FED_USER")
	os.Unsetenv("FED_PASS")
	unsetConfig()
}

// Test basic config creation and retrieval
func TestGetConfig(t *testing.T) {

	setup()

	// Ensure no error and non-empty vars
	conf, err := getConfig()
	assert.NotNil(t, conf)
	assert.Nil(t, err)
	assert.NotEqual(t, conf.User, "")
	assert.NotEqual(t, conf.Pass, "")
	assert.NotEqual(t, conf.Host, "")
	assert.NotEqual(t, conf.Port, "")
	assert.NotEqual(t, conf.confid, "")

	cleanup()

}

// Test that error is raised if required env vars not set
func TestGetConfigWithoutRequiredUser(t *testing.T) {
	setup()
	os.Unsetenv("FED_USER")
	conf, err := getConfig()
	assert.Nil(t, conf)
	assert.NotNil(t, err)
	cleanup()
}

// Test that error is raised if required env vars not set
func TestGetConfigWithoutRequiredPass(t *testing.T) {
	setup()
	os.Unsetenv("FED_PASS")
	conf, err := getConfig()
	assert.Nil(t, conf)
	assert.NotNil(t, err)
	cleanup()
}

// Test config isn't re-created
func TestConfigIsCreatedOnce(t *testing.T) {

	setup()

	// first
	conf1, err := getConfig()
	assert.NotNil(t, conf1)
	assert.Nil(t, err)

	// second
	conf2, err := getConfig()
	assert.NotNil(t, conf2)
	assert.Nil(t, err)

	assert.Equal(t, conf1.confid, conf2.confid)

	cleanup()
}

// Test config isn't re-created
func TestConfigIsRecreatedWhenUnset(t *testing.T) {

	setup()

	// first
	conf1, err := getConfig()
	assert.NotNil(t, conf1)
	assert.Nil(t, err)

	unsetConfig()

	// second
	conf2, err := getConfig()
	assert.NotNil(t, conf2)
	assert.Nil(t, err)

	assert.NotEqual(t, conf1.confid, conf2.confid)

	cleanup()
}
