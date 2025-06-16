package fil

type RawTx struct {
	Version    uint64  `json:"version"`
	To         string  `json:"to"`
	From       string  `json:"from"`
	Value      float64 `json:"value"`
	Nonce      uint64  `json:"nonce"`
	GasLimit   int64   `json:"gas_limit"`
	GasFeeCap  int64   `json:"gas_fee_cap"`
	GasPremium int64   `json:"gas_premium"`
	Method     uint64  `json:"method"`
	Params     []byte  `json:"params"`
}

type Signature struct {
	SignType string `json:"sign_type"` // secp256k1
	Data     []byte `json:"data"`
}

// ##################### plugin
type BroadcastRequest struct {
	ENV       string    `json:"env"` // main or test
	Raw       RawTx     `json:"raw"`
	Signature Signature `json:"signature"`
}

// ########################## sign
type SignRequest struct {
	ENV  string `json:"env"` // main or test
	Info RawTx  `json:"raw_tx"`
}
