package data

import (
	"errors"
	"task_manager/models"
	"sync"
)

type TaskService struct {
	tasks  map[string]models.Task
	mu     sync.RWMutex
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make(map[string]models.Task),
		nextID: 1,
	}
}

func (s *TaskService) GetAllTasks() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) CreateTask(task models.Task) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = string(rune(s.nextID + '0'))
	s.nextID++
	s.tasks[task.ID] = task
	return task
}

func (s *TaskService) UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return models.Task{}, errors.New("task not found")
	}

	updatedTask.ID = id
	s.tasks[id] = updatedTask
	return updatedTask, nil
}

func (s *TaskService) DeleteTask(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(s.tasks, id)
	return nil
}
