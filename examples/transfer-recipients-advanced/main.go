package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	transfer_recipients "github.com/huysamen/paystack-go/api/transfer-recipients"
)

// Employee represents an employee in our system
type Employee struct {
	ID           string
	Name         string
	Email        string
	Department   string
	BankAccount  string
	BankCode     string
	Salary       int
	PayrollGroup string
}

// RecipientManager manages transfer recipients for payroll processing
type RecipientManager struct {
	client *paystack.Client
}

func NewRecipientManager(client *paystack.Client) *RecipientManager {
	return &RecipientManager{client: client}
}

// CreateEmployeeRecipient creates a transfer recipient for an employee
func (rm *RecipientManager) CreateEmployeeRecipient(emp Employee) (*transfer_recipients.TransferRecipient, error) {
	req := &transfer_recipients.TransferRecipientCreateRequest{
		Type:          transfer_recipients.RecipientTypeNuban,
		Name:          emp.Name,
		AccountNumber: emp.BankAccount,
		BankCode:      emp.BankCode,
		Currency:      &[]string{"NGN"}[0],
		Description:   &[]string{fmt.Sprintf("%s - %s", emp.Department, emp.PayrollGroup)}[0],
		Metadata: map[string]any{
			"employee_id":   emp.ID,
			"department":    emp.Department,
			"payroll_group": emp.PayrollGroup,
			"salary":        emp.Salary,
			"email":         emp.Email,
			"created_by":    "payroll_system",
			"created_at":    time.Now().Format(time.RFC3339),
		},
	}

	resp, err := rm.client.TransferRecipients.Create(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to create recipient for employee %s: %w", emp.ID, err)
	}

	return &resp.Data, nil
}

// BulkCreateEmployeeRecipients creates recipients for multiple employees
func (rm *RecipientManager) BulkCreateEmployeeRecipients(employees []Employee) (*transfer_recipients.BulkCreateResult, error) {
	if len(employees) == 0 {
		return nil, fmt.Errorf("no employees provided")
	}

	batch := make([]transfer_recipients.BulkRecipientItem, len(employees))
	for i, emp := range employees {
		batch[i] = transfer_recipients.BulkRecipientItem{
			Type:          transfer_recipients.RecipientTypeNuban,
			Name:          emp.Name,
			AccountNumber: emp.BankAccount,
			BankCode:      emp.BankCode,
			Currency:      &[]string{"NGN"}[0],
			Description:   &[]string{fmt.Sprintf("%s - %s", emp.Department, emp.PayrollGroup)}[0],
			Metadata: map[string]any{
				"employee_id":   emp.ID,
				"department":    emp.Department,
				"payroll_group": emp.PayrollGroup,
				"salary":        emp.Salary,
				"email":         emp.Email,
			},
		}
	}

	req := &transfer_recipients.BulkCreateTransferRecipientRequest{
		Batch: batch,
	}

	resp, err := rm.client.TransferRecipients.BulkCreate(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk create recipients: %w", err)
	}

	return &resp.Data, nil
}

// FindRecipientsByDepartment finds all recipients from a specific department
func (rm *RecipientManager) FindRecipientsByDepartment(department string) ([]transfer_recipients.TransferRecipient, error) {
	// List all recipients (in a real system, you'd want to implement server-side filtering)
	req := &transfer_recipients.TransferRecipientListRequest{
		PerPage: &[]int{100}[0], // Adjust as needed
	}

	resp, err := rm.client.TransferRecipients.List(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to list recipients: %w", err)
	}

	var filtered []transfer_recipients.TransferRecipient
	for _, recipient := range resp.Data {
		if recipient.Metadata != nil {
			if dept, ok := recipient.Metadata["department"].(string); ok && dept == department {
				filtered = append(filtered, recipient)
			}
		}
	}

	return filtered, nil
}

// UpdateEmployeeDetails updates recipient information for an employee
func (rm *RecipientManager) UpdateEmployeeDetails(recipientCode, newName, newEmail string) error {
	req := &transfer_recipients.TransferRecipientUpdateRequest{
		Name:  newName,
		Email: &newEmail,
	}

	_, err := rm.client.TransferRecipients.Update(context.Background(), recipientCode, req)
	if err != nil {
		return fmt.Errorf("failed to update recipient %s: %w", recipientCode, err)
	}

	return nil
}

