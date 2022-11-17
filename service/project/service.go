package project

import "errors"

type Service interface {
	// todo: create UserService to make the connection
	// todo: separate into TaskService, ProjectService and UserProjectService
	// todo: do the same for repository
	ProjectService
	TaskService
}

type service struct {
	ProjectService
	TaskService
}

func NewService(projectService ProjectService, taskService TaskService) (Service, error) {

	if projectService == nil {
		return nil, errors.New("service: project service is nil")
	}
	if taskService == nil {
		return nil, errors.New("service: task service is nil")
	}

	return service{ProjectService: projectService, TaskService: taskService}, nil
}
