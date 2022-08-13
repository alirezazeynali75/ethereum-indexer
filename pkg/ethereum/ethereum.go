package ethereum

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Ethereum struct {
	NodeUrl string
	Headers http.Header
}

func (eth *Ethereum) GetTransactionByHash(id string) (*Transaction, error) {
	tx := &GetTransaction{}
	body := `{"jsonrpc":"2.0","method":"eth_getTransactionByHash", "params":["` + id + `"], "id":1}`
	bodyBytes := []byte(body)
	req, err := http.NewRequest("POST", eth.NodeUrl, bytes.NewBuffer(bodyBytes))
	req.Header = eth.Headers
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(data, tx); err != nil {
		return nil, err
	}
	return &tx.Result, nil
}
