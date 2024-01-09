// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: sf/arweave/type/v1/type.proto

package pbcodec

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BigInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *BigInt) Reset() {
	*x = BigInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_arweave_type_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigInt) ProtoMessage() {}

func (x *BigInt) ProtoReflect() protoreflect.Message {
	mi := &file_sf_arweave_type_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigInt.ProtoReflect.Descriptor instead.
func (*BigInt) Descriptor() ([]byte, []int) {
	return file_sf_arweave_type_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *BigInt) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Firehose block version (unrelated to Arweave block version)
	Ver uint32 `protobuf:"varint,1,opt,name=ver,proto3" json:"ver,omitempty"`
	// The block identifier
	IndepHash []byte `protobuf:"bytes,2,opt,name=indep_hash,json=indepHash,proto3" json:"indep_hash,omitempty"`
	// The nonce chosen to solve the mining problem
	Nonce []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// `indep_hash` of the previous block in the weave
	PreviousBlock []byte `protobuf:"bytes,4,opt,name=previous_block,json=previousBlock,proto3" json:"previous_block,omitempty"`
	// POSIX time of block discovery
	Timestamp uint64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// POSIX time of the last difficulty retarget
	LastRetarget uint64 `protobuf:"varint,6,opt,name=last_retarget,json=lastRetarget,proto3" json:"last_retarget,omitempty"`
	// Mining difficulty; the number `hash` must be greater than.
	Diff *BigInt `protobuf:"bytes,7,opt,name=diff,proto3" json:"diff,omitempty"`
	// How many blocks have passed since the genesis block
	Height uint64 `protobuf:"varint,8,opt,name=height,proto3" json:"height,omitempty"`
	// Mining solution hash of the block; must satisfy the mining difficulty
	Hash []byte `protobuf:"bytes,9,opt,name=hash,proto3" json:"hash,omitempty"`
	// Merkle root of the tree of Merkle roots of block's transactions' data.
	TxRoot []byte `protobuf:"bytes,10,opt,name=tx_root,json=txRoot,proto3" json:"tx_root,omitempty"`
	// Transactions contained within this block
	Txs []*Transaction `protobuf:"bytes,11,rep,name=txs,proto3" json:"txs,omitempty"`
	// The root hash of the Merkle Patricia Tree containing
	// all wallet (account) balances and the identifiers
	// of the last transactions posted by them; if any.
	WalletList []byte `protobuf:"bytes,12,opt,name=wallet_list,json=walletList,proto3" json:"wallet_list,omitempty"`
	// (string or) Address of the account to receive the block rewards. Can also be unclaimed which is encoded as a null byte
	RewardAddr []byte `protobuf:"bytes,13,opt,name=reward_addr,json=rewardAddr,proto3" json:"reward_addr,omitempty"`
	// Tags that a block producer can add to a block
	Tags []*Tag `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty"`
	// Size of reward pool
	RewardPool *BigInt `protobuf:"bytes,15,opt,name=reward_pool,json=rewardPool,proto3" json:"reward_pool,omitempty"`
	// Size of the weave in bytes
	WeaveSize *BigInt `protobuf:"bytes,16,opt,name=weave_size,json=weaveSize,proto3" json:"weave_size,omitempty"`
	// Size of this block in bytes
	BlockSize *BigInt `protobuf:"bytes,17,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	// Required after the version 1.8 fork. Zero otherwise.
	// The sum of the average number of hashes computed
	// by the network to produce the past blocks including this one.
	CumulativeDiff *BigInt `protobuf:"bytes,18,opt,name=cumulative_diff,json=cumulativeDiff,proto3" json:"cumulative_diff,omitempty"`
	// Required after the version 1.8 fork. Null byte otherwise.
	// The Merkle root of the block index - the list of {`indep_hash`; `weave_size`; `tx_root`} triplets
	HashListMerkle []byte `protobuf:"bytes,20,opt,name=hash_list_merkle,json=hashListMerkle,proto3" json:"hash_list_merkle,omitempty"`
	// The proof of access; Used after v2.4 only; set as defaults otherwise
	Poa *ProofOfAccess `protobuf:"bytes,21,opt,name=poa,proto3" json:"poa,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_arweave_type_v1_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_sf_arweave_type_v1_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_sf_arweave_type_v1_type_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetVer() uint32 {
	if x != nil {
		return x.Ver
	}
	return 0
}

func (x *Block) GetIndepHash() []byte {
	if x != nil {
		return x.IndepHash
	}
	return nil
}

func (x *Block) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *Block) GetPreviousBlock() []byte {
	if x != nil {
		return x.PreviousBlock
	}
	return nil
}

func (x *Block) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Block) GetLastRetarget() uint64 {
	if x != nil {
		return x.LastRetarget
	}
	return 0
}

func (x *Block) GetDiff() *BigInt {
	if x != nil {
		return x.Diff
	}
	return nil
}

func (x *Block) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Block) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Block) GetTxRoot() []byte {
	if x != nil {
		return x.TxRoot
	}
	return nil
}

