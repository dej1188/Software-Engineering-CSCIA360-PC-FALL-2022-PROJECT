package api

type CreateVideoRequest struct {
	YouTubeID   string `json:"youtube_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	UserHash    string `json:"user_hash" binding:"required"`
	Description string `json:"description"`
}

type VideoResponse struct {
	ID          string             `json:"video_id"`
	Title       string             `json:"title"`
	UserHash    string             `json:"user_hash"`
	Description string             `json:"description"`
	YouTubeID   string             `json:"youtube_id"`
	Notations   []NotationResponse `json:"notations"`
}

type CreateVideoResponse struct {
	ID          string             `json:"video_id"`
	Title       string             `json:"title"`
	UserHash    string             `json:"user_hash"`
	Description string             `json:"description"`
	YouTubeID   string             `json:"youtube_id"`
	EditKey     string             `json:"edit_key"'`
	Notations   []NotationResponse `json:"notations"`
}
