package jobs

import (
	"context"

	eiriniv1 "code.cloudfoundry.org/eirini-controller/pkg/apis/eirini/v1"
	"github.com/pkg/errors"
	batch "k8s.io/api/batch/v1"
)

//counterfeiter:generate . JobLister

type JobLister interface {
	List(ctx context.Context, includeCompleted bool) ([]batch.Job, error)
}

type Lister struct {
	jobLister JobLister
}

func NewLister(
	jobLister JobLister,
) Lister {
	return Lister{
		jobLister: jobLister,
	}
}

func (l *Lister) List(ctx context.Context) ([]*eiriniv1.Task, error) {
	jobs, err := l.jobLister.List(ctx, false)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list jobs")
	}

	tasks := make([]*eiriniv1.Task, 0, len(jobs))
	for _, job := range jobs {
		tasks = append(tasks, toTask(job))
	}

	return tasks, nil
}
