package repository

import (
	"context"
	"product-service/models"
)

func (r *ProductRepository) FindByProductId(ctx context.Context, id int64) (*models.Product, error) {
	var product models.Product
	err := r.Database.WithContext(ctx).Table("products").Where("id = ?", id).Last(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) FindByProductCategoryId(ctx context.Context, id int64) (*models.ProductCategory, error) {
	var category models.ProductCategory
	err := r.Database.WithContext(ctx).Table("product_categories").Where("id = ?", id).Last(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *ProductRepository) InsertProduct(ctx context.Context, product *models.Product) (int64, error) {
	err := r.Database.WithContext(ctx).Table("products").Create(product).Error
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (r *ProductRepository) InsertProductCategory(ctx context.Context, category *models.ProductCategory) (int64, error) {
	err := r.Database.WithContext(ctx).Table("product_categories").Create(category).Error
	if err != nil {
		return 0, err
	}
	return category.ID, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	err := r.Database.WithContext(ctx).Table("products").Save(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) UpdateProductCategory(ctx context.Context, category *models.ProductCategory) (*models.ProductCategory, error) {
	err := r.Database.WithContext(ctx).Table("product_categories").Save(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int64) error {
	err := r.Database.WithContext(ctx).Table("products").Where("id = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) DeleteProductCategory(ctx context.Context, id int64) error {
	err := r.Database.WithContext(ctx).Table("product_categories").Where("id = ?", id).Delete(&models.ProductCategory{}).Error
	if err != nil {
		return err
	}
	return nil
}
