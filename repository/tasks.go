package repository

import (
	"errors"
	"firstProject/model"
	"fmt"
)

var _tasks = []model.Task{
	{
		Id:          1,
		Name:        "Teste 1",
		DueDate:     "28/06/22",
		CreatedDate: "22/06/22",
	},
	{
		Id:          2,
		Name:        "Teste 2",
		DueDate:     "30/06/22",
		CreatedDate: "22/06/22",
	},
	{
		Id:          3,
		Name:        "Teste 3",
		DueDate:     "03/07/22",
		CreatedDate: "22/06/22",
	},
}

func GetTasks() []model.Task {
	return _tasks
}

func GetTask(id int) (model.Task, error) {
	var task model.Task

	for _, v := range _tasks {
		if v.Id == id {
			task = v
			break
		}
	}

	if task.Id == id {
		return task, nil
	} else {
		err := fmt.Sprintf("task com id %d não encontrada", id)
		return model.Task{}, errors.New(err)
	}
}

func CreateTask(task model.Task) {
	task.Id = len(_tasks) + 1
	_tasks = append(_tasks, task)
}

func DeleteTask(id int) error {
	var newSlice []model.Task
	for i, v := range _tasks {
		if v.Id == id {
			newSlice = append(_tasks[:i], _tasks[i+1:]...)
			break
		}
	}

	if len(newSlice)+1 == len(_tasks) {
		_tasks = newSlice
		return nil
	} else {
		return errors.New("task não encontrada")
	}
}

func UpdateTask(task model.Task) error {
	var oldTask model.Task
	for i, v := range _tasks {
		if v.Id == task.Id {
			oldTask = v
			_tasks[i] = task
			break
		}
	}

	if oldTask == task {
		return errors.New("task não atualizada")
	} else {
		return nil
	}
}
