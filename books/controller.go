package books

import (
	"belajarAPI/responses"
	"net/http"

	"github.com/labstack/echo"
)

var DataBooks = make(map[string][]Book, 0)

func Addbook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req Book
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, responses.CreateRespon(http.StatusBadRequest, nil, "BAD REQUEST", "Gagal Menambahkan Data Buku"))
		}
		DataBooks["books"] = append(DataBooks["books"], req)
		return c.JSON(http.StatusOK, responses.CreateRespon(http.StatusOK, nil, "Success", "data buku yang berhasil diinputkan"))
	}
}
func Getallbooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, responses.CreateRespon(http.StatusOK, DataBooks["books"], "Success", "data buku yang berhasil diinputkan"))
	}
}
