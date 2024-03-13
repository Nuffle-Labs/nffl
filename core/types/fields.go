package types

// we only use a single quorum (quorum 0) for sffl
var QUORUM_NUMBERS = []byte{0}

type BlockNumber = uint32
type TaskIndex = uint32
type RollupId = uint32
type MessageDigest = [32]byte
