package cachetogo

import (
	"sync"
	"time"
)

// CacheItem as cache item
type CacheItem struct {
	sync.RWMutex

	key      interface{}
	data     interface{}
	lifeSpan time.Duration

	createdOn   time.Time
	accessedOn  time.Time
	accessCount int64

	// Callback method triggered right before removing the item from the cache
	aboutToExpire []func(key interface{})
}

// NewCacheItem creates cache item
func NewCacheItem(key interface{}, lifeSpan time.Duration, data interface{}) *CacheItem {
	t := time.Now()
	return &CacheItem{
		key:           key,
		lifeSpan:      lifeSpan,
		createdOn:     t,
		accessedOn:    t,
		accessCount:   0,
		aboutToExpire: nil,
		data:          data,
	}
}

// KeepAlive marks an item to be kept for another expireDuration period.
func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()
	item.accessedOn = time.Now()
	item.accessCount++
}

// LifeSpan gets this item's expiration duration
func (item *CacheItem) LifeSpan() time.Duration {
	return item.lifeSpan
}

// AccessedOn gets when this item was last accessed
func (item *CacheItem) AccessedOn() time.Time {
	item.RLock()
	defer item.RUnlock()
	return item.accessedOn
}

// CreatedOn gets when this item was created
func (item *CacheItem) CreatedOn() time.Time {
	return item.createdOn
}

// AccessCount gets how may this item has been accessed
func (item *CacheItem) AccessCount() int64 {
	item.RLock()
	defer item.RUnlock()
	return item.accessCount
}

// Key gets the key of this item
func (item *CacheItem) Key() interface{} {
	return item.key
}

// Data gets the data of this item
func (item *CacheItem) Data() interface{} {
	return item.data
}

// SetAboutToExpireCallback configures a callback, which will be called right
// before the item is about to removed
func (item *CacheItem) SetAboutToExpireCallback(f func(interface{})) {
	if len(item.aboutToExpire) > 0 {
		item.RemoveAboutToExpireCallbacks()
	}
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = append(item.aboutToExpire, f)
}

// AddAboutToExpireCallback appends a new callback to the about to expire queue
func (item *CacheItem) AddAboutToExpireCallback(f func(interface{})) {
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = append(item.aboutToExpire, f)
}

// RemoveAboutToExpireCallback empties the about to expire queue
func (item *CacheItem) RemoveAboutToExpireCallbacks() {
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = nil
}
