package hsc

import (
	"time"

	"github.com/hschain/hschain/types"
)

// TxBody represents the body of a Cosmos transaction
// signed and ready to be sent over the LCD REST service.
type TxBody struct {
	Tx   SignedTransactionPayload `json:"tx"`
	Mode string                   `json:"mode"`
}

// AccountData holds informations about the account number and
// sequence number of a Cosmos account.
type AccountData struct {
	Result AccountDataResult `json:"result"`
}

// AccountDataResult is a wrapper struct for a call to auth/accounts/{address} LCD
// REST endpoint.
type AccountDataResult struct {
	Value AccountDataValue `json:"value"`
}

// AccountDataValue represents the real data obtained by calling /auth/accounts/{address} LCD
// REST endpoint.
type AccountDataValue struct {
	Address       string `json:"address"`
	AccountNumber string `json:"account_number"`
	Sequence      string `json:"sequence"`
}

// NodeInfo is the LCD REST response to a /node_info request,
// and contains the Network attribute (chain ID).
type NodeInfo struct {
	Info struct {
		Network string `json:"network"`
	} `json:"node_info"`
}

// TxResponse represents whatever data the LCD REST service returns to atomicwallet
// after a transaction gets forwarded to it.
type TxResponse struct {
	Height    string                `json:"height"`
	TxHash    string                `json:"txhash"`
	Code      uint32                `json:"code,omitempty"`
	Data      string                `json:"data,omitempty"`
	RawLog    string                `json:"raw_log,omitempty"`
	Logs      types.ABCIMessageLogs `json:"logs,omitempty"`
	Info      string                `json:"info,omitempty"`
	GasWanted string                `json:"gas_wanted,omitempty"`
	GasUsed   string                `json:"gas_used,omitempty"`
	Codespace string                `json:"codespace,omitempty"`
	Tx        types.Tx              `json:"tx,omitempty"`
	Timestamp string                `json:"timestamp,omitempty"`

	// DEPRECATED: Remove in the next next major release in favor of using the
	// ABCIMessageLog.Events field.
	Events types.StringEvents `json:"events,omitempty"`
}

// Error represents a JSON encoded error message sent whenever something
// goes wrong during the handler processing.
type Error struct {
	Error string `json:"error,omitempty"`
}

//BlockLatest
type BlockInfo struct {
	BlockMeta BlockMeta `json:"block_meta"`
	Block     Block     `json:"block"`
}
type Parts struct {
	Total string `json:"total"`
	Hash  string `json:"hash"`
}
type BlockID struct {
	Hash  string `json:"hash"`
	Parts Parts  `json:"parts"`
}
type Version struct {
	Block string `json:"block"`
	App   string `json:"app"`
}
type LastBlockID struct {
	Hash  string `json:"hash"`
	Parts Parts  `json:"parts"`
}
type Header struct {
	Version            Version     `json:"version"`
	ChainID            string      `json:"chain_id"`
	Height             string      `json:"height"`
	Time               time.Time   `json:"time"`
	NumTxs             string      `json:"num_txs"`
	TotalTxs           string      `json:"total_txs"`
	LastBlockID        LastBlockID `json:"last_block_id"`
	LastCommitHash     string      `json:"last_commit_hash"`
	DataHash           string      `json:"data_hash"`
	ValidatorsHash     string      `json:"validators_hash"`
	NextValidatorsHash string      `json:"next_validators_hash"`
	ConsensusHash      string      `json:"consensus_hash"`
	AppHash            string      `json:"app_hash"`
	LastResultsHash    string      `json:"last_results_hash"`
	EvidenceHash       string      `json:"evidence_hash"`
	ProposerAddress    string      `json:"proposer_address"`
}
type BlockMeta struct {
	BlockID BlockID `json:"block_id"`
	Header  Header  `json:"header"`
}
type Data struct {
	Txs []string `json:"txs"`
}
type Evidence struct {
	Evidence interface{} `json:"evidence"`
}
type Precommits struct {
	Type             int       `json:"type"`
	Height           string    `json:"height"`
	Round            string    `json:"round"`
	BlockID          BlockID   `json:"block_id"`
	Timestamp        time.Time `json:"timestamp"`
	ValidatorAddress string    `json:"validator_address"`
	ValidatorIndex   string    `json:"validator_index"`
	Signature        string    `json:"signature"`
}
type LastCommit struct {
	BlockID    BlockID      `json:"block_id"`
	Precommits []Precommits `json:"precommits"`
}
type Block struct {
	Header     Header     `json:"header"`
	Data       Data       `json:"data"`
	Evidence   Evidence   `json:"evidence"`
	LastCommit LastCommit `json:"last_commit"`
}
