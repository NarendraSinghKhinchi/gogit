package main

import (
	"fmt"
	"sync"
	"time"
)

// Rank of employee required for a call
type Rank int

const (
	OperatorRank Rank = iota
	SupervisorRank
	DirectorRank
)

func (r Rank) String() string {
	switch r {
	case OperatorRank:
		return "Operator"
	case SupervisorRank:
		return "Supervisor"
	case DirectorRank:
		return "Director"
	default:
		return "Unknown"
	}
}

// CallState enum
type CallState int

const (
	Ready CallState = iota
	InProgress
	Complete
)

// Call Event Structs for Channels
type CallEvent struct {
	Call *Call
}

// Call represents a customer call
type Call struct {
	ID       int
	Rank     Rank
	State    CallState
	Employee Employee
}

// Employee interface allows CallCenter to interact with all roles
type Employee interface {
	GetRank() Rank
	GetName() string
	IsFree() bool
	TakeCall(call *Call)
	CompleteCall()
	EscalateCall()
	GetActiveCall() *Call
}

// BaseEmployee contains shared logic for all employee types (Composition over Inheritance)
type BaseEmployee struct {
	ID             int
	Name           string
	Rank           Rank
	ActiveCall     *Call
	completionChan chan<- CallEvent
	escalationChan chan<- CallEvent
}

func (e *BaseEmployee) GetRank() Rank {
	return e.Rank
}

func (e *BaseEmployee) GetName() string {
	return e.Name
}

func (e *BaseEmployee) IsFree() bool {
	return e.ActiveCall == nil
}

func (e *BaseEmployee) GetActiveCall() *Call {
	return e.ActiveCall
}

func (e *BaseEmployee) TakeCall(call *Call) {
	fmt.Printf("[%s] %s taking call %d\n", e.Rank, e.Name, call.ID)
	e.ActiveCall = call
	call.State = InProgress
}

func (e *BaseEmployee) CompleteCall() {
	if e.ActiveCall != nil {
		e.ActiveCall.State = Complete
		call := e.ActiveCall
		e.ActiveCall = nil
		fmt.Printf("[%s] %s completed call %d\n", e.Rank, e.Name, call.ID)
		e.completionChan <- CallEvent{Call: call}
	}
}

func (e *BaseEmployee) internalEscalate(newRank Rank) {
	if e.ActiveCall != nil {
		e.ActiveCall.State = Ready
		e.ActiveCall.Rank = newRank
		call := e.ActiveCall
		e.ActiveCall = nil
		fmt.Printf("[%s] %s escalated call %d to %s\n", e.Rank, e.Name, call.ID, newRank)
		e.escalationChan <- CallEvent{Call: call}
	}
}

// Operator
type Operator struct {
	BaseEmployee
}

func NewOperator(id int, name string, compChan, escChan chan<- CallEvent) *Operator {
	return &Operator{
		BaseEmployee: BaseEmployee{
			ID:             id,
			Name:           name,
			Rank:           OperatorRank,
			completionChan: compChan,
			escalationChan: escChan,
		},
	}
}

func (o *Operator) EscalateCall() {
	o.internalEscalate(SupervisorRank)
}

// Supervisor
type Supervisor struct {
	BaseEmployee
}

func NewSupervisor(id int, name string, compChan, escChan chan<- CallEvent) *Supervisor {
	return &Supervisor{
		BaseEmployee: BaseEmployee{
			ID:             id,
			Name:           name,
			Rank:           SupervisorRank,
			completionChan: compChan,
			escalationChan: escChan,
		},
	}
}

func (s *Supervisor) EscalateCall() {
	s.internalEscalate(DirectorRank)
}

// Director
type Director struct {
	BaseEmployee
}

func NewDirector(id int, name string, compChan, escChan chan<- CallEvent) *Director {
	return &Director{
		BaseEmployee: BaseEmployee{
			ID:             id,
			Name:           name,
			Rank:           DirectorRank,
			completionChan: compChan,
			escalationChan: escChan,
		},
	}
}

func (d *Director) EscalateCall() {
	fmt.Printf("[Director] %s cannot escalate call %d further. Handling it.\n", d.Name, d.ActiveCall.ID)
	d.CompleteCall()
}

