package ch5

// 一个节点的度degree是指该节点的子树个数
// 树的度是树中所有节点的度的最大值。
// 度数为0的节点称为叶子节点
// 左儿子右兄弟树存储模式

// 在二叉树中，第i层的节点数最多为2^(i-1)
// 在深度为k的二叉树中，节点总数最多为2^k-1
// 对任何非空的二叉树T，如果叶节点的个数为n0,而度为2的节点数为n2,则n0=n2+1
// 证明，设度数为1的节点数为n1,总结点数为n 则n = n0+n1+n2。分支数为B，除了根节点，其他每一个节点对应一条分支，有 n=B+1。对于度数：B = n1+2*n2
// 		n0+n1+n2=n1+2*n2 => n0 = n2 + 1
// 满二叉树：深度为k的满二叉树是具有2^k-1个节点的二叉树
// 完全二叉树：深度为k有n个节点的完全二叉树的节点序号必须与深度为k的满二叉树的节点标号1到n相对应
// 遍历 inorder中序遍历、postorder后续遍历【节点的儿子将在节点之前输出】和preorder前序遍历
