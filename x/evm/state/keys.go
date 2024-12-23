package state

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common"
)

var (
	CodeKeyPrefix     = []byte("code")
	CodeSizeKeyPrefix = []byte("codesize")
	StateKeyPrefix    = []byte("state")
)

func codeKey(addr common.Address, codeHash []byte) []byte {
	return append(addr.Bytes(), append(CodeKeyPrefix, codeHash...)...)
}

func codeSizeKey(addr common.Address, codeHash []byte) []byte {
	return append(addr.Bytes(), append(CodeSizeKeyPrefix, codeHash...)...)
}

func stateKey(addr common.Address, slot common.Hash) []byte {
	return append(addr.Bytes(), append(StateKeyPrefix, slot.Bytes()...)...)
}

func stateKeyPrefix(addr common.Address) []byte {
	return append(addr.Bytes(), StateKeyPrefix...)
}

func uint64ToBytes(v uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, v)
	return bz
}

func bytesToUint64(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

// Keys for memStore
var (
	memStoreVMStorePrefix      = []byte{0x01} // prefix for memStore vm store
	memStoreCreatedPrefix      = []byte{0x02} // prefix for memStore created accounts
	memStoreSelfDestructPrefix = []byte{0x03} // prefix for memStore self destruct accounts
	memStoreLogsPrefix         = []byte{0x04} // prefix for memStore logs
	memStoreLogSizePrefix      = []byte{0x05} // prefix for memStore log size
	memStoreAccessListPrefix   = []byte{0x06} // prefix for memStore access list
	memStoreRefundPrefix       = []byte{0x07} // prefix for memStore refund
)
