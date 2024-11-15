package repo

import (
	"context"
	"encoding/json"
	"fmt"
	gopath "path"
	"strconv"
	"strings"
	"time"

	"github.com/duc-cnzj/mars/v5/internal/application"
	"github.com/duc-cnzj/mars/v5/internal/cache"
	"github.com/duc-cnzj/mars/v5/internal/data"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/util/serialize"
	"github.com/spf13/cast"
)

type GitRepo interface {
	AllProjects(ctx context.Context, forceFresh bool) (projects []*GitProject, err error)
	AllBranches(ctx context.Context, projectID int, forceFresh bool) (branches []*Branch, err error)
	ListCommits(ctx context.Context, projectID int, branch string) ([]*Commit, error)
	GetCommit(ctx context.Context, projectID int, sha string) (*Commit, error)
	GetCommitPipeline(ctx context.Context, projectID int, branch, sha string) (*Pipeline, error)
	GetByProjectID(ctx context.Context, id int) (project *GitProject, err error)
	GetFileContentWithBranch(ctx context.Context, projectID int, branch, path string) (string, error)
	GetProject(ctx context.Context, id int) (project *GitProject, err error)
	GetChartValuesYaml(ctx context.Context, localChartPath string) (string, error)
}

var _ GitRepo = (*gitRepo)(nil)

type gitRepo struct {
	logger mlog.Logger
	pl     application.PluginManger
	data   data.Data
	cache  cache.Cache
}

type Branch struct {
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	WebURL    string `json:"web_url"`
}

type GitProject struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	DefaultBranch string `json:"default_branch"`
	WebURL        string `json:"web_url"`
	Path          string `json:"path"`
	AvatarURL     string `json:"avatar_url"`
	Description   string `json:"description"`
}

type Commit struct {
	ID             string
	ShortID        string
	AuthorName     string
	AuthorEmail    string
	CommitterName  string
	CommitterEmail string
	Message        string
	Title          string
	WebURL         string
	CreatedAt      *time.Time
	CommittedDate  *time.Time
}

type Pipeline struct {
	Status application.Status
	WebURL string
}

func NewGitRepo(logger mlog.Logger, cache cache.Cache, pl application.PluginManger, data data.Data) GitRepo {
	return &gitRepo{
		logger: logger.WithModule("repo/git"),
		pl:     pl,
		cache:  cache,
		data:   data,
	}
}

func (g *gitRepo) AllProjects(ctx context.Context, forceFresh bool) ([]*GitProject, error) {
	fn := func() (projects []*GitProject, err error) {
		allProjects, err := g.pl.Git().AllProjects()
		if err != nil {
			return nil, ToError(400, err)
		}
		for _, gp := range allProjects {
			projects = append(projects, ToGitProject(gp))
		}
		return
	}
	if !g.data.Config().GitServerCached {
		return fn()
	}
	remember, err := g.cache.Remember(cache.NewKey("all_projects"), 600, func() ([]byte, error) {
		projects, err := fn()
		if err != nil {
			return nil, err
		}
		return json.Marshal(projects)
	}, forceFresh)
	var projects []*GitProject
	if err == nil {
		err = json.Unmarshal(remember, &projects)
	}
	return projects, err
}

func (g *gitRepo) GetChartValuesYaml(ctx context.Context, localChartPath string) (string, error) {
	if !IsRemoteLocalChartPath(localChartPath) {
		return "", nil
	}
	split := strings.Split(localChartPath, "|")
	pid := split[0]
	branch := split[1]
	filename := gopath.Join(split[2], "values.yaml")
	return g.GetFileContentWithBranch(ctx, cast.ToInt(pid), branch, filename)
}

func ToGitProject(gp application.Project) *GitProject {
	return &GitProject{
		ID:            gp.GetID(),
		Name:          gp.GetName(),
		DefaultBranch: gp.GetDefaultBranch(),
		WebURL:        gp.GetWebURL(),
		Path:          gp.GetPath(),
		AvatarURL:     gp.GetAvatarURL(),
		Description:   gp.GetDescription(),
	}
}

