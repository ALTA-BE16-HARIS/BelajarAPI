package users

import (
	"belajarAPI/responses"
	"net/http"

	"github.com/labstack/echo"
)

var DataUsers = make(map[string]User, 0)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req User
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, responses.CreateRespon(http.StatusBadRequest, nil, "BAD REQUEST", "Username/Password Salah"))
		}
		if data, ok := DataUsers[req.Phone]; ok {
			if data.Password == req.Password {
				return c.JSON(http.StatusOK, responses.CreateRespon(http.StatusOK, DataUsers[req.Phone], "Success", "Berhasil Login"))
			}
		}
		return c.JSON(http.StatusBadRequest, responses.CreateRespon(http.StatusBadRequest, nil, "BAD REQUEST", "Username/Password Salah"))
	}
}
func Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req User
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, responses.CreateRespon(http.StatusBadRequest, nil, "BAD REQUEST", "Gagal Mendaftar"))
		}
		DataUsers[req.Phone] = req
		return c.JSON(http.StatusOK, responses.CreateRespon(http.StatusOK, nil, "Success", "selamat anda telah terdaftar"))
	}
}
