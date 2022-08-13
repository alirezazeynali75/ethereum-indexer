package ethereum

type Transaction struct {
	AccessList []string `json:"accessList,omitempty"`
	BlockHash string `json:"blockHash"`
	BlockNumber string	`json:"blockNumber"`
	ChainId string `json:"chainId"`
	From string `json:"from"`
	Gas string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Hash string `json:"hash"`
	Input string `json:"input"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	MaxFeePerGas string `json:"maxFeePerGas"`
	Nonce string `json:"nonce"`
	R string `json:"r"`
	S string `json:"s"`
	To string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Type string `json:"type"`
	V string `json:"v"`
	Value string `json:"value"`
}

type GetTransaction struct {
	JsonRPC string `json:"jsonrpc"`
	Id int `json:"id"`
	Result Transaction `json:"result"`
}