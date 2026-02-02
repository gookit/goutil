package lcache

import (
	"container/list"
	"errors"
	"os"
	"sync"
	"time"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/x/stdio"
)

// Item represents a cached item with expiration
type Item struct {
	Val any
	Exp int64 // 过期时间 microtime. 0表示永不过期
}

// isExpired 检查是否已过期
func (i *Item) isExpired() bool {
	if i.Exp == 0 {
		return false
	}
	return time.Now().UnixMicro() > i.Exp
}

func (i *Item) isExpired1(nowUm int64) bool {
	if i.Exp == 0 {
		return false
	}
	return nowUm > i.Exp
}

// Cache represents a thread-safe local cache with TTL support
type Cache struct {
	mu      sync.RWMutex             // 读写锁
	items   map[string]*Item         // 存储 key-Val
	lruList *list.List               // LRU 链表管理访问顺序
	lruMap  map[string]*list.Element // LRU 链表节点索引，用于快速删除
	// 淘汰回调函数
	onEvicted func(key string, value any)
	// Maximum number of cached entries default is 1000
	capacity int
	// serializer name, use for save/load file.
	//
	// default is: "json". see JSONSerializer
	serializer string
}

// OptionFn option config func
type OptionFn func(*Cache)

// New create a new cache instance with options
func New(optFns ...OptionFn) *Cache {
	c := &Cache{
		items:   make(map[string]*Item),
		lruList: list.New(),
		lruMap:  make(map[string]*list.Element),
	}

	c.capacity = 1000
	c.serializer = "json"

	for _, optFn := range optFns {
		optFn(c)
	}
	return c
}

// WithCapacity set cache capacity
func WithCapacity(capacity int) OptionFn {
	return func(c *Cache) {
		c.capacity = capacity
	}
}

// WithSerializer specify serializer name. eg: "json", "gob"
func WithSerializer(serializer string) OptionFn {
	// check serializer name
	if _, ok := serializers[serializer]; !ok {
		panic("not registered serializer name: " + serializer)
	}

	return func(c *Cache) {
		c.serializer = serializer
	}
}

// WithOnEvictFn set cache item evicted callback function
func WithOnEvictFn(fn func(key string, value any)) OptionFn {
	return func(c *Cache) {
		c.onEvicted = fn
	}
}

// Set adds an item to the cache with a specified duration.
// If duration <= 0, the item will never Exp.
func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var exp int64
	if ttl > 0 {
		exp = time.Now().Add(ttl).UnixMicro()
	}

	// 如果 key 已存在，更新值并移动到 LRU 头部
	if elem, ok := c.lruMap[key]; ok {
		c.lruList.MoveToFront(elem)
		c.items[key] = &Item{Val: value, Exp: exp}
		return
	}

	// 检查容量并执行淘汰
	if c.lruList.Len() >= c.capacity {
		c.evict()
	}

	// 添加新项
	c.items[key] = &Item{Val: value, Exp: exp}
	elem := c.lruList.PushFront(key)
	c.lruMap[key] = elem
}

// Get retrieves an item from the cache.
// Returns the Val and true if found and not expired, otherwise nil and false.
func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	it, ok := c.items[key]
	if !ok {
		return nil, false
	}

	// 检查过期
	if it.isExpired() {
		c.removeElement(key)
		return nil, false
	}

	// 更新 LRU 位置
	if elem, ok := c.lruMap[key]; ok {
		c.lruList.MoveToFront(elem)
	}
	return it.Val, true
}

// Has checks if an item exists in the cache.
func (c *Cache) Has(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.items[key] != nil
}

// Keys Get a list of all valid keys in the current cache
//
// 注意：此操作会遍历所有数据，时间复杂度为 O(N)
// 如果数据量巨大，可能会短暂阻塞写操作
func (c *Cache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]string, 0, len(c.items))
	nowUm := time.Now().UnixMicro()

	// 遍历 map 过滤掉已过期的 key
	for k, v := range c.items {
		if !v.isExpired1(nowUm) {
			keys = append(keys, k)
		}
	}

	return keys
}

// Len get the number of items in the cache
//
// 返回的是 map 的大小，包含可能已过期但尚未被清理的“僵尸”数据
// 为了保证 O(1) 的高性能，这里不进行遍历去重
func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

// Clear removes all items from the cache
//
// 这会重置底层的 map 和 list，释放内存引用
// 注意：这不会触发 onEvicted 回调函数，因为那是针对单个元素淘汰的
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.reset()
}

// 直接重新初始化，比逐个 Delete 效率高得多
func (c *Cache) reset() {
	c.items = make(map[string]*Item)
	c.lruMap = make(map[string]*list.Element)
	c.lruList.Init()
}

// Delete removes an item from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.removeElement(key)
}

// removeElement 内部删除方法 (不加锁)
func (c *Cache) removeElement(key string) {
	if elem, ok := c.lruMap[key]; ok {
		c.lruList.Remove(elem)
		delete(c.lruMap, key)
	}
	if it, ok := c.items[key]; ok {
		delete(c.items, key)
		if c.onEvicted != nil {
			c.onEvicted(key, it.Val)
		}
	}
}

// evict 淘汰最久未使用的项
func (c *Cache) evict() {
	elem := c.lruList.Back()
	if elem != nil {
		key := elem.Value.(string)
		c.removeElement(key)
	}
}

// SaveFile Save the cache data to a file.
func (c *Cache) SaveFile(filename string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// 准备序列化数据，剔除已过期的
	data := make(map[string]any)
	nowUm := time.Now().UnixMicro()
	for k, v := range c.items {
		if !v.isExpired1(nowUm) {
			data[k] = v
		}
	}

	if len(data) == 0 {
		return nil
	}

	file, err := fsutil.OpenTruncFile(filename, 0644)
	if err != nil {
		return err
	}
	defer stdio.SafeClose(file)

	serializer, ok := serializers[c.serializer]
	if !ok {
		return errors.New("not registered serializer: " + c.serializer)
	}

	return serializer.EncodeTo(file, data)
}

// LoadFile Recover cache data from file load
func (c *Cache) LoadFile(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer stdio.SafeClose(file)

	serializer, ok := serializers[c.serializer]
	if !ok {
		return errors.New("not registered serializer: " + c.serializer)
	}

	var data map[string]Item
	err = serializer.DecodeFrom(file, &data)
	if err != nil {
		return err
	}

	// 恢复数据 (清空当前数据)
	c.reset()
	nowUm := time.Now().UnixMicro()

	for k, v := range data {
		// 加载时检查是否过期，避免加载即过期
		if v.isExpired1(nowUm) {
			c.items[k] = &v
			elem := c.lruList.PushFront(k)
			c.lruMap[k] = elem
		}
	}

	return nil
}
