package ratelimiter

import (
	"sync"
	"time"
)

type SlidingWindow struct {
	windowSize   time.Duration
	bucketCount  int
	buckets      []int
	currentIndex int
	lastTime     time.Time
	mutex        sync.Mutex
}

func NewSlidingWindow(windowSize time.Duration, bucketCount int) *SlidingWindow {
	return &SlidingWindow{
		windowSize:  windowSize,
		bucketCount: bucketCount,
		buckets:     make([]int, bucketCount),
		lastTime:    time.Now(),
	}
}

func (sw *SlidingWindow) AddRequest(limit int) bool {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(sw.lastTime)
	bucketDuration := sw.windowSize / time.Duration(sw.bucketCount)
	bucketsToAdvance := int(elapsed / bucketDuration)

	if bucketsToAdvance > 0 {
		for i := 0; i < bucketsToAdvance && i < sw.bucketCount; i++ {
			sw.currentIndex = (sw.currentIndex + 1) % sw.bucketCount
			sw.buckets[sw.currentIndex] = 0
		}
		sw.lastTime = now
	}

	totalRequests := 0
	for _, count := range sw.buckets {
		totalRequests += count
	}

	if totalRequests >= limit {
		return false
	}

	sw.buckets[sw.currentIndex]++
	return true
}
