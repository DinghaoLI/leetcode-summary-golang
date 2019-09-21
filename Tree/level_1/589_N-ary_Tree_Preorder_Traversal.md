# 589. N-ary Tree Preorder Traversal

Given an n-ary tree, return the preorder traversal of its nodes' values.

[589. N-ary Tree Preorder Traversal](https://leetcode.com/problems/n-ary-tree-preorder-traversal/)

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
    vector<int> preorder(Node* root) {
        std::vector<int> res;
        
        std::function<void (Node* root)> recur;
        recur = [&res, &recur] (Node* root) -> void {
            if (root) {
                res.push_back(root->val);
                for (Node* v : root->children) {
                    recur(v);
                }    
            }    
        };
        recur(root);
        return res;
    }
};
```
s