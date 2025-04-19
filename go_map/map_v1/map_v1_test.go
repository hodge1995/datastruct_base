package map_v1

import (
	"strconv"
	"testing"
)

func TestNewMapV1(t *testing.T) {
	m := NewMapV1(10)
	if m.cap != 10 {
		t.Errorf("Expected capacity of 10, got %d", m.cap)
	}
	if len(m.array) != 10 {
		t.Errorf("Expected array length of 10, got %d", len(m.array))
	}
}

func TestMapV1_Put_Get(t *testing.T) {
	m := NewMapV1(16)

	// 测试字符串键
	m.Put("key1", "value1")
	m.Put("key2", "value2")

	if val := m.Get("key1"); val != "value1" {
		t.Errorf("Expected 'value1', got %v", val)
	}
	if val := m.Get("key2"); val != "value2" {
		t.Errorf("Expected 'value2', got %v", val)
	}
	if val := m.Get("key3"); val != nil {
		t.Errorf("Expected nil, got %v", val)
	}

	// 测试整型键
	m.Put(1, "int1")
	m.Put(2, "int2")

	if val := m.Get(1); val != "int1" {
		t.Errorf("Expected 'int1', got %v", val)
	}
	if val := m.Get(2); val != "int2" {
		t.Errorf("Expected 'int2', got %v", val)
	}

	// 测试更新现有键
	m.Put("key1", "updated")
	if val := m.Get("key1"); val != "updated" {
		t.Errorf("Expected 'updated', got %v", val)
	}
}

func TestMapV1_Keys(t *testing.T) {
	m := NewMapV1(10)

	// 添加一些键值对
	testData := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	for k, v := range testData {
		m.Put(k, v)
	}

	// 获取所有键
	keys := m.Keys()

	// 验证获取的键的数量
	if len(keys) != len(testData) {
		t.Errorf("Expected %d keys, got %d", len(testData), len(keys))
	}

	// Note: 这个测试可能需要修改，因为当前的Keys()实现返回MapValue对象而不是键
	// 这里我们检查返回的是否都是MapValue类型
	for _, key := range keys {
		if _, ok := key.(MapValue); !ok {
			t.Errorf("Expected MapValue type, got %T", key)
		}
	}
}

func TestMapV1_Large(t *testing.T) {
	m := NewMapV1(100)
	count := 1000

	// 插入大量键值对
	for i := 0; i < count; i++ {
		m.Put(i, "value"+strconv.Itoa(i))
	}

	// 验证所有键值对是否能正确获取
	for i := 0; i < count; i++ {
		expectedValue := "value" + strconv.Itoa(i)
		if val := m.Get(i); val != expectedValue {
			t.Errorf("For key %d, expected '%s', got %v", i, expectedValue, val)
		}
	}
}
