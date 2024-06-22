//
// Created by huang on 2024/6/22.
//

//堆排序的实现(大堆)
#include "iostream"
#include "vector"

using namespace std;

//交换两个元素
void swap(vector<int>& arr , int i , int j){
    int  temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}

//堆调整
void heapify(vector<int>& arr , int n , int i){
    int largest = i;   //默认数组的第一个节点为最大值
    int left = 2 * i + 1;  //左孩子节点
    int right = 2 * i + 2; //右孩子节点

    //如果左节点要比根节点大
    if (left < n && arr[left] > arr[largest])
        largest = left;

    //如果右子树的节点值还比当前的最大值大
    if (right < n && arr[right] > arr[largest])
        largest = right;

    //如果不是最大值就需要往下进行交换直到找到最大的值为止
    if (largest != i){
        swap(arr , i , largest);

        heapify(arr , n , largest);
    }
}

//构建最大堆
void buildHeap(vector<int>& arr){
    int n = arr.size();

    //从最后一个非叶子节点开始向上遍历
    for (int i = n / 2; i > 0 ; i--) {
        heapify(arr,n,i);
    }
}

//堆排序
void HeapSort(vector<int> &arr){
    //构建最大堆
    buildHeap(arr);
    int n = arr.size();

    //一个从堆顶取出元素
    for (int i = n-1; i >= 0 ; i--) {
        swap(arr,0,i);

        heapify(arr,i,0);
    }
}

int main() {
    vector<int> arr = {12, 11, 13, 5, 6, 7};
    HeapSort(arr);

    for (int i : arr)
        cout << i << " ";
    cout << endl;

    return 0;
}