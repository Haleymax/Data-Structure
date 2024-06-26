//
// Created by huang on 2024/6/26.
//

#include "iostream"
#include "vector"

//定义哈希表大小
const int TABLE_SIZE = 128;

using namespace std;

//哈希集合类
class MyHashSet{
private:

    //哈希表
    vector<int> table;

    //记录元素数量
    int size;

    //求元素索引
    int hashFunction(int key){
        return key % TABLE_SIZE;
    }

public:
    //构造函数，初始化哈希表
    MyHashSet() : table(TABLE_SIZE,-1) , size(0){}

    //插入元素
    void insert(int key);

    //删除元素
    void remove(int key);

    //查找元素是否存在
    bool contains(int key);

    //获取当前集合的大小
    int getSize() const;
};

//插入元素
void MyHashSet::insert(int key) {
    int hash = hashFunction(key);   //获取位置索引
    int originalHash = hash;   //记录计算出的位置

    //检查冲突，采用线性探测处理冲突
    while (table[hash] != -1 && table[hash] != key){
        hash = (hash + 1) % TABLE_SIZE;  //发生冲突就往后移动一位

        //采用不可扩展的哈希表，如果满了就不能够插入了
        if (hash == originalHash){
            cout << "哈希表以满，无法插入新元素" << endl;
            return;
        }
    }

    //如果找到空位置就将元素插入
    if (table[hash] == -1){
        table[hash] = key;
        size++;
    }
}

//删除元素
void MyHashSet::remove(int key) {
    int hash = hashFunction(key);
    int orignalHash = hash;

    //线性探测查找元素
    while (table[hash] != -1){
        if (table[hash] == key){
            //找到元素，删除
            table[hash] = -1;
            size--;
            return;
        }
        hash = (hash + 1) % TABLE_SIZE;

        //如果探测回到原来的位置代表整个表以及满了以及没有这个元素
        if (hash == orignalHash)
            return;
    }
}

//查找元素是否存在
bool MyHashSet::contains(int key) {
    int hash = hashFunction(key);
    int originalHash = hash;

    //线性探测法
    while (table[hash] != -1){
        if (table[hash] == key)
            return true;

        hash = (hash + 1) % TABLE_SIZE;

        //如果探测回原来位置表示不存在
        if (hash == originalHash)
            return false;
    }

    return false;
}

int MyHashSet::getSize() const {
    return size;
}

// 测试自定义哈希集合
int main() {
    MyHashSet mySet;
    mySet.insert(1);
    mySet.insert(2);
    mySet.insert(3);

    std::cout << "集合中是否包含2: " << (mySet.contains(2) ? "是" : "否") << std::endl;
    std::cout << "集合大小: " << mySet.getSize() << std::endl;

    mySet.remove(2);
    std::cout << "删除2后，集合中是否包含2: " << (mySet.contains(2) ? "是" : "否") << std::endl;
    std::cout << "集合大小: " << mySet.getSize() << std::endl;

    return 0;
}