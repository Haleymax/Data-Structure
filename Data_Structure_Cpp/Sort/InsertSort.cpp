//
// Created by huang on 2024/6/18.
//

#include "iostream"
#include "vector"
using namespace std;

//插入排序

void InsertSort(vector<int> &nums){
    int n = nums.size();
    for (int i = 0; i < n; ++i) {
        int key = nums[i];  //选取后面的元素找到他的最终位置

        int j = i -1;  //遍历范围
        while (j >= 0 && nums[j] > key){
            nums[j + 1] = nums[j];  //大的数向后移动直到找到位置为止
            j = j -1;
        }
        nums[j + 1] = key;  //找到他的位置
    }
}

int main(){
    vector<int> nums;
    nums.push_back(8);
    nums.push_back(4);
    nums.push_back(1);
    nums.push_back(2);
    nums.push_back(5);
    nums.push_back(3);
    nums.push_back(9);

    cout << "元素数据为 : ";
    for (int i = 0; i < nums.size(); ++i) {
        cout << nums[i] << " ";
    }
    cout << endl;
    InsertSort(nums);
    cout << "冒泡排序后的结果 : " ;
    //使用迭代器来遍历
    vector<int>::iterator it = nums.begin();
    while (it != nums.end()){
        cout << *it << " ";
        it++;
    }
    cout << endl;
}