package buildInfo

import (
	"github.com/emicklei/go-restful"
	"net/http"
)

// return json feed of our version struct
func (b *BuildInfo) getVersion(request *restful.Request, response *restful.Response) {
	if err := response.WriteEntity(b); err != nil {
		if err = response.WriteServiceError(http.StatusInternalServerError,
			restful.NewError(http.StatusInternalServerError, err.Error())); err != nil {
			panic(err.Error())
		}
	}
}

// AddRoute adds a restful service route
func (b *BuildInfo) AddRoute(ws *restful.WebService) {
	// json route
	ws.Route(ws.GET("/version").
		To(b.getVersion).
		Doc("Get available version data").
		Notes("Return JSON formatted version data specific "+
			"to this application and it's dependencies.").
		Operation("getVersion").
		Returns(http.StatusAccepted, "Okay", nil).
		Returns(http.StatusInternalServerError, "Error", restful.ServiceError{}).
		Writes(BuildInfo{}).
		Produces(restful.MIME_JSON))
}
