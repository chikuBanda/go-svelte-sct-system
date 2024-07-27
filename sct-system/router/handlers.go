package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sct-system/models"
	"sct-system/repositories"
	"sct-system/utils"
	"strconv"
)

func registerNewBeneficiary(c *gin.Context) {
	var newBeneficiary models.Beneficiary

	if err := c.ShouldBindJSON(&newBeneficiary); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateBeneficiary(&newBeneficiary); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newBeneficiary)
}

func getBeneficiaries(c *gin.Context) {
	beneficiaries, err := repositories.GetBeneficiaries()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, beneficiaries)
}

func initiateTransfer(c *gin.Context) {
	var newTransfer models.Transfer

	if err := c.ShouldBindJSON(&newTransfer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTransfer.Status = utils.GenerateRandomTransferStatus()

	if err := repositories.CreateTransfer(&newTransfer); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newTransfer)
}

func getTransactionHistory(c *gin.Context) {
	body := struct {
		BeneficiaryId uint `json:"beneficiary_id" form:"beneficiary_id" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transfers, err := repositories.GetTransfersByBeneficiary(body.BeneficiaryId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, transfers)
}

func checkTransferStatus(c *gin.Context) {
	transferIdStr := c.Param("transfer_id")

	transferId, err := strconv.Atoi(transferIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transfer, err := repositories.GetTransfersById(transferId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": transfer.Status})
}
