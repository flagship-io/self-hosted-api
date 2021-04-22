package httputils

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

//nolint
type hit struct {
	EnvironmentID string `json:"cid"`
	VisitorID     string `json:"vid"`
	Type          string `json:"t"`
}

// @Summary Send a hit
// @Tags Hits
// @Description Send a hit to Flagship datacollect
// @ID send-hit
// @Accept  json
// @Produce  image/gif
// @Param request body hit true "Hit request"
// @Success 200
// @Router /hits [post]
// Proxy proxies the context to another domain
func Proxy(toURL string) func(c *gin.Context) {
	return func(c *gin.Context) {
		remote, err := url.Parse(toURL)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		//Define the director func
		//This is a good place to log, for example
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("proxyPath")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
