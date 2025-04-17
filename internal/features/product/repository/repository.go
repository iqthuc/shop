package repository

import (
	"context"
	"log/slog"
	"shop/internal/features/product/dto"
	"shop/internal/features/product/entity"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/database/store/db"
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
