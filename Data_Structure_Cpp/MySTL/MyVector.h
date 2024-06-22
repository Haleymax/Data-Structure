//
// Created by huang on 2024/6/22.
//

#ifndef DATA_STRUCTURE_MYVECTOR_H
#define DATA_STRUCTURE_MYVECTOR_H

#endif //DATA_STRUCTURE_MYVECTOR_H

#pragma once

#include "iostream"
#include "algorithm"


using namespace std;

template <class T>
class MyVector{
private:
    T* data;
    size_t size;
    size_t capacity;

public:

    //构造函数
    MyVector() : data(NULL) , size(0) , capacity(0) {}

    //拷贝构造
    MyVector(const MyVector& other);

    //赋值运算符
    MyVector& operator=(const MyVector& other);

    //重载=运算符
    T& operator[](size_t index);

    //析构函数
    ~MyVector(){
        delete[] data;
    }

    //resize 重新指定容量
    void resize(size_t capacity);

    //插入元素到末尾
    void push_back(const T& value);
};


template<class T>
void MyVector<T>::resize(size_t newcapacity){
    if(newcapacity == this->capacity)
        return;

    T* new_data = new T[newcapacity];  //重新分配数组

    if (newcapacity < size){
        size = newcapacity;
    }

    copy(data , data + min(size , newcapacity) , new_data);  //复制新到新到数组

    //将旧数组删除
    delete[] data;

    data = new_data;
    capacity = newcapacity;
}

//拷贝构造
template<class T>
MyVector<T>::MyVector(const MyVector<T> &other) {
    size = other.size;
    capacity = other.capacity;
    data = new T[capacity];
    copy(other.data , other.data + size , data);
}

//重载=运算符
template<class T>
MyVector<T> &MyVector<T>::operator=(const MyVector<T> &other) {
    if (this != &other){
        delete[] data;
        size = other.size;
        capacity = other.capacity;
        data = new T[capacity];
        copy(other.data , other.data + size , data);
    }
    return *this;
}

//重载=号运算符
template<class T>
T &MyVector<T>::operator[](size_t index) {
    return data[index];
}

//插入元素到末尾
template<class T>
void MyVector<T>::push_back(const T &value) {
    if (size == capacity){
        capacity *= 2;
        resize(capacity);
    }
    data[size++] = value;
}