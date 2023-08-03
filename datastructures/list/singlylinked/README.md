# Singly Linked List

Algorithms on a singly linked list

## Sort List in Ascending Order

The problem is to sort the linked list in O(nlogn) time and using only constant extra space. If we look at various sorting algorithms, Merge Sort is one of the efficient sorting algorithms that is popularly used for sorting the linked list. The merge sort algorithm runs in O(nlogn) time in all the cases. Let's discuss approaches to sort linked list using merge sort.

> Quicksort is also one of the efficient algorithms with the average time complexity of O(nlogn). But the worst-case time complexity is O(n^2). Also, variations of the quick sort like randomized quicksort are not efficient for the linked list because unlike arrays, random access in the linked list is not possible in O(1) time. If we sort the linked list using quicksort, we would end up using the head as a pivot element which may not be efficient in all scenarios.

Merge sort is a popularly known algorithm that follows the Divide and Conquer Strategy. The divide and conquer strategy can be split into 2 phases:

Divide phase: Divide the problem into subproblems.

Conquer phase: Repeatedly solve each subproblem independently and combine the result to form the original problem.

The Top Down approach for merge sort recursively splits the original list into sublists of equal sizes, sorts each sublist independently, and eventually merge the sorted lists. Let's look at the algorithm to implement merge sort in Top Down Fashion.

Algorithm

Recursively split the original list into two halves. The split continues until there is only one node in the linked list (Divide phase). To split the list into two halves, we find the middle of the linked list using the Fast and Slow pointer approach as mentioned in Find Middle Of Linked List.

Recursively sort each sublist and combine it into a single sorted list. (Merge Phase). This is similar to the problem Merge two sorted linked lists

The process continues until we get the original list in sorted order.

For the linked list = [10,1,60,30,5], the following figure illustrates the merge sort process using a top down approach.

![sort_list](./assets/sort_list.png)

### Complexity Analysis

- Time Complexity: O(nlogn), where nnn is the number of nodes in linked list. The algorithm can be split into 2 phases, Split and Merge.
Let's assume that nnn is power of 222. For n = 16, the split and merge operation in Top Down fashion can be visualized as follows

