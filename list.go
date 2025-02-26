package main

import (
	"sync"
)

// Node structure for the single-lock list
type Node struct {
	key  int
	next *Node
}

// ConcurrentLinkedList using a single lock for the whole list
type ConcurrentLinkedList struct {
	head *Node
	mu   sync.Mutex
}

// Insert adds a new node at the beginning
func (l *ConcurrentLinkedList) Insert(key int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := &Node{key: key, next: l.head}
	l.head = newNode
}

// Search locks for a key in the list
func (l *ConcurrentLinkedList) Search(key int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	curr := l.head
	for curr != nil {
		if curr.key == key {
			return true
		}
		curr = curr.next
	}
	return false
}

// Delete removes a node with the given key
func (l *ConcurrentLinkedList) Delete(key int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		return
	}

	if l.head.key == key {
		l.head = l.head.next
		return
	}

	prev, curr := l.head, l.head.next
	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			return
		}
		prev, curr = curr, curr.next
	}
}

// Node structure for hand-over-hand locking
type HOHNode struct {
	key  int
	next *HOHNode
	mu   sync.Mutex
}

// HandOverHandLinkedList uses fine-grained locking
type HandOverHandLinkedList struct {
	head *HOHNode
	mu   sync.Mutex
}

// Insert using hand-over-hand locking
func (l *HandOverHandLinkedList) Insert(key int) {
	newNode := &HOHNode{key: key}

	l.mu.Lock()
	newNode.next = l.head
	l.head = newNode
	l.mu.Unlock()
}

// Search using hand-over-hand locking
func (l *HandOverHandLinkedList) Search(key int) bool {
	l.mu.Lock()
	curr := l.head
	if curr != nil {
		curr.mu.Lock()
	}
	l.mu.Unlock()

	for curr != nil {
		if curr.key == key {
			curr.mu.Unlock()
			return true
		}

		next := curr.next
		if next != nil {
			next.mu.Lock()
		}
		curr.mu.Unlock()
		curr = next
	}
	return false
}

// Delete using hand-over-hand locking
func (l *HandOverHandLinkedList) Delete(key int) {
	l.mu.Lock()
	if l.head == nil {
		l.mu.Unlock()
		return
	}

	if l.head.key == key {
		l.head = l.head.next
		l.mu.Unlock()
		return
	}

	prev := l.head
	prev.mu.Lock()
	l.mu.Unlock()

	curr := prev.next
	for curr != nil {
		curr.mu.Lock()
		if curr.key == key {
			prev.next = curr.next
			curr.mu.Unlock()
			prev.mu.Unlock()
			return
		}
		prev.mu.Unlock()
		prev = curr
		curr = curr.next
	}
	prev.mu.Unlock()
}
