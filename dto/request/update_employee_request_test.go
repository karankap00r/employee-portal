package request

import "testing"

func TestUpdateEmployeeByEmployeeIDRequest_Validate(t *testing.T) {
	type testCase struct {
		name    string
		request UpdateEmployeeByEmployeeIDRequest
		wantErr bool
	}

	tests := []testCase{
		{name: "Valid request", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "E123456", Name: "John Doe", Position: "Engineer", Email: "john.doe@example.com", Salary: 50000}, wantErr: false},
		{name: "Empty employee_id", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "", Name: "John Doe", Position: "Engineer", Email: "john.doe@example.com", Salary: 50000}, wantErr: true},
		{name: "Empty name", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "E123456", Name: "", Position: "Engineer", Email: "john.doe@example.com", Salary: 50000}, wantErr: true},
		{name: "Long name", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "E123456", Name: "This is a very very very very very very very long name", Position: "Engineer", Email: "john.doe@example.com", Salary: 50000}, wantErr: true},
		{name: "Empty email", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "E123456", Name: "John Doe", Position: "Engineer", Email: "", Salary: 50000}, wantErr: true},
		{name: "Invalid email", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "E123456", Name: "John Doe", Position: "Engineer", Email: "john.doe@", Salary: 50000}, wantErr: true},
		{name: "Long email", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "E123456", Name: "John Doe", Position: "Engineer", Email: "thisisaveryveryveryveryveryveryveryveryverylongemailthisisaveryveryveryveryveryveryveryveryverylongemailthisisaveryveryveryveryveryveryveryveryverylongemailthisisaveryveryveryveryveryveryveryveryverylongemail@example.com", Salary: 50000}, wantErr: true},
		{name: "Empty position", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "E123456", Name: "John Doe", Position: "", Email: "john.doe@example.com", Salary: 50000}, wantErr: true},
		{name: "Long position", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "E123456", Name: "John Doe", Position: "This is a very very very very very very long position", Email: "john.doe@example.com", Salary: 50000}, wantErr: true},
		{name: "Negative salary", request: UpdateEmployeeByEmployeeIDRequest{EmployeeID: "E123456", Name: "John Doe", Position: "Engineer", Email: "john.doe@example.com", Salary: -1}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.request.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
