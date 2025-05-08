package types

type BorrowReq struct {
	BookName   string `json:"bookName"`
	ReturnPlan int64  `json:"returnPlan"`
}

type ReturnReq struct {
	BookName string `json:"bookName"`
}
