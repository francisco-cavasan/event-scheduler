package Controllers

import (
	"net/http"

	"where_my_pet_at/server/Models"
	"where_my_pet_at/server/Services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PetController struct {
	DB *gorm.DB
}

// Index retrieves a list of all pets
func (pc *PetController) Index(c *gin.Context) {
	var pets []Models.Pet
	pc.DB.Find(&pets)
	c.JSON(http.StatusOK, pets)
}

// Store creates a new pet
func (pc *PetController) Store(c *gin.Context) {
	var pet Models.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pc.DB.Create(&pet)
	c.JSON(http.StatusCreated, pet)
}

// Update updates an existing pet
func (pc *PetController) Update(c *gin.Context) {
	id := c.Param("id")
	var pet Models.Pet
	if err := pc.DB.First(&pet, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "pet not found"})
		return
	}

	if err := c.ShouldBindJSON(&pet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pc.DB.Save(&pet)
	c.JSON(http.StatusOK, pet)
}

// Get retrieves a specific pet by ID
func (pc *PetController) Get(c *gin.Context) {
	id := c.Param("id")
	var pet Models.Pet
	if err := pc.DB.First(&pet, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "pet not found"})
		return
	}

	c.JSON(http.StatusOK, pet)
}

// Delete removes a pet by ID
func (pc *PetController) Delete(c *gin.Context) {
	id := c.Param("id")
	var pet Models.Pet
	if err := pc.DB.First(&pet, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "pet not found"})
		return
	}

	pc.DB.Delete(&pet)
	c.Status(http.StatusNoContent)
}

func (pc *PetController) AddPetFoundLocation(c *gin.Context) {
	id := c.PostForm("id")
	var pet Models.Pet

	if err := pc.DB.First(&pet, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "pet not found"})
		return
	}

	// Parse request body into location object
	var location Models.Location
	if err := c.BindJSON(&location); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new pet location record
	petLocation := Models.PetLocation{Pet: pet, Location: location}

	// Save the new pet location record to the database
	if err := pc.DB.Create(&petLocation).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send an email to the pet owner

	Services.Handle("Your pet has been found!", pet.Owner.Email)

	c.JSON(http.StatusOK, gin.H{"message": "location added to pet"})
}
