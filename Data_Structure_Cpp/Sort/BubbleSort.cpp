//
// Created by huang on 2024/6/18.
//

//ð�������ʵ��
#include "iostream"
#include "vector"

using namespace std;


//ð���������
void BubbleSort(vector<int> &num){
    int n = num.size();
    for (int i = 0; i < n - 1; ++i) {
        for (int j = 0; j < n - i -1; ++j) {
            if (num[j] > num[j + 1])
                swap(num[j],num[j+1]);
        }
    }
}

int main(){
    vector<int> nums = {6,3,5,7,4,9,1};
    cout << "Ԫ������Ϊ : ";
    for (int i = 0; i < nums.size(); ++i) {
        cout << nums[i] << " ";
    }
    cout << endl;
    BubbleSort(nums);
    cout << "ð�������Ľ�� : " ;
    //ʹ�õ�����������
    vector<int>::iterator it = nums.begin();
    while (it != nums.end()){
        cout << *it << " ";
        it++;
    }
    cout << endl;
}
