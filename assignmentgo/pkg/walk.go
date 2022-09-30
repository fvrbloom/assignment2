package pkg

import "fmt"

type WalkStrategy struct {
}

func (r *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("Road A:[%d] to B:[%d] Average speed:[%d]  Total:[%d] Total Time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
