package algo

type SlidingWindowLogAlgo struct {
	CurrentLoggedRequestsTimestamp []int
}

func (a *SlidingWindowLogAlgo) IsRequestAllowed()
