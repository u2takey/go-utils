package tcache

// tcache is just another ttl lru cache with two ttl, if get a key of:
// LifeTime (soft ttl): return value and trigger update-key operation
// MaxLifeTime (hard ttl): do update-key operation first and then return updated value

import (
	"container/list"
	"sync"
	"sync/atomic"
	"time"
)

type UpdateKeyFunc func(key string) (value interface{}, err error)

// 1. Keep entries available, trigger a update-key operation instead of purging entries out-of-date.
// 2. concurrent access safe
// 3. LRU
type Cache struct {
	// LifeTime is lifespan for all entries
	LifeTime uint64
	// if exceed MaxLifeTime will force update
	MaxLifeTime uint64
	MaxEntries  int

	// update-key operation will be trigger when value is nil or value is out-of-date
	updateKeyFunc UpdateKeyFunc

	cache map[string]*list.Element
	ll    *list.List
	sync.RWMutex
}

const (
	ready int32 = iota
	updating
)

type entry struct {
	Key   string
	Born  int64 // Unix time
	State int32
	Value interface{}
}

// NewCache creates a new Cache.
// auto flush cache to file if flushFile not "" and flushInterval > 0
// for flush cache value must jsonable
func NewCache(lifeTime, maxlifeTime uint64, maxEntries int, updateKeyFunc UpdateKeyFunc) *Cache {
	c := &Cache{
		LifeTime:      lifeTime,
		MaxLifeTime:   maxlifeTime,
		MaxEntries:    maxEntries,
		ll:            list.New(),
		updateKeyFunc: updateKeyFunc,
		cache:         make(map[string]*list.Element),
	}

	if c.MaxLifeTime < 100*lifeTime {
		c.MaxLifeTime = 100 * lifeTime
	}
	return c
}

// Add adds a value to the cache. Update born and state if exsit
func (c *Cache) Add(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	if c.cache == nil {
		c.cache = make(map[string]*list.Element)
		c.ll = list.New()
	}
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value = &entry{
			Key:   key,
			Value: value,
			Born:  time.Now().Unix(),
			State: ready,
		}
		return
	}
	ele := c.ll.PushFront(&entry{
		Key:   key,
		Value: value,
		Born:  time.Now().Unix(),
		State: ready,
	})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.removeOldest()
	}
}

// Get looks up a key's value from the cache.
// if entry is nil will block util get value
// if entry is out of date will return old value and trigger an upate-key operation if needed
func (c *Cache) Get(key string) (interface{}, error) {
	c.RLock()
	if c.cache == nil {
		c.cache = make(map[string]*list.Element)
		c.ll = list.New()
	}
	now := time.Now().Unix()

	ee, hit := c.cache[key]
	var ele *entry
	if hit {
		ele = ee.Value.(*entry)
	}
	c.RUnlock()

	// hit
	if hit && uint64(now-ele.Born) < c.MaxLifeTime {

		// out of date and not updating
		if uint64(now-ele.Born) > c.LifeTime && atomic.CompareAndSwapInt32(&ele.State, ready, updating) {
			go func() {
				value, err := c.updateKeyFunc(key)
				if err == nil {
					c.Add(key, value)
				} else {
					atomic.SwapInt32(&ele.State, ready)
				}
			}()
		}
		return ele.Value, nil
	} else { // not hit
		value, err := c.updateKeyFunc(key)
		if err == nil {
			c.Add(key, value)
			return value, nil
		}
		return nil, err
	}
}

// Remove removes the provided key from the cache.
func (c *Cache) Remove(key string) {
	c.Lock()
	defer c.Unlock()
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

// Len returns the number of items in the cache.
func (c *Cache) Len() int {
	c.RLock()
	defer c.RUnlock()
	if c.cache == nil {
		return 0
	}
	return len(c.cache)
}

func (c *Cache) removeOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

// Remove specific element
func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.Key)
}
