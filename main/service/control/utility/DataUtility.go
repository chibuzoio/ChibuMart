package utility

import ( 

)

func DeleteCartProductArray(cartTableIdArray []int) {
    for _, value := range cartTableIdArray {
        DeleteCartProduct(value);
    }
}

func DeleteCartProduct(cartTableId int) {
    connector := GetConnection(); 
    
    defer connector.Close();
    
    query := "delete from chibumartcart where cartTableId = ?";
    
    stmt, error := connector.Prepare(query);
    
    Exception(error);
    
    _, error = stmt.Exec(cartTableId);
    
    Exception(error);
    
    stmt.Close();
    connector.Close();
}

func DoesEmailExists(emailAddress string) bool {
    connector := GetConnection(); 
    
    defer connector.Close();
    
    var chibuMartId int;
    
    query := "select chibuMartId from chibumart where emailAddress = ?";  
    
	resultSet, error := connector.Prepare(query);
	
	Exception(error);
	
	rows, error := resultSet.Query(emailAddress);
	
	Exception(error);
	
	for rows.Next() {
		error = rows.Scan(&chibuMartId);
		
		Exception(error);
	} 
	
	resultSet.Close();
	rows.Close();
	
    if (chibuMartId > 0) {
        return true;
    } else {
        return false;
    }
}


