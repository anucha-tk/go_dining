package channel

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

type Philosopher struct {
	id int
}

const numOfPhilosopher = 5

func (p Philosopher) think() {
	color.Blue("Philosopher #%d is thinking...", p.id)
	time.Sleep(time.Second)
}

func (p Philosopher) eat(leftFork, rightFork chan struct{}) {
	<-leftFork
	<-rightFork
	color.Red("Philosopher #%d is eating...", p.id)
	time.Sleep(time.Second)
	rightFork <- struct{}{}
	leftFork <- struct{}{}
}

func (p Philosopher) dine(wg *sync.WaitGroup, leftFork, rightFork chan struct{}) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.think()
		p.eat(leftFork, rightFork)
	}
}

func Channel_version() {
	color.Cyan("Channel version")
	color.Cyan("---------------")
	wg := sync.WaitGroup{}

	// create chan fork
	forks := make([]chan struct{}, numOfPhilosopher)
	for i := 0; i < numOfPhilosopher; i++ {
		forks[i] = make(chan struct{}, 1)
		forks[i] <- struct{}{}
	}

	// create Philosophers
	philosohers := make([]Philosopher, numOfPhilosopher)
	for i := 0; i < numOfPhilosopher; i++ {
		philosohers[i] = Philosopher{
			id: i + 1,
		}
	}

	// loop philosiohers for dinner
	for i := 0; i < numOfPhilosopher; i++ {
		wg.Add(1)
		leftFork := forks[i]
		rightFork := forks[(i+1)%numOfPhilosopher]
		go philosohers[i].dine(&wg, leftFork, rightFork)
	}

	wg.Wait()
	color.Green("Dinner of philosopher is over.")
}
