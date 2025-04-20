package signleflight_v1

import (
	"sync"
)

type SingleFlight struct {
	mu    sync.Mutex       // 保护 calls map
	calls map[string]*call // 正在处理的调用集合
}

type call struct {
	callWG sync.WaitGroup
	val    any
	err    error
}

func (sf *SingleFlight) Do(key string, fn func() (any, error)) (val any, err error) {
	sf.mu.Lock()

	if sf.calls == nil {
		sf.calls = make(map[string]*call)
	}

	// 如果已有相同 key 的调用
	if c, ok := sf.calls[key]; ok {
		sf.mu.Unlock()
		c.callWG.Wait()
		return c.val, c.err
	}

	// 创建新调用
	c := new(call)
	c.callWG.Add(1)
	sf.calls[key] = c
	sf.mu.Unlock()

	// 确保完成后删除调用记录
	defer func() {
		sf.mu.Lock()
		delete(sf.calls, key)
		sf.mu.Unlock()
	}()

	// 执行实际调用
	c.val, c.err = fn()
	c.callWG.Done()

	// 返回原始调用结果
	return c.val, c.err
}