func (g *gitRepo) AllBranches(ctx context.Context, projectID int, forceFresh bool) ([]*Branch, error) {
	fn := func() (branches []*Branch, err error) {
		res, err := g.pl.Git().AllBranches(fmt.Sprintf("%d", projectID))
		if err != nil {
			return nil, ToError(400, err)
		}
		for _, br := range res {
			branches = append(branches, &Branch{
				Name:      br.GetName(),
				IsDefault: br.IsDefault(),
				WebURL:    br.GetWebURL(),
			})
		}
		return
	}
	if !g.data.Config().GitServerCached {
		return fn()
	}
	remember, err := g.cache.Remember(cache.NewKey(fmt.Sprintf("all_branches_%d", projectID)), 600, func() ([]byte, error) {
		branches, err := fn()
		if err != nil {
			return nil, err
		}
		return json.Marshal(branches)
	}, forceFresh)
	var branches []*Branch
	if err == nil {
		err = json.Unmarshal(remember, &branches)
	}
	return branches, err
}

func (g *gitRepo) ListCommits(ctx context.Context, projectID int, branch string) ([]*Commit, error) {
	commits, err := g.pl.Git().ListCommits(fmt.Sprintf("%d", projectID), branch)
	if err != nil {
		return nil, ToError(404, err)
	}
	return serialize.Serialize(commits, ToCommit), nil
}

func (g *gitRepo) GetProject(ctx context.Context, id int) (project *GitProject, err error) {
	getProject, err := g.pl.Git().GetProject(fmt.Sprintf("%d", id))
	return ToGitProject(getProject), ToError(404, err)
}

func (g *gitRepo) GetFileContentWithBranch(ctx context.Context, projectID int, branch, path string) (string, error) {
	withBranch, err := g.pl.Git().GetFileContentWithBranch(fmt.Sprintf("%d", projectID), branch, path)
	return withBranch, ToError(404, err)
}

func (g *gitRepo) GetCommit(ctx context.Context, projectID int, sha string) (*Commit, error) {
	commit, err := g.pl.Git().GetCommit(fmt.Sprintf("%d", projectID), sha)
	if err != nil {
		return nil, ToError(404, err)
	}
	return ToCommit(commit), nil
}

func (g *gitRepo) GetCommitPipeline(ctx context.Context, projectID int, branch, sha string) (*Pipeline, error) {
	pipeline, err := g.pl.Git().GetCommitPipeline(fmt.Sprintf("%d", projectID), branch, sha)
	if err != nil {
		return nil, ToError(404, err)
	}
	return &Pipeline{
		Status: pipeline.GetStatus(),
		WebURL: pipeline.GetWebURL(),
	}, nil
}

func (g *gitRepo) GetByProjectID(ctx context.Context, id int) (project *GitProject, err error) {
	getProject, err := g.pl.Git().GetProject(fmt.Sprintf("%d", id))
	return ToGitProject(getProject), ToError(404, err)
}

func ToCommit(v application.Commit) *Commit {
	if v == nil {
		return nil
	}
	return &Commit{
		ID:             v.GetID(),
		ShortID:        v.GetShortID(),
		AuthorName:     v.GetAuthorName(),
		AuthorEmail:    v.GetAuthorEmail(),
		CommitterName:  v.GetCommitterName(),
		CommitterEmail: v.GetCommitterEmail(),
		Message:        v.GetMessage(),
		Title:          v.GetTitle(),
		WebURL:         v.GetWebURL(),
		CreatedAt:      v.GetCreatedAt(),
		CommittedDate:  v.GetCommittedDate(),
	}
}

func IsRemoteLocalChartPath(input string) bool {
	split := strings.Split(input, "|")

	return len(split) == 3 && intPid(split[0])
}

func intPid(pid string) bool {
	if _, err := strconv.ParseInt(pid, 10, 64); err == nil {
		return true
	}
	return false
}
