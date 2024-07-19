package storage

import (
	"context"
	"social-todo-list/module/item/model"
)

func (sql *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := sql.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
