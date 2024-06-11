cmake_minimum_required(VERSION 3.26)
project(Data_Structure_Cpp)

set(CMAKE_CXX_STANDARD 17)

add_executable(Data_Structure_Cpp main.cpp)

add_executable(Array Linklist/Array.cpp)
add_executable(Linklist Linklist/Linklist.cpp)
add_executable(RingLinklist Linklist/RingLinklist.cpp)