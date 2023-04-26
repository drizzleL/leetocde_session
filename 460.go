package main

type LFUCache struct {
	dict     map[int]*entry
	cap      int
	freqHead *freqEntry
}

type freqEntry struct {
	prev, next *freqEntry
	freq       int
	head, tail *entry
}

type entry struct {
	key  int
	val  int
	freq int

	next     *entry
	prev     *entry
	freqItem *freqEntry
}

func Constructor(capacity int) LFUCache {
	head := newFreqEntry()
	tail := newFreqEntry()
	head.next = tail
	tail.prev = head
	return LFUCache{
		cap:      capacity,
		dict:     make(map[int]*entry),
		freqHead: head,
	}
}

func newFreqEntry() *freqEntry {
	freq := new(freqEntry)
	head := new(entry)
	tail := new(entry)
	head.next = tail
	tail.prev = head
	freq.head = head
	freq.tail = tail
	return freq
}

// just find,
// add freq
func (this *LFUCache) Get(key int) int {
	item, ok := this.dict[key]
	if !ok {
		return -1
	}
	lastFreq := this.remove(item)
	item.freq++
	if lastFreq {
		this.insertFreq(item, item.freqItem.prev)
	} else {
		this.insertFreq(item, item.freqItem)
	}
	return item.val
}

func (this *LFUCache) Put(key int, value int) {
	item, ok := this.dict[key]
	if ok {
		item.val = value
		lastFreq := this.remove(item)
		item.freq++
		if lastFreq {
			this.insertFreq(item, item.freqItem.prev)
		} else {
			this.insertFreq(item, item.freqItem)
		}
		return
	}
	if this.cap == len(this.dict) { /// forget first
		this.forget()
	}
	item = &entry{
		key:  key,
		val:  value,
		freq: 1,
	}
	this.insertFreq(item, this.freqHead)

}

func (this *LFUCache) insertFreq(item *entry, lastFreqItem *freqEntry) {
	var freq *freqEntry
	if lastFreqItem.next.freq == item.freq { // 插入到队尾
		freq = lastFreqItem.next
	} else { // 不存在分组
		freq = newFreqEntry()

		freq.freq = item.freq
		freq.prev = lastFreqItem
		freq.next = lastFreqItem.next
		freq.prev.next = freq
		freq.next.prev = freq
	}
	item.freqItem = freq
	item.prev = freq.tail.prev
	item.next = freq.tail
	freq.tail.prev.next = item
	freq.tail.prev = item
	this.dict[item.key] = item
}

func (this *LFUCache) remove(item *entry) bool {
	delete(this.dict, item.key)
	// remove from item list
	item.prev.next = item.next
	item.next.prev = item.prev

	freq := item.freqItem
	if freq.head.next == freq.tail { // freqitem can be deleted
		freq.prev.next = freq.next
		freq.next.prev = freq.prev
		return true
	}
	return false
}

func (this *LFUCache) forget() {
	item := this.freqHead.next.head.next
	this.remove(item)
}
