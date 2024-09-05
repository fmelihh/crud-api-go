package cart

import (
	"fmt"

	"github.com/fmelihh/crud-api-go/types"
)

func getCartItemIDs(items []types.CartCheckoutItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}
		productIds[i] = item.ProductID
	}

	return productIds, nil
}

func (h *Handler) createOrder(ps []types.Product, items []types.CartCheckoutItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	if err := checkIfCartInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	totalPrice, _ := calculateTotalPrice(items, productMap)
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address",
	})
	if err != nil {
		return 0, 0, err
	}

	return orderID, totalPrice, nil
}

func checkIfCartInStock(items []types.CartCheckoutItem, productMap map[int]types.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range items {
		product, ok := productMap[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not avaiable in the store, plase refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not avaiable the quantity requested", product.Name)
		}
	}
	return nil
}

func calculateTotalPrice(items []types.CartCheckoutItem, productMap map[int]types.Product) (float64, error) {
	var totalPrice float64
	for _, item := range items {
		product, ok := productMap[item.ProductID]
		if !ok {
			return 0, fmt.Errorf("product %d is not avaiable in the store, plase refresh your cart", item.ProductID)
		}
		totalPrice += product.Price * float64(item.Quantity)
	}
	return totalPrice, nil
}
