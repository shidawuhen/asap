package algorithm
/*
原题：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof
剑指 Offer 36. 二叉搜索树与双向链表
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。



为了让您更好地理解问题，以下面的二叉搜索树为例：







我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。







特别地，我们希望可以就地完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。

分析：
幸亏还记得C++代码怎么写
总体上使用深度优先遍历，中心节点需要连接左子树的最右节点，连接右子树的最左节点
*/

/*
class Solution {
public:
    Node* treeToDoublyList(Node* root) {
        if(root == NULL){
            return NULL;
        }
        change(root);
        while(root->left != NULL){
           // cout<<root->val<<endl;
            root = root->left;
        }
        Node* head = root;
        while(root->right != NULL){
            root = root->right;
        }
        root->right = head;
        head->left = root;
        return head;
    }
    Node* change(Node* root){
        if(root == NULL){
            return NULL;
        }
       // cout<<root->val<<endl;
        Node* left = NULL;
        Node* right = NULL;
        if (root->left != NULL){
            left = change(root->left);
            //cout<<left->val<<endl;
            while(left->right != NULL){
                left = left->right;
            }
            root->left = left;
            left->right = root;
        }
        //cout<<root->val<<endl;
        if(root->right != NULL){
            right = change(root->right);
            while(right->left != NULL){
                right = right->left;
            }
            root->right = right;
                right->left = root;
            //cout<<right->val<<endl;
        }
        return root;
    }
};
*/
