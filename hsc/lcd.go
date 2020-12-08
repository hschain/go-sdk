package hsc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Retrieve the account data related to the given wallet address, like
// account number and sequence number.
func getAccountData(lcdEndpoint, address string) (AccountData, error) {
	endpoint := fmt.Sprintf("%s/auth/accounts/%s", lcdEndpoint, address)

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
func getNodeInfo(lcdEndpoint string) (NodeInfo, error) {
	endpoint := fmt.Sprintf("%s/node_info", lcdEndpoint)
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
