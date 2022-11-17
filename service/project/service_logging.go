package project

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"time"
)

type serviceLoggingMiddleware struct {
	next   Service
	logger log.Logger
}

func NewServiceLoggingMiddleware(logger log.Logger, next Service) (Service, error) {

	if logger == nil {
		return nil, errors.New("service-logging-middleware: logger is nil")
	}
	if next == nil {
		return nil, errors.New("service-logging-middleware: next is nil")
	}

	return &serviceLoggingMiddleware{logger: logger, next: next}, nil
}

func (s *serviceLoggingMiddleware) CreateProject(ctx context.Context, project Project) (newProjectID string, err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "create-project", "requested-project-id", project.ID, "real-project-id", newProjectID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.CreateProject(ctx, project)
}

func (s *serviceLoggingMiddleware) ReadProject(ctx context.Context, id string) (project Project, err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "read-project", "requested-project-id", id, "real-project-id", project.ID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.ReadProject(ctx, id)
}

func (s *serviceLoggingMiddleware) UpdateProject(ctx context.Context, project Project) (err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "update-project", "requested-project-id", project.ID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.UpdateProject(ctx, project)
}

func (s *serviceLoggingMiddleware) DeleteProject(ctx context.Context, projectID string) (err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "delete-project", "requested-project-id", projectID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.DeleteProject(ctx, projectID)
}

func (s *serviceLoggingMiddleware) ListProjects(ctx context.Context) (projects []Project, err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "list-projects", "project-count", len(projects), "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.ListProjects(ctx)
}

func (s *serviceLoggingMiddleware) CreateTask(ctx context.Context, projectID string, task Task) (newTaskID string, err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "create-task", "requested-project-id", projectID, "requested-task-id", task.ID, "real-task-id", newTaskID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.CreateTask(ctx, projectID, task)
}

func (s *serviceLoggingMiddleware) ReadTask(ctx context.Context, projectID string, taskID string) (task Task, err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "read-task", "requested-project-id", projectID, "requested-task-id", taskID, "real-task-id", task.ID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.ReadTask(ctx, projectID, taskID)
}

func (s *serviceLoggingMiddleware) UpdateTask(ctx context.Context, projectID string, task Task) (err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "update-task", "requested-task-id", projectID, "requested-task-id", task.ID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.UpdateTask(ctx, projectID, task)
}

func (s *serviceLoggingMiddleware) DeleteTask(ctx context.Context, projectID string, taskID string) (err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "delete-task", "requested-project-id", projectID, "requested-task-id", taskID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.DeleteTask(ctx, projectID, taskID)
}

func (s *serviceLoggingMiddleware) DeleteTasks(ctx context.Context, projectID string) (err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "delete-tasks", "requested-project-id", projectID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.DeleteTasks(ctx, projectID)
}

func (s *serviceLoggingMiddleware) ListTasks(ctx context.Context, projectID string) (tasks []Task, err error) {

	defer func(begin time.Time) {
		s.logger.Log("method", "list-tasks", "task-count", len(tasks), "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.next.ListTasks(ctx, projectID)
}