![top_down_sort_list](./assets//top_down_sort_list.png)

    Split

    The recursion tree expands in form of a complete binary tree, splitting the list into two halves recursively. The number of levels in a complete binary tree is given by log⁡2n. For n=16, number of splits = log2 16=4

    Merge

    At each level, we merge n nodes which takes O(n) time. For n=16, we perform merge operation on 16 nodes in each of the 4 levels.

    So the time complexity for split and merge operation is O(nlog⁡n)

- Space Complexity. O(logn) , where nnn is the number of nodes in linked list. Since the problem is recursive, we need additional space to store the recursive call stack. The maximum depth of the recursion tree is log n

---

## Delete the middle Node of a Linked List

You are given the head of a linked list. Delete the middle node, and return the middle node of the modified linked list.

The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing.

For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.

```plain
Example 1:

Input: head = [1,3,4,7,1,2,6]
Output: [1,3,4,1,2,6]
Explanation:
The above figure represents the given linked list. The indices of the nodes are written below.
Since n = 7, node 3 with value 7 is the middle node, which is marked in red.
We return the new list after removing this node. 
```

```plain
Example 2:

Input: head = [1,2,3,4]
Output: [1,2,4]
Explanation:
The above figure represents the given linked list.
For n = 4, node 2 with value 3 is the middle node, which is marked in red.
```

```plain
Input: head = [2,1]
Output: [2]
Explanation:
The above figure represents the given linked list.
For n = 2, node 1 with value 1 is the middle node, which is marked in red.
Node 0 with value 2 is the only node remaining after removing node 1.
```

### Solution(s)

#### Approach 1: 2 Passes

##### Intuition

Let's start with a simple approach.

We make the first iteration, starting from head, going through the entire linked list and getting the total number of
nodes (let's say `node_count`). According to the definition provided, the index of the middle node is:
`middleIndex=floor(node_count / 2)−1`.

Now we make a second iteration to the `predecessor` node of the middle node, which means that we stop at index
`middleIndex - 1`.

Once we reach the predecessor node of the middle node, we can remove the middle node from the linked list.

##### Algorithm

1. If there is only one node, return None.
2. Otherwise, initialize two pointers p1 = head and p2 = head.
3. Iterate the linked list with p1 and count the total number of nodes it has (count).
4. Let p2 move forward by floor(node_count / 2) - 1 nodes, now it is the predecessor of the middle node, delete the
   middle node. Return the middle node.

```go
func DeleteMiddle[T comparable](head *list.Node[T]) *list.Node[T] {
 if head == nil || head.Next == nil {
  return nil
 }

 pointer := head
 current := head
 nodeCount := 0

 for pointer != nil {
    pointer = pointer.Next
    nodeCount++
 }

 middleIndex := nodeCount / 2

 for i := 0; i < middleIndex-1; i++ {
  current = current.Next
 }

 middleNode := current.Next
 current.Next = current.Next.Next

 return middleNode
}
```

> Note that this has been implemented in the [SinglyLinkedList](./linkedlist.go) class but with slight
> modifications

##### Complexity Analysis

Let `n` be the length of the input linked list.

1. `Time complexity: O(n)`

    - We iterate over the linked list twice, the first time traversing the entire linked list and the second traversing
      half of it. Hence there are a total of `O(n)` steps.
    - In each step, we move a pointer forward by one node, which takes constant time.
    - Remove the middle node takes a constant amount of time.
    - In summary, the overall time complexity is `O(n)`.

2. `Space complexity: O(1)`

    - We only need two pointers, thus the space complexity is `O(1)`.

#### Approach 2: 2 Pointers

##### Intuition

The key of this approach is that we have two pointers `fast` and `slow` traversing the linked list at the same time, and
`fast` traverses twice as fast as `slow`. Therefore, when `fast` reaches the end of the linked list, `slow` is right in
the middle! We only need one iteration to reach the middle node!

All that needs to be determined are a few lookup details. If there is only one node in the linked list, this node is
also the one to be deleted and there are no nodes left after the deletion. Therefore, instead of initializing two
pointers for the following procedure, we can just return null.

> Why we initialize fast = head.next.next at the beginning?

The reason for this is that we want to delete the middle node instead of finding it. Therefore, we are actually looking
for the `predecessor` of the middle node, not the middle node itself, or rather, this is like moving slow backward one
node after the iteration stops.

Certainly, we can't move a pointer backward on a singly linked list, thus we can show this one less step on `slow` by
letting `fast` moves forward one more step (by two nodes, of course). Hence, `slow` will also point to the predecessor
node of the middle node (rather than the middle node) at the end of the iteration.

##### Algorithm

1. If there is only one node, return null.
2. Otherwise, initialize two pointers `slow` and `fast`, with slow pointing to head and fast pointing to the second
   successor node of head.
3. While neither `fast` and `fast.next` is `null`:
    - we move `fast` forward by 2 nodes.
    - we move `slow` forward by 1 node.
4. Now `slow` is the predecessor of the middle node, delete the middle node.
5. Return middle node.

``` go
func DeleteMiddle2Pointers[T comparable](head *list.Node[T]) *list.Node[T] {
 if head == nil || head.Next == nil {
  return nil
 }

 slowPointer, fastPointer := head, head.Next.Next

 for fastPointer != nil && fastPointer.Next != nil {
  slowPointer = slowPointer.Next
  fastPointer = fastPointer.Next.Next
 }

 middleNode := slowPointer.Next
 slowPointer.Next = slowPointer.Next.Next

 return middleNode
}
```

> Note that the implementation outlined here is what is denoted in the instance method `delete_middle_node_2_pointers`
> of the class [SinglyLinkedList](SinglyLinkedList.kt) with some slight modifications

##### Complexity Analysis

1. `Time Complexity O(n)`
    - We stop the iteration when the pointer fast reaches the end, fast moves forward 2 nodes per step, so there are at
      most `n/2` steps.
    - In each step, we move both `fast` and `slow`, which takes a constant amount of time.
    - Removing the middle node also takes constant time.
    - In summary, the overall time complexity is `O(n)`.

2. `Space Complexity O(1)`

    - We only need 2 pointers, so the space complexity is `O(1)`

---

### Related Topics

- Linked List
- Two Pointers

---
