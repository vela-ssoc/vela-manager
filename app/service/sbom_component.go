package service

import (
	"context"

	"github.com/vela-ssoc/vela-common-mb-itai/dal/model"
	"github.com/vela-ssoc/vela-common-mb-itai/dal/query"
	"github.com/vela-ssoc/vela-common-mb-itai/dynsql"
	"github.com/vela-ssoc/vela-manager/app/internal/param"
)

type SBOMComponentService interface {
	Page(ctx context.Context, page param.Pager, scope dynsql.Scope) (int64, []*model.SBOMComponent)
	Project(ctx context.Context, page param.Pager, scope dynsql.Scope) (int64, []*model.SBOMProject)
	Count(ctx context.Context, page param.Pager) (int64, []*param.NameCount)
}

func SBOMComponent() SBOMComponentService {
	return &sbomComponentService{}
}

type sbomComponentService struct{}

func (biz *sbomComponentService) Page(ctx context.Context, page param.Pager, scope dynsql.Scope) (int64, []*model.SBOMComponent) {
	tbl := query.SBOMComponent
	db := tbl.WithContext(ctx).
		Order(tbl.TotalScore.Desc()).
		UnderlyingDB().
		Scopes(scope.Where)

	var count int64
	if db.Count(&count); count == 0 {
		return 0, nil
	}

	var dats []*model.SBOMComponent
	db.Scopes(page.DBScope(count)).Find(&dats)

	return count, dats
}

func (biz *sbomComponentService) Project(ctx context.Context, page param.Pager, scope dynsql.Scope) (int64, []*model.SBOMProject) {
	tbl := query.SBOMComponent
	db := tbl.WithContext(ctx).UnderlyingDB()
	subSQL := db.Model(&model.SBOMComponent{}).
		Distinct("project_id").
		Scopes(scope.Where)

	tx := db.Model(&model.SBOMProject{}).Where("id IN (?)", subSQL)
	var count int64
	if tx.Count(&count); count == 0 {
		return 0, nil
	}
	var dats []*model.SBOMProject
	tx.Scopes(page.DBScope(count)).Find(&dats)

	return count, dats
}

func (biz *sbomComponentService) Count(ctx context.Context, page param.Pager) (int64, []*param.NameCount) {
	ret := make([]*param.NameCount, 0, 10)
	tbl := query.SBOMComponent
	count, _ := tbl.WithContext(ctx).Distinct(tbl.Name).Count()
	if count == 0 {
		return 0, ret
	}

	_ = tbl.WithContext(ctx).
		Select(tbl.Name, tbl.Name.Count().As("count")).
		Group(tbl.Name).
		Order(tbl.Name.Count().Desc()).
		Scopes(page.Scope(count)).
		Scan(&ret)

	return count, ret
}
