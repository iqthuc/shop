package repository

import (
	"context"
	"log/slog"
	"shop/internal/features/product/core"
	"shop/internal/features/product/core/dto"
	"shop/internal/features/product/core/entity"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"
	"shop/pkg/utils/errorx"
	safetype "shop/pkg/utils/safe_type"
)

type repository struct {
	store store.Store
}

func NewProductPostgreRepo(store store.Store) core.Repository {
	return repository{
		store: store,
	}
}

func (repo repository) FetchProducts(ctx context.Context, params dto.GetProductsParams) ([]entity.Product, int, error) {
	safeLimit, err1 := safetype.SafeIntToInt32(params.Limit)
	safeOffset, err2 := safetype.SafeIntToInt32(params.Offset)
	if err1 != nil || err2 != nil {
		return nil, 0, errorx.ErrOverflow
	}

	slog.Info("check", slog.Int("value", params.Offset))
	pr := db.GetProductsParams{
		Limit:         safeLimit,
		Offset:        safeOffset,
		KeyWord:       params.Filters.Keyword,
		SortColumn:    params.SortBy.Field,
		SortDirection: params.SortBy.Order,
	}
	result, err := repo.store.GetProducts(ctx, pr)
	if err != nil {
		slog.Debug("query db error", slog.String("error", err.Error()))
		return nil, 0, err
	}

	totalCount, err := repo.store.GetProductsCount(ctx)
	if err != nil {
		slog.Debug("query db error", slog.String("error", err.Error()))
	}

	products := make([]entity.Product, 0, len(result))
	for _, p := range result {
		products = append(products, entity.Product{
			ID:        int(p.ID),
			Name:      p.Name,
			BasePrice: p.BasePrice,
		})
	}

	return products, int(totalCount), nil
}

func (repo repository) GetProductByID(ctx context.Context, productID int) (*entity.ProductDetail, error) {
	safeProductID, err := safetype.SafeIntToInt32(productID)
	if err != nil {
		return nil, errorx.ErrOverflow
	}

	p, err := repo.store.GetProductDetails(ctx, safeProductID)
	if err != nil {
		slog.Debug("query db error", slog.String("error", err.Error()))
		return nil, err
	}

	product := entity.ProductDetail{
		ID:           int(p.ID),
		Name:         p.Name,
		Slug:         p.Slug,
		Desciprtion:  p.Desciprtion.String,
		MainImageUrl: p.MainImageUrl.String,
		BasePrice:    p.BasePrice,
		CategoryID:   int(p.CategoryID),
		CategoryName: p.CategoryName.String,
		BrandID:      int(p.BrandID),
		BrandName:    p.BrandName.String,
	}

	return &product, nil
}

func (repo repository) FetchProductVariantByID(ctx context.Context, productID int) ([]entity.ProductVariant, error) {
	safeProductID, err := safetype.SafeIntToInt32(productID)
	if err != nil {
		return nil, errorx.ErrOverflow
	}

	raws, err := repo.store.GetProductVariants(ctx, safeProductID)
	if err != nil {
		slog.Debug("query db error", slog.String("error", err.Error()))
		return nil, err
	}

	pv := make([]entity.ProductVariant, 0, len(raws))
	for _, v := range raws {
		p := entity.ProductVariant(v)
		pv = append(pv, p)
	}

	return pv, nil
}
