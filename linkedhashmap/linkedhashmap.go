package linkedhashmap

type node struct {
	key   interface{}
	value interface{}
	prev  *node
	next  *node
}

// LinkedHashMap is a hash map with linked list entries to maintain the insertion order.
type LinkedHashMap struct {
	head      *node
	tail      *node
	size      int
	entryMap  map[interface{}]*node
	tableSize int
}

// NewLinkedHashMap creates a new instance of LinkedHashMap with the specified table size.
func NewLinkedHashMap(tableSize int) *LinkedHashMap {
	entryMap := make(map[interface{}]*node, tableSize)
	return &LinkedHashMap{
		entryMap:  entryMap,
		tableSize: tableSize,
	}
}

// Put inserts a key-value pair into the map.
func (m *LinkedHashMap) Put(key interface{}, value interface{}) {
	entry := m.entryMap[key]
	if entry != nil {
		entry.value = value
		return
	}

	newNode := &node{
		key:   key,
		value: value,
	}

	if m.tail == nil {
		m.head = newNode
		m.tail = newNode
	} else {
		newNode.prev = m.tail
		m.tail.next = newNode
		m.tail = newNode
	}

	m.entryMap[key] = newNode
	m.size++
}

// Get retrieves the value associated with the given key.
func (m *LinkedHashMap) Get(key interface{}) interface{} {
	entry := m.entryMap[key]
	if entry != nil {
		return entry.value
	}
	return nil
}

// Visit applies a function to each entry in the map.
func (m *LinkedHashMap) Visit(visitFn func(key interface{}, value interface{}) bool) {
	entry := m.head
	for entry != nil {
		if !visitFn(entry.key, entry.value) {
			break
		}
		entry = entry.next
	}
}

// Size returns the number of entries in the map.
func (m *LinkedHashMap) Size() int {
	return m.size
}
