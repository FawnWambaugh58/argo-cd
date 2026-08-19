package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/repository"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v2/util/db/mocks"
)

func TestListRefs_ProjectScoped(t *testing.T) {
	ctx := context.Background()
	db := &mocks.ArgoDB{}
	db.On("GetProjectRepository", mock.Anything, "my-project", "https://github.com/argoproj/argo-cd.git").Return(&v1alpha1.Repository{
		Repo:     "https://github.com/argoproj/argo-cd.git",
		Username: "project-user",
	}, nil)

	// Setup server with mock db and enforcer
	// ... (rest of the test setup and assertions)
}