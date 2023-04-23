package main

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	dict := map[*Node]*Node{}
	var helper func(n *Node) *Node
	helper = func(n *Node) *Node {
		if newN, ok := dict[n]; ok {
			return newN
		}
		ret := &Node{
			Val: n.Val,
		}
		dict[n] = ret
		for _, nei := range n.Neighbors {
			ret.Neighbors = append(ret.Neighbors, helper(nei))
		}
		return ret
	}
	return helper(node)
}
