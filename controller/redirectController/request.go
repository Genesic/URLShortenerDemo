package redirectController

type Request struct {
	ShortenId string `uri:"shorten_id" binding:"required"`
}
