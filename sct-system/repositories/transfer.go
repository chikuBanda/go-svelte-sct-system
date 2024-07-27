package repositories

import (
	"gorm.io/gorm"
	"sct-system/database"
	"sct-system/models"
	"time"
)

func GetTransfers() ([]models.Transfer, error) {
	var transfers []models.Transfer
	err := database.DB.Model(&models.Transfer{}).Preload("Transfers").Find(&transfers).Error
	return transfers, err
}

func GetTransfersById(id int) (models.Transfer, error) {
	var transfer models.Transfer
	err := database.DB.Model(&models.Transfer{}).
		Where("id = ?", id).
		First(&transfer).Error
	return transfer, err
}

func GetTransfersByBeneficiary(beneficiaryId uint) ([]models.Transfer, error) {
	var transfers []models.Transfer
	err := database.DB.Model(&models.Transfer{}).
		Where("beneficiary_id = ?", beneficiaryId).
		Preload("TransactionHistory").
		Find(&transfers).Error
	return transfers, err
}

func CreateTransfer(newTransfer *models.Transfer) error {
	// Use an atomic transaction to create both the transfer and transaction history
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var beneficiary models.Beneficiary
		tx.Model(&models.Beneficiary{}).
			Where("id = ?", newTransfer.BeneficiaryId).
			First(&beneficiary)

		err := tx.Model(&beneficiary).Association("Transfers").Append(newTransfer)
		if err != nil {
			return err
		}

		var newTransactionHistory models.TransactionHistory
		newTransactionHistory.TransferID = newTransfer.ID
		newTransactionHistory.Timestamp = time.Now()
		newTransactionHistory.Status = "SENT"

		err = tx.Model(&models.TransactionHistory{}).Create(&newTransactionHistory).Error
		if err != nil {
			return err
		}

		newTransfer.TransactionHistory = newTransactionHistory

		// return nil will commit the whole transaction
		return nil
	})

	return err
}
