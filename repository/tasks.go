package repository

import "firstProject/model"

var tasks = []model.Task{
	{
		Name:        "Teste 1",
		DueDate:     "28/06/22",
		CreatedDate: "22/06/22",
	},
	{
		Name:        "Teste 2",
		DueDate:     "30/06/22",
		CreatedDate: "22/06/22",
	},
}

func GetTasks() []model.Task {
	return tasks
}
