package trie

type TrieNode struct {
	isPathEnd      bool               // 路径结束标识
	isTrieRootNode bool               // 跟结点标识
	Character      rune               // 标识
	Children       map[rune]*TrieNode // 子结点
}

func NewTrieNode(character rune, isRoot bool) *TrieNode {

	return &TrieNode{
		isTrieRootNode: isRoot,
		Character:      character,
		Children:       make(map[rune]*TrieNode, 0),
	}
}

// 该结点是叶子结点否？
func (n *TrieNode) IsLeafNode() bool {
	return len(n.Children) == 0
}

// 该结点是根结点否？
func (n *TrieNode) IsRootNode() bool {
	return n.isTrieRootNode
}

// 路径已结束否？
func (n *TrieNode) IsPathEnd() bool {
	return n.isPathEnd
}
