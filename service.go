package main

type Service interface {
	// todo: create UserService to make the connection
	// todo: separate into TaskService, ProjectService and UserProjectService
	// todo: do the same for repository
	ProjectService
	TaskService
}

type StdService struct {
	ProjectService
	TaskService
}
