package rolling

import (
	"sync"
	"time"
)

//滑动计数器, 不可配置, 10秒窗口的滑动计数
var (
	timeSpan int64 = 10
)

type Counter struct {
	Buckets map[int64]*Bucket
	Mutex   *sync.RWMutex
}

type Bucket struct {
	Value float64
}

func NewCounter() *Counter {
	r := &Counter{
		Buckets: make(map[int64]*Bucket),
		Mutex:   &sync.RWMutex{},
	}
	return r
}

func (r *Counter) GetBucket() *Bucket {
	now := time.Now().Unix()
	var bucket *Bucket
	var ok bool

	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &Bucket{}
		r.Buckets[now] = bucket
	}

	return bucket
}

func (r *Counter) RemoveBucket() {
	now := time.Now().Unix() - timeSpan

	for timestamp := range r.Buckets {
		if timestamp <= now {
			delete(r.Buckets, timestamp)
		}
	}
}

//添加当前bucket
func (r *Counter) Add(i float64) {
	if i == 0 {
		return
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.GetBucket()
	b.Value += i
	r.RemoveBucket()
}

//计算和
func (r *Counter) Sum(now time.Time) float64 {
	sum := float64(0)

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-timeSpan {
			sum += bucket.Value
		}
	}

	return sum
}

//计算最大值
func (r *Counter) Max(now time.Time) float64 {
	var max float64

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-timeSpan {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}

	return max
}

//计算窗口内平均值
func (r *Counter) Avg(now time.Time) float64 {
	return r.Sum(now) / float64(timeSpan)
}
