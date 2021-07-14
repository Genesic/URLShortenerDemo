package shortenController

type Request struct {
	Url       string `json:"url" binding:"required"`
	ExpiredAt string `json:"expiredAt" binding:"required"`
}
