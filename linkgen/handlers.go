package linkgen

import (
	"encoding/json"
	"io"
	"linkgen/core"
	"linkgen/pkg/baseresponse"
	"net/http"
)

type generateMinifiedLinkRequest struct {
	Link string `json:"link"`
}

type generateMinifiedLinkResponse struct {
	Link    string `json:"link"`
	ShortID string `json:"shortid"`
}

func generateMinifiedLinkRequestSerializer(body io.ReadCloser) (request generateMinifiedLinkRequest, err error) {
	decoder := json.NewDecoder(body)
	err = decoder.Decode(&request)
	return
}

// GenerateMinifiedLink - generates and save a shortid for a given URL
func (s *Server) GenerateMinifiedLink(w http.ResponseWriter, req *http.Request) {
	request, err := generateMinifiedLinkRequestSerializer(req.Body)
	if err != nil {
		response := baseresponse.BaseResponse{Errors: []string{"failed to decode request"}}
		response.Fail(w, nil)
		return
	}
	shortid, err := core.GenerateNewShortID()
	if err != nil {
		response := baseresponse.BaseResponse{Errors: []string{"failed to generate url shortid"}}
		response.Fail(w, nil)
		return
	}
	if err = s.LinkStore.AddLinkMapping(request.Link, shortid); err != nil {
		response := baseresponse.BaseResponse{Errors: []string{"failed to add link"}}
		response.Fail(w, nil)
		return
	}
	response := baseresponse.BaseResponse{Message: "Link Added!"}
	response.Succeed(w, &generateMinifiedLinkResponse{
		Link:    request.Link,
		ShortID: shortid,
	})
}

func (s *Server) RedirectToOriginalURL(w http.ResponseWriter, req *http.Request) {
	requestParams := req.Context().Value(paramsKey).(map[string]string)
	shortid := requestParams["code"]
	originalURL, err := s.LinkStore.GetOriginal(shortid)
	if err != nil {
		response := baseresponse.BaseResponse{Errors: []string{"failed to get link"}}
		response.Fail(w, nil)
		return
	}
	if originalURL == "" {
		response := baseresponse.BaseResponse{Errors: []string{"URL not found"}}
		response.FailWithCode(w, nil, 404)
		return
	}
	http.Redirect(w, req, originalURL, 302)
}
