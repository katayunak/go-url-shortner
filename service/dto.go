package service

type FindRequest struct {
	Short string
}

type FindResponse struct {
	Long string
}

type CreateRequest struct {
	Long string
}

type CreateResponse struct {
	Short string
}
