package repository

import (
	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/service/dto"
)

func InsertTemplate(req *dto.CreateTemplateRequest) bool {
	client := clients.GetMysqlClient()

	stmt, err1 := client.Prepare("insert into Templates(ContactType, Template) values(?, ?)")
	if err1 != nil {
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.ContactTypeId(), *req.Template)
	return err2 == nil
}

func GetTemplate(req *dto.TemplateIdRequest) (*dto.TemplateRecord, int) {
	client := clients.GetMysqlClient()

	stmt, err1 := client.Prepare("select * from Templates where Id=?")
	if err1 != nil {
		return nil, 1
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(*req.Id)
	if err2 != nil {
		return nil, 1
	}
	defer rows.Close()

	if rows.Next() {
		record := dto.TemplateRecord{ }
		if err3 := rows.Scan(&record.Id, &record.ContactType, &record.Template); err3 != nil {
			return nil, 2
		}
		return &record, 0
	} else {
		return nil, 2
	}
}

func UpdateTemplate(req *dto.UpdateTemplateRequest) (bool) {
	client := clients.GetMysqlClient()

	stmt, err1 := client.Prepare("update Templates set Template=? where Id=?")
	if err1 != nil {
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.Template, req.Id)
	return err2 == nil
}

func DeleteTemplate(req *dto.TemplateIdRequest) (bool) {
	client := clients.GetMysqlClient()

	stmt, err1 := client.Prepare("delete from Templates where Id=?")
	if err1 != nil {
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.Id)
	return err2 == nil
}
