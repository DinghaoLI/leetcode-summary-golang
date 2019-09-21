# 590. N-ary Tree Postorder Traversal

Given an n-ary tree, return the postorder traversal of its nodes' values.


[590. N-ary Tree Postorder Traversal](https://leetcode.com/problems/n-ary-tree-postorder-traversal/)

We cannot use golang in this problem, so we use c++ instead.

```c++
/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Solution {
public:
    vector<int> postorder(Node* root) {
        std::vector<int> res;
        
        std::function<void (Node* root)> recur;
        recur = [&res, &recur] (Node* root) -> void {
            if (root) {
                for (Node* v : root->children) {
                    recur(v);
                }
                res.push_back(root->val);
            }    
        };
        recur(root);
        return res;
    }
};
```
