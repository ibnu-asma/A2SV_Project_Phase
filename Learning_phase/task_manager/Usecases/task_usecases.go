package usecases

import (
	"task_manager/Domain"
	"task_manager/Repositories"
)

type TaskUsecase interface {
	GetAllTasks() ([]domain.Task, error)
	GetTaskByID(id string) (domain.Task, error)
	CreateTask(task domain.Task) (domain.Task, error)
	UpdateTask(id string, task domain.Task) (domain.Task, error)
	DeleteTask(id string) error
}

type taskUsecase struct {
	taskRepo repositories.TaskRepository
}

func NewTaskUsecase(taskRepo repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (u *taskUsecase) GetAllTasks() ([]domain.Task, error) {
	return u.taskRepo.GetAll()
}

func (u *taskUsecase) GetTaskByID(id string) (domain.Task, error) {
	return u.taskRepo.GetByID(id)
}

func (u *taskUsecase) CreateTask(task domain.Task) (domain.Task, error) {
	return u.taskRepo.Create(task)
}

func (u *taskUsecase) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	return u.taskRepo.Update(id, task)
}

func (u *taskUsecase) DeleteTask(id string) error {
	return u.taskRepo.Delete(id)
}
