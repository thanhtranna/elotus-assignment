package sumofdistancesintree

func sumOfDistancesInTree(N int, edges [][]int) []int {
	// count[i] stores i as the root node, the total number of all subtree nodes and the root node
	tree, visited, count, result := make([][]int, N), make([]bool, N), make([]int, N), make([]int, N)
	for _, e := range edges {
		i, j := e[0], e[1]
		tree[i] = append(tree[i], j)
		tree[j] = append(tree[j], i)
	}

	deepFirstSearch(0, visited, count, result, tree)
	// Reset the access state and do DFS again
	visited = make([]bool, N)
	// Before entering the second DFS, only the correct value is stored in result[0], because the first DFS calculated all the paths with 0 as the root node and
	// The purpose of the second DFS is to convert the path sum with 0 as the root node into the path sum with n as the root node
	deepSecondSearch(0, visited, count, result, tree)

	return result
}

func deepFirstSearch(root int, visited []bool, count, res []int, tree [][]int) {
	visited[root] = true
	for _, n := range tree[root] {
		if visited[n] {
			continue
		}
		deepFirstSearch(n, visited, count, res, tree)
		count[root] += count[n]
		// Sum of all paths from the root node to n = paths from the root node to all subtrees with n as the root node and the number of each node in res[n] + root to count[n] (the root node and the root node with n as the root node Each node of each adds a path)
		// add a path to the root node and each node with n as the root node = take n as the root node, the sum of the number of subtree nodes and the number of root nodes, that is, count[n]
		res[root] += res[n] + count[n]
	}
	count[root]++
}

// Starting from the root, set the child nodes of the root node as the new root node in turn
func deepSecondSearch(root int, visited []bool, count, res []int, tree [][]int) {
	N := len(visited)
	visited[root] = true
	for _, n := range tree[root] {
		if visited[n] {
			continue
		}
		// After the root node changes from root to n
		// res[root] stores the total length of the path from the root node to all nodes
		// 1. The length of the path increased from root to node n = the root node and each node with n as the root node adds a path = with n as the root node, the sum of the number of subtree nodes and the number of root nodes, that is, count[ n]
		// 2. The length of the path added from n to all subtree nodes with n as the root node = n nodes and each node of the non-n as the root node subtree adds a path = N - count[n]
		// So to transfer the root node from root to n, the path that needs to be increased is calculated in the second step of the above , and the path that needs to be reduced is calculated in the first step of the above
		res[n] = res[root] + (N - count[n]) - count[n]
		deepSecondSearch(n, visited, count, res, tree)
	}
}
