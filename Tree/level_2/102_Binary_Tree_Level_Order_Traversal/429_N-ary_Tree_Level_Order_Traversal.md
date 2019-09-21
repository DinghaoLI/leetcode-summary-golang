# 429. N-ary Tree Level Order Traversal

Given an n-ary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

[429. N-ary Tree Level Order Traversal](https://leetcode.com/problems/n-ary-tree-level-order-traversal/)

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
    vector<vector<int>> levelOrder(Node* root) {
        vector<vector<int>> result;
        std::function<void (Node* root, int depth)> helper;
        helper = [&helper, &result](Node* root, int depth) -> void {
            if (!root) { return; }
            if (depth > result.size()) {
                vector<int> tmp;
                tmp.push_back(root->val);
                result.push_back(tmp);
            } else {
                result[depth-1].push_back(root->val);
            }
            for (auto chil : root->children) {
                helper(chil, depth+1);
            }
        };
        helper(root, 1);
        return result;
    }
};
```
