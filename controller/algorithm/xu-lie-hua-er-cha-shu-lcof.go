package algorithm
/*
原题：https://leetcode-cn.com/problems/xu-lie-hua-er-cha-shu-lcof
剑指 Offer 37. 序列化二叉树
请实现两个函数，分别用来序列化和反序列化二叉树。

示例:

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"

分析：
主要是重组的时候，左右子节点位置需要计算准确
*/
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
/*
class Codec {
public:

    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        string s = "";
        if(root == NULL){
            return s;
        }
        //广度优先遍历
        queue<TreeNode*>  q;
        q.push(root);
        int larget = 3;
        int index = 0;
        while(!q.empty() && index <= larget){
            index++;
            TreeNode* temp = q.front();
            if(temp != NULL){
                q.push(temp->left);
                q.push(temp->right);
                if(temp->left != NULL || temp->right != NULL){
                    larget = 2*index+2;
                }else{
                    larget++;
                }
                //char c[8];
                //int length = sprintf(c, "%d", temp->val);
               // s.push_back(std::to_string(temp->val));
                s.append(to_string(temp->val));
                s.append(",");
            }else{
                s.append(" ,");
            }

            q.pop();
        }
        //cout<<s<<endl;
        return s;
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        vector<string> v = split(data,",");
        if(v.size()==0){
            return NULL;
        }
        vector<string>::iterator it;
        string strArray[v.size()];
        TreeNode* tree[v.size()];
        int i = 0;
        for(it=v.begin();it!=v.end();it++){
            strArray[i] = *it;
            if(*it != " "){
                int intStr = atoi(strArray[i].c_str());
                tree[i] = new TreeNode(intStr);
            }else{
                tree[i] = NULL;
            }
            i++;
        }
        int j = 1;
        for (i = 0; i < v.size();i++){
            if(tree[i] == NULL){
                continue;
            }
            if(j < v.size()){
                tree[i]->left = tree[j++];
            }
            if(j < v.size()){
                tree[i]->right = tree[j++];
            }
        }
        return tree[0];
    }


    vector<string> split(const string& str, const string& delim) {
	vector<string> res;
	if("" == str) return res;
	//先将要切割的字符串从string类型转换为char*类型
	char * strs = new char[str.length() + 1] ; //不要忘了
	strcpy(strs, str.c_str());

	char * d = new char[delim.length() + 1];
	strcpy(d, delim.c_str());

	char *p = strtok(strs, d);
	while(p) {
		string s = p; //分割得到的字符串转换为string类型
		res.push_back(s); //存入结果数组
		p = strtok(NULL, d);
	}

	return res;
    }
};
*/