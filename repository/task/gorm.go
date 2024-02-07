package task

import (
	"context"
	"fmt"

	"github.com/takahiromitsui/go-task-manager/model"
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

func (r *GormRepository) Insert(ctx context.Context, task model.Task) error {
	result := r.DB.WithContext(ctx).Create(&task)
	if result.Error != nil {
		return fmt.Errorf("error inserting task: %v", result.Error)
	}
	return nil
}
