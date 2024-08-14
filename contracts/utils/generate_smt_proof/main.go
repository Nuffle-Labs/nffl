package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/smt"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	poktsmt "github.com/pokt-network/smt"
)

func parseNode(data []byte) ([]byte, []byte) {
	if data == nil {
		return []byte{}, []byte{}
	}
	return data[1 : 32+1], data[1+32:]
}

type Hex32 [32]byte

func (h Hex32) String() string {
	return hex.EncodeToString(h[:])
}

func (h Hex32) MarshalJSON() ([]byte, error) {
	hexString := hex.EncodeToString(h[:])
	return json.Marshal(hexString)
}

type HexBytes []byte

func (h HexBytes) String() string {
	return hex.EncodeToString(h[:])
}

func (h HexBytes) MarshalJSON() ([]byte, error) {
	hexString := hex.EncodeToString(h[:])
	return json.Marshal(hexString)
}

func toHexBytesArray(arr [][]byte) []HexBytes {
	arr2 := make([]HexBytes, len(arr))
	for i, el := range arr {
		arr2[i] = el
	}
	return arr2
}

const (
	// Generates a proof that a key exists in the trie.
	MembershipProof = iota
	// Generates a proof that a key does not exist in the trie, and no leaf node was found while trying to follow the path.
	NonMembershipProofNoLeaf
	// Generates a proof that a key does not exist in the trie, and a leaf node was found while trying to follow the path.
	NonMembershipProofLeaf
)

const proofType = MembershipProof

func main() {
	msg := messages.StateRootUpdateMessage{
		RollupId:    uint32(20000),
		BlockHeight: uint64(20001),
		Timestamp:   uint64(20002),
	}

	var digest [32]byte
	var key [32]byte
	var count int

	if proofType == MembershipProof {
		digest, _ = msg.Digest()
		key = msg.Key()
		count = 1000
	} else if proofType == NonMembershipProofNoLeaf {
		digest = [32]byte{}
		key = core.Keccak256([]byte("non-existent-key"))
		count = 1000
	} else if proofType == NonMembershipProofLeaf {
		digest = [32]byte{}
		key = core.Keccak256([]byte("non-existent-key"))
		// adding more elements to make it likely to find a leaf node
		count = 10000
	}

	trie := smt.NewSMT()

	for i := 0; i < count; i++ {
		msg := messages.StateRootUpdateMessage{
			RollupId:    uint32(i),
			BlockHeight: uint64(i + 1),
			Timestamp:   uint64(i + 2),
		}

		err := trie.AddMessage(msg)
		if err != nil {
			panic(err)
		}
	}

	err := trie.AddMessage(msg)
	if err != nil {
		panic(err)
	}

	err = trie.Commit()
	if err != nil {
		panic(err)
	}

	zero := [32]byte{}
	value := []byte(nil)
	if !bytes.Equal(digest[:], zero[:]) {
		value = digest[:]
	}

	proof, err := trie.ProveCompact(key[:])
	if err != nil {
		panic(err)
	}

	success, err := poktsmt.VerifyCompactProof(proof, trie.Root(), key[:], value, trie.Spec())
	if err != nil {
		panic(err)
	}
	if !success {
		val, _ := trie.Get(key[:])
		fmt.Println("val", val)
		panic("Not successful")
	}

	type SolidityProof struct {
		Key                    Hex32
		Value                  Hex32
		BitMask                *big.Int
		SideNodes              []HexBytes
		NumSideNodes           uint
		NonMembershipLeafPath  HexBytes
		NonMembershipLeafValue HexBytes
	}

	type SolidityArgs struct {
		Root  HexBytes
		Proof SolidityProof
	}

	path, val := parseNode(proof.NonMembershipLeafData)

	verifierProof := smt.NewSMTVerifierProof(key, digest, proof)

	args, err := json.MarshalIndent(SolidityArgs{
		Root: HexBytes(trie.Root()),
		Proof: SolidityProof{
			Key:                    key,
			Value:                  digest,
			BitMask:                verifierProof.BitMask,
			SideNodes:              toHexBytesArray(proof.SideNodes),
			NumSideNodes:           uint(proof.NumSideNodes),
			NonMembershipLeafPath:  HexBytes(path),
			NonMembershipLeafValue: HexBytes(val),
		},
	}, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(args))
}