func (x *Block) GetTxs() []*Transaction {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *Block) GetWalletList() []byte {
	if x != nil {
		return x.WalletList
	}
	return nil
}

func (x *Block) GetRewardAddr() []byte {
	if x != nil {
		return x.RewardAddr
	}
	return nil
}

func (x *Block) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Block) GetRewardPool() *BigInt {
	if x != nil {
		return x.RewardPool
	}
	return nil
}

func (x *Block) GetWeaveSize() *BigInt {
	if x != nil {
		return x.WeaveSize
	}
	return nil
}

func (x *Block) GetBlockSize() *BigInt {
	if x != nil {
		return x.BlockSize
	}
	return nil
}

func (x *Block) GetCumulativeDiff() *BigInt {
	if x != nil {
		return x.CumulativeDiff
	}
	return nil
}

func (x *Block) GetHashListMerkle() []byte {
	if x != nil {
		return x.HashListMerkle
	}
	return nil
}

func (x *Block) GetPoa() *ProofOfAccess {
	if x != nil {
		return x.Poa
	}
	return nil
}

// A succinct proof of access to a recall byte found in a TX
type ProofOfAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The recall byte option chosen; global offset of index byte
	Option string `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
	// The path through the Merkle tree of transactions' `data_root`s;
	// from the `data_root` being proven to the corresponding `tx_root`
	TxPath []byte `protobuf:"bytes,2,opt,name=tx_path,json=txPath,proto3" json:"tx_path,omitempty"`
	// The path through the Merkle tree of identifiers of chunks of the
	// corresponding transaction; from the chunk being proven to the
	// corresponding `data_root`.
	DataPath []byte `protobuf:"bytes,3,opt,name=data_path,json=dataPath,proto3" json:"data_path,omitempty"`
	// The data chunk.
	Chunk []byte `protobuf:"bytes,4,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *ProofOfAccess) Reset() {
	*x = ProofOfAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_arweave_type_v1_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofOfAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofOfAccess) ProtoMessage() {}

func (x *ProofOfAccess) ProtoReflect() protoreflect.Message {
	mi := &file_sf_arweave_type_v1_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofOfAccess.ProtoReflect.Descriptor instead.
func (*ProofOfAccess) Descriptor() ([]byte, []int) {
	return file_sf_arweave_type_v1_type_proto_rawDescGZIP(), []int{2}
}

func (x *ProofOfAccess) GetOption() string {
	if x != nil {
		return x.Option
	}
	return ""
}

func (x *ProofOfAccess) GetTxPath() []byte {
	if x != nil {
		return x.TxPath
	}
	return nil
}

func (x *ProofOfAccess) GetDataPath() []byte {
	if x != nil {
		return x.DataPath
	}
	return nil
}

func (x *ProofOfAccess) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1 or 2 for v1 or v2 transactions. More allowable in the future
	Format uint32 `protobuf:"varint,1,opt,name=format,proto3" json:"format,omitempty"`
	// The transaction identifier.
	Id []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Either the identifier of the previous transaction from the same
	// wallet or the identifier of one of the last ?MAX_TX_ANCHOR_DEPTH blocks.
	LastTx []byte `protobuf:"bytes,3,opt,name=last_tx,json=lastTx,proto3" json:"last_tx,omitempty"`
	// The public key the transaction is signed with.
	Owner []byte `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	// A list of arbitrary key-value pairs
	Tags []*Tag `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// The address of the recipient; if any. The SHA2-256 hash of the public key.
	Target []byte `protobuf:"bytes,6,opt,name=target,proto3" json:"target,omitempty"`
	// The amount of Winstons to send to the recipient; if any.
	Quantity *BigInt `protobuf:"bytes,7,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// The data to upload; if any. For v2 transactions; the field is optional
	//   - a fee is charged based on the `data_size` field;
	//     data may be uploaded any time later in chunks.
	Data []byte `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	// Size in bytes of the transaction data.
	DataSize *BigInt `protobuf:"bytes,9,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
	// The Merkle root of the Merkle tree of data chunks.
	DataRoot []byte `protobuf:"bytes,10,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	// The signature.
	Signature []byte `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
	// The fee in Winstons.
	Reward *BigInt `protobuf:"bytes,12,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_arweave_type_v1_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_sf_arweave_type_v1_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_sf_arweave_type_v1_type_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetFormat() uint32 {
	if x != nil {
		return x.Format
	}
	return 0
}

func (x *Transaction) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Transaction) GetLastTx() []byte {
	if x != nil {
		return x.LastTx
	}
	return nil
}

func (x *Transaction) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Transaction) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Transaction) GetTarget() []byte {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Transaction) GetQuantity() *BigInt {
	if x != nil {
		return x.Quantity
	}
	return nil
}

func (x *Transaction) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Transaction) GetDataSize() *BigInt {
	if x != nil {
		return x.DataSize
	}
	return nil
}

func (x *Transaction) GetDataRoot() []byte {
	if x != nil {
		return x.DataRoot
	}
	return nil
}

