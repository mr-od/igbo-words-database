// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: search.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const searchProduct = `-- name: SearchProduct :many
SELECT id, name, owner, price, description, imgs_url, imgs_name, tsv, created_at
FROM products
WHERE tsv @@ to_tsquery($1)
ORDER BY created_at DESC
LIMIT 100
`

func (q *Queries) SearchProduct(ctx context.Context, searchQuery string) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, searchProduct, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Owner,
			&i.Price,
			&i.Description,
			pq.Array(&i.ImgsUrl),
			pq.Array(&i.ImgsName),
			&i.Tsv,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
