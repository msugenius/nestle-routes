package query

import (
	"context"
	"nestle/internal/models"

	"github.com/Masterminds/squirrel"
)

func (c Core) SelectAgents(ctx context.Context, region string, day int) (routes []models.Route) {
	query := squirrel.Select("Route_id", "RouteOwner").
		From("[boba].[dbo].[NestleRoutes]").Distinct().
		OrderBy("RouteOwner")

	if region != "" {
		query = query.Where(squirrel.Eq{"Territory": region})
	}

	if day != 8 {
		query = query.Where(squirrel.Eq{"Weekday": day})
	}

	sql, args, err := query.PlaceholderFormat(squirrel.AtP).ToSql()
	if err != nil {
		c.Logger.Fatal(err.Error())
	}

	rows, err := c.DB.QueryxContext(ctx, sql, args...)
	if err != nil {
		c.Logger.Fatal(err.Error())
	}

	var route models.Route
	for rows.Next() {
		err := rows.StructScan(&route)
		if err != nil {
			c.Logger.Fatal(err.Error())
		}
		routes = append(routes, route)
	}
	return
}
