package core

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"shop/internal/features/product/core/dto"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type ProductUseCase interface {
	GetProducts(ctx context.Context, input dto.GetProductsRequest) (*dto.GetProductsResult, error)
	GetProductDetail(ctx context.Context, productID int) (*dto.GetProductDetailResult, error)
}

type useCase struct {
	repo  ProductRepository
	redis *redis.Client
}

func NewProductUseCase(repo ProductRepository, redis *redis.Client) useCase {
	return useCase{
		repo:  repo,
		redis: redis,
	}
}

func (u useCase) GetProducts(ctx context.Context, input dto.GetProductsRequest) (*dto.GetProductsResult, error) {
	currentPage := max(input.Page, 1)
	cacheKey := fmt.Sprintf("products:page=%d", currentPage)
	cached, err := u.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var result dto.GetProductsResult
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return &result, nil
		}
	}
	perpage := 15
	offset := currentPage * perpage
	sortDirection := strings.ToLower(input.SortBy.Order)
	if sortDirection == "" {
		sortDirection = "asc"
	}
	sb := dto.SortBy{
		Field: input.SortBy.Field,
		Order: sortDirection,
	}
	params := dto.GetProductsParams{
		Limit:   perpage,
		Offset:  offset,
		Filters: input.Filters,
		SortBy:  sb,
	}

	rawProducts, totalCount, err := u.repo.FetchProducts(ctx, params)
	if err != nil {
		return nil, err
	}

	products := make([]dto.Product, len(rawProducts))
	for i, p := range rawProducts {
		products[i] = dto.Product(p)
	}

	// calculate pagination
	totalPages := (totalCount + perpage - 1) / perpage
	prevPage := max(currentPage-1, 0)
	nextPage := min(currentPage+1, totalPages)

	result := &dto.GetProductsResult{
		Items: products,
		Pagination: dto.Pagination{
			CurrentPage: currentPage,
			PrevPage:    prevPage,
			NextPage:    nextPage,
			PerPage:     perpage,
			Total:       totalCount,
			TotalPages:  totalPages,
		},
	}

	bytes, err := json.Marshal(products)
	if err == nil {
		const productsCacheTTL = 5 * time.Minute
		u.redis.Set(ctx, cacheKey, bytes, productsCacheTTL)
	}

	return result, nil
}

func (u useCase) GetProductDetail(ctx context.Context, productID int) (*dto.GetProductDetailResult, error) {
	rawProduct, err := u.repo.GetProductByID(ctx, productID)
	if err != nil {
		slog.Debug("failed to get productByID", slog.String("error", err.Error()))
		return nil, err
	}

	product := dto.ProductDetail(*rawProduct)
	detail := dto.GetProductDetailResult{
		Detail: &product,
	}

	rawVariants, err := u.repo.FetchProductVariantByID(ctx, productID)
	if err != nil {
		slog.Debug("failed to get productByID", slog.String("error", err.Error()))
		return &detail, err
	}

	variants := make([]dto.ProductVariant, len(rawVariants))
	for i, v := range rawVariants {
		variants[i] = dto.ProductVariant{
			ID:            v.ID,
			ProductID:     v.ProductID,
			Sku:           v.Sku,
			Price:         v.Price,
			StockQuantity: v.StockQuantity,
			Sold:          v.Sold,
			ImageUrl:      v.ImageUrl.String,
			IsDefault:     v.IsDefault,
		}
	}

	detail.Variants = variants

	return &detail, nil
}
