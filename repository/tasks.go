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
}

func GetTasks() []model.Task {
	return _tasks
}

func GetTask(id int) (model.Task, error) {
	var task model.Task

	for _, v := range _tasks {
		if v.Id == id {
			task = v
		}
	}

	if task.Id == id {
		return task, nil
	} else {
		err := fmt.Sprintf("Task com id %d", id)
		return model.Task{}, errors.New(err)
	}
}

func CreateTask(task model.Task) {
	_tasks = append(_tasks, task)
}