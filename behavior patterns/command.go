package main
//https://github.com/yksz/go-design-patterns/blob/master/behavior/command.go
//https://sourcemaking.com/design_patterns/command
//Command pattern : a request is wrapped under an object as command and passed to invoker object.
// Invoker object looks for the appropriate object which can handle this command and passes the command
// to the corresponding object which executes the command.

import (
	"fmt"
)

type Command interface {
	Execute()
}

type ConcreteCommandA struct {
	msg string
	receiver *Receiver
}
func (c *ConcreteCommandA) String() string {
	return "msg A >> "+c.msg+" , re>> "+c.receiver.r+" / "
}

func (c *ConcreteCommandA) Execute() {
	c.receiver.Action("CommandA")
}

type ConcreteCommandB struct {
	msg string
	receiver *Receiver
}
func (c *ConcreteCommandB) String() string {
	return "msg B "+c.msg+" , res >> "+c.receiver.r
}

func (c *ConcreteCommandB) Execute() {
	c.receiver.Action("CommandB")
}

type Receiver struct{
	r string
}

func (r *Receiver) Action(msg string) {
	fmt.Println("excude action : ", msg)
}

type Invoker struct {
	history []Command
}

func (i *Invoker) StoreAndExecute(cmd Command) {
	i.history = append(i.history, cmd)


		fmt.Println("new : ",cmd)
		cmd.Execute()


}
func (i *Invoker) showhistory() {


	fmt.Println("history ",i.history)



}

func main() {
	receiver := &Receiver{"reciever"}
	commandA := &ConcreteCommandA{"Command A",receiver}
	commandB := &ConcreteCommandB{"Command B",receiver}
	invoker := new(Invoker)
	invoker.StoreAndExecute(commandA)
	invoker.StoreAndExecute(commandB)
	invoker.showhistory()
}