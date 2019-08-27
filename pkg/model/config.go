package model

import "time"

const (
	// BorrowingPeriods 為圖書館的借閱期限，預設為 30 天
	BorrowingPeriods = time.Hour * 24 * 30
)
