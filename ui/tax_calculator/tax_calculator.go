package taxcalculator

import (
	"fmt"
	"strconv"
	nationalinsurance "tax_calculator/engine/internal/valueobjects/hmrc_valueobjects/national_insurance"
	taxliability "tax_calculator/engine/internal/valueobjects/hmrc_valueobjects/tax_liability"
	"tax_calculator/engine/lib/app"

	"github.com/rivo/tview"
)

func GetLayout(app *app.Application) tview.Primitive {
	totalProfit := float32(0)

	taxDueDisplay := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Income tax charged: ")

	nationalInsuranceDueDisplay := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Class 4 National Insurance contributions: ")

	totalDueDisplay := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Income tax and Class 4 National Insurance contributions due:")

	form := tview.NewForm()

	form.AddInputField("Yearly Total Profit", "", 20, nil, func(text string) {
		if s, err := strconv.ParseFloat(text, 32); err == nil {
			totalProfit = float32(s)
		}
	})

	form.AddButton("Calculate", func() {
		taxDue := taxliability.CalculateTaxLiability(totalProfit)
		nationalInsuranceDue := nationalinsurance.CalculateNationalInsurance(totalProfit)

		taxDueDisplay.SetText(fmt.Sprintf("£%.2f", taxDue))
		nationalInsuranceDueDisplay.SetText(fmt.Sprintf("£%.2f", nationalInsuranceDue))
		totalDueDisplay.SetText(fmt.Sprintf("£%.2f", taxDue+nationalInsuranceDue))
	}).SetFieldBackgroundColor(000).SetButtonBackgroundColor(000)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(taxDueDisplay, 3, 0, false).
		AddItem(nationalInsuranceDueDisplay, 3, 0, false).
		AddItem(totalDueDisplay, 3, 0, false)

	return layout
}
