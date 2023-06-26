package main

import (
	"github.com/arriqaaq/flashdb/pkg/flashdb"
	"github.com/gin-gonic/gin"
)

func main() {
	db := flashdb.NewFlashDB()

	router := gin.Default()

	router.POST("/add", func(c *gin.Context) {
		var word flashdb.Word
		c.BindJSON(&word)

		// pour ADD dans la base de données FlashDB
		db.Add(word.Key, word.Value)

		c.JSON(200, gin.H{
			"message": "Mot ajouté avec succes",
		})
	})

	router.GET("/define", func(c *gin.Context) {
		word := c.Query("word")

		// pour DEFINE dans la base de données FlashDB
		definition, ok := db.Get(word)
		if !ok {
			definition = "Mot non trouvé"
		}

		c.JSON(200, gin.H{
			"definition": definition,
		})
	})

	router.POST("/remove", func(c *gin.Context) {
		var word flashdb.Word
		c.BindJSON(&word)

		// pour REMOVE dans la base de données FlashDB
		db.Remove(word.Key)

		c.JSON(200, gin.H{
			"message": "Mot retiré avec succès",
		})
	})

	router.GET("/list", func(c *gin.Context) {
		// pour LIST dans la base de données FlashDB et récupérer la liste des mots
		words := db.List()

		c.JSON(200, gin.H{
			"words": words,
		})
	})

	router.Run(":8080")
}
