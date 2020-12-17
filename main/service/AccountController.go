package service  

import ( 
    "net/http";     
    
	"./control";
    "./control/model";
    "./control/utility";
	
	"github.com/gin-gonic/gin";   
)
     
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
        
    response := control.PostUserData(userDataJSON)
    
//         context.JSON(http.StatusOK, userDataJSON);                       
    context.JSON(http.StatusOK, gin.H{"theResponse" : response});                 
}
    
    
