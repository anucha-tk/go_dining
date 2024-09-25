package mutex

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *sync.Mutex
}

const numOfPhilosopher = 5

func (p Philosopher) think() {
	color.Blue("Philosopher #%d is thinking...", p.id)
	time.Sleep(time.Second)
}

func (p Philosopher) eat() {
	p.leftFork.Lock()
	p.rightFork.Lock()
	color.Red("Philosopher #%d is eating...", p.id)
	time.Sleep(time.Second)
	p.rightFork.Unlock()
	p.leftFork.Unlock()
}

func (p Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.think()
		p.eat()
	}
}

func Mutex_version() {
	color.Cyan("Mutex version")
	color.Cyan("---------------")

	wg := sync.WaitGroup{}

	// create mutex fork
	forks := make([]*sync.Mutex, numOfPhilosopher)
	for i := 0; i < numOfPhilosopher; i++ {
		forks[i] = &sync.Mutex{}
	}

	// create Philosophers
	philosohers := make([]*Philosopher, numOfPhilosopher)
	for i := 0; i < numOfPhilosopher; i++ {
		philosohers[i] = &Philosopher{
			id:        i + 1,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numOfPhilosopher],
		}
	}

	// loop philosiohers for dinner
	for i := 0; i < numOfPhilosopher; i++ {
		wg.Add(1)
		go philosohers[i].dine(&wg)
	}

	wg.Wait()
	color.Green("Dinner of philosopher is over.")
}
