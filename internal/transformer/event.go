package transformer

import (
	"github.com/duc-cnzj/mars/api/v5/types"
	"github.com/duc-cnzj/mars/v5/internal/repo"
	"github.com/duc-cnzj/mars/v5/internal/util/date"
	"github.com/samber/lo"
)

func FromEvent(e *repo.Event) *types.EventModel {
	if e == nil {
		return nil
	}

	return &types.EventModel{
		Id:        int32(e.ID),
		Action:    e.Action,
		Username:  e.Username,
		Message:   e.Message,
		Old:       e.Old,
		New:       e.New,
		Duration:  e.Duration,
		FileId:    int32(lo.FromPtr(e.FileID)),
		File:      FromFile(e.File),
		HasDiff:   e.HasDiff,
		EventAt:   date.ToHumanizeDatetimeString(&e.CreatedAt),
		CreatedAt: date.ToRFC3339DatetimeString(&e.CreatedAt),
		UpdatedAt: date.ToRFC3339DatetimeString(&e.UpdatedAt),
		DeletedAt: date.ToRFC3339DatetimeString(e.DeletedAt),
	}
}
