package utility

import ( 

)

func DoesEmailExists(emailAddress string) bool {
    connector := GetConnection(); 
    
    defer connector.Close();
    
    var chibuMartId int;
    
    query := "select chibuMartId from chibumart where emailAddress = ?";  
    
	resultSet, error := connector.Prepare(query);
	
	utility.Exception(error);
	
	rows, error := resultSet.Query(emailAddress);
	
	utility.Exception(error);
	
	for rows.Next() {
		error = rows.Scan(&chibuMartId);
		
		utility.Exception(error);
	} 
	
	resultSet.Close();
	rows.Close();
	
    if (chibuMartId > 0) {
        return true;
    } else {
        return false;
    }
}


