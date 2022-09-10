package trie

import (
	"errors"

	"github.com/NethermindEth/juno/pkg/collections"
	"github.com/NethermindEth/juno/pkg/felt"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidValue  = errors.New("invalid value")
	ErrUnexistingKey = errors.New("unexisting key")
)

type Trie interface {
	Root() *felt.Felt
	Get(key *felt.Felt) (*felt.Felt, error)
	Put(key *felt.Felt, value *felt.Felt) error
	Del(key *felt.Felt) error
	Commit() error
}

type TrieManager interface {
	GetTrieNode(hash *felt.Felt) (TrieNode, error)
	StoreTrieNode(node TrieNode) error
	StoreTrieNodeBatch(nodes []TrieNode) error
}

type memoryMap struct {
	memory map[felt.Felt]TrieNode
}

func (m *memoryMap) Put(key *felt.Felt, value TrieNode) {
	m.memory[*key] = value
}

func (m *memoryMap) Get(key *felt.Felt) (TrieNode, bool) {
	node, ok := m.memory[*key]
	return node, ok
}

func (m *memoryMap) Clear() {
	m.memory = make(map[felt.Felt]TrieNode, len(m.memory))
}

type trie struct {
	height  int
	root    *felt.Felt
	manager TrieManager
	memory  memoryMap
}

// New creates a new trie, pass zero as root hash to initialize an empty trie
func New(manager TrieManager, root *felt.Felt, height int) Trie {
	return &trie{
		height:  height,
		root:    root,
		manager: manager,
		memory:  memoryMap{make(map[felt.Felt]TrieNode)},
	}
}

// Root returns the hash of the root node of the trie.
func (t *trie) Root() *felt.Felt {
	return t.root
}

// Get gets the value for a key stored in the trie.
func (t *trie) Get(key *felt.Felt) (*felt.Felt, error) {
	path := collections.NewBitSet(t.height, key.ByteSlice())
	node, _, err := t.get(path, false)
	return node, err
}

// Put inserts a new key/value pair into the trie.
func (t *trie) Put(key *felt.Felt, value *felt.Felt) error {
	path := collections.NewBitSet(t.height, key.ByteSlice())
	_, siblings, err := t.get(path, true)
	if err != nil {
		return err
	}
	return t.put(path, value, siblings)
}

// Del deltes the value associated with the given key.
func (t *trie) Del(key *felt.Felt) error {
	return t.Put(key, new(felt.Felt))
}

func (t *trie) Commit() error {
	defer t.memory.Clear()
	node, ok := t.memory.Get(t.root)
	if !ok {
		return errors.New("root not found in memory")
	}
	toStore := []TrieNode{node}
	for i := 0; i < len(toStore); i++ {
		node := toStore[i]
		switch n := node.(type) {
		case *EdgeNode:
			child, ok := t.memory.Get(n.Bottom())
			if ok {
				toStore = append(toStore, child)
			}
		case *BinaryNode:
			leftChild, ok := t.memory.Get(n.LeftH)
			if ok {
				toStore = append(toStore, leftChild)
			}
			rightChild, ok := t.memory.Get(n.RightH)
			if ok {
				toStore = append(toStore, rightChild)
			}

		default:
			return errors.New("unexpected node type on memory")
		}
	}
	return t.manager.StoreTrieNodeBatch(toStore)
}

func (t *trie) getFromMemoryFirst(key *felt.Felt) (TrieNode, error) {
	if node, ok := t.memory.Get(key); ok {
		return node, nil
	}
	return t.manager.GetTrieNode(key)
}

