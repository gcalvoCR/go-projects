package model

const (
	ReqAdd        = iota
	ReqAvg        = iota
	ReqRandom     = iota
	ReqSpellCheck = iota
	ReqSearch     = iota
)

// ReqDataSize is max bytes per ClientReq.Data byte array
const ReqDataSize = 1 * 1024 // 1kb

type (
	ClientReq struct {
		ID      int               `json:"id"`       // one of ReqX defined above
		ReqType [ReqDataSize]byte `json:"req_type"` // request specific encoded data
		Size    int               `json:"size"`     // how many byte in data
	}
)
