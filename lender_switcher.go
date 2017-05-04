package main

import (
	"encoding/json"
	"errors"
	"fmt"
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)



 // FFP is a high level smart contract that FFPs together business artifact based smart contracts
type FFP struct {

}


// UserDetails is for storing User Details

type UserDetails struct{	
	FfId string `json:"ffId"`
	Title string `json:"title"`
	Gender string `json:"gender"`
	
}

// Transaction is for storing transaction Details


// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
}
	



// Init initializes the smart contracts
func (t *FFP) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("UserDetails")
	fmt.Println("creating table !!!!! %b",err)
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("UserDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ffId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		
	})
	if err != nil {
		return nil, errors.New("Failed creating UserDetails.")
	}
		
	
	
	return nil, nil
}
	

	
//registerUser to register a user
func (t *FFP) registerUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 12 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 12. Got: %d.", len(args))
		}
		
		ffId:=args[0]
		title:=args[1]
		gender:=args[2]
		
		
	
fmt.Println("Iserted rows with columsn ::::: "+ffId+"   "+title+"   "+gender )
		
		
		


		// Insert a row
		ok, err := stub.InsertRow("UserDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: ffId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				
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



// to get the deatils of a user against ffid (for internal testing, irrespective of org)
func (t *FFP) getUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("Hello, World!")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this ffId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: ffId}}
	columns = append(columns, col1)
fmt.Println("Hello, World!!!")
	row, err := stub.GetRow("UserDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed subbu is changins the msg " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	fmt.Println("columns are  !!!!! %d",len(row.Columns))
	
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	res2E := UserDetails{}
	
	res2E.FfId = row.Columns[0].GetString_()
	res2E.Title = row.Columns[1].GetString_()
	res2E.Gender = row.Columns[2].GetString_()
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}




// Invoke invokes the chaincode
func (t *FFP) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerUser" {
		t := FFP{}
		return t.registerUser(stub, args)	
	} 

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *FFP) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getUser" { 
		t := FFP{}
		return t.getUser(stub, args)
	}
	
	return nil, nil
}

func main() {
	
	err := shim.Start(new(FFP))
	if err != nil {
		fmt.Printf("Error starting FFP: %s", err)
	}
} 