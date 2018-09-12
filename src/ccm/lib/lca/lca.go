package lca

type node struct {
	name string
	children []*node
	ancestor *node
	rank int
	color int
	parent *node
}

type Data struct {
	Name string
	Children []*Data
}

var root, target_node_1, target_node_2 *node

func makeset(n *node) {
	n.parent =  n
	n.rank = 1
}

func union(nodex, nodey *node) {
	xroot := find(nodex)
	yroot := find(nodey)
	if xroot.rank > yroot.rank {
		yroot.parent = xroot
	} else if xroot.rank < yroot.rank {
		xroot.parent = yroot
	} else if xroot.rank == yroot.rank {
		yroot.parent = xroot
		xroot.rank += 1
	}
}

func find(node *node) *node {
	if node.parent != node {
		node.parent = find(node.parent)
	}
	return node.parent
}

func lca(node *node) (*node, bool) {
	makeset(node)
	node.ancestor = node
	for _, child := range node.children {
		return_value, found := lca(child)
		if found == true {
			return return_value, found
		}
		union(node, child)
		this_node := find(node)
		this_node.ancestor = node
	}
	node.color = 1
	if node.name == target_node_2.name && target_node_2.color == 1 && target_node_1.color == 1 {
		return_val := find(target_node_1).ancestor
		return return_val, true
	} else if node.name == target_node_1.name && target_node_2.color == 1 && target_node_1.color == 1 {
		return_val := find(target_node_2).ancestor
		return return_val, true
	} else {
		return nil, false
	}
}

func find_node_by_name(name string, parent *node) (*node, bool) {
	if parent.name == name {
		return parent, true
	}
	for _, child := range parent.children {
		node, found := find_node_by_name(name, child)
		if found == true {
			return node, true
		}
	}
	return nil, false
}

func build_tree(data *Data) *node {
	node := &node{name: data.Name}
	for _,child := range data.Children {
		child_node := build_tree(child)
		node.children = append(node.children, child_node)
	}
	return node
}

func Search(name1, name2 string) string {
	target_node_1, _ = find_node_by_name(name1, root)
	target_node_2, _ = find_node_by_name(name2, root)
	ancestor, _ := lca(root)
	return ancestor.name
}

func Init(data *Data) {
	root = build_tree(data)
}
