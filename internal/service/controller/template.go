package controller

import (
	"net/http"

	"notification-service.com/packages/internal/dto"
	"notification-service.com/packages/internal/util"
	"notification-service.com/packages/internal/repository"
)

type template struct {
	util.Controller
}

func NewTemplateRepository() *template {
	return &template{}
}

func (t *template) Handle(res http.ResponseWriter, req *http.Request) {
	brw := util.ConvertResponseWriter(&res)

	switch(req.Method) {
		case "POST": {
			t.create(brw, req)
		}
		case "GET": {
			t.get(brw, req)
		}
		case "PATCH": {
			t.update(brw, req)
		}
		case "DELETE": {
			t.delete(brw, req)
		}
	}
}

func (t *template) create(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.CreateTemplateRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	result := repository.NewTemplateRepository().Insert(&reqObj)
	if result {
		// Maybe return metadata such as id
		res.Status(http.StatusOK).Text("Created successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to add template to the database. Try again!")
	}
}

func (t *template) get(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.TemplateIdRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	record, statusCode := repository.NewTemplateRepository().Get(&reqObj)
	if statusCode == 1 {
		res.Status(http.StatusBadRequest).Text("Failed to get the requested template. Try again!")
		return
	} else if statusCode == 2 {
		res.Status(http.StatusNotFound).Text("Couldn't find the template you were looking for!")
		return
	} else {
		res.Status(http.StatusOK).Json(record.ToReadable())
	}
}

func (t *template) update(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.UpdateTemplateRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	status := repository.NewTemplateRepository().Update(&reqObj)
	if status {
		res.Status(http.StatusOK).Text("Updated successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to update it. Try again!")
	}
}

func (t *template) delete(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.TemplateIdRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	status := repository.NewTemplateRepository().Delete(&reqObj)
	if status {
		res.Status(http.StatusOK).Text("Deleted successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to delete it. Try again!")
	}
}