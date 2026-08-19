package repository

import (
	"context"
	"fmt"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient/repository"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v2/server/rbacpolicy"
	"github.com/argoproj/argo-cd/v2/util/db"
	"github.com/argoproj/argo-cd/v2/util/git"
)

// ... (rest of the imports and Server struct definition)

func (s *Server) ListRefs(ctx context.Context, q *repository.RepoQuery) (*repository.RefsQueryResult, error) {
	if err := s.enf.EnforceErr(ctx.Value("claims"), rbacpolicy.ResourceRepositories, rbacpolicy.ActionGet, q.Repo); err != nil {
		return nil, err
	}
	if q.Project != "" {
		if err := s.enf.EnforceErr(ctx.Value("claims"), rbacpolicy.ResourceProjects, rbacpolicy.ActionGet, q.Project); err != nil {
			return nil, err
		}
	}
	var repo *v1alpha1.Repository
	var err error
	if q.Project != "" {
		repo, err = s.db.GetProjectRepository(ctx, q.Project, q.Repo)
	} else {
		repo, err = s.db.GetRepository(ctx, q.Repo)
	}
	if err != nil {
		return nil, err
	}
	// ... (rest of the ListRefs implementation remains unchanged)
}