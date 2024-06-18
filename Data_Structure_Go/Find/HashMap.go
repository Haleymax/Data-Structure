package main

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
	"math"
	"sync"
)

//长度等于 2^x

const expand = 0.75 //扩容因子，当哈希表中百分之七十五的元素都被占用时就需要进行扩容

// 哈希表
type HashMap struct {
	array        []*keyPairs //哈希表中的键值对数组
	capacity     int         //数组容量
	len          int         //键值对元素的数量
	capacityMask int         //掩码参与与运算计算位置(capacity - 1)
	lock         sync.Mutex
}

// 键值对,采用链表形式
type keyPairs struct {
	key   string
	value interface{}
	next  *keyPairs
}

// 初始化哈希表
func NewHashMap(capactity int) *HashMap {
	defaultCapacity := 1 << 4 //将1向右移动四位  二进制 ： 10000
	if capactity <= defaultCapacity {
		capactity = defaultCapacity
	} else {
		capactity = 1 << (int(math.Ceil(math.Log(float64(capactity))))) //取容量的对数（ln e）再取整右移   因为表长为 2^x次方，容量就需要求对数
	}

	m := new(HashMap)
	m.array = make([]*keyPairs, capactity, capactity)
	m.capacity = capactity
	m.capacityMask = capactity - 1
	return m
}

// 求哈希值(匿名函数)
var hashAlgorithm = func(key []byte) uint64 {
	//创建哈希对象
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64() //返回哈希值
}

// 通过哈希值求哈表的下标
func (m *HashMap) hashIndex(key string, mask int) int {

	hash := hashAlgorithm([]byte(key)) //哈希值

	index := hash & uint64(mask) //进行的是与运算，求出哈希表中的坐标（二进制）
	return int(index)            //返回整数坐标
}

// 添加键值对
func (m *HashMap) Put(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 计算哈希表中的坐标
	index := m.hashIndex(key, m.capacityMask) // 掩码就是容量-1

	// 检查插入位置的元素
	element := m.array[index]

	// 元素为空不冲突直接插入
	if element == nil {
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		// 检查遍历后续链表中是否存在空位
		var last *keyPairs
		for element != nil {
			if element.key == key {
				element.value = value // 哈希表中不允许存在相同的键（键肯定是唯一的）
				return
			}
			last = element
			element = element.next
		}
		// 插入到最后这个键值对后面
		last.next = &keyPairs{
			key:   key,
			value: value,
		}
	}

	// 更新哈希表
	newLen := m.len + 1

	// 检查新的哈希表长度是否超出扩容因子的范围
	if float64(newLen)/float64(m.capacity) >= expand {
		// 创建新的哈希表
		newM := new(HashMap)
		newM.array = make([]*keyPairs, 2*m.capacity, 2*m.capacity)
		newM.capacity = 2 * m.capacity
		newM.capacityMask = newM.capacity - 1

		// 遍历旧哈希表将元素转移过来
		for _, pairs := range m.array {
			for pairs != nil {
				// 直接递归 Put
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		// 更新当前哈希表的信息
		m.array = newM.array
		m.capacity = newM.capacity
		m.capacityMask = newM.capacityMask
		m.len = newLen
	}
}

// 哈希表获取键值对
func (m *HashMap) Get(key string) (value interface{}, ret bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	//查找健的位置
	index := m.hashIndex(key, m.capacityMask)

	//遍历哈希表检查是否存在这个键若存在则返回
	element := m.array[index]

	for element != nil {
		if element.key == key {
			return element.value, true
		}

		element = element.next
	}
	return nil, false
}

// 通过键删除哈希表中的元素
func (m *HashMap) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capacityMask) //求出删除键的索引

	element := m.array[index]

	if element == nil {
		return //如果index位置没有元素代表该表中不存在对应的键
	}

	//直接找到
	if element.key == key {
		m.array[index] = element.next //链接后面的键
		m.len = m.len - 1
		return
	}

	//直接查找index的位置不是对应的key(发生过冲突)
	nextElement := element.next //需要用element记录删除前的元素防止断链
	for nextElement != nil {
		if nextElement.key == key {
			element.next = nextElement.next
			m.len = m.len - 1
			return
		}
		element = nextElement
		nextElement = nextElement.next
	}
}

// 哈希表遍历
func (m *HashMap) Range() {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, pairs := range m.array {
		for pairs != nil {
			fmt.Printf("'%v'='%v',", pairs.key, pairs.value)
			pairs = pairs.next
		}
	}
	fmt.Println()
}

func Hash() {
	// 新建一个哈希表
	hashMap := NewHashMap(16)
	// 放35个值
	for i := 0; i < 35; i++ {
		hashMap.Put(fmt.Sprintf("%d", i), fmt.Sprintf("v%d", i))
	}
	fmt.Println("cap:", hashMap.capacity, "len:", hashMap.len)
	// 打印全部键值对
	hashMap.Range()
	key := "4"
	value, ok := hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}
	// 删除键
	hashMap.Delete(key)
	fmt.Println("after delete cap:", hashMap.capacity, "len:", hashMap.len)
	value, ok = hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}
}
