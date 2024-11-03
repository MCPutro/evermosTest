package product

import (
	"context"
	"evermosTest/internal/entity"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

func (p *productRepositoryImpl) Save(ctx context.Context, newProduct *entity.Product) (*entity.Product, error) {
	err := p.db.WithContext(ctx).Create(&newProduct).Error
	if err != nil {
		return nil, err
	}
	return newProduct, err
}

func (p *productRepositoryImpl) FindById(ctx context.Context, ID int) (*entity.Product, error) {
	var product entity.Product

	result := p.db.WithContext(ctx).First(&product, "id = ?", ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (p *productRepositoryImpl) FindAll(ctx context.Context) (entity.ProductList, error) {
	var products []*entity.Product
	result := p.db.WithContext(ctx).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected > 0 {
		return products, nil
	}
	// fmt.Println("Found ", result.RowsAffected, "(s) data")

	return nil, nil
}

func (p *productRepositoryImpl) Update(ctx context.Context, newProduct *entity.Product) (*entity.Product, error) {
	result := p.db.WithContext(ctx).Save(&newProduct)
	if result.Error != nil {
		return nil, result.Error
	}

	return newProduct, nil
}

func (p *productRepositoryImpl) Remove(ctx context.Context, ID int) error {

	//find product
	var product entity.Product
	result := p.db.WithContext(ctx).First(&product, "id = ?", ID)
	if result.Error != nil {
		return result.Error
	}

	//delete product
	return p.db.Delete(&product).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &productRepositoryImpl{db: db}
}
