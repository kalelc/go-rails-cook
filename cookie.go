package rails5cook

type Cookie struct {
	Value         string
	Data          []byte
	Iv            []byte
	SecretKeyBase string
	Salt          string
	Content       *Content
}

type Content struct {
	SessionID string `json:"session_id"`
	Csrf      string `json:"_csrf_token"`
}
