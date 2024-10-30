package main

import (
	"errors"
	"fmt"
)

// 0 = closed
// 1 = open
// 2 = filled
type Percolation struct {
	sites [][]int
}

func NewPercolation (size int) *Percolation {
	sites := make([][]int, size)
	for i := range sites {
		sites[i] = make([]int, size)
	}
	return &Percolation{sites: sites}
}

func (p *Percolation) logsSite(){
	for i := range p.sites {
		fmt.Println(p.sites[i])
	}
}
func (p *Percolation) open (row int, col int) error {
	if row < 1 || row > len(p.sites) || col < 1 || col > len(p.sites) {
		return errors.New("Invalid range")
	}
	p.sites[row-1][col-1] = 1
	return nil
}

func (p *Percolation) isOpen (row int, col int) (bool,  error){
	if row < 1 || row > len(p.sites) || col < 1 || col > len(p.sites) {
		return false, errors.New("Invalid range")
	}
	return p.sites[row-1][col-1] == 1, nil
}

func (p *Percolation) isFull (row int, col int) (bool, error){
	if row < 1 || row > len(p.sites) || col < 1 || col > len(p.sites) {
		return false, errors.New("Invalid range")
	}
	return p.sites[row-1][col-1] == 2, nil
}

func (p *Percolation) numberOfOpenSites () int {
	count := 0
	for i := range p.sites {
		for j := range p.sites[i-1] {
			if p.sites[i-1][j-1] == 1 {
				count++
			}
		}
	}
	return count
}

func (p *Percolation) percolates () {

}

func main() {
	test := NewPercolation(10)
	a, err := test.isFull(1,11)
	if err != nil {
		fmt.Println("yo error", err)
	}
	fmt.Println("result", a)

}