package service

import (
	"context"

	"github.com/vela-ssoc/vela-common-mb/dal/dyncond"
	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
	"github.com/vela-ssoc/vela-common-mb/param/request"
	"github.com/vela-ssoc/vela-common-mb/param/response"
)

func NewTaskExecuteItem(qry *query.Query) (*TaskExecuteItem, error) {
	tbl, err := dyncond.ParseModels(qry, []any{model.TaskExecuteItem{}}, dyncond.Options{})
	if err != nil {
		return nil, err
	}

	return &TaskExecuteItem{
		qry: qry,
		tbl: tbl,
	}, nil
}

type TaskExecuteItem struct {
	qry *query.Query
	tbl *dyncond.Tables
}

func (tex *TaskExecuteItem) Cond() *response.Cond {
	return response.ParseCond(tex.tbl)
}

func (tex *TaskExecuteItem) Page(ctx context.Context, args *request.PageConditions) (*response.Pages[*model.TaskExecuteItem], error) {
	wheres, _, err := tex.tbl.CompileWhere(args.CondWhereInputs.Inputs(), false)
	if err != nil {
		return nil, err
	}

	executeItem := tex.qry.TaskExecuteItem
	executeItemDo := executeItem.WithContext(ctx).Where(wheres...)

	pages := response.NewPages[*model.TaskExecuteItem](args.PageSize())
	cnt, err := executeItemDo.Count()
	if err != nil {
		return nil, err
	} else if cnt == 0 {
		return pages.Empty(), nil
	}

	items, err := executeItemDo.Scopes(pages.FP(cnt)).Find()
	if err != nil {
		return nil, err
	}

	return pages.SetRecords(items), nil
}
