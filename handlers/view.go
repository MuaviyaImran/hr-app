package handlers

import (
	"bytes"
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed templates/dashboard.html
var dashboard string

//go:embed templates/index.html
var index string

//go:embed templates/employees.html
var employees string

//go:embed templates/employee-details.html
var employeeDetails string

//go:embed templates/employee-salary.html
var employeeSalary string

//go:embed templates/employee-documents.html
var employeeDocuments string

//go:embed templates/employee-technical.html
var employeeTechnical string

//go:embed templates/employee-personal.html
var employeePersonal string

//go:embed templates/transact-salaries.html
var transactSalaries string

func H1(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	temp,_:= template.New("").Parse(dashboard)
	temp.Execute(&body, struct{}{})

	baseTemplate,_ :=	template.New("").Parse(index)
	baseTemplate.Execute(w, struct{
		Body template.HTML
	}{
		Body: template.HTML(body.String()),
	})
    
}
func H2(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	temp,_:= template.New("").Parse(employees)
	temp.Execute(&body, struct{}{})

	baseTemplate,_:=	template.New("").Parse(index)
	baseTemplate.Execute(w, struct{
		Body template.HTML
	}{
		Body: template.HTML(body.String()),
	})
}

func H3(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	temp,_:= template.New("").Parse(employeeDetails)
	temp.Execute(&body, struct{}{})

	baseTemplate,_:=	template.New("").Parse(index)
	baseTemplate.Execute(w, struct{
		Body template.HTML
	}{
		Body: template.HTML(body.String()),
	})
}

func H4(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	temp,_:= template.New("").Parse(employeePersonal)
	temp.Execute(&body, struct{}{})

	baseTemplate,_:=	template.New("").Parse(employeeDetails)
	baseTemplate.Execute(w, struct{
		EmployeeBody template.HTML
	}{
		EmployeeBody: template.HTML(body.String()),
	})
}
func H5(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	temp,_:= template.New("").Parse(employeeTechnical)
	temp.Execute(&body, struct{}{})

	baseTemplate,_:=	template.New("").Parse(employeeDetails)
	baseTemplate.Execute(w, struct{
		EmployeeBody template.HTML
	}{
		EmployeeBody: template.HTML(body.String()),
	})
}

func H6(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	temp,_:= template.New("").Parse(employeeDocuments)
	temp.Execute(&body, struct{}{})

	baseTemplate,_:=	template.New("").Parse(employeeDetails)
	baseTemplate.Execute(w, struct{
		EmployeeBody template.HTML
	}{
		EmployeeBody: template.HTML(body.String()),
	})
}

func H7(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	temp,_:= template.New("").Parse(employeeSalary)
	temp.Execute(&body, struct{}{})

	baseTemplate,_:=	template.New("").Parse(employeeDetails)
	baseTemplate.Execute(w, struct{
		EmployeeBody template.HTML
	}{
		EmployeeBody: template.HTML(body.String()),
	})
}

func H8(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer
	temp,_:= template.New("").Parse(transactSalaries)
	temp.Execute(&body, struct{}{})

	baseTemplate,_:=	template.New("").Parse(index)
	baseTemplate.Execute(w, struct{
		Body template.HTML
	}{
		Body: template.HTML(body.String()),
	})
}