package controllers

import (
	"net/http"

	config "bioskop-management-gin/configs"
	"bioskop-management-gin/models"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(ctx *gin.Context) {
	var newBioskop models.Bioskop

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if newBioskop.Nama == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "nama tidak boleh kosong",
		})
		return
	}

	if newBioskop.Lokasi == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "lokasi tidak boleh kosong",
		})
		return
	}

	query := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := config.DB.QueryRow(
		query,
		newBioskop.Nama,
		newBioskop.Lokasi,
		newBioskop.Rating,
	).Scan(&newBioskop.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"bioskop": newBioskop,
	})
}

func ShowAllBioskop(ctx *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, nama, lokasi, rating FROM bioskop`,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var bioskops []models.Bioskop

	for rows.Next() {
		var bioskop models.Bioskop
		if err := rows.Scan(
			&bioskop.ID,
			&bioskop.Nama,
			&bioskop.Lokasi,
			&bioskop.Rating,
		); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		bioskops = append(bioskops, bioskop)
	}

	ctx.JSON(http.StatusOK, bioskops)
}

func ShowBioskopByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var bioskop models.Bioskop
	err := config.DB.QueryRow(
		`SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1`,
		id,
	).Scan(
		&bioskop.ID,
		&bioskop.Nama,
		&bioskop.Lokasi,
		&bioskop.Rating,
	)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "bioskop tidak ditemukan",
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, bioskop)
}

func UpdateBioskop(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedBioskop models.Bioskop
	if err := ctx.ShouldBindJSON(&updatedBioskop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if updatedBioskop.Nama == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "nama tidak boleh kosong",
		})
		return
	}

	if updatedBioskop.Lokasi == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "lokasi tidak boleh kosong",
		})
		return
	}

	err := config.DB.QueryRow(
		`UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4 RETURNING id`,
		updatedBioskop.Nama,
		updatedBioskop.Lokasi,
		updatedBioskop.Rating,
		id,
	).Scan(&updatedBioskop.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"bioskop": updatedBioskop,
	})

}

func DeleteBioskop(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := config.DB.Exec(
		`DELETE FROM bioskop WHERE id = $1`,
		id,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "bioskop tidak ditemukan",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "bioskop berhasil dihapus",
	})

}