// CallCenter struct holds the employees, queue, and event channels
type CallCenter struct {
	operators      []Employee
	supervisors    []Employee
	directors      []Employee
	queuedCalls    []*Call
	newCallChan    chan *Call
	completionChan chan CallEvent
	escalationChan chan CallEvent
	wg             *sync.WaitGroup
}

func NewCallCenter(wg *sync.WaitGroup) *CallCenter {
	return &CallCenter{
		operators:      make([]Employee, 0),
		supervisors:    make([]Employee, 0),
		directors:      make([]Employee, 0),
		queuedCalls:    make([]*Call, 0),
		newCallChan:    make(chan *Call, 100),
		completionChan: make(chan CallEvent, 100),
		escalationChan: make(chan CallEvent, 100),
		wg:             wg,
	}
}

func (cc *CallCenter) AddOperator(name string) {
	cc.operators = append(cc.operators, NewOperator(len(cc.operators)+1, name, cc.completionChan, cc.escalationChan))
}

func (cc *CallCenter) AddSupervisor(name string) {
	cc.supervisors = append(cc.supervisors, NewSupervisor(len(cc.supervisors)+1, name, cc.completionChan, cc.escalationChan))
}

func (cc *CallCenter) AddDirector(name string) {
	cc.directors = append(cc.directors, NewDirector(len(cc.directors)+1, name, cc.completionChan, cc.escalationChan))
}

func (cc *CallCenter) DispatchNewCall(call *Call) {
	cc.wg.Add(1)
	cc.newCallChan <- call
}

func (cc *CallCenter) dispatch(call *Call) {
	var emp Employee

	// Find the first free employee of the appropriate rank or higher
	if call.Rank == OperatorRank {
		emp = cc.findFree(cc.operators)
	}
	if call.Rank <= SupervisorRank && emp == nil {
		emp = cc.findFree(cc.supervisors)
	}
	if call.Rank <= DirectorRank && emp == nil {
		emp = cc.findFree(cc.directors)
	}

	if emp == nil {
		fmt.Printf("[Queue] Nobody available for call %d (Rank req: %s). Queueing...\n", call.ID, call.Rank)
		cc.queuedCalls = append(cc.queuedCalls, call)
	} else {
		emp.TakeCall(call)
		call.Employee = emp

		// Simulate asynchronous call duration
		go func(e Employee, c *Call) {
			time.Sleep(300 * time.Millisecond) // Simulated handle time

			// Arbitrarily escalate every 3rd call for simulation testing
			if c.ID%3 == 0 && e.GetRank() != DirectorRank {
				e.EscalateCall()
			} else {
				e.CompleteCall()
			}
		}(emp, call)
	}
}

func (cc *CallCenter) findFree(employees []Employee) Employee {
	for _, e := range employees {
		if e.IsFree() {
			return e
		}
	}
	return nil
}

// processQueue re-evaluates the queue when an employee frees up
func (cc *CallCenter) processQueue() {
	if len(cc.queuedCalls) == 0 {
		return
	}
	// Copy queue and reset
	pending := cc.queuedCalls
	cc.queuedCalls = make([]*Call, 0)

	for _, call := range pending {
		cc.dispatch(call) // might be requeued if still no one is free
	}
}

// Run is the central event loop for the CallCenter
func (cc *CallCenter) Run() {
	for {
		select {
		case call := <-cc.newCallChan:
			cc.dispatch(call)
		case <-cc.completionChan:
			cc.processQueue() // someone is now free, check queue
			cc.wg.Done()
		case event := <-cc.escalationChan:
			// escalate call gets dispatched again
			cc.dispatch(event.Call)
			cc.processQueue() // previous employee is now free, check queue
		}
	}
}

func main() {
	var wg sync.WaitGroup
	cc := NewCallCenter(&wg)

	// Setup our employees
	cc.AddOperator("Alice")
	cc.AddOperator("Bob")
	cc.AddSupervisor("Charlie")
	cc.AddDirector("Diana")

	// Start the CallCenter event loop
	go cc.Run()

	// Dispatch 10 customer calls into the system
	for i := 1; i <= 10; i++ {
		cc.DispatchNewCall(&Call{
			ID:    i,
			Rank:  OperatorRank,
			State: Ready,
		})
		time.Sleep(100 * time.Millisecond) // Calls arrive slightly apart
	}

	fmt.Println("All calls dispatched, waiting for processing to complete...")
	wg.Wait()
	fmt.Println("All calls have been handled! Terminating.")
}
