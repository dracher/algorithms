# 关于avl
每个节点维护一个h,即该节点的高度，左右高度差在[-1,1]之间（左减右）
## 插入和删除
如果超过高度差范围，自底向上做旋转
如果为正，即左边比右边多，做右旋转，如果为负，即右边比左边多，做左旋转

# 关于treap
同时具有bst和最小堆的性质，key满足bst性质，priority满足最小堆性质。
## 插入和删除
priority破坏最小堆性质后，按照节点priority自底向上回复最小堆性质
如果为左孩子，左旋转，如果为右孩子，右旋转，直到满足最小堆性质

# 关于MAP和hashtable
二者都可以实现key-value的存贮
hashtable是线性数据结构，不可排序，插入删除查找时间复杂度O(1)，但空间不定，需要参数capacity，要至少大于n，比例达到一定程度要扩大
capacity，扩容代价大，要copy之前的键值对，要重新计算hash。
map一般用rbtree实现，可排序，插入删除查找时间复杂度O(lgn)，（最大高度2lg(n+1)），空间复杂度为lg(n)，没有扩容的问题，不需要计算hash

# 关于reduced-space van Emde Boas Tree
用list存储value,支持重复的key。在Insert操作考虑重复key的情况。

用hashtable 存储 summary和cluster，减小存储空间。cluster的创建采用lazy方式，在Insert和Delete操作时考虑cluster的创建。

对于u，直接存储bit数，即lgu，这样的话，最小单元为lgu=1的单元，操作可以转化为位操作。

对于key的类型问题，通过mixin的方式，剥离和key的类型相关的操作。

[代码](https://github.com/shady831213/algorithms/blob/master/tree/vEBTree/rsVEBTree.go)

[测试](https://github.com/shady831213/algorithms/blob/master/tree/vEBTree/rsVEBTree_test.go)

# 关于disjointSet

## 离线最小值问题

### 问题
![](https://github.com/shady831213/algorithms/blob/master/tree/statics/offlineMinimum.PNG)
![](https://github.com/shady831213/algorithms/blob/master/tree/statics/offlineMinimum1.PNG)

[代码](https://github.com/shady831213/algorithms/blob/master/tree/disjointSetTree/offLineMinimum.go)

[测试](https://github.com/shady831213/algorithms/blob/master/tree/disjointSetTree/offLineMinimum_test.go)

--------
### 复杂度
  初始化: n + m个MakeSet + m个Union = O(2m+n)
  
  offLineMinimum: n 个 FindSet + sigma(k)(k = 1...m)个Union = O(n + m^2)
