package statesbolt

// https://github.com/etcd-io/bbolt
// Implement states.States interface

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/coreos/bbolt"
	"github.com/factorysh/gyumao/states"
)

type Cache struct {
	Content interface{}
}

type Cachedb struct {
	db *bbolt.DB
}

const (
	DBNAME = "DB"
)

var (
	BUCKET = []byte(DBNAME)
)

func New(path string) (*Cachedb, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("can't open db, %v", err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(BUCKET)
		if err != nil {
			return fmt.Errorf("can't create DB bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("set up bucket error, %v", err)
	}
	return &Cachedb{
		db: db,
	}, nil
}

func (c *Cachedb) Close() {
	c.db.Close()
}

func encode(data Cache) ([]byte, error) {
	buffer := new(bytes.Buffer)
	gob.Register(states.States{})
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, fmt.Errorf("can't encode with gob: %v", err)
	}
	return buffer.Bytes(), nil
}

func decode(data []byte) (*Cache, error) {
	var cache Cache
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&cache)
	if err != nil {
		return nil, fmt.Errorf("can't decode with gob: %v", err)
	}
	return &cache, nil
}

func (c *Cachedb) Set(key string, value interface{}) error {
	cache := Cache{
		Content: value,
	}
	encoded, err := encode(cache)
	if err != nil {
		return err
	}
	err = c.db.Update(func(tx *bbolt.Tx) error {
		err = tx.Bucket([]byte("DB")).Put([]byte(key), encoded)
		if err != nil {
			return fmt.Errorf("can't set cache: %v", err)
		}
		return nil
	})
	return err
}

// Get key return value, is expired, error
func (c *Cachedb) Get(key string) (interface{}, error) {
	var encoded []byte
	err := c.db.View(func(tx *bbolt.Tx) error {
		bk := tx.Bucket(BUCKET)
		if bk == nil {
			return fmt.Errorf("failed to get bucket DB")
		}
		encoded = bk.Get([]byte(key))
		return nil
	})
	if err != nil {
		return nil, err
	}
	if encoded == nil {
		return nil, nil
	}
	cache, err := decode(encoded)
	if err != nil {
		return nil, err
	}
	return cache.Content, nil
}
