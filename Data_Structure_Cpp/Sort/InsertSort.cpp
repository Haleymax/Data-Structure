//
// Created by huang on 2024/6/18.
//

#include "iostream"
#include "vector"
using namespace std;

//��������

void InsertSort(vector<int> &nums){
    int n = nums.size();
    for (int i = 0; i < n; ++i) {
        int key = nums[i];  //ѡȡ�����Ԫ���ҵ���������λ��

        int j = i -1;  //������Χ
        while (j >= 0 && nums[j] > key){
            nums[j + 1] = nums[j];  //���������ƶ�ֱ���ҵ�λ��Ϊֹ
            j = j -1;
        }
        nums[j + 1] = key;  //�ҵ�����λ��
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

    cout << "Ԫ������Ϊ : ";
    for (int i = 0; i < nums.size(); ++i) {
        cout << nums[i] << " ";
    }
    cout << endl;
    InsertSort(nums);
    cout << "ð�������Ľ�� : " ;
    //ʹ�õ�����������
    vector<int>::iterator it = nums.begin();
    while (it != nums.end()){
        cout << *it << " ";
        it++;
    }
    cout << endl;
}