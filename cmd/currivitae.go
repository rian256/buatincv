package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/rian256/buatincv/helper"
	"github.com/spf13/cobra"
)

var pdf *gofpdf.Fpdf

const (
	confirmPrompt = "\nPlease type either y/n\n"
	fontName      = "Arial"
)

func init() {
	rootCmd.AddCommand(createResumeCmd)
}

var createResumeCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new Resume",
	Run:   wrapper,
}

// Wrapping all functions required to render and create a PDF
func wrapper(cmd *cobra.Command, args []string) {
	createPDF()
	generateName()
	generateRole()
	generateAddress()
	renderSummaryTitle()
	summarySection()
	renderExperienceTitle()
	profExpSection()
	profExpInfoSection()
	jobResponsibilitiySection()
	renderEducationalTitle()
	educationExperienceSection()
	educationalDegreeSection()
	renderSkillsTitle()
	skillsSection()
	pdfOutput()
}

func createPDF() {
	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
}

func renderSummaryTitle() {
	pdf.Ln(5)
	pdf.SetFont(fontName, "B", 13)
	pdf.WriteAligned(0, 10, "SUMMARY:", "L")
}

func renderExperienceTitle() {
	pdf.Ln(10)
	pdf.SetFont(fontName, "B", 13)
	pdf.WriteAligned(0, 10, "EXPERIENCE:", "L")
}

func renderEducationalTitle() {
	pdf.Ln(10)
	pdf.SetFont(fontName, "B", 13)
	pdf.WriteAligned(0, 10, "EDUCATION:", "L")
}

func renderSkillsTitle() {
	pdf.Ln(10)
	pdf.SetFont(fontName, "B", 13)
	pdf.WriteAligned(0, 10, "SKILLS:", "L")
}

func pdfOutput() {
	err := pdf.OutputFileAndClose("Resume.pdf")
	helper.ErrorHelper(err)

	if err != nil {
		helper.ErrorHelper(err)
	} else {
		log.Printf("PDF successfully created: Resume.pdf")
	}

}

func generateName() {
	name, err := helper.CaptureInput("\nEnter your name\nExample: Jane Doe\nYour Name:\n")
	helper.ErrorHelper(err)

	IsNameInputed := true
	pdf.SetFont(fontName, "B", 20)
	pdf.SetTitle(name, IsNameInputed)
	pdf.WriteAligned(0, 10, name, "C")
	pdf.Ln(3)
}

func generateRole() {
	role, err := helper.CaptureInput("\nEnter your role\nExample: Backend Developer\n")
	helper.ErrorHelper(err)
	pdf.Ln(3)
	pdf.SetFont(fontName, "B", 12)
	pdf.WriteAligned(0, 10, role, "C")
}

func generateAddress() {
	address, err := helper.CaptureInput("\nEnter your current address\nExample: Banten, Indonesia\n")
	helper.ErrorHelper(err)
	pdf.Ln(10)
	pdf.SetFont(fontName, "", 10)
	pdf.WriteAligned(0, 2, address, "C")
}

func summarySection() {
	summary, err := helper.CaptureInput("\nEnter your summary\nI'am Backend developer with n years experience\n")
	helper.ErrorHelper(err)
	pdf.Ln(7)
	pdf.SetFont(fontName, "", 10)
	pdf.WriteAligned(0, 6, summary, "L")
}
func profExpSection() {
	profExp, err := helper.CaptureInput("\nEnter the name and location of the company where you previously worked or currently work.\nExample: PT XYZ - Surabaya, Indonesia\n")
	helper.ErrorHelper(err)

	pdf.Ln(10)
	pdf.SetFont(fontName, "B", 10)
	pdf.WriteAligned(0, 2, profExp, "L")
	pdf.Ln(5)
}

func profExpInfoSection() {

	ProfExpInformation, err := helper.CaptureInput("\nEnter your experience and indicate the start and end dates of your employment.:\nExample: Backend Developer (2017-2023)\n")
	helper.ErrorHelper(err)

	pdf.SetFont(fontName, "", 10)
	pdf.WriteAligned(0, 2, ProfExpInformation, "L")
}

func jobResponsibilitiySection() {
	for {
		jobResponbility, err := helper.CaptureInput("\nDescribe your job responsibilities\n")
		helper.ErrorHelper(err)
		pdf.Ln(5)
		pdf.SetFont(fontName, "", 10)
		jobResponbilitySummary := fmt.Sprintf("%s\n", jobResponbility)
		pdf.WriteAligned(0, 5, jobResponbilitySummary, "L")

		input, err := helper.CaptureInput("\nWould you like to add more information about the job responsibilities? y/n\n")
		helper.ErrorHelper(err)
		input = strings.TrimSpace(input)
		if input == "y" {
			continue
		} else if input == "n" {
			break
		} else {
			println(confirmPrompt)
		}
	}
}

func educationExperienceSection() {
	formalEducation, err := helper.CaptureInput("\nEnter your formal education experience\nExample: XYZ University - (2010-2014)\n")
	helper.ErrorHelper(err)

	pdf.Ln(10)
	pdf.SetFont(fontName, "B", 10)
	pdf.WriteAligned(0, 2, formalEducation, "L")
}

func educationalDegreeSection() {
	educationalDegree, err := helper.CaptureInput("\nEnter your Educational Degree\nExample: Bachelor of Computer Science\n")
	helper.ErrorHelper(err)
	pdf.Ln(5)
	pdf.SetFont(fontName, "", 9)
	pdf.WriteAligned(0, 2, educationalDegree, "L")
}

func skillsSection() {
	for {
		skills, err := helper.CaptureInput("\nEnter your skill:\nExample: - In-depth understanding of microservices concepts\n")
		helper.ErrorHelper(err)
		pdf.Ln(5)

		input, err := helper.CaptureInput("\nDo you want to add more skills? y/n\n")
		helper.ErrorHelper(err)
		input = strings.TrimSpace(input)

		helper.ErrorHelper(err)
		if input == "y" {
			pdf.SetFont(fontName, "", 10)
			pdf.WriteAligned(0, 10, skills, "L")
			continue
		} else if input == "n" {
			pdf.SetFont(fontName, "", 10)
			pdf.WriteAligned(0, 10, skills, "L")
			break
		} else {
			println(confirmPrompt)
		}
	}
}
