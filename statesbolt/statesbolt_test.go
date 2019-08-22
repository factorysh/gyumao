package statesbolt

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"testing"

	"github.com/coreos/bbolt"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	str := "string for test"
	test_cache := Cache{str}

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	encoder.Encode(test_cache)
	expected := buffer.Bytes()

	result, err := encode(test_cache)
	assert.NoError(t, err)
	assert.Equal(t, expected, result, "The []byte should be the same")
}

func TestDecode(t *testing.T) {
	str := "string for test"
	test_cache := &Cache{str}

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	encoder.Encode(test_cache)
	encoded := buffer.Bytes()

	result, err := decode(encoded)
	assert.NoError(t, err)
	assert.Equal(t, test_cache, result, "The decoded string should be the same")
}

func TestNewCache(t *testing.T) {
	c, err := New("test_statesbolt.db")
	assert.NoError(t, err)
	assert.NotNil(t, *c)
	assert.NotNil(t, *c.db)
	err = c.Set("testnew", "test_string")
	assert.NoError(t, err)
	c.Close()
	err = c.Set("testnew", "test_string")
	assert.Error(t, err)
	// Remove the DB file
	os.Remove("test_statesbolt.db")
}

// Test for the Get function
func TestGetCache(t *testing.T) {
	// Setup the DB
	c, err := New("test_statesbolt.db")
	assert.NoError(t, err)
	defer c.Close()
	// Create an artificial cache in the DB
	err = c.db.Update(func(tx *bbolt.Tx) error {
		cache := Cache{
			Content: "test_get_string",
		}
		encoded, err := encode(cache)
		if err != nil {
			return err
		}
		err = tx.Bucket([]byte(BUCKET)).Put([]byte("testget"), encoded)
		if err != nil {
			return fmt.Errorf("can't set cache: %v", err)
		}
		return nil
	})
	assert.NoError(t, err)
	// get the cache with normal expiration time -> should be good
	cached, err := c.Get("testget")
	assert.NoError(t, err)
	assert.Equal(t, "test_get_string", cached, "The two interfaces should be the same")
	// Remove the DB file
	os.Remove("test_statesbolt.db")
}

// Test for the Set function
func TestSetCache(t *testing.T) {
	// Setup the DB
	c, err := New("test_statesbolt.db")
	assert.NoError(t, err)
	defer c.Close()

	err = c.Set("testset", "test_set_string")
	assert.NoError(t, err)
	// Get the cache with the normal (previously tested) Get
	cached, err := c.Get("testset")
	assert.NoError(t, err)
	assert.Equal(t, "test_set_string", cached, "The two interfaces should be the same")
	// Remove the DB file
	os.Remove("test_statesbolt.db")
}