// GeneratePayrollReport generates a report of all recipients by department
func (rm *RecipientManager) GeneratePayrollReport() error {
	req := &transfer_recipients.TransferRecipientListRequest{
		PerPage: &[]int{100}[0],
	}

	resp, err := rm.client.TransferRecipients.List(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to list recipients for report: %w", err)
	}

	// Group by department
	departments := make(map[string][]transfer_recipients.TransferRecipient)
	totalSalary := make(map[string]int)

	for _, recipient := range resp.Data {
		if recipient.Metadata != nil {
			dept, hasDept := recipient.Metadata["department"].(string)
			salary, hasSalary := recipient.Metadata["salary"].(float64)

			if hasDept {
				departments[dept] = append(departments[dept], recipient)
				if hasSalary {
					totalSalary[dept] += int(salary)
				}
			}
		}
	}

	// Print report
	fmt.Println("\n=== PAYROLL REPORT ===")
	fmt.Printf("Generated: %s\n", time.Now().Format("January 2, 2006 15:04:05"))
	fmt.Printf("Total Recipients: %d\n\n", len(resp.Data))

	for dept, recipients := range departments {
		fmt.Printf("📋 %s Department\n", dept)
		fmt.Printf("   Employees: %d\n", len(recipients))
		fmt.Printf("   Total Salary: ₦%s\n", formatMoney(totalSalary[dept]))
		fmt.Printf("   Active Recipients: %d\n", countActive(recipients))

		if len(recipients) <= 5 {
			for _, recipient := range recipients {
				status := "✅"
				if !recipient.Active {
					status = "❌"
				}
				fmt.Printf("   %s %s (%s)\n", status, recipient.Name, recipient.RecipientCode)
			}
		} else {
			for i := 0; i < 3; i++ {
				recipient := recipients[i]
				status := "✅"
				if !recipient.Active {
					status = "❌"
				}
				fmt.Printf("   %s %s (%s)\n", status, recipient.Name, recipient.RecipientCode)
			}
			fmt.Printf("   ... and %d more\n", len(recipients)-3)
		}
		fmt.Println()
	}

	return nil
}

// Helper functions
func formatMoney(amount int) string {
	return fmt.Sprintf("%d", amount)
}

