package seeding

import (
	"context"
	"fmt"
	"lexium-utility/config"
	"lexium-utility/datahandler"
	"log"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Define the struct that represents the schema of your MongoDB documents

const totalColumns = 44 // Define the total number of columns expected

func Seed() {
	f, err := excelize.OpenFile("RuleAuthor-OnlineUtilitywithEnhancements/AY 23-24/ITR1/0.4/RuleAuthorSorted_ITR1 V2 R4A12P15M2 26.06.2024.xlsm")
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
		return
	}
	file := excelize.NewFile()
	file.NewSheet("Empty Json Tags")
	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		fmt.Println("Error converting coordinates to cell name:", err)
		return
	}
	file.SetCellValue("Empty Json Tags", cell, "Element Id")

	cell, err = excelize.CoordinatesToCellName(2, 1)
	if err != nil {
		fmt.Println("Error converting coordinates to cell name:", err)
		return
	}
	file.SetCellValue("Empty Json Tags", cell, "Sheet Name")
	cell, err = excelize.CoordinatesToCellName(3, 1)
	if err != nil {
		fmt.Println("Error converting coordinates to cell name:", err)
		return
	}
	file.SetCellValue("Empty Json Tags", cell, "Row No")
	rowJsonTagEmpty := 2
	client := config.GetMongoClient()
	collection := client.Database("your_database_name").Collection("your_collection_name")
	for _, sheetName := range f.GetSheetMap() {

		rows, err := f.GetRows(sheetName)
		if err != nil {
			log.Fatalf("Failed to get rows from sheet %s: %v", sheetName, err)
			return
		}

		ctx := context.TODO()

		// Accumulate documents
		var documents []interface{}
		for rowIndex, _ := range rows {
			if rowIndex == 0 {
				continue
			}
			row := constructCompleteRow(f, sheetName, rowIndex+1, totalColumns)
			document := mapExcelRowToDocument(row)
			if document.JsonTagName == "" {
				cell, err := excelize.CoordinatesToCellName(1, rowJsonTagEmpty)
				if err != nil {
					fmt.Println("Error converting coordinates to cell name:", err)
					return
				}
				file.SetCellValue("Empty Json Tags", cell, document.ElementID)

				cell, err = excelize.CoordinatesToCellName(2, rowJsonTagEmpty)
				if err != nil {
					fmt.Println("Error converting coordinates to cell name:", err)
					return
				}
				file.SetCellValue("Empty Json Tags", cell, sheetName)
				cell, err = excelize.CoordinatesToCellName(3, rowJsonTagEmpty)
				if err != nil {
					fmt.Println("Error converting coordinates to cell name:", err)
					return
				}
				file.SetCellValue("Empty Json Tags", cell, rowIndex+1)
				rowJsonTagEmpty++

			}
			documents = append(documents, document)
		}

		// Insert documents in bulk
		if len(documents) > 0 {
			_, err := collection.InsertMany(ctx, documents)
			if err != nil {
				log.Printf("Failed to insert documents: %v", err)
				return
			}
			fmt.Printf("Inserted %d documents of sheet:%s\n", len(documents), sheetName)
		}
		// Save the XLSM file
		err = file.SaveAs("example.xlsm")
		if err != nil {
			fmt.Println("Error saving file:", err)
			return
		}

	}
}

// Construct a complete row with empty columns included
func constructCompleteRow(f *excelize.File, sheetName string, rowNum int, colCount int) []string {
	row := make([]string, colCount)
	for colIndex := 0; colIndex < colCount; colIndex++ {
		cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowNum)
		value, err := f.GetCellValue(sheetName, cell)
		if err != nil {
			log.Printf("Failed to get cell value for %s: %v", cell, err)
		}
		row[colIndex] = strings.Trim(value, " ")
	}
	return row
}

// Example function to map an Excel row to a Document struct
func mapExcelRowToDocument(row []string) datahandler.Element {
	parent := row[0]
	if row[1] != "" {
		parent += "." + row[1]
	}
	if row[2] != "" {
		parent += "." + row[2]
	}
	if row[3] != "" {
		parent += "." + row[3]
	}
	if row[4] != "" {
		parent += "." + row[4]
	}
	if row[5] != "" {
		parent += "." + row[5]
	}
	groupID, err := strconv.Atoi(row[24])
	if err != nil {
		fmt.Println("error converting groupID")
	}
	seqID, err := strconv.Atoi(row[25])
	if err != nil {
		fmt.Println("error converting seqID")
	}
	document := datahandler.Element{
		Parent:                          parent,
		ElementID:                       row[6],
		FieldName:                       row[7],
		JsonTagName:                     row[8],
		FieldType:                       row[9],
		DataType:                        row[10],
		IsMandatory:                     row[11],
		IsEditable:                      row[12],
		IsAvailable:                     row[13],
		Formula:                         row[14],
		PrefillTag:                      row[15],
		Validation:                      row[16],
		Pattern:                         row[17],
		Enum:                            row[18],
		ErrorMessage:                    row[19],
		NoteComments:                    row[20],
		MinValue:                        row[21],
		MaxValue:                        row[22],
		Depth:                           row[23],
		GroupID:                         groupID,
		SeqID:                           seqID,
		GroupHeading:                    row[26],
		AffectedIDs:                     row[27],
		PDFFieldType:                    row[28],
		ErrorCode:                       row[29],
		ErrorCategory:                   row[30],
		ErrorCategoryDescription:        row[31],
		ErrorCategoryType:               row[32],
		Suggestion:                      row[33],
		RuleNumber:                      row[34],
		ChangeValidation:                row[35],
		SaveValidation:                  row[36],
		Classes:                         row[37],
		WarningMessages:                 row[38],
		MandatoryMinMaxPatternEnumError: row[39],
		AffectedIDsFormula:              row[40],
		FormulaType:                     row[41],
		CSV:                             row[42],
		EnumCondition:                   row[43],
	}
	return document
}
