package constant

import "time"

var (
	CheckPointCycle = 5 * time.Second
)

const (
	Cancel = iota // must be zero
	Delete
)

const (
	MaxKeySize         = 4074
	MaxValueSize       = 1 << 16 // 64KB
	MaxDataFileSize    = 1 << 32 // 4GB - avoiding the use of triple-indirect blocks
	MaxTransactionSize = 1 << 26 // 64MB
	//MaxDataFileSize = 1 << 40 // 1TB
)

const (
	BlockSize = 4096 // 4k
)

const (
	PN = iota // prefix node
	SN        // suffix node
	MS        // mixed suffix node
	ES        // empty suffix node
)

const (
	TypeOff  = uint64(56)
	TypeMask = uint64(0xFF)
	Mask     = uint64(0xFFFFFFFFFFFFFF)
)