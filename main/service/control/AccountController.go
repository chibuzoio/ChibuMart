package control

import (  
	"./utility";                  
    "./model";
)
         
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
      

