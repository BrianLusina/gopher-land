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
