# Jump Game

## Version IV

Given an array of integers arr, you are initially positioned at the first index of the array.

In one step you can jump from index i to index:

i + 1 where: i + 1 < arr.length. i - 1 where: i - 1 >= 0. j where: arr[i] == arr[j] and i != j. Return the minimum
number of steps to reach the last index of the array.

Notice that you can not jump outside of the array at any time.

Example 1:

```plain
Input: arr = [100,-23,-23,404,100,23,23,23,3,404]
Output: 3 Explanation: You need three jumps from index 0 --> 4 --> 3 --> 9. Note that index 9 is the last index of the
array. 
```

Example 2:

```plain
Input: arr = [7]
Output: 0 Explanation: Start index is the last index. You do not need to jump. 
```

Example 3:

```plain
Input: arr = [7,6,9,6,9,6,9,7]
Output: 1 Explanation: You can jump directly from index 0 to index 7 which is last index of the array.
```

### Solution

Most solutions start from a brute force approach and are optimized by removing unnecessary calculations. Same as this
one.

A naive brute force approach is to iterate all the possible routes and check if there is one reaches the last index.
However, if we already checked one index, we do not need to check it again. We can mark the index as visited by storing
them in a visited set.

From convenience, we can store nodes with the same value together in a graph dictionary. With this method, when
searching, we do not need to iterate the whole list to find the nodes with the same value as the next steps, but only
need to ask the precomputed dictionary. However, to prevent stepping back, we need to clear the dictionary after we get
to that value.

#### Complexity Analysis

Assume NN is the length of arr.

Time complexity: O(N) since we will visit every node at most once. Space complexity: O(N) since it needs curs and nex to
store nodes.
