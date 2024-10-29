package main

type QuickUnion struct {
	arr [] int
	sz [] int
}

func NewQuickUnion (size int) *QuickUnion {
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return &QuickUnion{arr: arr, sz: make([]int, size)}
}

func (u *QuickUnion) connected (q int, p int) bool {
	return u.root(q) == u.root(p)
} 

func (u *QuickUnion) union (q int, p int) {
	rootQ := u.root(q)
	rootP := u.root(p)
	if rootQ == rootP {
		return
	}
	if u.sz[p] > u.sz[q] {
		u.arr[rootP] = rootQ
		u.sz[p] = u.sz[q];
	}else{
		u.arr[rootQ] = rootP
		u.sz[q] = u.sz[p];
	}
}

func (u *QuickUnion) root (i int) int {
	for i != u.arr[i] {
		u.arr[i] = u.arr[u.arr[i]] // path compression
		i = u.arr[i]
	}
	return i
}

func main(){
	quickUnion := NewQuickUnion(10)
	quickUnion.union(4,3)
	quickUnion.union(3,8)
	quickUnion.union(6,5)
	quickUnion.union(9,4)
	quickUnion.union(2,1)
	quickUnion.union(5,0)
	quickUnion.union(7,2)
	quickUnion.union(6,1)
	quickUnion.union(7,3)
	
	quickUnion.connected(5,4)
	quickUnion.connected(8,9)
}