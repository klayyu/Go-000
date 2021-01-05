package main

import (
	"log"
	"sync"
	"time"
)

var windowTime int64 = 5//秒
var bucketTime int64 = 1//秒

type window struct {
	Buckets map[int64]*Bucket
	Mutex   *sync.RWMutex
}

type Bucket struct {
	Value int64
}

func NewWindow() *window {
	w := &window{
		Buckets: make(map[int64]*Bucket),
		Mutex:   &sync.RWMutex{},
	}

	return w
}

func (w *window) getCurrentBucket() *Bucket {
	now := time.Now().Unix()
	var bucket *Bucket
	var ok bool

	if bucket, ok = w.Buckets[now]; !ok {
		bucket = &Bucket{}
		w.Buckets[now] = bucket
	}

	return bucket
}


func (w *window) removeOldBucket() {
	now := time.Now().Unix() - windowTime
	for timeStamp := range w.Buckets {
		if timeStamp <= now {
			delete(w.Buckets, timeStamp)
		}
	}
}

func (w *window) Increment(i int64) {
	if i == 0 {
		return
	}

	log.Println("Increment... i=",i)
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	b := w.getCurrentBucket()
	b.Value = i
	w.removeOldBucket()
}

func (w *window) Sum(now int64) int64 {
	sum := int64(0)
	w.Mutex.RLock()
	defer w.Mutex.RUnlock()

	for timeStamp, bucket := range w.Buckets {
		if timeStamp >= now - windowTime {
			log.Println("Sum... bucket.Value=",bucket.Value)
			sum += bucket.Value
		}
	}
	return sum
}

func (w *window) Avg(now int64) int64 {
	return w.Sum(now) / (windowTime / bucketTime)
}

func main() {
	log.Println("Hello world")
	w := NewWindow()
	for _, x := range []int64{15, 14, 13, 14, 15,6,7,4,6,8} {
		w.Increment(x)
		time.Sleep(time.Second * time.Duration(bucketTime))
	}
	log.Println("calu avg...")
	log.Println(w.Avg(time.Now().Unix()))

}
