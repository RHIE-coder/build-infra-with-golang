package handler

import "github.com/gofiber/fiber/v2"

func AddAccount(ctx *fiber.Ctx) error {
	return ctx.SendString("AddAccount")
}

/*
ShowAccount godoc

@Summary      Show an account
@Description  get string by ID
@Tags         accounts
@Accept       json
@Produce      json
@Param        id   path      int  true  "Account ID"
@Success      200  {object}  model.Account
@Failure      400  {object}  httputil.HTTPError
@Failure      404  {object}  httputil.HTTPError
@Failure      500  {object}  httputil.HTTPError
@Router       /accounts/{id} [get]
*/
func ShowAccount(ctx *fiber.Ctx) error {
	return ctx.SendString("ShowAccount")
}

// ListAccounts lists all existing accounts
//
//  @Summary      List accounts
//  @Description  get accounts
//  @Tags         accounts
//  @Accept       json
//  @Produce      json
//  @Param        q    query     string  false  "name search by q"  Format(email)
//  @Success      200  {array}   model.Account
//  @Failure      400  {object}  httputil.HTTPError
//  @Failure      404  {object}  httputil.HTTPError
//  @Failure      500  {object}  httputil.HTTPError
//  @Router       /accounts [get]
func ListAccount(ctx *fiber.Ctx) error {
	return ctx.SendString("ListAccount")
}

func UpdateAccount(ctx *fiber.Ctx) error {
	return ctx.SendString("UpdateAccount")
}

func DeleteAccount(ctx *fiber.Ctx) error {
	return ctx.SendString("DeleteAccount")
}
