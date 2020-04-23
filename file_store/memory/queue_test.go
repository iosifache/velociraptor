package memory

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/suite"
	"www.velocidex.com/golang/velociraptor/config"
	"www.velocidex.com/golang/velociraptor/file_store/api"
)

func TestMemoryQueueManager(t *testing.T) {
	dir, err := ioutil.TempDir("", "file_store_test")
	assert.NoError(t, err)

	defer os.RemoveAll(dir) // clean up

	//dir = "/tmp/file_store_test"

	config_obj := config.GetDefaultConfig()
	config_obj.Datastore.FilestoreDirectory = dir
	config_obj.Datastore.Location = dir

	manager := NewMemoryQueueManager(config_obj, Test_memory_file_store)
	suite.Run(t, api.NewQueueManagerTestSuite(config_obj, manager))
}
