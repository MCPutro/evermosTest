package product

import (
	"context"
	"evermosTest/internal/entity"
	"evermosTest/internal/handler/request"
	"evermosTest/internal/handler/response"
	"evermosTest/internal/repository/order"
	"evermosTest/internal/repository/product"
	"fmt"
	"sync"
	"time"
)

type productServiceImpl struct {
	product product.Repository
	order   order.Repository
	mu      sync.Mutex
}

// Checkout for deduct product stock and create order
func (p *productServiceImpl) Checkout(ctx context.Context, checkoutRequest *request.Checkout) (*response.Order, error) {
	p.mu.Lock()         // lock transaction
	defer p.mu.Unlock() // unlock

	// find the product details
	existingProduct, err := p.product.FindById(ctx, checkoutRequest.ProductId)
	if err != nil {
		return nil, fmt.Errorf("failed to find product id %d. %w", checkoutRequest.ProductId, err)
	}

	// Check available stock and return an error if available stock is less than the requested quantity
	if existingProduct.Stock < checkoutRequest.Quantity {
		return nil, fmt.Errorf("product stock is lower than quantity")
	}

	// Simulating processing time
	time.Sleep(200 * time.Millisecond)

	// deduct stock
	existingProduct.Stock -= checkoutRequest.Quantity

	// store new available stock in the database after deduction
	_, err = p.product.Update(ctx, existingProduct)
	if err != nil {
		return nil, fmt.Errorf("failed to Deduct product %d. %w", checkoutRequest.ProductId, err)
	}

	// create new order
	// count total price
	totalPrice := existingProduct.Price * checkoutRequest.Quantity
	// store new order in the database
	orderSaved, err := p.order.Save(ctx, &entity.Order{ProductId: checkoutRequest.ProductId, Quantity: checkoutRequest.Quantity, TotalPrice: totalPrice})
	if err != nil {
		return nil, fmt.Errorf("failed to save order. %w", err)
	}

	// return response
	respMessage := orderSaved.ToResponse()
	respMessage.ProductName = existingProduct.Name
	return respMessage, nil
}

func (p *productServiceImpl) GetAllProducts(ctx context.Context) ([]*response.Product, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	products, err := p.product.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return products.ToResponseProductList(), nil
}

// Create for store new product to database
func (p *productServiceImpl) Create(ctx context.Context, product *request.Product) (*response.Product, error) {
	productSaved, err := p.product.Save(ctx, product.ToEntity())
	if err != nil {
		return nil, err
	}
	return productSaved.ToResponse(), nil
}

// PriceAndStockAdjustment for stock in/out and update price
func (p *productServiceImpl) PriceAndStockAdjustment(ctx context.Context, productId int, isDeduct bool, n int, newPrice int) (*response.Product, error) {
	p.mu.Lock()         // lock transaction
	defer p.mu.Unlock() // unlock

	// check for existing product ID
	existingProduct, err := p.product.FindById(ctx, productId)
	if err != nil {
		return nil, err
	}

	// handling when case update price, set -1 to no update product stock
	if n > 0 {
		// update stock
		if isDeduct {
			// Check if the deduction is less than the existing stock. If the deduction is greater than current stock, return an error.
			if (existingProduct.Stock - n) >= 0 {
				existingProduct.Stock -= n
			} else {
				return nil, fmt.Errorf("product stock out of range. stock cannot be negative")
			}
		} else {
			existingProduct.Stock += n
		}
	}

	// handling when update product stock, set -1 to no update product price
	// update product price
	if newPrice > 0 {
		existingProduct.Price = newPrice
	}

	// store new product details in the database
	updated, err := p.product.Update(ctx, existingProduct)
	if err != nil {
		return nil, err
	}

	return updated.ToResponse(), nil
}

// Delete for remove product
func (p *productServiceImpl) Delete(ctx context.Context, productId int) error {
	p.mu.Lock()         // lock transaction
	defer p.mu.Unlock() // unlock

	return p.product.Remove(ctx, productId)
}

// NewService is a constructor function for creating product instances
func NewService(product product.Repository, order order.Repository) Service {
	return &productServiceImpl{product: product, order: order}
}
