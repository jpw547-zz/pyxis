package handlers

import (
	"net/http"

	"github.com/jpw547/pyxis/server/helpers"
	"github.com/labstack/echo"
)

// SearchFoodPlaces searches the possible food places and returns the list based on any given parameters.
func SearchFoodPlaces(ctx echo.Context) error {
	x := make(map[string][]string)

	x["categories"] = append(x["categories"], "hotdogs")

	helpers.BusinessSearch(40.2338, -111.6585, x)

	return ctx.String(http.StatusOK, "Ok")
}
