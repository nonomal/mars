package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/duc-cnzj/mars-client/v4/mars"
	"github.com/duc-cnzj/mars-client/v4/types"
	"github.com/duc-cnzj/mars/internal/utils/date"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGitProject_GlobalMarsConfig(t *testing.T) {
	marsCfg := mars.Config{
		ConfigFile:       "cfgfile",
		ConfigFileValues: "values",
		ConfigField:      "conf",
		IsSimpleEnv:      true,
		ConfigFileType:   "yaml",
		LocalChartPath:   "./charts",
		Branches:         []string{"master", "dev"},
		ValuesYaml:       "name: duc\nage: 27",
		Elements: []*mars.Element{
			{
				Path:         "conf->env",
				Type:         mars.ElementType_ElementTypeSelect,
				Default:      "dev",
				Description:  "environment",
				SelectValues: []string{"dev", "master", "*"},
			},
		},
	}
	marshal, _ := json.Marshal(&marsCfg)
	m := GitProject{
		ID:           1,
		GlobalConfig: string(marshal),
	}
	assert.Equal(t, &marsCfg, m.GlobalMarsConfig())
}

func TestGitProject_PrettyYaml(t *testing.T) {
	marsCfg := mars.Config{
		ConfigFile:       "cfgfile",
		ConfigFileValues: "values",
		ConfigField:      "conf",
		IsSimpleEnv:      true,
		ConfigFileType:   "yaml",
		LocalChartPath:   "./charts",
		Branches:         []string{"master", "dev"},
		ValuesYaml:       "name: duc\nage: 27",
		Elements: []*mars.Element{
			{
				Path:         "conf->env",
				Type:         mars.ElementType_ElementTypeSelect,
				Default:      "dev",
				Description:  "environment",
				SelectValues: []string{"dev", "master", "*"},
			},
		},
	}
	marshal, _ := json.Marshal(&marsCfg)
	m := GitProject{
		ID:           1,
		GlobalConfig: string(marshal),
	}

	assert.Equal(t, `config_file: cfgfile
config_file_values: values
config_field: conf
is_simple_env: true
config_file_type: yaml
local_chart_path: ./charts
branches:
- master
- dev
values_yaml:
  age: 27
  name: duc
elements:
- path: conf->env
  type: 3
  default: dev
  description: environment
  selectvalues:
  - dev
  - master
  - '*'
`, m.PrettyYaml())
}

func TestGitProject_ProtoTransform(t *testing.T) {
	m := GitProject{
		ID:            1,
		DefaultBranch: "dev",
		Name:          "mars",
		GitProjectId:  100,
		Enabled:       true,
		GlobalEnabled: true,
		GlobalConfig:  "xxx",
		CreatedAt:     time.Now().Add(15 * time.Minute),
		UpdatedAt:     time.Now().Add(30 * time.Minute),
		DeletedAt: gorm.DeletedAt{
			Time:  time.Now().Add(-10 * time.Second),
			Valid: true,
		},
	}
	assert.Equal(t, &types.GitProjectModel{
		Id:            int64(m.ID),
		DefaultBranch: m.DefaultBranch,
		Name:          m.Name,
		GitProjectId:  int64(m.GitProjectId),
		Enabled:       m.Enabled,
		GlobalEnabled: m.GlobalEnabled,
		GlobalConfig:  m.GlobalConfig,
		CreatedAt:     date.ToRFC3339DatetimeString(&m.CreatedAt),
		UpdatedAt:     date.ToRFC3339DatetimeString(&m.UpdatedAt),
		DeletedAt:     date.ToRFC3339DatetimeString(&m.DeletedAt.Time),
	}, m.ProtoTransform())
}
