package main

import (
	"fmt"
	//"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func soma(c *fiber.Ctx) error {
	op1Str := c.Params("op1")
	op2Str := c.Params("op2")
	op1, err := strconv.ParseFloat(op1Str, 64) 
	 if err != nil {
		return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Parametro Invalido:\":%s\"", op1Str))
	}
	op2, err := strconv.ParseFloat(op2Str, 64) 
	 if err != nil {
		return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Parametro Invalido:\":%s\"", op2Str))
	 }
	soma := op1+op2
	return c.SendString(fmt.Sprintf("%.2f", soma))

}
