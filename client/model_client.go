package client

import (
	"context"
	"fmt"
)

type modelClient[M REST] struct {
	HTTP *HTTP
}

func (mc modelClient[M]) BaseURL() string {
	var model M
	return fmt.Sprintf("/api/v1/%s", model.PluralName())
}

func (mc modelClient[M]) Get(ctx context.Context, id int64) (M, error) {
	var data M

	_, err := mc.HTTP.Do(ctx, requestSpec{
		method: "GET",
		url:    fmt.Sprintf("%s/%d", mc.BaseURL(), id),
	}, &data)

	return data, err
}

func (mc modelClient[M]) All(ctx context.Context) ([]M, error) {
	var data []M

	_, err := mc.HTTP.Do(ctx, requestSpec{
		method: "GET",
		url:    mc.BaseURL(),
	}, &data)

	return data, err
}

func (mc modelClient[M]) Delete(ctx context.Context, id int64) error {
	_, err := mc.HTTP.Do(ctx, requestSpec{
		method: "DELETE",
		url:    fmt.Sprintf("%s/%d", mc.BaseURL(), id),
	}, nil)

	return err
}

func (mc modelClient[M]) Update(ctx context.Context, data M) (M, error) {
	var updated M

	_, err := mc.HTTP.Do(ctx, requestSpec{
		method: "PUT",
		url:    fmt.Sprintf("%s/%d", mc.BaseURL(), data.ID()),
		body:   data,
	}, &updated)

	return data, err
}

func (mc modelClient[M]) Create(ctx context.Context, data M) (M, error) {
	var updated M

	_, err := mc.HTTP.Do(ctx, requestSpec{
		method: "POST",
		url:    mc.BaseURL(),
		body:   data,
	}, &updated)

	return data, err
}
