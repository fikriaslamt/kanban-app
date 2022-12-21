package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	results := []entity.Category{}

	if err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("user_id = ?", id).Find(&results).Error; err != nil {
		return []entity.Category{}, err
	}

	return results, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {

	err = r.db.WithContext(ctx).Create(&category).Error
	if err != nil {
		return 0, err
	}

	return category.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {

	if err := r.db.WithContext(ctx).Create(&categories).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {

	result := entity.Category{}

	if err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("id = ?", id).Find(&result).Error; err != nil {
		return entity.Category{}, err
	}
	return result, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {

	if err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("id = ?", category.ID).Updates(&category).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Category{}).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
