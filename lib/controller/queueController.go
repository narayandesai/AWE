package controller

import (
	"github.com/MG-RAST/AWE/lib/core"
	"github.com/MG-RAST/golib/goweb"
	"net/http"
)

type QueueController struct{}

// POST: /queue
func (cr *QueueController) Create(cx *goweb.Context) {
	LogRequest(cx.Request)
	cx.RespondWithError(http.StatusNotImplemented)
}

// GET: /queue/{id}
func (cr *QueueController) Read(id string, cx *goweb.Context) {
	LogRequest(cx.Request)
	cx.RespondWithError(http.StatusNotImplemented)
}

// GET: /queue
// get status from queue manager
func (cr *QueueController) ReadMany(cx *goweb.Context) {
	LogRequest(cx.Request)

	// Gather query params
	//	query := &Query{list: cx.Request.URL.Query()}

	msg, err := json.Marshal(core.QMgr.ShowStatus())
	if err == nil {
		cx.RespondWithData(string(msg))
	} else {
		cx.RespondWithData("Err")
	}
}

// PUT: /queue/{id} -> status update
func (cr *QueueController) Update(id string, cx *goweb.Context) {
	LogRequest(cx.Request)
	cx.RespondWithError(http.StatusNotImplemented)
}

// PUT: /queue
func (cr *QueueController) UpdateMany(cx *goweb.Context) {
	LogRequest(cx.Request)
	cx.RespondWithError(http.StatusNotImplemented)
}

// DELETE: /queue/{id}
func (cr *QueueController) Delete(id string, cx *goweb.Context) {
	LogRequest(cx.Request)
	cx.RespondWithError(http.StatusNotImplemented)
}

// DELETE: /queue
func (cr *QueueController) DeleteMany(cx *goweb.Context) {
	LogRequest(cx.Request)
	cx.RespondWithError(http.StatusNotImplemented)
}
