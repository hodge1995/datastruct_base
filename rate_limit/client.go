package main

//
//import (
//	"fmt"
//	"net/http"
//	"sync"
//	"time"
//)
//
//// 批量测试滑动窗口
//func testSlidingWindow(count int) {
//	fmt.Printf("\n🧪 测试滑动窗口限流 - 发送 %d 个请求\n", count)
//	fmt.Println("========================================")
//
//	client := &http.Client{Timeout: 5 * time.Second}
//
//	// 创建WaitGroup来等待所有请求完成
//	var wg sync.WaitGroup
//
//	for i := 0; i < count; i++ {
//		wg.Add(1)
//		go func(id int) {
//			defer wg.Done()
//
//			resp, err := client.Get("http://localhost:8080/sliding")
//			if err != nil {
//				fmt.Printf("[请求 %d] ❌ 请求失败: %v\n", id, err)
//				return
//			}
//			defer resp.Body.Close()
//
//			if resp.StatusCode == http.StatusOK {
//				fmt.Printf("[请求 %d] ✅ 请求成功 (HTTP 200)\n", id)
//			} else if resp.StatusCode == http.StatusTooManyRequests {
//				fmt.Printf("[请求 %d] 🚫 请求被限流 (HTTP 429)\n", id)
//			} else {
//				fmt.Printf("[请求 %d] ⚠️  未知状态 (HTTP %d)\n", id, resp.StatusCode)
//			}
//
//			// 添加小延迟以便观察效果
//			time.Sleep(100 * time.Millisecond)
//		}(i)
//	}
//
//	wg.Wait()
//	fmt.Println("========================================")
//}
//
//// 批量测试漏桶
//func testLeakyBucket(count int) {
//	fmt.Printf("\n🧪 测试漏桶限流 - 发送 %d 个请求\n", count)
//	fmt.Println("========================================")
//
//	client := &http.Client{Timeout: 5 * time.Second}
//
//	var wg sync.WaitGroup
//
//	for i := 0; i < count; i++ {
//		wg.Add(1)
//		go func(id int) {
//			defer wg.Done()
//
//			resp, err := client.Get("http://localhost:8080/leaky")
//			if err != nil {
//				fmt.Printf("[请求 %d] ❌ 请求失败: %v\n", id, err)
//				return
//			}
//			defer resp.Body.Close()
//
//			if resp.StatusCode == http.StatusOK {
//				fmt.Printf("[请求 %d] ✅ 请求成功 (HTTP 200)\n", id)
//			} else if resp.StatusCode == http.StatusTooManyRequests {
//				fmt.Printf("[请求 %d] 🚫 请求被限流 (HTTP 429)\n", id)
//			} else {
//				fmt.Printf("[请求 %d] ⚠️  未知状态 (HTTP %d)\n", id, resp.StatusCode)
//			}
//
//			time.Sleep(100 * time.Millisecond)
//		}(i)
//	}
//
//	wg.Wait()
//	fmt.Println("========================================")
//}
//
//// 混合测试两种算法
//func testMixed(count int) {
//	fmt.Printf("\n🧪 混合测试 - 发送 %d 个请求到两个端点\n", count)
//	fmt.Println("========================================")
//
//	client := &http.Client{Timeout: 5 * time.Second}
//
//	var wg sync.WaitGroup
//
//	// 发送一半请求到滑动窗口，一半到漏桶
//	for i := 0; i < count; i++ {
//		wg.Add(2) // 每个请求都发送两个请求（滑动窗口和漏桶）
//
//		go func(id int) {
//			defer wg.Done()
//			resp, _ := client.Get("http://localhost:8080/sliding")
//			if resp != nil {
//				defer resp.Body.Close()
//				if resp.StatusCode == http.StatusOK {
//					fmt.Printf("[滑动 %d] ✅\n", id)
//				} else {
//					fmt.Printf("[滑动 %d] 🚫\n", id)
//				}
//			}
//		}(i)
//
//		go func(id int) {
//			defer wg.Done()
//			resp, _ := client.Get("http://localhost:8080/leaky")
//			if resp != nil {
//				defer resp.Body.Close()
//				if resp.StatusCode == http.StatusOK {
//					fmt.Printf("[漏桶 %d] ✅\n", id)
//				} else {
//					fmt.Printf("[漏桶 %d] 🚫\n", id)
//				}
//			}
//		}(i)
//
//		time.Sleep(50 * time.Millisecond)
//	}
//
//	wg.Wait()
//	fmt.Println("========================================")
//}
//
//// 持续压力测试
//func stressTest(algorithm string, duration time.Duration) {
//	fmt.Printf("\n🔥 压力测试 - %s 算法，持续 %v\n", algorithm, duration)
//	fmt.Println("========================================")
//
//	startTime := time.Now()
//	requestCount := 0
//	successCount := 0
//	blockedCount := 0
//
//	client := &http.Client{Timeout: 2 * time.Second}
//
//	// 创建goroutine池
//	var wg sync.WaitGroup
//	ticker := time.NewTicker(100 * time.Millisecond)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			requestCount++
//			wg.Add(1)
//
//			go func() {
//				defer wg.Done()
//
//				var url string
//				if algorithm == "sliding" {
//					url = "http://localhost:8080/sliding"
//				} else {
//					url = "http://localhost:8080/leaky"
//				}
//
//				resp, err := client.Get(url)
//				if err == nil {
//					defer resp.Body.Close()
//					if resp.StatusCode == http.StatusOK {
//						successCount++
//					} else {
//						blockedCount++
//					}
//				}
//			}()
//
//		case <-time.After(duration):
//			elapsed := time.Since(startTime)
//			fmt.Printf("\n📊 压力测试结果 (%v):\n", elapsed)
//			fmt.Printf("   总请求: %d\n", requestCount)
//			fmt.Printf("   成功: %d (%.2f%%)\n", successCount, float64(successCount)*100/float64(requestCount))
//			fmt.Printf("   被限流: %d (%.2f%%)\n", blockedCount, float64(blockedCount)*100/float64(requestCount))
//			fmt.Println("========================================")
//			return
//		}
//	}
//}
//
//// 展示菜单
//func showMenu() {
//	fmt.Println("\n" + "=".repeat(50))
//	fmt.Println("          限流算法测试客户端")
//	fmt.Println("=".repeat(50))
//	fmt.Println("1. 测试滑动窗口 (发送10个请求)")
//	fmt.Println("2. 测试漏桶 (发送10个请求)")
//	fmt.Println("3. 混合测试 (发送10个请求)")
//	fmt.Println("4. 滑动窗口压力测试 (5秒)")
//	fmt.Println("5. 漏桶压力测试 (5秒)")
//	fmt.Println("6. 查看服务器状态")
//	fmt.Println("0. 退出")
//	fmt.Println("=".repeat(50))
//}
//
//func main() {
//	fmt.Println("🚀 限流算法测试客户端启动")
//	fmt.Println("请确保服务器已在 http://localhost:8080 运行")
//
//	for {
//		showMenu()
//
//		var choice int
//		fmt.Print("请选择 (0-6): ")
//		fmt.Scan(&choice)
//
//		switch choice {
//		case 1:
//			testSlidingWindow(10)
//		case 2:
//			testLeakyBucket(10)
//		case 3:
//			testMixed(10)
//		case 4:
//			stressTest("sliding", 5*time.Second)
//		case 5:
//			stressTest("leaky", 5*time.Second)
//		case 6:
//			client := &http.Client{Timeout: 5 * time.Second}
//			resp, err := client.Get("http://localhost:8080/status")
//			if err != nil {
//				fmt.Printf("❌ 无法获取状态: %v\n", err)
//			} else {
//				defer resp.Body.Close()
//				buf := make([]byte, 2048)
//				n, _ := resp.Body.Read(buf)
//				fmt.Println(string(buf[:n]))
//			}
//		case 0:
//			fmt.Println("👋 再见!")
//			return
//		default:
//			fmt.Println("❌ 无效选择，请重试")
//		}
//
//		fmt.Println("\n按回车键继续...")
//		fmt.Scanln()
//	}
//}
//
//// 添加strings包的repeat方法扩展
//type stringRepeater string
//
//func (s stringRepeater) repeat(count int) string {
//	result := ""
//	for i := 0; i < count; i++ {
//		result += string(s)
//	}
//	return result
//}
//
//// 为string类型添加repeat方法
//func (s string) repeat(count int) string {
//	result := ""
//	for i := 0; i < count; i++ {
//		result += s
//	}
//	return result
//}