func countActive(recipients []transfer_recipients.TransferRecipient) int {
	count := 0
	for _, r := range recipients {
		if r.Active {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("=== Advanced Transfer Recipients Management ===")
	fmt.Println()

	// Create client
	client := paystack.DefaultClient("sk_test_your_secret_key_here")
	manager := NewRecipientManager(client)

	// Sample employees data
	employees := []Employee{
		{
			ID:           "EMP001",
			Name:         "Alice Johnson",
			Email:        "alice.johnson@company.com",
			Department:   "Engineering",
			BankAccount:  "1234567890",
			BankCode:     "058",  // GTBank
			Salary:       350000, // ₦350,000
			PayrollGroup: "Senior",
		},
		{
			ID:           "EMP002",
			Name:         "Bob Smith",
			Email:        "bob.smith@company.com",
			Department:   "Marketing",
			BankAccount:  "2345678901",
			BankCode:     "044",  // Access Bank
			Salary:       280000, // ₦280,000
			PayrollGroup: "Mid-level",
		},
		{
			ID:           "EMP003",
			Name:         "Carol Davis",
			Email:        "carol.davis@company.com",
			Department:   "Engineering",
			BankAccount:  "3456789012",
			BankCode:     "033",  // UBA
			Salary:       420000, // ₦420,000
			PayrollGroup: "Senior",
		},
		{
			ID:           "EMP004",
			Name:         "David Wilson",
			Email:        "david.wilson@company.com",
			Department:   "Finance",
			BankAccount:  "4567890123",
			BankCode:     "058",  // GTBank
			Salary:       320000, // ₦320,000
			PayrollGroup: "Mid-level",
		},
		{
			ID:           "EMP005",
			Name:         "Eva Brown",
			Email:        "eva.brown@company.com",
			Department:   "HR",
			BankAccount:  "5678901234",
			BankCode:     "044",  // Access Bank
			Salary:       300000, // ₦300,000
			PayrollGroup: "Mid-level",
		},
	}

	// Scenario 1: Bulk create recipients for new employees
	fmt.Println("🚀 Scenario 1: Onboarding new employees...")
	bulkResult, err := manager.BulkCreateEmployeeRecipients(employees)
	if err != nil {
		log.Printf("Error in bulk create: %v", err)
		return
	}

	fmt.Printf("✅ Onboarding complete:\n")
	fmt.Printf("   Successfully created: %d recipients\n", len(bulkResult.Success))
	fmt.Printf("   Failed: %d recipients\n", len(bulkResult.Errors))

	// Show created recipients
	for _, recipient := range bulkResult.Success {
		emp := findEmployeeByName(employees, recipient.Name)
		fmt.Printf("   + %s (%s) - %s Department\n",
			recipient.Name,
			recipient.RecipientCode,
			emp.Department)
	}

	// Show any errors
	for _, err := range bulkResult.Errors {
		fmt.Printf("   - Failed: %s (%s)\n", err.Name, err.Message)
	}
	fmt.Println()

	// Store some recipient codes for further operations
	var engineeringRecipient string
	if len(bulkResult.Success) >= 1 {
		engineeringRecipient = bulkResult.Success[0].RecipientCode
	}

	// Scenario 2: Find recipients by department
	fmt.Println("🔍 Scenario 2: Department analysis...")
	engineeringRecipients, err := manager.FindRecipientsByDepartment("Engineering")
	if err != nil {
		log.Printf("Error finding Engineering recipients: %v", err)
		return
	}

	fmt.Printf("Engineering Department: %d recipients\n", len(engineeringRecipients))
	for _, recipient := range engineeringRecipients {
		salary := "Unknown"
		if recipient.Metadata != nil {
			if s, ok := recipient.Metadata["salary"].(float64); ok {
				salary = fmt.Sprintf("₦%s", formatMoney(int(s)))
			}
		}
		fmt.Printf("   - %s (%s) - Salary: %s\n",
			recipient.Name,
			recipient.RecipientCode,
			salary)
	}
	fmt.Println()

	// Scenario 3: Update employee details (promotion/name change)
	if engineeringRecipient != "" {
		fmt.Println("📝 Scenario 3: Employee promotion...")
		err = manager.UpdateEmployeeDetails(
			engineeringRecipient,
			"Alice Johnson-Smith", // Name change after marriage
			"alice.johnson-smith@company.com",
		)
		if err != nil {
			log.Printf("Error updating employee: %v", err)
		} else {
			fmt.Printf("✅ Updated employee details for recipient: %s\n", engineeringRecipient)

			// Fetch updated details
			updated, err := client.TransferRecipients.Fetch(context.Background(), engineeringRecipient)
			if err != nil {
				log.Printf("Error fetching updated recipient: %v", err)
			} else {
				fmt.Printf("   New name: %s\n", updated.Data.Name)
				if updated.Data.Email != nil {
					fmt.Printf("   New email: %s\n", *updated.Data.Email)
				}
			}
		}
		fmt.Println()
	}

	// Scenario 4: Generate comprehensive payroll report
	fmt.Println("📊 Scenario 4: Generating payroll report...")
	err = manager.GeneratePayrollReport()
	if err != nil {
		log.Printf("Error generating report: %v", err)
		return
	}

	// Scenario 5: Recipient lifecycle management
	fmt.Println("🔄 Scenario 5: Recipient lifecycle management...")

	// Create a temporary recipient for testing deletion
	tempEmployee := Employee{
		ID:           "TEMP001",
		Name:         "Temporary Employee",
		Email:        "temp@company.com",
		Department:   "Testing",
		BankAccount:  "9999999999",
		BankCode:     "058",
		Salary:       200000,
		PayrollGroup: "Contract",
	}

	tempRecipient, err := manager.CreateEmployeeRecipient(tempEmployee)
	if err != nil {
		log.Printf("Error creating temp recipient: %v", err)
	} else {
		fmt.Printf("✅ Created temporary recipient: %s\n", tempRecipient.RecipientCode)

		// Delete (deactivate) the temporary recipient
		deleteResp, err := client.TransferRecipients.Delete(context.Background(), tempRecipient.RecipientCode)
		if err != nil {
			log.Printf("Error deleting temp recipient: %v", err)
		} else {
			fmt.Printf("✅ %s\n", deleteResp.Message)

			// Verify deactivation
			verifyResp, err := client.TransferRecipients.Fetch(context.Background(), tempRecipient.RecipientCode)
			if err != nil {
				log.Printf("Error verifying deletion: %v", err)
			} else {
				fmt.Printf("   Status after deletion: Active=%t\n", verifyResp.Data.Active)
			}
		}
	}
	fmt.Println()

	// Scenario 6: Recipient validation and error handling
	fmt.Println("⚠️  Scenario 6: Error handling demonstration...")

	// Try to create recipient with invalid data
	invalidReq := &transfer_recipients.TransferRecipientCreateRequest{
		Type:          transfer_recipients.RecipientTypeNuban,
		Name:          "", // Invalid: empty name
		AccountNumber: "1234567890",
		BankCode:      "058",
	}

	_, err = client.TransferRecipients.Create(context.Background(), invalidReq)
	if err != nil {
		fmt.Printf("✅ Validation error caught: %v\n", err)
	}

	// Try to fetch non-existent recipient
	_, err = client.TransferRecipients.Fetch(context.Background(), "RCP_nonexistent")
	if err != nil {
		fmt.Printf("✅ Not found error handled: %v\n", err)
	}

	fmt.Println()
	fmt.Println("=== Advanced Transfer Recipients Management Complete ===")
}

// Helper function to find employee by name
func findEmployeeByName(employees []Employee, name string) Employee {
	for _, emp := range employees {
		if emp.Name == name {
			return emp
		}
	}
	return Employee{Department: "Unknown"}
}
