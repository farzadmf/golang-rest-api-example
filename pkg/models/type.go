package models

const (
	// Contractor represents a contractor.
	Contractor = 1

	// Employee represents an employee.
	Employee = 2
)

// GetMemberTypes returns the list of valid member types as a mep.
// Example use case: display valid choices in a dropdown in the frontend.
func GetMemberTypes() map[int]string {
	return map[int]string{
		1: "Contractor",
		2: "Employee",
	}
}
