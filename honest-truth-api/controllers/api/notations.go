package api

type CreateNotationRequest struct {
	VideoID     string `json:"video_id" binding:"required"`
	Time        uint64 `json:"time" binding:"required"`
	Description string `json:"description" binding:"required"`
	EditKey     string `json:"edit_key" binding:"required"`
}

type NotationResponse struct {
	ID          string `json:"notation_id"`
	VideoID     string `json:"video_id"`
	Time        uint64 `json:"timeStamp"`
	Description string `json:"description"`
}
