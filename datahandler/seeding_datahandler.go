package datahandler

type Element struct {
	Parent                          string `bson:"Parent"`
	ElementID                       string `bson:"elementID"`
	FieldName                       string `bson:"Field_Name_Header"`
	JsonTagName                     string `bson:"Json_Tag_Name"`
	FieldType                       string `bson:"Field_Type"`
	DataType                        string `bson:"Data_Type"`
	IsMandatory                     string `bson:"IsMandatory"`
	IsEditable                      string `bson:"IsEditable"`
	IsAvailable                     string `bson:"IsAvailable"`
	Formula                         string `bson:"Formula"`
	PrefillTag                      string `bson:"Prefill_Tag"`
	Validation                      string `bson:"Validation"`
	Pattern                         string `bson:"Pattern"`
	Enum                            string `bson:"Enum"`
	ErrorMessage                    string `bson:"ErrorMessage"`
	NoteComments                    string `bson:"Note_Comments"`
	MinValue                        string `bson:"Min_Value"`
	MaxValue                        string `bson:"Max_Value"`
	Depth                           string `bson:"Depth"`
	GroupID                         int    `bson:"GroupID"`
	SeqID                           int    `bson:"SeqID"`
	GroupHeading                    string `bson:"GroupHeading"`
	AffectedIDs                     string `bson:"AffectedIDs"`
	PDFFieldType                    string `bson:"PDFFieldType"`
	ErrorCode                       string `bson:"ErrorCode"`
	ErrorCategory                   string `bson:"Error_Category"`
	ErrorCategoryDescription        string `bson:"Error_Category_Description"`
	ErrorCategoryType               string `bson:"Error_Category_Type"`
	Suggestion                      string `bson:"Suggestion"`
	RuleNumber                      string `bson:"Rule_Number"`
	ChangeValidation                string `bson:"Change_Validation"`
	SaveValidation                  string `bson:"Save_Validation"`
	Classes                         string `bson:"Classes"`
	WarningMessages                 string `bson:"Warning_Messages"`
	MandatoryMinMaxPatternEnumError string `bson:"Mandatory_Min_Max_Pattern_EnumError"`
	AffectedIDsFormula              string `bson:"AffectedIDs_Formula"`
	FormulaType                     string `bson:"Formula_Type"`
	CSV                             string `bson:"CSV"`
	EnumCondition                   string `bson:"EnumCondition"`
}


