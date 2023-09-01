package main


import (
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("NewHomeLoanCC")

type NewHomeLoanCC struct {
}

type UserData struct {
    
	ApplicationId               string            `json:"applicationId"`
	BeneficiaryFullName        	string        	  `json:"beneficiaryfullName"`
	BeneficiaryID       	    string        	  `json:"beneficiaryfullId"`
	Status                      string        	  `json:"status"`
	CreatedDate			        string            `json:"createdDate"`
	CreatedBy				    string            `json:"createdBy"`
	BuyerFullName        	    string            `json:"buyerFullName"`
	Price                       string            `json:"price"`
	BuyerPassportNumber         string        	  `json:"buyerPassportNumber"`
	BuyerEmail		            string            `json:"buyerEmail"`
	BuyerPhoneNumber			string           `json:"buyerPhoneNumber"`
	AgentFullName        	    string            `json:"agentFullName"`
	AgentPassportNumber         string        	  `json:"agentPassportNumber"`
	AgentEmail		            string            `json:"agentEmail"`
	AgentPhoneNumber			string           `json:"agentPhoneNumber"`
	Remark                      string            `json:"remark"`
	Transaction			       []TransactionSchema`json:"transaction"`
	
	}



type TransactionSchema struct {
	TransactionId		   string            `json:"transactionId"`
	TimeStamp		   	   string            `json:"timeStamp"`
	TransactionType        string            `json:"transactionType"`
	UserId			       string		     `json:"userId"`
	Organization		   string            `json:"organization"`
}


type ResponceUser struct {
	Key           			string			  `json:"key"`	
	Record					UserData		      `json:"record"`
}


func (t *NewHomeLoanCC) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	logger.Info("########### NewHomeLoanCC Init ###########")
	
	retStr := fmt.Sprintf("NewHomeLoanCC Initialized successfully")
	return shim.Success([]byte(retStr))
}

// Transaction makes payment of X units from A to B
func (t *NewHomeLoanCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### NewHomeLoanCC Invoke ###########")
	function, args := stub.GetFunctionAndParameters()
	logger.Info("########### NewHomeLoanCC Invoke function :: ", function," :: args :: ", args)

	if(function == "addApplicationForm"){
		var userData UserData
	        json.Unmarshal([]byte(args[0]), &userData)
		err := stub.PutState(userData.ApplicationId, []byte(args[0]))
		if err != nil {
			logger.Info("########### NewHomeLoanCC Invoke :: err :: ", err.Error())
			return shim.Error(err.Error())
		}

		retStr := fmt.Sprintf("Data Added Sucessfully")
		return shim.Success([]byte(retStr))
	}

	if(function == "getUserFormByApplicationId"){

		queryString := fmt.Sprintf("{\"selector\":{\"applicationId\":\"%s\"}}", args[0])

		finalCaseData, err := getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(finalCaseData)	
	}

if(function == "findBySellerMail"){
		
		
			queryString := fmt.Sprintf("{\"selector\":{\"email\":\"%s\"}}", args[0])

		finalCaseData, err := getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(finalCaseData)
	}
	
	if(function == "findByBuyerMail"){
		
		
			queryString := fmt.Sprintf("{\"selector\":{\"buyerEmail\":\"%s\"}}", args[0])

		finalCaseData, err := getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(finalCaseData)
	}
	
	
	
	if(function == "findAll"){
		
		
			queryString := fmt.Sprintf("{\"selector\":{\"createdBy\":\"%s\"}}", args[0])

		finalCaseData, err := getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(finalCaseData)
	}

	retStr := fmt.Sprintf("No Valid Function")
	return shim.Error(retStr)
	
}



func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	fmt.Println("=======before loop===");
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		fmt.Println("=======inside loop===");
		fmt.Println(queryResponse);
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		fmt.Println(queryResponse.Key);
		fmt.Println(queryResponse.Value);
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return []byte(buffer.String()), nil
}

func main() {
	err := shim.Start(new(NewHomeLoanCC))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}