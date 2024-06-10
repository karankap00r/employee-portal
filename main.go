package main

import (
	_ "github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/api"
	_ "github.com/swaggo/http-swagger"

	"github.com/karankap00r/employee_portal/config"
	"github.com/karankap00r/employee_portal/server"
)

func main() {

	// Mock data
	api.Employees = append(api.Employees, api.Employee{ID: "1", Name: "John Doe", Age: 30, Dept: "Engineering"})
	api.Employees = append(api.Employees, api.Employee{ID: "2", Name: "Jane Doe", Age: 25, Dept: "Marketing"})

	cfg := config.LoadConfig()
	server.Start(cfg)
}
