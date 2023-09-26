package jumptable

type skipListNode struct {
	val        int //存放数据
	key        int
	levelNodes []*skipListNode //后续节点
}

type SkipList struct {
	head  *skipListNode
	level int
}

func CreateSkipList() *SkipList {
	return &SkipList{
		head:  nil,
		level: 0,
	}
}

func (sl *SkipList) Get(key int) (int, bool) {
	if sl.head == nil {
		return -1, false
	}

	checkNode := sl.head
	for level := sl.level - 1; level >= 0; level-- {
		for checkNode.levelNodes[level] != nil && checkNode.levelNodes[level].key < key {
			
		}
	}

}
