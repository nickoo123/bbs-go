package repositories

import (
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"gorm.io/gorm"

	"bbs/model"
)

var ProjectRepository = newProjectRepository()

func newProjectRepository() *projectRepository {
	return &projectRepository{}
}

type projectRepository struct {
}

func (r *projectRepository) Get(db *gorm.DB, id int64) *model.Project {
	ret := &model.Project{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *projectRepository) Take(db *gorm.DB, where ...interface{}) *model.Project {
	ret := &model.Project{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *projectRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.Project) {
	cnd.Find(db, &list)
	return
}

func (r *projectRepository) FindOne(db *gorm.DB, cnd *sqls.Cnd) *model.Project {
	ret := &model.Project{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *projectRepository) FindPageByParams(db *gorm.DB, params *params.QueryParams) (list []model.Project, paging *sqls.Paging) {
	return r.FindPageByCnd(db, &params.Cnd)
}

func (r *projectRepository) FindPageByCnd(db *gorm.DB, cnd *sqls.Cnd) (list []model.Project, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Project{})

	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *projectRepository) Create(db *gorm.DB, t *model.Project) (err error) {
	err = db.Create(t).Error
	return
}

func (r *projectRepository) Update(db *gorm.DB, t *model.Project) (err error) {
	err = db.Save(t).Error
	return
}

func (r *projectRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Project{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *projectRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Project{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *projectRepository) Delete(db *gorm.DB, id int64) {
	db.Model(&model.Project{}).Delete("id", id)
}
