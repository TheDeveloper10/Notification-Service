package repositories

import (
	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/service/dtos"
)

func InsertTemplate(req *dtos.CreateTemplateRequest) bool {
	client := clients.GetMysqlClient()

	stmt, err1 := client.Prepare("insert into Templates(ContactType, Template) values(?, ?)")
	if err1 != nil {
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.ContactTypeId(), *req.Template)
	return err2 == nil
}

func DeleteTemplate(req *dtos.TemplateIdRequest) (bool) {
	client := clients.GetMysqlClient()

	stmt, err1 := client.Prepare("delete from Templates where Id=?")
	if err1 != nil {
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.Id)
	return err2 == nil
}
