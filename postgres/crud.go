package postgres

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jeremybower/go-common/pagination"
)

func CountT(
	ctx context.Context,
	querier Querier,
	templ *Template,
	data map[string]any,
) (int64, error) {
	// Execute the template to build the SQL.
	sql, args, err := templ.ExecuteCount(data)
	if err != nil {
		return 0, NormalizeError(err)
	}

	return Count(ctx, querier, sql, args...)
}

func Count(
	ctx context.Context,
	querier Querier,
	sql string,
	args ...any,
) (int64, error) {
	// Execute the SQL to count the items.
	var count int64
	if err := pgxscan.Get(ctx, querier, &count, sql, args...); err != nil {
		return 0, NormalizeError(err)
	}

	// Success.
	return count, nil
}

func ExecT(
	ctx context.Context,
	querier Querier,
	templ *Template,
	data map[string]any,
) (int64, error) {
	// Execute the template to build the SQL.
	sql, args, err := templ.Execute(data)
	if err != nil {
		return 0, NormalizeError(err)
	}

	// Success.
	return Exec(ctx, querier, sql, args...)
}

func Exec(
	ctx context.Context,
	querier Querier,
	sql string,
	args ...any,
) (int64, error) {
	// Execute the SQL.
	commandTag, err := querier.Exec(ctx, sql, args...)
	if err != nil {
		return 0, NormalizeError(err)
	}

	// Success.
	return commandTag.RowsAffected(), nil
}

func ReadOneT[T any](
	ctx context.Context,
	querier Querier,
	templ *Template,
	data map[string]any,
) (*T, error) {
	// Execute the template to build the SQL.
	sql, args, err := templ.Execute(data)
	if err != nil {
		return nil, NormalizeError(err)
	}

	// Success.
	return ReadOne[T](ctx, querier, sql, args...)
}

func ReadOne[T any](
	ctx context.Context,
	querier Querier,
	sql string,
	args ...any,
) (*T, error) {
	// Read the item.
	var item T
	if err := pgxscan.Get(ctx, querier, &item, sql, args...); err != nil {
		return nil, NormalizeError(err)
	}

	// Success.
	return &item, nil
}

func ReadManyT[T any](
	ctx context.Context,
	querier Querier,
	templ *Template,
	data map[string]any,
) ([]*T, error) {
	// Execute the template to build the SQL.
	sql, args, err := templ.Execute(data)
	if err != nil {
		return nil, NormalizeError(err)
	}

	// Success.
	return ReadMany[T](ctx, querier, sql, args...)
}

func ReadMany[T any](
	ctx context.Context,
	querier Querier,
	sql string,
	args ...any,
) ([]*T, error) {
	// Read the items.
	var items []*T
	if err := pgxscan.Select(ctx, querier, &items, sql, args...); err != nil {
		return nil, NormalizeError(err)
	}

	// Success.
	return items, nil
}

func ListT[T any](
	ctx context.Context,
	querier Querier,
	templ *Template,
	data map[string]any,
	pageIndex int64,
	pageSize int64,
	minimumPageSize int64,
	defaultPageSize int64,
	maximumPageSize int64,
) (*pagination.Result[*T], error) {
	// Count the total items.
	totalItems, err := CountT(ctx, querier, templ, data)
	if err != nil {
		return nil, err
	}

	// Normalize pagination.
	norm := pagination.Normalize(totalItems, pageIndex, pageSize, minimumPageSize, defaultPageSize, maximumPageSize)

	// Execute the template to build the SQL.
	sql, args, err := templ.ExecuteList(data, norm.FirstItemIndex, norm.PageSize)
	if err != nil {
		return nil, NormalizeError(err)
	}

	// Execute the SQL to list the items.
	var items []*T
	if err := pgxscan.Select(ctx, querier, &items, sql, args...); err != nil {
		return nil, NormalizeError(err)
	}

	// Success.
	return &pagination.Result[*T]{
		PageIndex:      norm.PageIndex,
		PageSize:       norm.PageSize,
		FirstItemIndex: norm.FirstItemIndex,
		TotalItems:     norm.TotalItems,
		TotalPages:     norm.TotalPages,
		Items:          items,
	}, nil
}
