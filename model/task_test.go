package model

import (
	json2 "github.com/goccy/go-json"
	"testing"
)

func TestStruct(t *testing.T) {
	var task Task

	json := "{\"id\":1, \"name\":\"testing\", \"createdDate\":\"01/07/22\", \"dueDate\":\"01/07/22\"}"

	err := json2.Unmarshal([]byte(json), &task)

	if err != nil {
		t.Error(err.Error())
	}

	if task.Name == "" {
		t.Errorf("json parsing failed")
	}
}

func TestSetDone(t *testing.T) {
	task := Task{}

	if task.done == true {
		t.Errorf("task already filled")
	}

	task.SetDone(true)

	if task.done != true {
		t.Errorf("task could not be setted as done")
	}
}

func TestGetDone(t *testing.T) {
	task := Task{done: true}

	got := task.GetDone()

	if got != true {
		t.Errorf("got the wrong value of done")
	}
}
