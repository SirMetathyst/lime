package main

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type StdServiceLoggingMiddleware struct {
	Next   Service
	Logger log.Logger
}

func (s *StdServiceLoggingMiddleware) CreateProject(ctx context.Context, project Project) (newProjectID string, err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "create-project", "requested-project-id", project.ID, "real-project-id", newProjectID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.CreateProject(ctx, project)
}

func (s *StdServiceLoggingMiddleware) ReadProject(ctx context.Context, id string) (project Project, err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "read-project", "requested-project-id", id, "real-project-id", project.ID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.ReadProject(ctx, id)
}

func (s *StdServiceLoggingMiddleware) UpdateProject(ctx context.Context, project Project) (err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "update-project", "requested-project-id", project.ID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.UpdateProject(ctx, project)
}

func (s *StdServiceLoggingMiddleware) DeleteProject(ctx context.Context, projectID string) (err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "delete-project", "requested-project-id", projectID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.DeleteProject(ctx, projectID)
}

func (s *StdServiceLoggingMiddleware) ListProjects(ctx context.Context) (projects []Project, err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "list-projects", "project-count", len(projects), "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.ListProjects(ctx)
}

func (s *StdServiceLoggingMiddleware) CreateTask(ctx context.Context, projectID string, task Task) (newTaskID string, err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "create-task", "requested-project-id", projectID, "requested-task-id", task.ID, "real-task-id", newTaskID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.CreateTask(ctx, projectID, task)
}

func (s *StdServiceLoggingMiddleware) ReadTask(ctx context.Context, projectID string, taskID string) (task Task, err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "read-task", "requested-project-id", projectID, "requested-task-id", taskID, "real-task-id", task.ID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.ReadTask(ctx, projectID, taskID)
}

func (s *StdServiceLoggingMiddleware) UpdateTask(ctx context.Context, projectID string, task Task) (err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "update-task", "requested-task-id", projectID, "requested-task-id", task.ID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.UpdateTask(ctx, projectID, task)
}

func (s *StdServiceLoggingMiddleware) DeleteTask(ctx context.Context, projectID string, taskID string) (err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "delete-task", "requested-project-id", projectID, "requested-task-id", taskID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.DeleteTask(ctx, projectID, taskID)
}

func (s *StdServiceLoggingMiddleware) DeleteTasks(ctx context.Context, projectID string) (err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "delete-tasks", "requested-project-id", projectID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.DeleteTasks(ctx, projectID)
}

func (s *StdServiceLoggingMiddleware) ListTasks(ctx context.Context, projectID string) (tasks []Task, err error) {

	defer func(begin time.Time) {
		s.Logger.Log("method", "list-tasks", "task-count", len(tasks), "took", time.Since(begin), "err", err)
	}(time.Now())

	return s.Next.ListTasks(ctx, projectID)
}
