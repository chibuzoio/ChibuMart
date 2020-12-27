package service  

import ( 
    "strings";
    "net/http";     
    
	"./control";
    "./control/model";
    "./control/utility";
	
	"github.com/gin-gonic/gin";    
    "github.com/gin-contrib/sessions";
)

func LoginUser(context *gin.Context) {   
    var loginRequest model.LoginRequest;
    var loginResponse model.LoginResponse;
    
    session := sessions.Default(context);
    
    context.Bind(&loginRequest);
    
    loginRequest.EmailAddress = strings.TrimSpace(loginRequest.EmailAddress);
    
    if (utility.DoesEmailExists(loginRequest.EmailAddress)) {
        if (control.IsPasswordValid(loginRequest)) {
            session.Set("emailAddress", loginRequest.EmailAddress);
            session.Save();
        
            loginResponse.Data = control.GetUserLoginData(loginRequest.EmailAddress);
            loginResponse.Message = "Login successful";
            loginResponse.Success = true;
        } else {
            loginResponse.Message = "Email or password is wrong";
            loginResponse.Success = false;
        }
    } else {
        loginResponse.Message = "Email or password is wrong";
        loginResponse.Success = false;
    }
    
    context.JSON(http.StatusOK, loginResponse);
}

func RegisterUser(context *gin.Context) { 
    var registrationRequest model.RegistrationRequest;
    var registrationResponse model.RegistrationResponse;
    
    session := sessions.Default(context);
    
    context.Bind(&registrationRequest);
 
    registrationRequest.EmailAddress = strings.TrimSpace(registrationRequest.EmailAddress);
    
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
    
    
