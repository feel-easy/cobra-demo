package tree

type MyNode struct {
	node *Node
}

type A interface {
	Get(url string) string
}

type Aa struct {
}

func (a Aa) Get(url string) string {
	return url
}

func (myNode *MyNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

}
