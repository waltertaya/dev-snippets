package main

import "fmt"

// chaining hash table for executing top priority

type TaskManager struct {
	taskToUser map[int]int
	taskToPrirority map[int]int
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		taskToUser: make(map[int]int),
		taskToPrirority: make(map[int]int),
	}

	for _, task := range tasks {
		tm.taskToUser[task[1]] = task[0]
		tm.taskToPrirority[task[1]] = task[2]
	}

	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.taskToUser[taskId] = userId
	this.taskToPrirority[taskId] = priority
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	this.taskToPrirority[taskId] = newPriority
}

func (this *TaskManager) Rmv(taskId int) {
	delete(this.taskToUser, taskId)
	delete(this.taskToPrirority, taskId)
}

func (this *TaskManager) ExecTop() int {

}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */

func main() {

	// Testcase
	tasks := [][]int{
		{1, 101, 10}, {2, 102, 20}, {3, 103, 15},
	}
	obj := Constructor(tasks)
	fmt.Println(obj)
	obj.Add(4, 104, 5)
	obj.Edit(102, 8)
	fmt.Println(obj.ExecTop()) // returns 3
	obj.Rmv(101)
	obj.Add(5, 105, 15)
	fmt.Println(obj.ExecTop()) // returns 5

	// fmt.Println(obj)
}
