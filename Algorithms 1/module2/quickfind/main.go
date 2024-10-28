package main

import "fmt"

type UnionFind struct {
	arr []int
}

func NewUnionFind(size int) *UnionFind {
	newArr := make([]int, size)
	for i := 0; i < size; i++ {
		newArr[i] = i
	}
	return &UnionFind{arr: newArr}
}

func (u *UnionFind) union(p int, q int)  {
	pid := u.arr[p]
	qid := u.arr[q]
	for i := 0; i < len(u.arr); i++ {
		if u.arr[i] == pid {
			u.arr[i] = qid
		}
	}
}

func (u *UnionFind) connected(p int, q int) bool {
	return u.arr[p] == u.arr[q]
}

func main() {
	unionFind := NewUnionFind(10)
	fmt.Println("Inital ", unionFind.arr)  
	unionFind.union(0,5)
	unionFind.union(0,6)
	unionFind.union(1,2)
	unionFind.union(1,7)
	unionFind.union(3,4)
	unionFind.union(8,3)
	unionFind.union(4,9)
	fmt.Println("test 1:", unionFind.connected(0,6))
	fmt.Println("test 2:", unionFind.connected(1,7))  
	fmt.Println("test 3:", unionFind.connected(4,3))  
	fmt.Println("test 4:", unionFind.connected(9,7))  
}
