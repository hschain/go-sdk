package hsc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func NewHsc(lcdEndpoint string) *Hsc {
	return &Hsc{
		LcdEndpoint: lcdEndpoint,
	}
}

// Retrieve the account data related to the given wallet address, like
// account number and sequence number.
func (h *Hsc) GetAccountData(address string) (AccountData, error) {
	endpoint := fmt.Sprintf("%s/auth/accounts/%s", h.LcdEndpoint, address)

	resp, err := http.Get(endpoint)
	if err != nil {
		return AccountData{}, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AccountData{}, fmt.Errorf("read body error: %w", err)
	}

	var accountData AccountData
	err = json.Unmarshal(data, &accountData)
	if err != nil {
		return AccountData{}, fmt.Errorf("could not unmarshal node response: %w", err)
	}

	if accountData.Result.Value.Address == "" {
		return AccountData{}, fmt.Errorf("account with address %s is not online", address)
	}
	return accountData, nil
}

// Return useful information of the full node, like the Network
// (chain) name.
func (h *Hsc) getNodeInfo() (NodeInfo, error) {
	endpoint := fmt.Sprintf("%s/node_info", h.LcdEndpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return NodeInfo{}, err
	}

	var nodeInfo NodeInfo
	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&nodeInfo)
	if err != nil {
		return NodeInfo{}, err
	}

	return nodeInfo, nil
}

func (h *Hsc) GetNewestBlockHeight() (int64, error) {
	endpoint := fmt.Sprintf("%s/blocks/latest", h.LcdEndpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return 0, err
	}

	var blockInfo BlockInfo
	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&blockInfo)
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(blockInfo.Block.Header.Height, 10, 64)
}

func (h *Hsc) GetBlockHeightInfo(height int64) (BlockMeta, error) {
	endpoint := fmt.Sprintf("%s/blocks/%d", h.LcdEndpoint, height)
	fmt.Println(endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return BlockMeta{}, err
	}

	var blockInfo BlockInfo
	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&blockInfo)
	if err != nil {
		return BlockMeta{}, err
	}

	return blockInfo.BlockMeta, nil
}

func (tx *TransferInfo) UserTransferInfo() (UserTransferInfo, error) {
	userTransferInfo := UserTransferInfo{}
	userTransferInfo.Count = tx.Count
	userTransferInfo.Limit = tx.Limit
	userTransferInfo.PageNumber = tx.PageNumber
	userTransferInfo.PageTotal = tx.PageTotal
	userTransferInfo.TotalCount = tx.TotalCount
	userTransferInfo.Txs = make([]UserTxs, 0)
	for _, val := range tx.Txs {
		userTxs := UserTxs{}
		userTxs.Txhash = val.Txhash
		userTxs.Height = val.Height
		if len(val.Logs) < 1 {
			return userTransferInfo, fmt.Errorf("not find logs info")
		}
		userTxs.Success = val.Logs[0].Success
		userTxs.Log = val.Logs[0].Log
		if len(val.Tx.Value.Msg) < 1 {
			return userTransferInfo, fmt.Errorf("not find tx info")
		}
		userTxs.FromAddress = val.Tx.Value.Msg[0].Value.FromAddress
		userTxs.ToAddress = val.Tx.Value.Msg[0].Value.ToAddress
		if len(val.Tx.Value.Msg[0].Value.Amount) < 1 {
			return userTransferInfo, fmt.Errorf("not find amount info")
		}
		userTxs.Amount = val.Tx.Value.Msg[0].Value.Amount[0].Amount
		userTxs.Denom = val.Tx.Value.Msg[0].Value.Amount[0].Denom
		userTxs.Memo = val.Tx.Value.Memo
		userTransferInfo.Txs = append(userTransferInfo.Txs, userTxs)
	}
	return userTransferInfo, nil
}

func (h *Hsc) GetBlockHeightTxInfo(height, limit, page int64) (TransferInfo, error) {
	endpoint := fmt.Sprintf("%s/txs?tx.minheight=%d&tx.maxheight=%d&page=%d&limit=%d", h.LcdEndpoint, height, height, page, limit)
	fmt.Println(endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return TransferInfo{}, err
	}

	var info TransferInfo
	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&info)
	if err != nil {
		return TransferInfo{}, err
	}

	return info, nil
}

func (h *Hsc) GetAddressTxInfo(role, address string, limit, page int64) (TransferInfo, error) {
	endpoint := ""
	switch role {
	case "sender":
		endpoint = fmt.Sprintf("%s/txs?message.action=send&message.%s=%s&page=%d&limit=%d", h.LcdEndpoint, role, address, page, limit)
		break
	case "recipient":
		endpoint = fmt.Sprintf("%s/txs?message.action=send&transfer.%s=%s&page=%d&limit=%d", h.LcdEndpoint, role, address, page, limit)
		break
	default:
		return TransferInfo{}, fmt.Errorf("not find tx message!")
	}

	fmt.Println(endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return TransferInfo{}, err
	}

	var info TransferInfo
	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&info)
	if err != nil {
		return TransferInfo{}, err
	}

	return info, nil
}

func (tx *Txs) UserTxs() (UserTxs, error) {
	userTxs := UserTxs{}
	userTxs.Txhash = tx.Txhash
	userTxs.Height = tx.Height
	if len(tx.Logs) < 1 {
		return userTxs, fmt.Errorf("not find logs info")
	}
	userTxs.Success = tx.Logs[0].Success
	userTxs.Log = tx.Logs[0].Log
	if len(tx.Tx.Value.Msg) < 1 {
		return userTxs, fmt.Errorf("not find tx info")
	}
	userTxs.FromAddress = tx.Tx.Value.Msg[0].Value.FromAddress
	userTxs.ToAddress = tx.Tx.Value.Msg[0].Value.ToAddress
	if len(tx.Tx.Value.Msg[0].Value.Amount) < 1 {
		return userTxs, fmt.Errorf("not find amount info")
	}
	userTxs.Amount = tx.Tx.Value.Msg[0].Value.Amount[0].Amount
	userTxs.Denom = tx.Tx.Value.Msg[0].Value.Amount[0].Denom
	userTxs.Memo = tx.Tx.Value.Memo
	return userTxs, nil
}

func (h *Hsc) GetTxHashTxInfo(hash string) (Txs, error) {
	endpoint := fmt.Sprintf("%s/txs/%s", h.LcdEndpoint, hash)
	fmt.Println(endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return Txs{}, err
	}

	var info Txs
	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&info)
	if err != nil {
		return Txs{}, err
	}

	if info.Error != nil {
		return Txs{}, fmt.Errorf("%+v", *info.Error)
	}
	return info, nil
}
