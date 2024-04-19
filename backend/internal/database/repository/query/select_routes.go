package query

import (
	"context"
	"fmt"
	"nestle/internal/models"

	"github.com/Masterminds/squirrel"
)

func (c Core) SelectAgents(ctx context.Context, region string, days []int) (routes []models.Route) {
	query := squirrel.Select("Route_id", "RouteOwner", "Weekday").
		From("[boba].[dbo].[NestleRoutes]").Distinct().
		OrderBy("RouteOwner", "Weekday")

	if region != "" {
		query = query.Where(squirrel.Eq{"Territory": region})
	}

	if len(days) == 1 {
		if days[0] != 8 {
			query = query.Where(squirrel.Eq{"Weekday": days[0]})
		}
	} else if len(days) > 1 {
		orStatement := squirrel.Or{}
		for _, day := range days {
			orStatement = append(orStatement, squirrel.Eq{"Weekday": day})
		}
		query = query.Where(orStatement)
	}

	sql, args, err := query.PlaceholderFormat(squirrel.AtP).ToSql()
	if err != nil {
		c.Logger.Fatal(err.Error())
	}

	fmt.Println(sql)

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
