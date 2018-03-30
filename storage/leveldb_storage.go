// Copyright 2018 The go-medibloc Authors
// This file is part of the go-medibloc library.
//
// The go-medibloc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-medibloc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-medibloc library. If not, see <http://www.gnu.org/licenses/>.

package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// LeveldbStorage storage which backend is leveldb
type LeveldbStorage struct {
	db *leveldb.DB
}

// NewLeveldbStorage init a LeveldbStorage
func NewLeveldbStorage(path string) (*LeveldbStorage, error) {
	// TODO path & parameters may be passed within Config struct
	db, err := leveldb.OpenFile(path, &opt.Options{
		BlockCacheCapacity:     8 * opt.MiB,
		Filter:                 filter.NewBloomFilter(10),
		OpenFilesCacheCapacity: 500,
		WriteBuffer:            4 * opt.MiB,
	})

	if err != nil {
		return nil, err
	}

	return &LeveldbStorage{
		db: db,
	}, nil
}

// Del delete the key entry in Storage.
func (storage *LeveldbStorage) Delete(key []byte) error {
	return storage.db.Delete(key, nil)
}

// Get return the value to the key in Storage.
func (storage *LeveldbStorage) Get(key []byte) ([]byte, error) {
	value, err := storage.db.Get(key, nil)
	if err != nil && err == leveldb.ErrNotFound {
		return nil, ErrKeyNotFound
	}
	return value, err
}

// Put put the key-value entry to Storage.
func (storage *LeveldbStorage) Put(key []byte, value []byte) error {
	return storage.db.Put(key, value, nil)
}
