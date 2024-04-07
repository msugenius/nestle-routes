package query

import (
	"context"
	"nestle/internal/models"

	"github.com/Masterminds/squirrel"
)

func (c Core) SelectPoints(ctx context.Context, route string) (points []models.Point) {
	query := squirrel.
		Select("Latitude", "Longitude", "OLDeliveryAddress", "OlName", "OL_id", "Weekday").
		From("[boba].[dbo].[NestleRoutes]")

	if route != "0" {
		query = query.Where(squirrel.Eq{"Route_id": route})
	}

	if route == "-1" {
		return
	}

	sql, args, err := query.PlaceholderFormat(squirrel.AtP).ToSql()
	if err != nil {
		c.Logger.Fatal(err.Error())
	}

	rows, err := c.DB.QueryxContext(ctx, sql, args...)
	if err != nil {
		c.Logger.Fatal(err.Error())
	}

	var point models.Point
	for rows.Next() {
		err := rows.StructScan(&point)
		if err != nil {
			c.Logger.Fatal(err.Error())
		}
		points = append(points, point)
	}
	return
}
