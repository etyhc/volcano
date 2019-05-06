package message

type HiContent struct {
	UID  int32  `json:"uid"`
	Addr string `json:"addr"`
}

func (hc *HiContent) Topic() string {
	return "HiContent"
}
