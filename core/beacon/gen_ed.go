// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package beacon

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*executableDataMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (e ExecutableDataV1) MarshalJSON() ([]byte, error) {
	type ExecutableDataV1 struct {
		ParentHash    common.Hash     `json:"parentHash"    gencodec:"required"`
		FeeRecipient  common.Address  `json:"feeRecipient"  gencodec:"required"`
		StateRoot     common.Hash     `json:"stateRoot"     gencodec:"required"`
		ReceiptsRoot  common.Hash     `json:"receiptsRoot"  gencodec:"required"`
		LogsBloom     hexutil.Bytes   `json:"logsBloom"     gencodec:"required"`
		Random        common.Hash     `json:"prevRandao"    gencodec:"required"`
		Number        hexutil.Uint64  `json:"blockNumber"   gencodec:"required"`
		GasLimit      hexutil.Uint64  `json:"gasLimit"      gencodec:"required"`
		GasUsed       hexutil.Uint64  `json:"gasUsed"       gencodec:"required"`
		Timestamp     hexutil.Uint64  `json:"timestamp"     gencodec:"required"`
		ExtraData     hexutil.Bytes   `json:"extraData"     gencodec:"required"`
		BaseFeePerGas *hexutil.Big    `json:"baseFeePerGas" gencodec:"required"`
		BlockHash     common.Hash     `json:"blockHash"     gencodec:"required"`
		Transactions  []hexutil.Bytes `json:"transactions"  gencodec:"required"`
		//Codeword      hexutil.Bytes   `json:"codeword" rlp:"optional"`
	}
	var enc ExecutableDataV1
	enc.ParentHash = e.ParentHash
	enc.FeeRecipient = e.FeeRecipient
	enc.StateRoot = e.StateRoot
	enc.ReceiptsRoot = e.ReceiptsRoot
	enc.LogsBloom = e.LogsBloom
	enc.Random = e.Random
	enc.Number = hexutil.Uint64(e.Number)
	enc.GasLimit = hexutil.Uint64(e.GasLimit)
	enc.GasUsed = hexutil.Uint64(e.GasUsed)
	enc.Timestamp = hexutil.Uint64(e.Timestamp)
	enc.ExtraData = e.ExtraData
	enc.BaseFeePerGas = (*hexutil.Big)(e.BaseFeePerGas)
	//enc.Codeword = e.Codeword
	enc.BlockHash = e.BlockHash
	if e.Transactions != nil {
		enc.Transactions = make([]hexutil.Bytes, len(e.Transactions))
		for k, v := range e.Transactions {
			enc.Transactions[k] = v
		}
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (e *ExecutableDataV1) UnmarshalJSON(input []byte) error {
	type ExecutableDataV1 struct {
		ParentHash    *common.Hash    `json:"parentHash"    gencodec:"required"`
		FeeRecipient  *common.Address `json:"feeRecipient"  gencodec:"required"`
		StateRoot     *common.Hash    `json:"stateRoot"     gencodec:"required"`
		ReceiptsRoot  *common.Hash    `json:"receiptsRoot"  gencodec:"required"`
		LogsBloom     *hexutil.Bytes  `json:"logsBloom"     gencodec:"required"`
		Random        *common.Hash    `json:"prevRandao"    gencodec:"required"`
		Number        *hexutil.Uint64 `json:"blockNumber"   gencodec:"required"`
		GasLimit      *hexutil.Uint64 `json:"gasLimit"      gencodec:"required"`
		GasUsed       *hexutil.Uint64 `json:"gasUsed"       gencodec:"required"`
		Timestamp     *hexutil.Uint64 `json:"timestamp"     gencodec:"required"`
		ExtraData     *hexutil.Bytes  `json:"extraData"     gencodec:"required"`
		BaseFeePerGas *hexutil.Big    `json:"baseFeePerGas" gencodec:"required"`
		BlockHash     *common.Hash    `json:"blockHash"     gencodec:"required"`
		Transactions  []hexutil.Bytes `json:"transactions"  gencodec:"required"`
		//Codeword      *hexutil.Bytes  `json:"codeword" rlp:"optional"`
	}
	var dec ExecutableDataV1
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for ExecutableDataV1")
	}
	e.ParentHash = *dec.ParentHash
	if dec.FeeRecipient == nil {
		return errors.New("missing required field 'feeRecipient' for ExecutableDataV1")
	}
	e.FeeRecipient = *dec.FeeRecipient
	if dec.StateRoot == nil {
		return errors.New("missing required field 'stateRoot' for ExecutableDataV1")
	}
	e.StateRoot = *dec.StateRoot
	if dec.ReceiptsRoot == nil {
		return errors.New("missing required field 'receiptsRoot' for ExecutableDataV1")
	}
	e.ReceiptsRoot = *dec.ReceiptsRoot
	if dec.LogsBloom == nil {
		return errors.New("missing required field 'logsBloom' for ExecutableDataV1")
	}
	e.LogsBloom = *dec.LogsBloom
	if dec.Random == nil {
		return errors.New("missing required field 'prevRandao' for ExecutableDataV1")
	}
	e.Random = *dec.Random
	if dec.Number == nil {
		return errors.New("missing required field 'blockNumber' for ExecutableDataV1")
	}
	e.Number = uint64(*dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for ExecutableDataV1")
	}
	e.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for ExecutableDataV1")
	}
	e.GasUsed = uint64(*dec.GasUsed)
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for ExecutableDataV1")
	}
	e.Timestamp = uint64(*dec.Timestamp)
	if dec.ExtraData == nil {
		return errors.New("missing required field 'extraData' for ExecutableDataV1")
	}
	e.ExtraData = *dec.ExtraData
	if dec.BaseFeePerGas == nil {
		return errors.New("missing required field 'baseFeePerGas' for ExecutableDataV1")
	}
	e.BaseFeePerGas = (*big.Int)(dec.BaseFeePerGas)
	if dec.BlockHash == nil {
		return errors.New("missing required field 'blockHash' for ExecutableDataV1")
	}
	e.BlockHash = *dec.BlockHash
	/*codeword
	if dec.Codeword == nil {
		return errors.New("missing required field 'Codeword' for ExecutableDataV1")
	}
	e.Codeword = *dec.Codeword */
	if dec.Transactions == nil {
		return errors.New("missing required field 'transactions' for ExecutableDataV1")
	}
	e.Transactions = make([][]byte, len(dec.Transactions))
	for k, v := range dec.Transactions {
		e.Transactions[k] = v
	}
	return nil
}
