// Copyright 2014 The Metabase Authors
// This file is part of vasuki p2p library.
//
// The vasuki library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The vasuki library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package stor

import (
	"github.com/syndtr/goleveldb/leveldb"
)

// S  struct for S
type S struct {
	StorPath string
	Stordb   *leveldb.DB
}

// start function to initilize nodep
func start(path string) (S, error) {
	var ret S
	db, err := leveldb.OpenFile(path, nil)
	if err == nil {
		ret = S{StorPath: path, Stordb: db}
	} else {
		ret = S{}
	}
	return ret, err
}
