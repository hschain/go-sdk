package hsc

import "time"

type TransferInfo struct {
	TotalCount string `json:"total_count"`
	Count      string `json:"count"`
	PageNumber string `json:"page_number"`
	PageTotal  string `json:"page_total"`
	Limit      string `json:"limit"`
	Txs        []Txs  `json:"txs"`
}
type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Events struct {
	Type       string       `json:"type"`
	Attributes []Attributes `json:"attributes"`
}
type Logs struct {
	MsgIndex int      `json:"msg_index"`
	Success  bool     `json:"success"`
	Log      string   `json:"log"`
	Events   []Events `json:"events"`
}
type Amount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
type MsgValue struct {
	FromAddress string   `json:"from_address"`
	ToAddress   string   `json:"to_address"`
	Amount      []Amount `json:"amount"`
}
type Msg struct {
	Type  string   `json:"type"`
	Value MsgValue `json:"value"`
}

type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Signatures struct {
	PubKey    PubKey `json:"pub_key"`
	Signature string `json:"signature"`
}
type Value struct {
	Msg        []Msg        `json:"msg"`
	Fee        Fee          `json:"fee"`
	Signatures []Signatures `json:"signatures"`
	Memo       string       `json:"memo"`
}
type Tx struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}
type Txs struct {
	Height    string    `json:"height"`
	Txhash    string    `json:"txhash"`
	RawLog    string    `json:"raw_log"`
	Logs      []Logs    `json:"logs"`
	GasWanted string    `json:"gas_wanted"`
	GasUsed   string    `json:"gas_used"`
	Tx        Tx        `json:"tx"`
	Timestamp time.Time `json:"timestamp"`
	Events    []Events  `json:"events"`
	Error     *string   `json:"error"`
}

type Hsc struct {
	LcdEndpoint string `json:"lcd_endpoint"`
}
