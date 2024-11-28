package service

import (
	"context"
	"fmt"
	"microservice_grpc_cart/models"
	"microservice_grpc_cart/pb/cart"

	"gorm.io/gorm"
)

type CartSeerviceServer struct {
	cart.UnimplementedCartSeerviceServer
	DB *gorm.DB
}

func (c *CartSeerviceServer) AddToCart(ctx context.Context, req *cart.AddToCartRequest) (*cart.AddToCartResponse, error) {
	carts := models.Cart{
		UserID:    int(req.GetUserID()),
		ProductID: int(req.GetProductID()),
		Quantity:  int(req.GetQuantity()),
	}

	if err := c.DB.Create(&carts).Error; err != nil {
		return &cart.AddToCartResponse{
			Status: "failed to add product to the cart",
		}, err
	}

	return &cart.AddToCartResponse{
		Id:     uint32(carts.ID),
		Status: "Item added to cart successfully",
	}, nil
}

func (c *CartSeerviceServer) RemoveFromCart(ctx context.Context, req *cart.RemoveFromCartRequest) (*cart.RemoveFromCartResponse, error) {
	var carts models.Cart

	if err := c.DB.Where("id = ?", req.GetId()).First(&carts).Error; err != nil {
		return &cart.RemoveFromCartResponse{
			Status: "cannot found the item in this id",
		}, err
	}

	if err := c.DB.Delete(&carts).Error; err != nil {
		return &cart.RemoveFromCartResponse{
			Status: "failed to delete the item",
		}, err
	}

	return &cart.RemoveFromCartResponse{
		Id:     uint32(carts.ID),
		Status: "Item removed successfully",
	}, nil

}

func (c *CartSeerviceServer) UpdateCartQuantity(ctx context.Context, req *cart.UpdateCartQuantityRequest) (*cart.UpdateCartQuantiyResponse, error) {
	var carts models.Cart

	if err := c.DB.Where("id = ?", req.GetId()).First(&carts).Error; err != nil {
		return &cart.UpdateCartQuantiyResponse{
			Status: "cannot find the item",
		}, err
	}
	carts.Quantity = int(req.GetQuantity())
	if err := c.DB.Save(&carts).Error; err != nil {
		return &cart.UpdateCartQuantiyResponse{
			Status: "failed to update the cart quantity",
		}, err
	}

	return &cart.UpdateCartQuantiyResponse{
		Id:     uint32(carts.ID),
		Status: "successfully update the cart quantity",
	}, nil
}

func (c *CartSeerviceServer) ViewCart(ctx context.Context, req *cart.ViewCartRequest) (*cart.ViewCartResponse, error) {
	var carts []models.Cart
	if err := c.DB.Where("user_id = ?", req.GetUserID()).Find(&carts).Error; err != nil {
		return nil, fmt.Errorf("cannot find the userID")
	}

	var cartResponse []*cart.Cart
	for _, crt := range carts {
		cartResponse = append(cartResponse, &cart.Cart{
			Id:        uint32(crt.ID),
			UserID:    int64(crt.UserID),
			ProductID: int64(crt.ProductID),
			Quantity:  int64(crt.Quantity),
		})
	}

	return &cart.ViewCartResponse{
		Carts: cartResponse,
	}, nil
}

func NewCartServer(db *gorm.DB) *CartSeerviceServer {
return &CartSeerviceServer{
	DB: db,
}
}
