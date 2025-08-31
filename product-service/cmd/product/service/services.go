package services

import (
	"context"
	"product-service/cmd/product/repository"
	"product-service/models"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepo,
	}
}

func (s *ProductService) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {
	product, err := s.ProductRepository.FindByProductId(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProductCategoryByID(ctx context.Context, id int64) (*models.ProductCategory, error) {
	category, err := s.ProductRepository.FindByProductCategoryId(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) (int64, error) {
	id, err := s.ProductRepository.InsertProduct(ctx, product)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ProductService) CreateProductCategory(ctx context.Context, category *models.ProductCategory) (int64, error) {
	id, err := s.ProductRepository.InsertProductCategory(ctx, category)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product, err := s.ProductRepository.UpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) UpdateProductCategory(ctx context.Context, category *models.ProductCategory) (*models.ProductCategory, error) {
	category, err := s.ProductRepository.UpdateProductCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	err := s.ProductRepository.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProductCategory(ctx context.Context, id int64) error {
	err := s.ProductRepository.DeleteProductCategory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
