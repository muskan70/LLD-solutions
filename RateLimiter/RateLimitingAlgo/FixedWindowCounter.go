package algo

type FixedWindowCounterAlgo struct {
	CurrentRequests int
	CurrentWindow   int
}

func (a *FixedWindowCounterAlgo) IsRequestAllowed()
