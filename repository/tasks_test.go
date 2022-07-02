package repository

import (
	"firstProject/model"
	"testing"
)

func init() {
	_tasks = []model.Task{
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
}

func TestGetTasks(t *testing.T) {
	tasks := GetTasks()

	if len(tasks) != len(_tasks) {
		t.Errorf("get tasks not working")
	}
}

func TestGetTask(t *testing.T) {
	id := 1
	task, err := GetTask(id)

	if err != nil {
		t.Errorf(err.Error())
	}

	if task.Id != id || task.Name == "" {
		t.Errorf("task getter failed")
	}
}

func TestCreateTask(t *testing.T) {
	task := model.Task{
		Id:          2,
		Name:        "testing",
		DueDate:     "01/07/22",
		CreatedDate: "01/07/22",
	}

	oldLen := len(_tasks)

	CreateTask(task)

	if len(_tasks) == oldLen {
		t.Errorf("create task failed")
	}
}

func TestDeleteTask(t *testing.T) {
	oldLen := len(_tasks)

	err := DeleteTask(1)

	if err != nil {
		t.Errorf(err.Error())
	}

	if oldLen == len(_tasks) {
		t.Errorf("delete task failed")
	}
}

func TestUpdateTask(t *testing.T) {
	task := model.Task{
		Id:          0,
		Name:        "testing updated",
		DueDate:     "01/07/22",
		CreatedDate: "01/07/22",
	}

	err := UpdateTask(task)

	if err != nil {
		t.Errorf(err.Error())
	}

	got, _ := GetTask(0)

	if got.Name == task.Name {
		t.Errorf("update task failed")
	}
}
