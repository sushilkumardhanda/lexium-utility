package datahandler

type SelectedITR struct {
	ITR string `json:"ITR" binding:"required"`
}
type SelectedITR_Schema struct {
	ITR    string `json:"ITR" binding:"required"`
	Schema string `json:"Schema" binding:"required"`
}
type SelectedElement struct {
	ITR    string `json:"ITR" binding:"required"`
	Schema string `json:"Schema" binding:"required"`
	ElementID string `json:"ElementID" binding:"required"`
}
