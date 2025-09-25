package exercise1

import (
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	ID      uint `gorm:"primaryKey"`
	Balance int  // 账户余额（单位：分或元，根据业务定）
}

type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	FromAccountID uint // 转出账户ID
	ToAccountID   uint // 转入账户ID
	Amount        int  // 转账金额
}

func TransferMoney(db *gorm.DB, fromID, toID uint, amount int) error {
	// 使用Transaction方法自动处理事务
	return db.Debug().Transaction(func(tx *gorm.DB) error {
		// 1. 检查转出账户余额
		var fromAccount Account
		if err := tx.First(&fromAccount, fromID).Error; err != nil {
			return fmt.Errorf("查询转出账户失败: %v", err)
		}
		if fromAccount.Balance < amount {
			return fmt.Errorf("账户余额不足")
		}

		// 2. 扣减转出账户余额
		if err := tx.Model(&Account{}).
			Where("id = ?", fromID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return fmt.Errorf("扣减余额失败: %v", err)
		}

		// 3. 增加转入账户余额
		if err := tx.Model(&Account{}).
			Where("id = ?", toID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return fmt.Errorf("增加余额失败: %v", err)
		}

		// 4. 记录转账交易
		transaction := Transaction{
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return fmt.Errorf("记录交易失败: %v", err)
		}
		var account []Account
		findAccount := tx.Debug().Find(&account)
		if findAccount.Error != nil {
			return fmt.Errorf("查询账户失败: %v", findAccount.Error)
		}
		for _, acc := range account {
			fmt.Printf("ID: %d, Balance: %d\n", acc.ID, acc.Balance)
		}
		// 返回nil自动提交事务
		return nil
	})
}

func Run1(db *gorm.DB) {
	db.AutoMigrate(&Account{}, &Transaction{})
	db.Create(&Account{ID: 1, Balance: 1000})
	db.Create(&Account{ID: 2, Balance: 500})
}
