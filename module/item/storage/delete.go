package storage

import (
	"context"
	"social-todo-list/module/item/model"
)

func (sql *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := model.ItemStatusDeleted
	if err := sql.db.Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deletedStatus.String(),
		}).Error; err != nil {
		return err
	}
	return nil
}
