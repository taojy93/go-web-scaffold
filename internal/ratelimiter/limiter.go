package ratelimiter

import (
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	requestLimit   int           // 每个窗口允许的请求数
	windowSize     time.Duration // 窗口的大小
	bucketCount    int           // 每个窗口划分成的份数
	GslidingWindow = NewSlidingWindow(windowSize, bucketCount)
)

func init() {
	windowSize = viper.GetDuration("BUCKET_COUNT")
	bucketCount = viper.GetInt("BUCKET_COUNT")
	requestLimit = viper.GetInt("REQUEST_LIMIT")
}

// SlidingWindow 窗口结构
type SlidingWindow struct {
	windowSize   time.Duration // 总窗口大小
	bucketCount  int           // 时间片数量（把总窗口分为几片）
	buckets      []int         // 各时间片的请求数
	currentIndex int           // 当前时间片索引
	lastTime     time.Time     // 最后更新时间
	mutex        sync.Mutex    // 保护临界区
}

// NewSlidingWindow 构造窗口
func NewSlidingWindow(windowSize time.Duration, bucketCount int) *SlidingWindow {
	return &SlidingWindow{
		windowSize:  windowSize,
		bucketCount: bucketCount,
		buckets:     make([]int, bucketCount),
		lastTime:    time.Now(),
	}
}

// AddRequest 请求计数
func (sw *SlidingWindow) AddRequest() bool {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(sw.lastTime)
	bucketDuration := sw.windowSize / time.Duration(sw.bucketCount)
	bucketsToAdvance := int(elapsed / bucketDuration)

	if bucketsToAdvance > 0 {
		// 滑动窗口
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

	if totalRequests >= requestLimit {
		return false
	}

	sw.buckets[sw.currentIndex]++
	return true
}
