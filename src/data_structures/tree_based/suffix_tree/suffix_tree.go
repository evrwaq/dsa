package data_structures

type SuffixTreeNode struct {
    children map[rune]*SuffixTreeNode
    suffixLink *SuffixTreeNode
    start, *end *int
    suffixIndex int
}

type SuffixTree struct {
    root *SuffixTreeNode
    text string
    activeNode *SuffixTreeNode
    activeEdge int
    activeLength int
    remainingSuffixCount int
    leafEnd *int
    size int
}

func NewSuffixTree(text string) *SuffixTree {
    size := len(text)
    leafEnd := -1
    root := &SuffixTreeNode{
        children:    make(map[rune]*SuffixTreeNode),
        suffixLink:  nil,
        start:       nil,
        end:         nil,
        suffixIndex: -1,
    }
    st := &SuffixTree{
        root:        root,
        text:        text,
        activeNode:  root,
        activeEdge:  -1,
        activeLength: 0,
        remainingSuffixCount: 0,
        leafEnd:     &leafEnd,
        size:        size,
    }
    st.buildSuffixTree()
    return st
}

func (st *SuffixTree) buildSuffixTree() {
    for i := 0; i < st.size; i++ {
        st.extendSuffixTree(i)
    }
}

func (st *SuffixTree) extendSuffixTree(pos int) {
    st.leafEnd = &pos
    st.remainingSuffixCount++
    
    var lastNewNode *SuffixTreeNode
    
    for st.remainingSuffixCount > 0 {
        if st.activeLength == 0 {
            st.activeEdge = pos
        }
        
        edge := rune(st.text[st.activeEdge])
        
        if st.activeNode.children[edge] == nil {
            st.activeNode.children[edge] = &SuffixTreeNode{
                children:    make(map[rune]*SuffixTreeNode),
                start:       &pos,
                end:         st.leafEnd,
                suffixIndex: -1,
            }
            
            if lastNewNode != nil {
                lastNewNode.suffixLink = st.activeNode
                lastNewNode = nil
            }
        } else {
            next := st.activeNode.children[edge]
            if st.walkDown(next) {
                continue
            }
            
            if st.text[next.start + st.activeLength] == st.text[pos] {
                if lastNewNode != nil && st.activeNode != st.root {
                    lastNewNode.suffixLink = st.activeNode
                    lastNewNode = nil
                }
                
                st.activeLength++
                break
            }
            
            splitEnd := next.start + st.activeLength - 1
            split := &SuffixTreeNode{
                children:    make(map[rune]*SuffixTreeNode),
                start:       next.start,
                end:         &splitEnd,
                suffixIndex: -1,
            }
            st.activeNode.children[edge] = split
            
            split.children[rune(st.text[pos])] = &SuffixTreeNode{
                children:    make(map[rune]*SuffixTreeNode),
                start:       &pos,
                end:         st.leafEnd,
                suffixIndex: -1,
            }
            next.start += st.activeLength
            split.children[rune(st.text[next.start])] = next
            
            if lastNewNode != nil {
                lastNewNode.suffixLink = split
            }
            
            lastNewNode = split
        }
        
        st.remainingSuffixCount--
        
        if st.activeNode == st.root && st.activeLength > 0 {
            st.activeLength--
            st.activeEdge = pos - st.remainingSuffixCount + 1
        } else if st.activeNode != st.root {
            st.activeNode = st.activeNode.suffixLink
        }
    }
}

func (st *SuffixTree) walkDown(currNode *SuffixTreeNode) bool {
    edgeLength := *currNode.end - *currNode.start + 1
    if st.activeLength >= edgeLength {
        st.activeEdge += edgeLength
        st.activeLength -= edgeLength
        st.activeNode = currNode
        return true
    }
    return false
}
