package transformer

import (
	"github.com/duc-cnzj/mars/api/v5/types"
	"github.com/duc-cnzj/mars/v5/internal/repo"
	"github.com/duc-cnzj/mars/v5/internal/util/date"
)

func FromProject(project *repo.Project) *types.ProjectModel {
	if project == nil {
		return nil
	}
	return &types.ProjectModel{
		Id:                int32(project.ID),
		Name:              project.Name,
		GitProjectId:      int32(project.GitProjectID),
		GitBranch:         project.GitBranch,
		GitCommit:         project.GitCommit,
		Config:            project.Config,
		OverrideValues:    project.OverrideValues,
		DockerImage:       project.DockerImage,
		PodSelectors:      project.PodSelectors,
		NamespaceId:       int32(project.NamespaceID),
		Atomic:            project.Atomic,
		EnvValues:         project.EnvValues,
		ExtraValues:       project.ExtraValues,
		FinalExtraValues:  project.FinalExtraValues,
		DeployStatus:      project.DeployStatus,
		HumanizeCreatedAt: date.ToHumanizeDatetimeString(&project.CreatedAt),
		HumanizeUpdatedAt: date.ToHumanizeDatetimeString(&project.UpdatedAt),
		ConfigType:        project.ConfigType,
		GitCommitWebUrl:   project.GitCommitWebURL,
		GitCommitTitle:    project.GitCommitTitle,
		GitCommitAuthor:   project.GitCommitAuthor,
		GitCommitDate:     date.ToHumanizeDatetimeString(project.GitCommitDate),
		Version:           int32(project.Version),
		RepoId:            int32(project.RepoID),
		Repo:              FromRepo(project.Repo),
		Namespace:         FromNamespace(project.Namespace),
		CreatedAt:         date.ToRFC3339DatetimeString(&project.CreatedAt),
		UpdatedAt:         date.ToRFC3339DatetimeString(&project.UpdatedAt),
		DeletedAt:         date.ToRFC3339DatetimeString(project.DeletedAt),
	}
}
