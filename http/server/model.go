package server

type State int32

const (
	Success State = 1000

	ParameterError State = -1003
	NetworkError   State = -1004
)

type ResultInfo struct {
	//状态码
	State State `json:"state"`
	//错误或者返回信息
	MessageInfo string `json:"messageinfo"`
}

type Server struct {
	ListenAddress string
	Lcd           string
}

type HscInfo struct {
	Mnemonic string `json:"mnemonic"`
}

type Transfer struct {
	HscInfo
	To     string `json:"to"`
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
	Memo   string `json:"memo"`
}

type Destory struct {
	HscInfo
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
	Memo   string `json:"memo"`
}
