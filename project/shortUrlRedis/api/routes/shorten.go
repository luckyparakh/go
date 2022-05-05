package routes

import "time"

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining int           `json:"rate_remaining"`
	XRateReset     time.Duration `json:"rate_reset"`
}

func ShortenURL(ctx *fiber.ctx) error{
		body :=new(request)
		if err := c.BodyParser(&body); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.map("error:Cannot Parse JSON"))
        }

		//implement rate limit


		//check if the input is valid URL
		if !govalidator.IsURL(body.URL){
			return c.Status(fiber.StatusBadRequest).JSON(fiber.map("error:Provide correct URL"))
		}

		//check for domain error
		if !helper.RemoveDomainError(body.URL){
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.map("error:Cannot serve localhost"))
		}

		//enforce ssl,https
		body.URL = helper.EnforceHTTP(body.URL)
}