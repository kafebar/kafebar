package postgres

import (
	"context"
	"fmt"
	"slices"

	sq "github.com/Masterminds/squirrel"
	kafebar "github.com/kafebar/kafebar/api/kafebar"
)

type ProductService struct {
	builder sq.StatementBuilderType
}

var _ kafebar.ProductService = (*ProductService)(nil)

func NewProductService(db sq.StdSqlCtx) *ProductService {
	builder := sq.StatementBuilder.RunWith(sq.WrapStdSqlCtx(db)).PlaceholderFormat(sq.Dollar)
	return &ProductService{
		builder: builder,
	}
}

func (o *ProductService) CreateProduct(ctx context.Context, product kafebar.Product) (kafebar.Product, error) {
	err := o.builder.Insert(tableProducts).
		Columns(columnName, columnPrice).
		Values(product.Name, product.Price).
		Suffix("RETURNING id").
		QueryRowContext(ctx).Scan(&product.Id)

	if err != nil {
		return product, fmt.Errorf("cannot insert product record: %w", err)
	}

	if len(product.AvailableOptions) > 0 {
		err := o.createProductOptions(ctx, product.Id, product.AvailableOptions)
		if err != nil {
			return product, err
		}
	}

	return product, nil
}

func (o *ProductService) UpdateProduct(ctx context.Context, product kafebar.Product) (kafebar.Product, error) {
	_, err := o.builder.Update(tableProducts).
		Set(columnName, product.Name).
		Set(columnPrice, product.Price).
		Where(sq.Eq{columnId: product.Id}).
		ExecContext(ctx)

	if err != nil {
		return product, fmt.Errorf("cannot update product record: %w", err)
	}

	if len(product.AvailableOptions) == 0 {
		return product, nil
	}

	productOptions, err := o.getProductOptions(ctx, sq.Eq{columnProductId: product.Id})
	if err != nil {
		return product, err
	}

	optsToRemove := []int{}
	optsToCreate := []string{}

	for _, opt := range productOptions {
		if !slices.Contains(product.AvailableOptions, opt.option) {
			optsToRemove = append(optsToRemove, opt.Id)
		}
	}

	for _, opt := range product.AvailableOptions {
		if !slices.ContainsFunc(productOptions, func(po productOption) bool {
			return po.option == opt
		}) {
			optsToCreate = append(optsToCreate, opt)
		}
	}

	if len(optsToRemove) > 0 {
		_, err := o.builder.Delete(tableProducts).
			Where(sq.Eq{columnId: optsToRemove}).
			Exec()
		if err != nil {
			return product, fmt.Errorf("cannot remove existing options: %w", err)
		}
	}

	if len(optsToCreate) > 0 {
		err := o.createProductOptions(ctx, product.Id, optsToCreate)
		if err != nil {
			return product, err
		}
	}

	return product, nil
}

func (o *ProductService) DeleteProduct(ctx context.Context, id int) error {
	_, err := o.builder.Delete(tableProducts).
		Where(sq.Eq{columnId: id}).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf("cannot delete product record: %w", err)
	}

	_, err = o.builder.Delete(tableProductOptions).
		Where(sq.Eq{columnProductId: id}).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf("cannot delete product options: %w", err)
	}

	return nil
}

func (o *ProductService) GetProducts(ctx context.Context) ([]kafebar.Product, error) {
	productRows, err := o.builder.
		Select(columnId, columnName, columnPrice).
		From(tableProducts).
		OrderBy(columnId).
		QueryContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("cannot get products: %w", err)
	}

	products := []kafebar.Product{}

	for productRows.Next() {
		var product kafebar.Product
		err := productRows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			return products, fmt.Errorf("cannot scan order item: %w", err)
		}
		product.AvailableOptions = []string{}
		products = append(products, product)
	}

	productOptions, err := o.getProductOptions(ctx, nil)
	if err != nil {
		return products, fmt.Errorf("cannot get product_options: %w", err)
	}

	for _, po := range productOptions {
		productIdx := slices.IndexFunc(products, func(i kafebar.Product) bool { return i.Id == po.productId })
		if productIdx == -1 {
			return products, fmt.Errorf("found option for non existing item")
		}
		products[productIdx].AvailableOptions = append(products[productIdx].AvailableOptions, po.option)
	}

	return products, nil
}

type productOption struct {
	Id        int
	productId int
	option    string
}

func (o *ProductService) getProductOptions(ctx context.Context, predicate any) ([]productOption, error) {
	productOptionRows, err := o.builder.
		Select(columnId, columnProductId, columnOption).
		From(tableProductOptions).
		Where(predicate).
		OrderBy(columnId).
		QueryContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("cannot get product_options: %w", err)
	}

	productOptions := []productOption{}

	for productOptionRows.Next() {
		var po productOption
		err := productOptionRows.Scan(&po.Id, &po.productId, &po.option)
		if err != nil {
			return nil, fmt.Errorf("cannot scan order item option: %w", err)
		}

		productOptions = append(productOptions, po)
	}

	return productOptions, nil
}

func (o *ProductService) createProductOptions(ctx context.Context, productId int, options []string) error {
	insertOptionsStmt := o.builder.Insert(tableProductOptions).
		Columns(columnProductId, columnOption)

	for _, opt := range options {
		insertOptionsStmt = insertOptionsStmt.Values(productId, opt)
	}

	_, err := insertOptionsStmt.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("cannot insert product options: %w", err)
	}
	return nil
}
