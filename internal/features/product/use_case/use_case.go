package use_case

import (
	"context"
	"shop/internal/features/product/dto"
	"shop/internal/features/product/entity"
	"strings"
)

type useCase struct {
	repo repository
}

func NewUseCase(repo repository) useCase {
	return useCase{
		repo: repo,
	}
}

func (u useCase) GetProducts(ctx context.Context, input dto.GetProductsRequest) (*dto.GetProductsResult[entity.Product], error) {
	currentPage := input.Page
	perpage := 15
	offset := (currentPage - 1) * perpage
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

	products, totalCount, err := u.repo.FetchProducts(ctx, params)
	if err != nil {
		return nil, err
	}

	// calculate pagination
	totalPages := (totalCount + perpage - 1) / perpage
	prevPage := max(currentPage-1, 1)
	nextPage := min(currentPage+1, totalPages)

	result := &dto.GetProductsResult[entity.Product]{
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

	return result, nil
}
