# LeetCode Go代码

【易】：简单，能做出；【中】：需要思考做出；【难】：做不出；

## 经典题目

## 题型分类

### 单调队列

* [239. 滑动窗口最大值](leetcode/239.go)【难，重要】

### 二叉树

* [919. 完全二叉树插入器](leetcode/919.go)【中】
* [124. 二叉树中的最大路径和](leetcode/124.go)【中】
* [1206. 设计跳表](leetcode/1206.go)【中，重要】
* [1161. 最大层内元素和](leetcode/1161.go)【中】
* [623. 在二叉树中增加一行](leetcode/623.go)【易】
* [1302. 层数最深叶子节点的和](leetcode/1302.go)【易】
* [6163. 给定条件下构造矩阵](leetcode/6163.go)【难，拓扑排序，重要】
* [662. 二叉树最大宽度](leetcode/662.go)【难，重要】
* [1372. 二叉树中的最长交错路径](leetcode/1372.go)【中，重要】

### 前缀树

* [6183. 字符串的前缀分数和](leetcode/6186.go)【中，重要】

### 分数、最大公倍数、最小公约数

* [592. 分数加减运算](leetcode/592.go)【易】

### 并查集

* [952. 按公因数计算最大组件大小](leetcode/952.go)【中】
* [6191. 好路径的数目](leetcode/6191.go)【难】

### 图

* [6134. 找到离给定两个节点最近的节点](leetcode/6134.go)【中，基环树】
* [6135. 图中的最长环](leetcode/6135.go)【中，基环树】
* [6235. 逐层排序二叉树所需的最少操作数目](leetcode/6235.go)【难，环图】
* [753. 破解保险箱](leetcode/753.go)【难，深度优先搜索】

### 基本数据结构

* [622. 设计循环队列](leetcode/622.go)

### 贪心

* [6144. 将数组排序的最少替换次数](leetcode/6144.go)【易】
* [768. 最多能完成排序的块 II](leetcode/768.go)【中，单调栈】
* [32. 最长有效括号](leetcode/32.go)【中+，重要】
* [76. 最小覆盖子串](leetcode/76.go)【中】
* [658. 找到 K 个最接近的元素](leetcode/658.go)【中】
* [646. 最长数对链](leetcode/646.go)【中，重要】

### 动态规划

* [6137. 检查数组是否存在有效划分](leetcode/6138.go)【中】
* [6138. 最长理想子序列](leetcode/6138.go)【中】
* [801. 使序列递增的最小交换次数](leetcode/801.go)【难】
* [813. 最大平均值和的分组](leetcode/813.go)【难】

### 分治

* [761. 特殊的二进制序列](leetcode/761.go)【难，重要】

### 哈希

* [1224. 最大相等频率](leetcode/1224.go)【难，重要】

### 数论

* [172. 阶乘后的零](leetcode/172.go)【难，重要】
* [793. 阶乘函数后 K 个零](leetcode/793.go)【难，重要】
* [1359. 有效的快递序列数目](leetcode/1359.go)【难，重要】

### 桶

* [1370. 上升下降字符串](leetcode/1370.go)【易，重要】

### 状态压缩

* [1371. 每个元音包含偶数次的最长子字符串](leetcode/1371.go)【难，重要】
* [698. 划分为k个相等的子集](leetcode/698.go)【难，重要】

### 其他

* [667. 优美的排列 II](leetcode/667.go)【中，重要】
* [850. 矩形面积 II](leetcode/850.go)【难，重要，扫描线，线段树】
* [6207. 统计定界子数组的数目](leetcode/6207.go)【难，双指针】




## 知识点总结

1. 遇到**nums[i]-nums[j] == i-j**类似的等式可以转换成**nums[i] - i == nums[j] - j**，然后令数组**a[i] = nums[i] - i**即可，例如题目[6142. 统计坏数对的数目](leetcode/6142.go)；
1. 如果需要同时记录**元素值**和**该值在数组中的位置**，可以直接记录**位置**，这样就可以通过**位置**找到**该元素**，例题：[239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)
