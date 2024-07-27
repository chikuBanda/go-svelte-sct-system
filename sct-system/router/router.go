package router

import (
	"github.com/gin-gonic/gin"
	"sct-system/middleware"
)

func StartRouting() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.POST("/register_beneficiary", registerNewBeneficiary)
	router.GET("/get_beneficiaries", getBeneficiaries)
	router.POST("/initiate_transfer", initiateTransfer)
	router.GET("/check_transfer_status/:transfer_id", checkTransferStatus)
	router.POST("/get_transaction_history", getTransactionHistory)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
