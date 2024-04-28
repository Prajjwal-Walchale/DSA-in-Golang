Problem statement
A doubly-linked list is a data structure that consists of sequentially linked nodes, and the nodes have reference to both the previous and the next nodes in the sequence of nodes.

Given a doubly-linked list, delete the node at the end of the doubly linked list.

Note:
You need not print anything. You’re given the head of the linked list, just return the head of the modified list.

Example:
Input: Linked List:  4 <-> 10 <-> 3 <-> 5 <-> 20

Output: Modified Linked List: 4 <-> 10 <-> 3 <-> 5

Explanation: The last node having ‘data’ = 20 is removed from the linked list.
Detailed explanation ( Input/output format, Notes, Images )
Sample Input 1:
5
4 10 3 5 20

Sample Output 1:
4 10 3 5 NULL

Explanation Of Sample Input 1 :
The last node having ‘data’ = 20 is removed from the linked list.

Sample Input 2 :
1
5

Sample Output 2 :
NULL

Explanation Of Sample Input 2 :
The linked list has only one node, so the modified linked list is empty.

Expected time complexity :
The expected time complexity is O(N).