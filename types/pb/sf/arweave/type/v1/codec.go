// Copyright 2019 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pbarweave

import (
	"encoding/hex"
	"time"

	"github.com/streamingfast/bstream"
)

const CONFIRMS uint64 = 20

// TODO: We should probably memoize all fields that requires computation
//       like Time() and likes.

func (b *Block) ID() string {
	return hex.EncodeToString(b.IndepHash)
}

func (b *Block) Num() uint64 {
	return b.Height
}

func (b *Block) PreviousID() string {
	var previousId string
	if b.Height != 0 {
		previousId = hex.EncodeToString(b.PreviousBlock)
	} else {
		var empty_hash [64]byte
		previousId = hex.EncodeToString(empty_hash[:])
	}
	return previousId
}

// ???
func (b *Block) PreviousNumber() uint64 {
	if b.PreviousID() == "" {
		return 0
	}
	return uint64(b.Height) - 1
}

func (b *Block) Time() time.Time {
	return time.Unix(int64(b.Timestamp), 0)
}

// firecore

func (b *Block) GetFirehoseBlockID() string {
	return b.ID()
}

func (b *Block) GetFirehoseBlockNumber() uint64 {
	return b.Num()
}

func (b *Block) GetFirehoseBlockParentID() string {
	return b.PreviousID()
}

func (b *Block) GetFirehoseBlockParentNumber() uint64 {
	return b.PreviousNumber()
}

func (b *Block) GetFirehoseBlockTime() time.Time {
	return b.Time()
}

func (b *Block) GetFirehoseBlockVersion() int32 {
	return int32(b.Ver)
}

func (b *Block) GetFirehoseBlockLIBNum() uint64 {
	return b.LIBNum()
}

func (b *Block) LIBNum() uint64 {
	var libNum uint64
	if b.Height > CONFIRMS {
		libNum = b.Height - CONFIRMS
	} else {
		libNum = 0
	}
	return libNum
}

func (b *Block) AsRef() bstream.BlockRef {
	return bstream.NewBlockRef(hex.EncodeToString(b.IndepHash), uint64(b.Height))
}
