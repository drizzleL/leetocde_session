package main

type LFUCache struct {
	dict     map[int]*entry
	cap      int
	freqHead *entry
	freqTail *entry
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
	freqHead := &entry{}
	freqTail := &entry{}
	freqHead.next = freqTail
	freqTail.prev = freqHead
	return LFUCache{
		cap:      capacity,
		freqHead: freqHead,
		freqTail: freqTail,
	}
}

// just find,
// add freq
func (this *LFUCache) Get(key int) int {
	item, ok := this.dict[key]
	if !ok {
		return -1
	}
	// add freq
	return item.val
}

// not exist + full, expire old
// push
func (this *LFUCache) Put(key int, value int) {
	oldItem, ok := this.dict[key]
	if ok {
		oldItem.val = value
		return
	}
	// newly inserted
	// if this.freqTail.prev.times == 1 { // 可以添加在这个后面

	// } else { // 新创建一个再添加在这里
	// }
}

// 新的lastItem为head
// 旧的lastItem为当前所属的item
func (this *LFUCache) insertFreq(item *entry, lastFreqItem *freqEntry) {
	if lastFreqItem.next.freq == item.freq { // 插入到队尾
		group := lastFreqItem.next
		item.freqItem = group
		item.prev = group.tail.prev
		item.next = group.tail
		group.tail.prev.next = item
		group.tail.prev = item
	} else { // 不存在分组
		newFreq := &freqEntry{
			freq: item.freq,
			prev: lastFreqItem,
			next: lastFreqItem.next,
			head: item,
			tail: item,
		}
		item.freqItem = newFreq
		item.next = newFreq.head
		item.prev = newFreq.tail
	}
}

func (this *LFUCache) remove(item *entry) {
	// 如果他是group head
	// if item.prevFreq.next == item {

	// } else { // tail, just remove
	// 	item.prev.next = item.next
	// 	item.next.prev = item.prev
	// }
}
