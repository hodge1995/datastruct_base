package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// 滑动窗口限流器
type SlidingWindowLimiter struct {
	maxRequests int           // 窗口内最大请求数
	windowSize  time.Duration // 窗口大小
	requests    []time.Time   // 请求时间队列
	mu          sync.Mutex
}

func NewSlidingWindowLimiter(maxRequests int, windowSize time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		maxRequests: maxRequests,
		windowSize:  windowSize,
		requests:    make([]time.Time, 0),
	}
}

// 检查是否允许请求
func (l *SlidingWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	// 清理过期请求
	l.requests = l.removeOldRequests(now)

	// 检查是否超过限制
	if len(l.requests) >= l.maxRequests {
		return false
	}

	// 记录当前请求
	l.requests = append(l.requests, now)
	return true
}

// 移除过期请求
func (l *SlidingWindowLimiter) removeOldRequests(now time.Time) []time.Time {
	cutoff := now.Add(-l.windowSize)
	idx := 0
	for idx < len(l.requests) && l.requests[idx].Before(cutoff) {
		idx++
	}
	return l.requests[idx:]
}

// 获取当前窗口内的请求数
func (l *SlidingWindowLimiter) GetCurrentCount() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	l.requests = l.removeOldRequests(now)
	return len(l.requests)
}

// 漏桶限流器
type LeakyBucketLimiter struct {
	capacity      int           // 桶的容量
	rate          float64       // 漏桶速率 (请求/秒)
	water         int           // 当前水量
	lastTime      time.Time     // 上次漏水时间
	mu            sync.Mutex
}

func NewLeakyBucketLimiter(capacity int, rate float64) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		capacity:  capacity,
		rate:      rate,
		water:     0,
		lastTime:  time.Now(),
	}
}

// 检查是否允许请求
func (l *LeakyBucketLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	// 计算从上次漏水到现在的间隔
	elapsed := now.Sub(l.lastTime).Seconds()
	// 计算漏出的水量
	l.water = int(float64(l.water) - elapsed*l.rate)
	if l.water < 0 {
		l.water = 0
	}
	l.lastTime = now

	// 检查桶是否已满
	if l.water >= l.capacity {
		return false
	}

	// 加入新请求
	l.water++
	return true
}

// 获取当前桶中的水量
func (l *LeakyBucketLimiter) GetCurrentWater() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(l.lastTime).Seconds()
	l.water = int(float64(l.water) - elapsed*l.rate)
	if l.water < 0 {
		l.water = 0
	}
	l.lastTime = now
	return l.water
}

// 全局限流器实例
var (
	slidingLimiter = NewSlidingWindowLimiter(5, time.Second*10)
	leakyLimiter   = NewLeakyBucketLimiter(10, 1) // 容量10，每秒漏出1个请求
)

// 滑动窗口限流handler
func slidingWindowHandler(w http.ResponseWriter, r *http.Request) {
	// 记录请求时间
	timestamp := time.Now().Format("15:04:05.000")

	if slidingLimiter.Allow() {
		count := slidingLimiter.GetCurrentCount()
		// 可视化显示：当前窗口内请求数
		bar := ""
		for i := 0; i < count; i++ {
			bar += "█"
		}
		for i := count; i < 5; i++ {
			bar += "░"
		}

		msg := fmt.Sprintf("[%s] ✅ 滑动窗口: 请求通过 | 当前: %s (%d/5) | 窗口: 10秒", timestamp, bar, count)
		fmt.Println(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success", "message": "请求通过", "count": ` + fmt.Sprintf("%d", count) + `}`))
	} else {
		count := slidingLimiter.GetCurrentCount()
		bar := ""
		for i := 0; i < 5; i++ {
			bar += "█"
		}

		msg := fmt.Sprintf("[%s] ❌ 滑动窗口: 请求被限流 | 当前: %s (5/5) | 窗口: 10秒", timestamp, bar)
		fmt.Println(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"status": "error", "message": "请求被限流", "count": ` + fmt.Sprintf("%d", count) + `}`))
	}
}

// 漏桶限流handler
func leakyBucketHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Format("15:04:05.000")

	if leakyLimiter.Allow() {
		water := leakyLimiter.GetCurrentWater()
		// 可视化显示：当前水量
		bar := ""
		for i := 0; i < water; i++ {
			bar += "◆"
		}
		for i := water; i < 10; i++ {
			bar += "◇"
		}

		msg := fmt.Sprintf("[%s] ✅ 漏桶: 请求通过 | 当前水量: %s (%d/10) | 速率: 1/秒", timestamp, bar, water)
		fmt.Println(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success", "message": "请求通过", "water": ` + fmt.Sprintf("%d", water) + `}`))
	} else {
		water := leakyLimiter.GetCurrentWater()
		bar := ""
		for i := 0; i < 10; i++ {
			bar += "◆"
		}

		msg := fmt.Sprintf("[%s] ❌ 漏桶: 请求被限流 | 当前水量: %s (10/10) | 速率: 1/秒", timestamp, bar)
		fmt.Println(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"status": "error", "message": "请求被限流", "water": ` + fmt.Sprintf("%d", water) + `}`))
	}
}

// 状态展示handler
func statusHandler(w http.ResponseWriter, r *http.Request) {
	slidingCount := slidingLimiter.GetCurrentCount()
	leakyWater := leakyLimiter.GetCurrentWater()

	// 状态栏可视化
	slidingBar := ""
	for i := 0; i < slidingCount; i++ {
		slidingBar += "█"
	}
	for i := slidingCount; i < 5; i++ {
		slidingBar += "░"
	}

	leakyBar := ""
	for i := 0; i < leakyWater; i++ {
		leakyBar += "◆"
	}
	for i := leakyWater; i < 10; i++ {
		leakyBar += "◇"
	}

	status := fmt.Sprintf(`
========================================
         限流算法状态监控面板
========================================
🪟 滑动窗口 (10秒窗口, 最多5请求)
   状态: %s (%d/5)

🪣 漏桶 (容量10, 速率1/秒)
   状态: %s (%d/10)

测试接口:
  - 滑动窗口: http://localhost:8080/sliding
  - 漏桶: http://localhost:8080/leaky
  - 状态: http://localhost:8080/status
========================================
`, slidingBar, slidingCount, leakyBar, leakyWater)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(status))
}

func main() {
	// 注册路由
	http.HandleFunc("/sliding", slidingWindowHandler)
	http.HandleFunc("/leaky", leakyBucketHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("🚀 限流算法服务器启动中...")
	fmt.Println("📊 访问 http://localhost:8080/status 查看状态")
	fmt.Println("🧪 测试命令:")
	fmt.Println("   滑动窗口: curl http://localhost:8080/sliding")
	fmt.Println("   漏桶:     curl http://localhost:8080/leaky")
	fmt.Println("")

	// 启动服务器
	log.Fatal(http.ListenAndServe(":8080", nil))
}
