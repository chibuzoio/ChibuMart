package control

import (  
    "fmt";         
    "time";
	"math/big";          
    
    "./model";             
	"./utility";                  
    
	"golang.org/x/crypto/bcrypt";
)
     
func StoreGeneratedTableNames(userTableJSON model.UserTable) {
    connector := utility.GetConnection();
    
	defer connector.Close();  
	 
    query := "insert into usertable (userTableId, chibuMartId, notificationTableName, " +  
        "productReceptionTable, productWishListTable, stockroomCartTable) values (?, ?, ?, ?, ?, ?)";
	
	stmt, error := connector.Prepare(query);
	
	utility.Exception(error);
	
	_, error = stmt.Exec(0, userTableJSON.ChibuMartId, userTableJSON.NotificationTableName, 
        userTableJSON.ProductReceptionTable, userTableJSON.ProductWishListTable, userTableJSON.StockroomCartTable);
	
	utility.Exception(error);
	
	stmt.Close();                  
	connector.Close();  
}

func GenerateTableNames(chibuMartId int) *model.UserTable {
	timeNow := time.Now(); 
    var userTableJSON model.UserTable;
	chibuMartIdPart := fmt.Sprintf("%05d", chibuMartId);
	currentUnixTime := fmt.Sprintf("%d", timeNow.Unix());
	chibuMartIdPart = chibuMartIdPart[len(chibuMartIdPart) - 5 : len(chibuMartIdPart)];      
    stockroomCartTable := "cart" + currentUnixTime + chibuMartIdPart; 
    productWishListTable := "productwish" + currentUnixTime + chibuMartIdPart;
    notificationTableName := "notification" + currentUnixTime + chibuMartIdPart; 
    productReceptionTable := "productreception" + currentUnixTime + chibuMartIdPart; 
          
    userTableJSON.ChibuMartId = chibuMartId;
    userTableJSON.StockroomCartTable = stockroomCartTable;
    userTableJSON.ProductWishListTable = productWishListTable;
    userTableJSON.NotificationTableName = notificationTableName;
    userTableJSON.ProductReceptionTable = productReceptionTable;
    
	return userTableJSON;
}

func StoreRegistrationData(registrationJSON) int {
	connector := utility.GetConnection(); 
	password := registrationJSON.Password;
	emailAddress := registrationJSON.EmailAddress; 

	defer connector.Close();
     
    query := "insert into chibumart (chibuMartId, firstName, lastName, " + 
        "gender, emailAddress, phoneNumber, town, state, country, password, " + 
        "regDate, profilePicture, billingAddressId, passwordTimestamp) values " + 
        "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)";
    
	timeNow := time.Now();              
    theDate := timeNow.Format("2006-01-02 15:04:05");               
    theTime := timeNow.Unix();               
    firstSalt := emailAddress[0 : 5];                   
    thePassword := string(firstSalt) + string(password) + fmt.Sprintf("%d", theTime);             

    passwordHash, error := bcrypt.GenerateFromPassword([]byte(thePassword), bcrypt.DefaultCost);                 

    utility.Exception(error);
    
    stmt, error := connector.Prepare(query);
    
    utility.Exception(error);
    
    _, error = stmt.Exec(0, "", "", "", emailAddress, "", "", "", "", 
        passwordHash, theDate, "", 0, theTime);
    
    utility.Exception(error);

    stmt.Close();
    
    var chibuMartId int;
    query = "select chibuMartId from chibumart where emailAddress = ?";
    
    resultSet, error := connector.Prepare(query);
    
    utility.Exception(error);
    
    rows, error := resultSet.Query(emailAddress);
    
    utility.Exception(error);
    
    for (rows.Next()) {
        error = rows.Scan(&chibuMartId);
        
        utility.Exception(error);
    }
    
    resultSet.Close();
    rows.Close();
    connector.Close();
    
    return chibuMartId;
}    

func PostUserData(userDataJSON model.UserData) string {
    connector := utility.GetConnection();
    
	defer connector.Close();  
	 
	query := "insert into userdatatable (userDataId, firstName, lastName) values (?, ?, ?)";    
	
	stmt, error := connector.Prepare(query);
	
	utility.Exception(error);
	
	_, error = stmt.Exec(0, userDataJSON.FirstName, userDataJSON.LastName);
	
	utility.Exception(error);
	
	stmt.Close();                  
	connector.Close();  
	 
    return "Successfully registered user " + userDataJSON.FirstName + " " + userDataJSON.LastName;
} 
      

