package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/item/model"
)

func (sql *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
