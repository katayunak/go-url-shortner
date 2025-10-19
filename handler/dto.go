package handler

type CreateRequest struct {
	Long string `json:"long"`
}

type CreateResponse struct {
	Short string `json:"shot"`
}
