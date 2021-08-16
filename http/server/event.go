package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hschain/go-sdk/hsc"
)

func ReadBody(c *gin.Context) ([]byte, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Reading body failed!"})
		return body, err
	}
	fmt.Printf("The received body is : %+v", string(body[:]))
	return body, err
}

func (s *Server) TransferAccounts(c *gin.Context) {

	body, err := ReadBody(c)
	if err != nil {
		return
	}

	var infos Transfer
	err = json.Unmarshal(body, &infos)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Failed to parse parameters!"})
		return
	}

	hschain := hsc.NewHsc(s.Lcd)
	wallet, err := hsc.FromMnemonic("hsc", infos.Mnemonic, "m/44'/532'/0'/0/0", hschain)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	log.Printf("wallet is %+v", wallet)

	amount, err := strconv.ParseFloat(infos.Amount, 64)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Currency quantity conversion error!"})
		return
	}

	txHash, err := wallet.TransferAccounts(infos.To, amount, infos.Denom, infos.Memo)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}

	result := map[string]interface{}{
		"txHash": txHash,
	}
	Response(c, result, ResultInfo{Success, "Transfer transaction broadcast successful!"})
}

func (s *Server) TransferDestory(c *gin.Context) {

	body, err := ReadBody(c)
	if err != nil {
		return
	}

	var infos Destory
	err = json.Unmarshal(body, &infos)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Failed to parse parameters!"})
		return
	}

	hschain := hsc.NewHsc(s.Lcd)
	wallet, err := hsc.FromMnemonic("hsc", infos.Mnemonic, "m/44'/532'/0'/0/0", hschain)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	log.Printf("wallet is %+v", wallet)

	amount, err := strconv.ParseFloat(infos.Amount, 64)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Currency quantity conversion error!"})
		return
	}

	txHash, err := wallet.TransferDestory(amount, infos.Denom, infos.Memo)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}

	result := map[string]interface{}{
		"txHash": txHash,
	}
	Response(c, result, ResultInfo{Success, "Destroy transaction broadcast successful!"})
}

func (s *Server) QueryMnemonicAccounts(c *gin.Context) {
	mnemonic := c.DefaultQuery("mnemonic", "")
	if mnemonic == "" {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing mnemonic parameter!"})
		return
	}

	hschain := hsc.NewHsc(s.Lcd)

	wallet, err := hsc.FromMnemonic("hsc", mnemonic, "m/44'/532'/0'/0/0", hschain)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}

	log.Printf("wallet is %+v", wallet)
	account, err := hschain.GetAccountData(wallet.Address)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	result := map[string]interface{}{
		"account": account,
	}
	Response(c, result, ResultInfo{Success, "Get account successful!"})
}

func (s *Server) QueryAddressAccounts(c *gin.Context) {
	address := c.DefaultQuery("address", "")
	if address == "" {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing address parameter!"})
		return
	}

	hschain := hsc.NewHsc(s.Lcd)
	account, err := hschain.GetAccountData(address)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	result := map[string]interface{}{
		"account": account,
	}
	Response(c, result, ResultInfo{Success, "Get account successful!"})
}

func (s *Server) CreateMnemonic(c *gin.Context) {

	mnemonic, err := hsc.GenerateMnemonic()
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	hschain := hsc.NewHsc(s.Lcd)

	wallet, err := hsc.FromMnemonic("hsc", mnemonic, "m/44'/532'/0'/0/0", hschain)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}

	privatekey, err := wallet.GetPrivateKey()
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	result := map[string]interface{}{
		"mnemonic":   mnemonic,
		"privatekey": privatekey,
		"address":    wallet.Address,
	}
	Response(c, result, ResultInfo{Success, "Create mnemonic successful!"})
}

func (s *Server) GetHashTxInfo(c *gin.Context) {

	hash := c.DefaultQuery("hash", "")
	if hash == "" {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing address parameter!"})
		return
	}
	hschain := hsc.NewHsc(s.Lcd)
	transferInfo, err := hschain.GetTxHashTxInfo(hash)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	userTransferInfo, err := transferInfo.UserTxs()
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	result := map[string]interface{}{
		"transferinfo": userTransferInfo,
	}
	Response(c, result, ResultInfo{Success, "get tx successful!"})
}

func (s *Server) GetHeightTxInfo(c *gin.Context) {

	height, err := strconv.ParseInt(c.DefaultQuery("height", "1"), 10, 64)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing height parameter!"})
		return
	}

	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "30"), 10, 64)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing limit parameter!"})
		return
	}

	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing page parameter!"})
		return
	}

	hschain := hsc.NewHsc(s.Lcd)
	transferInfo, err := hschain.GetBlockHeightTxInfo(height, limit, page)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	userTransferInfo, err := transferInfo.UserTransferInfo()
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	result := map[string]interface{}{
		"transferinfo": userTransferInfo,
	}
	Response(c, result, ResultInfo{Success, "get tx successful!"})
}

func (s *Server) GetAddressTxInfo(c *gin.Context) {

	role := c.DefaultQuery("role", "")
	if role == "" {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing role parameter!"})
		return
	}

	address := c.DefaultQuery("address", "")
	if address == "" {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing address parameter!"})
		return
	}

	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "30"), 10, 64)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing limit parameter!"})
		return
	}

	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, "Parameter error: missing page parameter!"})
		return
	}

	hschain := hsc.NewHsc(s.Lcd)
	transferInfo, err := hschain.GetAddressTxInfo(role, address, limit, page)
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}

	userTransferInfo, err := transferInfo.UserTransferInfo()
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	result := map[string]interface{}{
		"transferinfo": userTransferInfo,
	}
	Response(c, result, ResultInfo{Success, "get tx successful!"})
}

func (s *Server) GetLastHeight(c *gin.Context) {
	hschain := hsc.NewHsc(s.Lcd)
	height, err := hschain.GetNewestBlockHeight()
	if err != nil {
		Response(c, nil, ResultInfo{ParameterError, err.Error()})
		return
	}
	result := map[string]interface{}{
		"height": height,
	}
	Response(c, result, ResultInfo{Success, "get last height successful!"})
}