func (x *Transaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Transaction) GetReward() *BigInt {
	if x != nil {
		return x.Reward
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_arweave_type_v1_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_sf_arweave_type_v1_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_sf_arweave_type_v1_type_proto_rawDescGZIP(), []int{4}
}

func (x *Tag) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Tag) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_sf_arweave_type_v1_type_proto protoreflect.FileDescriptor

var file_sf_arweave_type_v1_type_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x66, 0x2f, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x1e, 0x0a, 0x06, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x22, 0xa6, 0x06, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x0a,
	0x03, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x70, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x70, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2e,
	0x0a, 0x04, 0x64, 0x69, 0x66, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x66, 0x2e, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x69, 0x66, 0x66, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x31, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x03, 0x74, 0x78, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x41, 0x64, 0x64, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65,
	0x61, 0x76, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x66, 0x2e,
	0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x77, 0x65, 0x61, 0x76, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65,
	0x61, 0x76, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x49,
	0x6e, 0x74, 0x52, 0x09, 0x77, 0x65, 0x61, 0x76, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x63, 0x75, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x0e, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x44, 0x69, 0x66, 0x66, 0x12, 0x28, 0x0a,
	0x10, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x03, 0x70, 0x6f, 0x61, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4f,
	0x66, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x03, 0x70, 0x6f, 0x61, 0x22, 0x73, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4f, 0x66, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x22, 0x9d, 0x03, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x74, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74,
	0x54, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65,
	0x61, 0x76, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x36, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x66, 0x2e, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x32, 0x0a,
	0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x66, 0x2e, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x22, 0x2f, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x61, 0x66, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x68,
	0x6f, 0x73, 0x65, 0x2d, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x73,
	0x66, 0x2f, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x62, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sf_arweave_type_v1_type_proto_rawDescOnce sync.Once
	file_sf_arweave_type_v1_type_proto_rawDescData = file_sf_arweave_type_v1_type_proto_rawDesc
)

func file_sf_arweave_type_v1_type_proto_rawDescGZIP() []byte {
	file_sf_arweave_type_v1_type_proto_rawDescOnce.Do(func() {
		file_sf_arweave_type_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_arweave_type_v1_type_proto_rawDescData)
	})
	return file_sf_arweave_type_v1_type_proto_rawDescData
}

var file_sf_arweave_type_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sf_arweave_type_v1_type_proto_goTypes = []interface{}{
	(*BigInt)(nil),        // 0: sf.arweave.type.v1.BigInt
	(*Block)(nil),         // 1: sf.arweave.type.v1.Block
	(*ProofOfAccess)(nil), // 2: sf.arweave.type.v1.ProofOfAccess
	(*Transaction)(nil),   // 3: sf.arweave.type.v1.Transaction
	(*Tag)(nil),           // 4: sf.arweave.type.v1.Tag
}
var file_sf_arweave_type_v1_type_proto_depIdxs = []int32{
	0,  // 0: sf.arweave.type.v1.Block.diff:type_name -> sf.arweave.type.v1.BigInt
	3,  // 1: sf.arweave.type.v1.Block.txs:type_name -> sf.arweave.type.v1.Transaction
	4,  // 2: sf.arweave.type.v1.Block.tags:type_name -> sf.arweave.type.v1.Tag
	0,  // 3: sf.arweave.type.v1.Block.reward_pool:type_name -> sf.arweave.type.v1.BigInt
	0,  // 4: sf.arweave.type.v1.Block.weave_size:type_name -> sf.arweave.type.v1.BigInt
	0,  // 5: sf.arweave.type.v1.Block.block_size:type_name -> sf.arweave.type.v1.BigInt
	0,  // 6: sf.arweave.type.v1.Block.cumulative_diff:type_name -> sf.arweave.type.v1.BigInt
	2,  // 7: sf.arweave.type.v1.Block.poa:type_name -> sf.arweave.type.v1.ProofOfAccess
	4,  // 8: sf.arweave.type.v1.Transaction.tags:type_name -> sf.arweave.type.v1.Tag
	0,  // 9: sf.arweave.type.v1.Transaction.quantity:type_name -> sf.arweave.type.v1.BigInt
	0,  // 10: sf.arweave.type.v1.Transaction.data_size:type_name -> sf.arweave.type.v1.BigInt
	0,  // 11: sf.arweave.type.v1.Transaction.reward:type_name -> sf.arweave.type.v1.BigInt
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_sf_arweave_type_v1_type_proto_init() }
func file_sf_arweave_type_v1_type_proto_init() {
	if File_sf_arweave_type_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_arweave_type_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigInt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_arweave_type_v1_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_arweave_type_v1_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofOfAccess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_arweave_type_v1_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_arweave_type_v1_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sf_arweave_type_v1_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_arweave_type_v1_type_proto_goTypes,
		DependencyIndexes: file_sf_arweave_type_v1_type_proto_depIdxs,
		MessageInfos:      file_sf_arweave_type_v1_type_proto_msgTypes,
	}.Build()
	File_sf_arweave_type_v1_type_proto = out.File
	file_sf_arweave_type_v1_type_proto_rawDesc = nil
	file_sf_arweave_type_v1_type_proto_goTypes = nil
	file_sf_arweave_type_v1_type_proto_depIdxs = nil
}
