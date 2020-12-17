package control

import (      
	"./utility";
    "database/sql";
    _ "github.com/go-sql-driver/mysql";
)
 
func CreateProjectTables() {
	tableCollection := []string{               
		"create table if not exists userdatatable (" +
			"userDataId bigint(20) unsigned not null auto_increment, " + 
			"firstName varchar(23), " +
			"lastName varchar(23), " +     
			"primary key(userDataId) " + 
			") engine = InnoDB default charset = utf8", 
	};                                  

	connector := utility.GetConnection();

	defer connector.Close();

	for _, table := range tableCollection {
	    _, err := connector.Exec(table);
	    
	    utility.Exception(err);     
	}

	connector.Close(); 
}

func CreateProjectDatabase() {        
	connector, connectionError := sql.Open("mysql", "computebone:2352C/C++solu+++;@/");

	if connectionError != nil {
	 	utility.Exception(connectionError);
	} else {
		utility.Println("Connection to mysql server gotten successfully");
	}
	 
	_, executionError := connector.Exec("create database if not exists assignment");

	if executionError != nil {
		utility.Exception(executionError);
	} else {
		utility.Println("Successfully created database.....");
	}
	
	connector.Close();
}


