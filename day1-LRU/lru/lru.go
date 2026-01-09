package lru

import "container/list"

type lru struct {
	maxBytes  int64
	nBytes    int64
	ll        *list.List
	uMap      map[string]*list.Element
	onEvicted func(key string, value Value)
}

type Value interface {
	Len() int
}

type entry struct {
	k string
	v Value
}

func New(maxBytes int64, onEvicted func(key string, value Value)) *lru {
	return &lru{
		maxBytes:  maxBytes,
		ll:        list.New(),
		uMap:      make(map[string]*list.Element),
		onEvicted: onEvicted,
	}
}

func (l *lru) Get(key string) (value Value, ok bool) {
	if ele, ok := l.uMap[key]; ok {
		l.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.v, true
	}
	return nil, false
}

func (l *lru) RemoveOldest() {
	ele := l.ll.Back()
	if ele != nil {
		l.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(l.uMap, kv.k)
		l.nBytes -= int64(len(kv.k)) + int64(kv.v.Len())
		if l.onEvicted != nil {
			l.onEvicted(kv.k, kv.v)
		}
	}
}

func (l *lru) Add(key string, value Value) {
	if ele, ok := l.uMap[key]; ok {
		l.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		l.nBytes += int64(value.Len()) - int64(kv.v.Len())
		kv.v = value
	} else {
		// l.uMap[key] = l.ll.PushFront(&entry{key, value})
		l.ll.PushFront(&entry{key, value})
		l.uMap[key] = l.ll.Front()
		l.nBytes += int64(len(key)) + int64(value.Len())
	}
	// if l.maxBytes == 0, which means no limit
	for l.maxBytes != 0 && l.nBytes > l.maxBytes {
		l.RemoveOldest()
	}
}

func (l *lru) Len() int {
	return len(l.uMap)
}
