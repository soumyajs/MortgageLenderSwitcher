package main

import (
	"encoding/json"
	"errors"
	"fmt"
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)



// MORTGAGE is a high level smart contract that MORTGAGEs together business artifact based smart contracts
type MORTGAGE struct {

}

// BorrowerDetails is for storing User Details

type BorrowerDetails struct{	
	uid string `json:"uid"`	
	Gender string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Dob string `json:"dob"`
	Email string `json:"email"`
	Phone string `json:"phone"`	
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
	LenderId string `json:"lenderId"`
	LenderName string `json:"lenderName"`
	ProductName string `json:"productName"`
	LoanAmount string `json:"loanAmount"`
	InterestRate string `json:"interestRate"`	
	DocumentsSubmitted string `json:"documentsSubmitted"`
	SwitchUserFlag string `json:"switchUserFlag"`
	SwitchLenderId string `json:"switchLenderId"`
}

// LenderDetails is for storing Lender Details

type LenderDetails struct{	
	LenderId string `json:"lenderId"`
	LenderName string `json:"lenderName"`
	ProductId string   `json:"productId"`
	ProductName string `json:"productName"`
	ProductType string `json:"productType"`
	Rate string `json:"rate"`
	
	
}

// Transaction is for storing transaction Details


// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
}
	

func createLenderDetails(stub shim.ChaincodeStubInterface) error {
	// Create table one
	var columnDefsTableOne []*shim.ColumnDefinition
	columnOneTableOneDef := shim.ColumnDefinition{Name: "lenderId", Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableOneDef := shim.ColumnDefinition{Name: "lenderName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableOneDef := shim.ColumnDefinition{Name: "productId", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFourTableOneDef := shim.ColumnDefinition{Name: "productName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFiveTableOneDef := shim.ColumnDefinition{Name: "productType", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSixTableOneDef := shim.ColumnDefinition{Name: "rate", Type: shim.ColumnDefinition_STRING, Key: false}
	columnDefsTableOne = append(columnDefsTableOne, &columnOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwoTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThreeTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFourTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFiveTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSixTableOneDef)
	fmt.Println("creating table !!!!! %s",columnDefsTableOne)
	return stub.CreateTable("LenderDetails", columnDefsTableOne)
}


// Init initializes the smart contracts
func (t *MORTGAGE) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

fmt.Println("creating table !!!!! ")
	// Check if table already exists
	_, err := stub.GetTable("LenderDetails")
	fmt.Println("err  table !!!!! &b",err)
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = createLenderDetails(stub)
	
	
	
	
	fmt.Println("creating table created!!!!! %b",err)
	if err != nil {
		return nil, errors.New("Failed creating UserDetails.")
	}
		
	
	
	
	
	
	// Check if table already exists
	_, err = stub.GetTable("BorrowerDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("BorrowerDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "uid", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "phone", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lenderId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lenderName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "productName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "loanAmount", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "interestRate", Type: shim.ColumnDefinition_STRING, Key: false},		
		&shim.ColumnDefinition{Name: "documentsSubmitted", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "switchUserFlag", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "switchLenderId", Type: shim.ColumnDefinition_STRING, Key: false},		
	})
	
	if err != nil {
		return nil, errors.New("Failed creating BorrowerDetails.")
	}
	
	return nil, nil
}
	

	
//registerUser to register a user
func (t *MORTGAGE) registerLender(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("no of argumrents  %d",len(args) )

		
	
		lenderId:=args[0]
		lenderName:=args[1]
		productId:=args[2]
		productName:=args[3]
		productType:=args[4]
		rate:=args[5]
		

		
		
		


		// Insert a row
		ok, err := stub.InsertRow("LenderDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: lenderId}},
				&shim.Column{Value: &shim.Column_String_{String_: lenderName}},
				&shim.Column{Value: &shim.Column_String_{String_: productId}},
				&shim.Column{Value: &shim.Column_String_{String_: productName}},
				&shim.Column{Value: &shim.Column_String_{String_: productType}},
				&shim.Column{Value: &shim.Column_String_{String_: rate}},
				
			}})

			fmt.Println("Iserted rows !!!!! %b",ok)
		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}



// to get the deatils of a product against productId 
func (t *MORTGAGE) getProductRatesFromLender(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("Hello, World!")
	

	productId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this productId
	var columns []shim.Column
	//col3 := shim.Column{Value: &shim.Column_String_{String_: productId}}
	
	//columns = append(columns, col3)
fmt.Println("Hello, World!!!")
	rows, err := stub.GetRows("LenderDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the product with id  " + productId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	
	

	
	res2E :=  []*LenderDetails{}
	for row := range rows { 
	newApp:= new(LenderDetails)
	newApp.LenderId = row.Columns[0].GetString_()
	newApp.LenderName = row.Columns[1].GetString_()
	newApp.ProductId = row.Columns[2].GetString_()
	newApp.ProductName = row.Columns[3].GetString_()
	newApp.ProductType = row.Columns[4].GetString_()
	newApp.Rate = row.Columns[5].GetString_()
	if newApp.ProductId == productId{
		res2E=append(res2E,newApp)		
		}	
		
		}
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// to get the deatils of a product against lenderId 
func (t *MORTGAGE) getLenderDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("Hello, World!")
	

	lenderId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this lenderId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: lenderId}}
	
	columns = append(columns, col1)
fmt.Println("Hello, World!!!")
	row, err := stub.GetRow("LenderDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the lender with id  " + lenderId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	
	

	
	res2E :=  LenderDetails{}
	
	res2E.LenderId = row.Columns[0].GetString_()
	res2E.LenderName = row.Columns[1].GetString_()
	res2E.ProductId = row.Columns[2].GetString_()
	res2E.ProductName = row.Columns[3].GetString_()
	res2E.ProductType = row.Columns[4].GetString_()
	res2E.Rate = row.Columns[5].GetString_()
	
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}
func (t *MORTGAGE) switchLenders(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("Hello, World!")
	

	currentLenderId := args[0]
	switchLenderId := args[1]
	borrowerId := args[2]
	fmt.Println("Hello, World!!")

	var columns []*shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: borrowerId}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: switchLenderId}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: currentLenderId}}
		columns = append(columns, &col1)
		columns = append(columns, &col2)
		columns = append(columns, &col3)

		row := shim.Row{Columns: columns}
	//row, err := stub.GetRow("BorrowerDetails", columns)
	ok, err := stub.ReplaceRow("BorrowerDetails", row)
		if err != nil {
			return nil, fmt.Errorf("replaceRowTableOne operation failed. %s", err)
		}
		if !ok {
			return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
		}

return nil, nil
}


// Invoke invokes the chaincode
func (t *MORTGAGE) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerLender" {
		t := MORTGAGE{}
		return t.registerLender(stub, args)	
	}

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *MORTGAGE) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getProductRatesFromLender" { 
		t := MORTGAGE{}
		return t.getProductRatesFromLender(stub, args)
	}else if function == "getLenderDetails" {
		t := MORTGAGE{}
		return t.getLenderDetails(stub, args)	
	} 	

	
	return nil, nil
}

func main() {
	
	err := shim.Start(new(MORTGAGE))
	if err != nil {
		fmt.Printf("Error starting MORTGAGE: %s", err)
	}
} 