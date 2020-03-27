package todo-app

import "fmt"

type taskList struct {
	task []*task
}

func (taskList *taskList) addItemToList(newTask *task) {
	taskList.task = append(taskList.task, newTask)
}

func (taskList *taskList) deleteItemToList(index int) {
	taskList.task = append(taskList.task[:index], taskList.task[index+1:]...)
}

type task struct {
	name        string
	description string
	isCompleted bool
}

func (t *task) checkAsCompleted() {
	t.isCompleted = true
}

func (t *task) updateTaskName(name string) {
	t.name = name
}

func (t *task) updateTaskDescription(description string) {
	t.description = description
}

func main() {

	t1 := &task{
		name:        "Completar mi curso de GO",
		description: "Completar mi curso de GO en Platzi hoy mismo!",
		isCompleted: false,
	}
	t2 := &task{
		name:        "Completar mi curso de Python",
		description: "Completar mi curso de Python en Platzi hoy mismo!",
		isCompleted: false,
	}
	t3 := &task{
		name:        "Completar mi curso de C#",
		description: "Completar mi curso de C# en Platzi hoy mismo!",
		isCompleted: false,
	}
	t4 := &task{
		name:        "Completar mi curso de DevOps",
		description: "Completar mi curso de DevOps en Platzi hoy mismo!",
		isCompleted: false,
	}

	tList := &taskList{
		task: []*task{
			t1,
			t2,
			t3,
		},
	}

	fmt.Printf("%+v\n", tList)
	tList.addItemToList(t4)
	fmt.Printf("%+v\n", tList)
	tList.deleteItemToList(1)
	fmt.Printf("%+v\n", tList)

}
