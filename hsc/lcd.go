package hsc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func NewTransfer(lcdEndpoint string) *Transfer {
	return &Transfer{
		LcdEndpoint: lcdEndpoint,
	}
}

// Retrieve the account data related to the given wallet address, like
// account number and sequence number.
func (T Transfer) GetAccountData(address string) (AccountData, error) {
	endpoint := fmt.Sprintf("%s/auth/accounts/%s", T.LcdEndpoint, address)

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
func (T Transfer) getNodeInfo() (NodeInfo, error) {
	endpoint := fmt.Sprintf("%s/node_info", T.LcdEndpoint)
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

func (T Transfer) GetNewestBlockHeight() (int64, error) {
	endpoint := fmt.Sprintf("%s/blocks/latest", T.LcdEndpoint)
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

func (T Transfer) GetBlockHeightInfo(height int64) (BlockMeta, error) {
	endpoint := fmt.Sprintf("%s/blocks/%d", T.LcdEndpoint, height)
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

func (T Transfer) GetBlockHeightTxInfo(height, limit, page int64) (TransferInfo, error) {
	endpoint := fmt.Sprintf("%s/txs?tx.minheight=%d&tx.maxheight=%d&page=%d&limit=%d", T.LcdEndpoint, height, height, limit, page)
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
	fmt.Println(info)

	return info, nil
}

func (T Transfer) GetTxHashTxInfo(hash string) (Txs, error) {
	endpoint := fmt.Sprintf("%s/txs/%s", T.LcdEndpoint, hash)
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
