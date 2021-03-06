// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package plasmachain

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*headerMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (h Header) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash      common.Hash    `json:"parentHash"      gencodec:"required"`
		BlockNumber     hexutil.Uint64 `json:"blockNumber"     gencodec:"required`
		Time            hexutil.Uint64 `json:"time"            gencodec:"required`
		BloomID         common.Hash    `json:"bloomID"         gencodec:"required`
		TransactionRoot common.Hash    `json:"transactionRoot" gencodec:"required`
		TokenRoot       common.Hash    `json:"tokenRoot"       gencodec:"required`
		AccountRoot     common.Hash    `json:"accountRoot"     gencodec:"required`
		L3ChainRoot     common.Hash    `json:"l3ChainRoot"     gencodec:"required`
		AnchorRoot      common.Hash    `json:"anchorRoot"      gencodec:"required`
		Minter          common.Address `json:"minter"          gencodec:"required`
		Sig             hexutil.Bytes  `json:"sig"             gencodec:"required`
		HeaderHash      common.Hash    `json:"headerHash"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.BlockNumber = hexutil.Uint64(h.BlockNumber)
	enc.Time = hexutil.Uint64(h.Time)
	enc.BloomID = h.BloomID
	enc.TransactionRoot = h.TransactionRoot
	enc.TokenRoot = h.TokenRoot
	enc.AccountRoot = h.AccountRoot
	enc.L3ChainRoot = h.L3ChainRoot
	enc.AnchorRoot = h.AnchorRoot
	enc.Minter = h.Minter
	enc.Sig = h.Sig
	enc.HeaderHash = h.HeaderHash()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (h *Header) UnmarshalJSON(input []byte) error {
	type Header struct {
		ParentHash      *common.Hash    `json:"parentHash"      gencodec:"required"`
		BlockNumber     *hexutil.Uint64 `json:"blockNumber"     gencodec:"required`
		Time            *hexutil.Uint64 `json:"time"            gencodec:"required`
		BloomID         *common.Hash    `json:"bloomID"         gencodec:"required`
		TransactionRoot *common.Hash    `json:"transactionRoot" gencodec:"required`
		TokenRoot       *common.Hash    `json:"tokenRoot"       gencodec:"required`
		AccountRoot     *common.Hash    `json:"accountRoot"     gencodec:"required`
		L3ChainRoot     *common.Hash    `json:"l3ChainRoot"     gencodec:"required`
		AnchorRoot      *common.Hash    `json:"anchorRoot"      gencodec:"required`
		Minter          *common.Address `json:"minter"          gencodec:"required`
		Sig             *hexutil.Bytes  `json:"sig"             gencodec:"required`
	}
	var dec Header
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for Header")
	}
	h.ParentHash = *dec.ParentHash
	if dec.BlockNumber != nil {
		h.BlockNumber = uint64(*dec.BlockNumber)
	}
	if dec.Time != nil {
		h.Time = uint64(*dec.Time)
	}
	if dec.BloomID != nil {
		h.BloomID = *dec.BloomID
	}
	if dec.TransactionRoot != nil {
		h.TransactionRoot = *dec.TransactionRoot
	}
	if dec.TokenRoot != nil {
		h.TokenRoot = *dec.TokenRoot
	}
	if dec.AccountRoot != nil {
		h.AccountRoot = *dec.AccountRoot
	}
	if dec.L3ChainRoot != nil {
		h.L3ChainRoot = *dec.L3ChainRoot
	}
	if dec.AnchorRoot != nil {
		h.AnchorRoot = *dec.AnchorRoot
	}
	if dec.Minter != nil {
		h.Minter = *dec.Minter
	}
	if dec.Sig != nil {
		h.Sig = *dec.Sig
	}
	return nil
}
