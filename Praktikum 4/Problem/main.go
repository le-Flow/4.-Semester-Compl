package main

//implementing philospher's dining problem
//the problem is not solved to simulate the behavior of phylosophers

import (
	"fmt"
	"sync"
	"time"
)

const N = 1000000

type Fork struct {
	id int
	sync.Mutex
}

func (f *Fork) Lock() {
	f.Mutex.Lock()
}

func (f *Fork) Unlock() {
	f.Mutex.Unlock()
}

type Phylosopher struct {
	id int
}

func (p Phylosopher) Eat(rFork *Fork, lFork *Fork) {
	rFork.Lock()
	lFork.Lock()

	fmt.Printf("Phylosopher %d is eating\n", p.id)

	//simulating eating
	time.Sleep(10 * time.Second)

	rFork.Unlock()
	lFork.Unlock()
	fmt.Printf("Phylosopher %d is thinking\n", p.id)
}

type Table struct {
	forks        [N]Fork
	phylosophers [N]Phylosopher
}

// create constructor for Table
func NewTable() *Table {
	table := Table{}
	for i := 0; i < N; i++ {
		table.phylosophers[i] = Phylosopher{id: i}
	}

	for i := 0; i < N; i++ {
		table.forks[i] = Fork{id: i}
	}

	return &table
}

func (t *Table) StartDinner() {

	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 0; i < N; i++ {
		wg.Add(1)
		//anonymous function to simulate phylosopher's behavior
		go func(idx int) {
			defer wg.Done()
			t.phylosophers[idx].Eat(t.LeftFork(idx), t.RightFork(idx))
		}(i)

	}
}

func (t *Table) LeftFork(idx int) *Fork {
	return &t.forks[idx]
}

func (t *Table) RightFork(idx int) *Fork {
	return &t.forks[(idx+1)%N]
}

func main() {
	table := NewTable()
	table.StartDinner()
}
