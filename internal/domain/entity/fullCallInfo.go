package entity

type FullCallInfo struct {
	Adcel *CelTable `json:"cel"`
	Adcrd *CdrTable `json:"cdr"`
}
