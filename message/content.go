package message

type HiContent struct {
	Uid int32 `json:"uid"`
}

func (hc *HiContent) Topic() string {
	return "HiContent"
}
