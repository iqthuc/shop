package repository

import (
	"context"
	"log/slog"
	"shop/internal/features/product/dto"
	"shop/internal/features/product/entity"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type repository struct {
	store store.Store
}

func NewRepository(store store.Store) repository {
	return repository{
		store: store,
	}
}

func (repo repository) FetchProducts(ctx context.Context, params dto.GetProductsParams) ([]entity.Product, int, error) {
	pr := db.GetProductsParams{
		Limit:         int32(params.Limit),
		Offset:        int32(params.Offset),
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
	p, err := repo.store.GetProductDetails(ctx, int32(productID))
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
		BasePrice:    p.PBasePrice,
		CategoryID:   int(p.CategoryID.Int32),
		CategoryName: p.CategoryName.String,
		BrandID:      int(p.BrandID.Int32),
		BrandName:    p.BrandName.String,
	}

	return &product, nil
}

func (repo repository) FetchProductVariantByID(
	ctx context.Context,
	productID int,
) ([]entity.ProductVariant, error) {
	raws, err := repo.store.GetProductVariants(ctx, pgtype.Int4{Int32: int32(productID), Valid: true})
	if err != nil {
		slog.Debug("query db error", slog.String("error", err.Error()))
		return nil, err
	}
	pv := make([]entity.ProductVariant, 0, len(raws))
	for _, v := range raws {
		p := entity.ProductVariant{
			ID: int(v.ID),
		}
		pv = append(pv, p)
	}

	return pv, nil
}
