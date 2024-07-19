package storage

import (
	"context"
	"social-todo-list/module/item/model"
)

func (sql *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := sql.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
