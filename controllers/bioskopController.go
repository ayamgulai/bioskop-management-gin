package controllers

import (
	"net/http"

	config "bioskop-management-gin/config"
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
		ctx.AbortWithError(http.StatusInternalServerError, err)
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
