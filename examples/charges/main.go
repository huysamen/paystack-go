
package main

import (
"context"
"fmt"
"log"

"github.com/huysamen/paystack-go"
"github.com/huysamen/paystack-go/api/charges"
)

func main() {
	// Create client
	client := paystack.DefaultClient("sk_test_your_secret_key_here")
	ctx := context.Background()

	fmt.Println("=== Charges API Example ===\n")

	// 1. Create a charge with bank details using builder pattern
	fmt.Println("1. Creating charge with bank details...")
	createBuilder := charges.NewCreateChargeRequest("customer@example.com", "50000"). // ₦500.00 in kobo
		Bank(&charges.BankDetails{
			Code:          "057", // Zenith Bank
			AccountNumber: "0000000000",
		}).
		Reference("charge-" + generateReference()).
		Metadata(map[string]any{
"payment_type":  "bank_charge",
"customer_note": "Direct bank charge example",
})

	charge, err := client.Charges.Create(ctx, createBuilder)
	if err != nil {
		log.Printf("Failed to create charge: %v", err)
		return
	}

	fmt.Printf("✅ Charge created: %s\n", charge.Data.Reference)
	fmt.Printf("   Status: %s\n", charge.Data.Status)
	fmt.Printf("   Amount: ₦%.2f\n", float64(charge.Data.Amount)/100)
	fmt.Printf("   Channel: %s\n", charge.Data.Channel)

	// Handle different charge statuses
	switch charge.Data.Status {
	case "pending":
		fmt.Println("   ⏳ Charge is pending - might need additional verification")

		// Check pending charge status after a delay (in real scenario, wait 10+ seconds)
		fmt.Println("\n   Checking pending charge status...")
		pendingCharge, err := client.Charges.CheckPending(ctx, charge.Data.Reference)
		if err != nil {
			log.Printf("   Failed to check pending charge: %v", err)
		} else {
			fmt.Printf("   Updated status: %s\n", pendingCharge.Data.Status)
		}

	case "send_pin":
		fmt.Println("   🔐 PIN required - submitting PIN...")

		// Example PIN submission using builder (in real scenario, get PIN from user)
		pinBuilder := charges.NewSubmitPINRequest("1234", charge.Data.Reference)

		pinCharge, err := client.Charges.SubmitPIN(ctx, pinBuilder)
		if err != nil {
			log.Printf("   Failed to submit PIN: %v", err)
		} else {
			fmt.Printf("   ✅ PIN submitted, new status: %s\n", pinCharge.Data.Status)
		}

	case "send_otp":
		fmt.Println("   📱 OTP required - submitting OTP...")

		// Example OTP submission using builder (in real scenario, get OTP from user)
		otpBuilder := charges.NewSubmitOTPRequest("123456", charge.Data.Reference)

		otpCharge, err := client.Charges.SubmitOTP(ctx, otpBuilder)
		if err != nil {
			log.Printf("   Failed to submit OTP: %v", err)
		} else {
			fmt.Printf("   ✅ OTP submitted, new status: %s\n", otpCharge.Data.Status)
		}

	case "send_phone":
		fmt.Println("   📞 Phone number required - submitting phone...")

		phoneBuilder := charges.NewSubmitPhoneRequest("08012345678", charge.Data.Reference)

		phoneCharge, err := client.Charges.SubmitPhone(ctx, phoneBuilder)
		if err != nil {
			log.Printf("   Failed to submit phone: %v", err)
		} else {
			fmt.Printf("   ✅ Phone submitted, new status: %s\n", phoneCharge.Data.Status)
		}

	case "send_birthday":
		fmt.Println("   🎂 Birthday required - submitting birthday...")

		birthdayBuilder := charges.NewSubmitBirthdayRequest("1990-12-25", charge.Data.Reference)

		birthdayCharge, err := client.Charges.SubmitBirthday(ctx, birthdayBuilder)
		if err != nil {
			log.Printf("   Failed to submit birthday: %v", err)
		} else {
			fmt.Printf("   ✅ Birthday submitted, new status: %s\n", birthdayCharge.Data.Status)
		}

	case "send_address":
		fmt.Println("   🏠 Address required - submitting address...")

		addressBuilder := charges.NewSubmitAddressRequest(
"123 Main Street, Apartment 4B",
charge.Data.Reference,
"Lagos",
"Lagos",
"100001",
)

		addressCharge, err := client.Charges.SubmitAddress(ctx, addressBuilder)
		if err != nil {
			log.Printf("   Failed to submit address: %v", err)
		} else {
			fmt.Printf("   ✅ Address submitted, new status: %s\n", addressCharge.Data.Status)
		}

	case "success":
		fmt.Println("   ✅ Charge successful!")
		if charge.Data.Authorization != nil {
			fmt.Printf("   Authorization: %s\n", charge.Data.Authorization.AuthorizationCode)
		}

	case "failed":
		fmt.Println("   ❌ Charge failed")
		fmt.Printf("   Message: %s\n", charge.Data.Message)

	default:
		fmt.Printf("   ❓ Unknown status: %s\n", charge.Data.Status)
	}

	fmt.Println("\n=== Charges API Example Complete ===")
}

// Helper functions
func generateReference() string {
	// Simple reference generator for example
	return fmt.Sprintf("%d", 1000000000+int(1000000000*0.5))
}
