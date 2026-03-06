package client1

import "github.com/codercollo/go-mistakes/ch02/mistake05/store2"

//FIX (consumer side): The client defines only the interface it actually needs
//It's minimal, unexported and perfectly sized for this package's use case
//The stro package has no idea this interface eists - no coupling back

// ✅ Interface lives here in the consumer package.
// Only the one method this package actually uses — Interface Segregation Principle.
type customersGetter interface {
	GetAllCustomers() ([]store2.Customer, error) // minimal required behavior
}

// ReportService builds customer reports
type ReportService struct {
	getter customersGetter // dependency providing customer list
}

// Constructor for ReportService
func NewReportService(getter customersGetter) ReportService {
	return ReportService{getter: getter} // inject dependency
}

// BuildReport retrieves customers for reporting
func (rs ReportService) BuildReport() ([]store2.Customer, error) {
	// This package only needs to list customers — nothing else.
	// The interface enforces that boundary.
	return rs.getter.GetAllCustomers() // fetch customers
}