func (t *trie) get(path *collections.BitSet, withSiblings bool) (*felt.Felt, []TrieNode, error) {
	// list of siblings we need to hash with to get to the root
	var siblings []TrieNode
	if withSiblings {
		siblings = make([]TrieNode, t.height)
	}
	curr := t.root // curr is the current node's hash in the traversal
	for walked := 0; walked < t.height && curr.Cmp(EmptyNode.Hash()) != 0; {
		retrieved, err := t.getFromMemoryFirst(curr)
		if err != nil {
			return nil, nil, err
		}

		// switch on the type of the node
		switch node := retrieved.(type) {
		case *EdgeNode:
			// longest common prefix of the key and the edge's path
			lcp := longestCommonPrefix(node.Path(), path.Slice(walked, path.Len()))

			if lcp == node.Path().Len() {
				// if the lcp is the length of the path, we need to go down the edge
				// the node we jump to is either a leaf or a binary node, hence its
				// hash is stored in the edge's bottom
				curr = node.Bottom()
			} else {
				// our path diverges with the edge's path
				if withSiblings {
					// we need to collect the node lcp+1 steps down the edge
					if lcp+1 < node.Path().Len() {
						// sibling is still an edge node
						edgePath := node.Path().Slice(lcp+1, node.Path().Len())
						siblings[walked+lcp] = &EdgeNode{nil, edgePath, node.Bottom()}
					} else if lcp+1 < path.Len()-walked {
						// notest
						// sibling is a binary node, we need to retrieve it from the store
						sibling, err := t.getFromMemoryFirst(node.Bottom())
						if err != nil {
							return nil, nil, err
						}
						// add sibling to the list of siblings
						// notest
						siblings[walked+lcp] = sibling
					} else {
						// sibling is a leaf node
						siblings[walked+lcp] = &leafNode{node.Bottom()}
					}
				}

				// we jump to an empty node since we didn't match the path in the edge
				curr = EmptyNode.Hash()
			}

			// we just walk down lcp steps
			walked += lcp

		case *BinaryNode:
			var nextH, siblingH *felt.Felt
			// walk left or right depending on the bit
			if path.Get(walked) {
				// next is right node
				nextH, siblingH = node.RightH, node.LeftH
			} else {
				// next is left node
				nextH, siblingH = node.LeftH, node.RightH
			}

			if withSiblings {
				if path.Len()-walked > 1 {
					// sibling is a binary node, we need to retrieve it from the store
					sibling, err := t.getFromMemoryFirst(siblingH)
					if err != nil {
						return nil, nil, err
					}
					// add sibling to the list of siblings
					siblings[walked] = sibling
				} else {
					// sibling is a leaf node
					siblings[walked] = &leafNode{siblingH}
				}
			}

			// get the next node
			curr = nextH
			// increment the walked counter
			walked++
		default:
			panic("invalid node type") // should never happen
		}
	}
	return curr, siblings, nil
}

// put inserts a node in a given path in the trie.
func (t *trie) put(path *collections.BitSet, value *felt.Felt, siblings []TrieNode) error {
	var node TrieNode
	node = &leafNode{value}
	// reverse walk the key
	for i := path.Len() - 1; i >= 0; i-- {
		sibling := siblings[i]
		if sibling == nil {
			sibling = EmptyNode
		}

		var left, right TrieNode
		if path.Get(i) {
			left, right = sibling, node
		} else {
			left, right = node, sibling
		}

		leftIsEmpty := left.Hash().Cmp(EmptyNode.Hash()) == 0
		rightIsEmpty := right.Hash().Cmp(EmptyNode.Hash()) == 0

		// compute parent
		if leftIsEmpty && rightIsEmpty {
			node = EmptyNode
		} else if leftIsEmpty {
			edgePath := collections.NewBitSet(right.Path().Len()+1, right.Path().Bytes())
			edgePath.Set(0)
			node = &EdgeNode{nil, edgePath, right.Bottom()}
		} else if rightIsEmpty {
			edgePath := collections.NewBitSet(left.Path().Len()+1, left.Path().Bytes())
			node = &EdgeNode{nil, edgePath, left.Bottom()}
		} else {
			node = &BinaryNode{nil, left.Hash(), right.Hash()}
			t.memory.Put(node.Hash(), node)
			if i < path.Len()-1 {
				// don't store leafs
				t.memory.Put(left.Hash(), left)
				t.memory.Put(right.Hash(), right)
			}
		}
	}

	if t.height > 0 && node.Hash().Cmp(EmptyNode.Hash()) != 0 {
		// only store root if it's neither empty nor a leaf
		t.memory.Put(node.Hash(), node)
	}

	t.root = node.Hash()
	return nil
}

func longestCommonPrefix(path, other *collections.BitSet) int {
	n := 0
	for ; n < path.Len() && n < other.Len(); n++ {
		if path.Get(n) != other.Get(n) {
			break
		}
	}
	return n
}
