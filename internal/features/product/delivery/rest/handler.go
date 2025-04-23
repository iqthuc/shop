package rest

import (
	"log/slog"
	"shop/internal/features/product/core"
	"shop/internal/features/product/core/dto"
	"shop/pkg/utils/errorx"
	"shop/pkg/utils/messages"
	"shop/pkg/utils/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	useCase   core.ProductUseCase
	validator validator.Validate
}

func NewHandler(useCase core.ProductUseCase, validator validator.Validate) handler {
	return handler{
		useCase:   useCase,
		validator: validator,
	}
}
func (h handler) GetProductDetail(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("id")
	if err != nil {
		slog.Debug("convert params error", slog.String("value", c.Params("product_id")))
		return response.ErrorJson(c, errorx.ErrGetProductDetailConvertParamError, fiber.StatusBadRequest)
	}

	result, err := h.useCase.GetProductDetail(c.Context(), productID)
	if err != nil {
		return response.ErrorJson(c, errorx.ErrGetProductDetailFailed, fiber.StatusBadRequest)
	}

	return response.SuccessJson(c, result, messages.GetProductDetailSuccess.String())
}

func (h handler) GetProducts(c *fiber.Ctx) error {
	page := c.QueryInt("page")
	sortField := c.Query("sort_field")
	sortOrder := c.Query("sort_order")
	keyWord := c.Query("key_word")

	req := dto.GetProductsRequest{
		Page: page,
		Filters: dto.Filter{
			Keyword: keyWord,
		},
		SortBy: dto.SortBy{
			Field: sortField,
			Order: sortOrder,
		},
	}

	rawResult, err := h.useCase.GetProducts(c.Context(), req)
	if err != nil {
		slog.Debug("failed to get products ", slog.String("error", err.Error()))
		return response.ErrorJson(c, errorx.ErrGetProductsFailed, fiber.StatusInternalServerError)
	}

	products := make([]dto.Product, 0, len(rawResult.Items))
	for _, p := range rawResult.Items {
		products = append(products, dto.Product{
			ID:        p.ID,
			Name:      p.Name,
			BasePrice: p.BasePrice,
		})
	}

	var sortBy *dto.SortBy
	if (req.SortBy != dto.SortBy{}) {
		sortBy = &req.SortBy
	}
	var filter *dto.Filter
	if (req.Filters != dto.Filter{}) {
		filter = &req.Filters
	}

	result := dto.GetProductsResult{
		Items:      products,
		Filter:     filter,
		Pagination: rawResult.Pagination,
		SortBy:     sortBy,
	}

	return response.SuccessJson(c, result, messages.GetProductsSuccess.String())
}
