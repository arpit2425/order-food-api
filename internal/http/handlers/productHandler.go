package handlers

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/helpers"
	"oilio.com/internal/model"
)

func GetProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		products := []model.Product{
			{
				ID: "1",
				Image: model.Image{
					Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
					Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
					Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
					Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
				},
				Name:     "Waffle with Berries",
				Category: "Waffle",
				Price:    6.5,
			},
			{
				ID: "2",
				Image: model.Image{
					Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-thumbnail.jpg",
					Mobile:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-mobile.jpg",
					Tablet:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-tablet.jpg",
					Desktop:   "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-desktop.jpg",
				},
				Name:     "Vanilla Bean Crème Brûlée",
				Category: "Crème Brûlée",
				Price:    7,
			},
			{
				ID: "3",
				Image: model.Image{
					Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-macaron-thumbnail.jpg",
					Mobile:    "https://orderfoodonline.deno.dev/public/images/image-macaron-mobile.jpg",
					Tablet:    "https://orderfoodonline.deno.dev/public/images/image-macaron-tablet.jpg",
					Desktop:   "https://orderfoodonline.deno.dev/public/images/image-macaron-desktop.jpg",
				},
				Name:     "Macaron Mix of Five",
				Category: "Macaron",
				Price:    8,
			},
			{
				ID: "4",
				Image: model.Image{
					Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-tiramisu-thumbnail.jpg",
					Mobile:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-mobile.jpg",
					Tablet:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-tablet.jpg",
					Desktop:   "https://orderfoodonline.deno.dev/public/images/image-tiramisu-desktop.jpg",
				},
				Name:     "Classic Tiramisu",
				Category: "Tiramisu",
				Price:    5.5,
			},
			{
				ID: "5",
				Image: model.Image{
					Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-baklava-thumbnail.jpg",
					Mobile:    "https://orderfoodonline.deno.dev/public/images/image-baklava-mobile.jpg",
					Tablet:    "https://orderfoodonline.deno.dev/public/images/image-baklava-tablet.jpg",
					Desktop:   "https://orderfoodonline.deno.dev/public/images/image-baklava-desktop.jpg",
				},
				Name:     "Pistachio Baklava",
				Category: "Baklava",
				Price:    4,
			},
			{
				ID: "6",
				Image: model.Image{
					Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-meringue-thumbnail.jpg",
					Mobile:    "https://orderfoodonline.deno.dev/public/images/image-meringue-mobile.jpg",
					Tablet:    "https://orderfoodonline.deno.dev/public/images/image-meringue-tablet.jpg",
					Desktop:   "https://orderfoodonline.deno.dev/public/images/image-meringue-desktop.jpg",
				},
				Name:     "Lemon Meringue Pie",
				Category: "Pie",
				Price:    5,
			},
			{
				ID: "7",
				Image: model.Image{
					Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-cake-thumbnail.jpg",
					Mobile:    "https://orderfoodonline.deno.dev/public/images/image-cake-mobile.jpg",
					Tablet:    "https://orderfoodonline.deno.dev/public/images/image-cake-tablet.jpg",
					Desktop:   "https://orderfoodonline.deno.dev/public/images/image-cake-desktop.jpg",
				},
				Name:     "Red Velvet Cake",
				Category: "Cake",
				Price:    4.5,
			},
			{
				ID: "8",
				Image: model.Image{
					Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-brownie-thumbnail.jpg",
					Mobile:    "https://orderfoodonline.deno.dev/public/images/image-brownie-mobile.jpg",
					Tablet:    "https://orderfoodonline.deno.dev/public/images/image-brownie-tablet.jpg",
					Desktop:   "https://orderfoodonline.deno.dev/public/images/image-brownie-desktop.jpg",
				},
				Name:     "Salted Caramel Brownie",
				Category: "Brownie",
				Price:    4.5,
			},
			{
				ID: "9",
				Image: model.Image{
					Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-thumbnail.jpg",
					Mobile:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-mobile.jpg",
					Tablet:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-tablet.jpg",
					Desktop:   "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-desktop.jpg",
				},
				Name:     "Vanilla Panna Cotta",
				Category: "Panna Cotta",
				Price:    6.5,
			},
		}
		return helpers.Success(c, products)
	}
}
