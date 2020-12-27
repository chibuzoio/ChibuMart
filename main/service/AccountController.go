package service  

import ( 
    "net/http";     
    
	"./control";
    "./control/model";
    "./control/utility";
	
	"github.com/gin-gonic/gin";    
    "github.com/gin-contrib/sessions";
)

func RegisterUser(context *gin.Context) { 
    var registrationRequest model.RegistrationRequest;
    var registrationResponse model.RegistrationResponse;
    
    session := sessions.Default(context);
    
    context.Bind(&registrationRequest);
    
    if (utility.DoesEmailExists(registrationRequest.EmailAddress)) {
        registrationResponse.Success = false;
        registrationResponse.Message = "Registration failed!";
    } else {
        // proceed with registration   
        session.Set("emailAddress", registrationRequest.EmailAddress);
        session.Save();
        
        // return registered memberId
        chibuMartId := control.StoreRegistrationData(registrationRequest);
        
        if (chibuMartId > 0) {
            userTableJSON := control.GenerateTableNames(chibuMartId);
            control.StoreGeneratedTableNames(userTableJSON);
            
            registrationResponse.Success = true;
            registrationResponse.Message = "Registration successful!";
        } else {
            registrationResponse.Success = false;
            registrationResponse.Message = "Registration failed!";   
        }
    }
    
    context.JSON(http.StatusOK, registrationResponse);
} 

func FetchUserData(context *gin.Context) {
    var userDataId int;
    var firstName, lastName string;
    var userDataJSON model.AllUserData;
    var userDataArray []model.AllUserData;
    connector := utility.GetConnection();
    
    defer connector.Close();
    
    query := "select * from userdatatable";
    
    rows, error := connector.Query(query);
    
    utility.Exception(error);
    
    for rows.Next() {
        error = rows.Scan(&userDataId, &firstName, &lastName);
        
        utility.Exception(error);
        
        userDataJSON.LastName = lastName;
        userDataJSON.FirstName = firstName;
        userDataJSON.UserDataId = userDataId;
        
        userDataArray = append(userDataArray, userDataJSON);
    }
    
    rows.Close();
    connector.Close();
    
    context.JSON(http.StatusOK, userDataArray);
}

func PostUserData(context *gin.Context) { 
    var userDataJSON model.UserData;    
        
    context.Bind(&userDataJSON);
        
    response := control.PostUserData(userDataJSON);
    
//         context.JSON(http.StatusOK, userDataJSON);                       
    context.JSON(http.StatusOK, gin.H{"theResponse" : response});                 
}
    
    
