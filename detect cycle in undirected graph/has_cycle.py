def hasCycle(numNodes, edges):
    graph = buildGraph(numNodes, edges)
    print("adjacency list", graph)
    print("")
    visited = set()
    for node in range(numNodes):
        print("----------------------------------------------")
        print("dfs node", node)
        if node not in visited and detectCycle(graph, node, visited, None):
            return True
    
    return False
        
def detectCycle(graph, node, visited, parent):
    if node in visited:
        print("node ", node, "has been visited")
        return True
    
    visited.add(node)
    for descendant in graph[node]:
        print("current node: node", node)
        print("visited", visited)
        print("parent: node", parent)
        print("visit descendant: node", descendant) if descendant != parent else "not visit descendant node because it is parent"
        print("")
        if descendant != parent and detectCycle(graph, descendant, visited, node):
            return True

    return False
    
    
def buildGraph(numNodes, edges):
    graph = {x: [] for x in range(numNodes)}
    for edge in edges:
        a, b = edge
        # bidirection
        graph[b].append(a)
        graph[a].append(b)
    return graph
        

if __name__ == "__main__":
    numNodes = 6
    edges = [[1, 0], [3, 0], [5, 0], [2, 1], [3, 1], [4, 3], [5, 4], [3, 5]]
    print("has cycle?", hasCycle(numNodes, edges))

"""
output:
('adjacency list', {0: [1, 3, 5], 1: [0, 2, 3], 2: [1], 3: [0, 1, 4, 5], 4: [3, 5], 5: [0, 4, 3]})

----------------------------------------------
('dfs node', 0)
('current node: node', 0)
('visited', set([0]))
('parent: node', None)
('visit descendant: node', 1)

('current node: node', 1)
('visited', set([0, 1]))
('parent: node', 0)
not visit descendant node because it is parent

('current node: node', 1)
('visited', set([0, 1]))
('parent: node', 0)
('visit descendant: node', 2)

('current node: node', 2)
('visited', set([0, 1, 2]))
('parent: node', 1)
not visit descendant node because it is parent

('current node: node', 1)
('visited', set([0, 1, 2]))
('parent: node', 0)
('visit descendant: node', 3)

('current node: node', 3)
('visited', set([0, 1, 2, 3]))
('parent: node', 1)
('visit descendant: node', 0)

('node ', 0, 'has been visited')
('has cycle?', True)
"""