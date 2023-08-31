# Binary Trees

This is to implement a function to check whether two binary trees store the same sequence This uses Go’s concurrency and
channels to write a simple solution.

---

## Leaf-Similar Trees

Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value
sequence.

For example, in the given tree [3,5,1,6,2,9,8,null,null,7,4], the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

```plain
Example:

Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
Output: true
```

```plain
Example:

Input: root1 = [1,2,3], root2 = [1,3,2]
Output: false
```

### Approach 1: Depth First Search

#### Intuition and Algorithm

Let's find the leaf value sequence for both given trees. Afterwards, we can compare them to see if they are equal or
not.

To find the leaf value sequence of a tree, we use a depth first search. Our dfs function writes the node's value if it
is a leaf, and then recursively explores each child. This is guaranteed to visit each leaf in left-to-right order, as
left-children are fully explored before right-children.

#### Complexity Analysis

1. `Time Complexity: O(T1+T2)` where T1, T2, are the lengths of the given trees.
2. `Space Complexity: O(T1+T2)`, the space used in storing the leaf values.

### Related Topics

- Tree
- Depth-First Search
- Binary Tree

---
## Right View of Binary Tree

Given a binary tree A of integers. Return an array of integers representing the right view of the Binary tree.

Right view of a Binary Tree: is a set of nodes visible when the tree is visited from Right side.

Return an integer array denoting the right view of the binary tree A.

> Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you
> can see ordered from top to bottom.

```plain
Example Input
Input 1:

        1
      /   \
     2    3
    / \  / \
   4   5 6  7
  /
 8 
 
Output 1:

[1, 3, 7, 8]

Input 2:

    1
   /  \
  2    3
   \
    4
     \
      5
Output 2:
[1, 3, 4, 5]
```

### Solution

#### Approach: Breadth-First Search

A level order traversal, selecting the far right node in each level makes a lot of sense. We can perform a level order
traversal using a queue and performing a breadth-first search.

A level order traversal can be started by placing the root into the queue. Then for each iteration, we can loop over the
length of the queue,n. By looping over n it means we only ever loop over the current level, meaning we can add nodes to
the queue, and will never reach them as
our for loop will stop, maintaining a perfect level order traversal.

We can place the nodes from left to right, or right to left. If we place nodes right first, then left. Then on each
iteration, our rightmost node will be first in the queue. If we place them left to right, then on each iteration the
rightmost node will be last in our queue.

Time Complexity: O(n) we have to process each node once. If we only tried to process the right node, and skip nodes when
the rightmost node exists, we would be skipping nodes in the left subtree in the cases where the right subtree is
shorter in height than the left subtree.

Space Complexity (O(n). In the worst case, that is a full binary tree, the last level of our traversal will fill our
queue with n/2 nodes, leading us to a O(n) space complexity.

### Related Topics

- Binary Tree
- Tree
- Dynamic Programming
- Depth First Search
- Breadth-first Search

---
