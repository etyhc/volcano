package message

type HiContent struct {
	UID int32  `json:"uid"`
	AID uint32 `json:"aid"`
}

func (hc *HiContent) Topic() string {
	return "HiContent"
}
