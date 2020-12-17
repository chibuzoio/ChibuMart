package control

import (      
	"./utility";
    "database/sql";
    _ "github.com/go-sql-driver/mysql";
)
 
func CreateProjectTables() {
	tableCollection := []string{               
		"create table if not exists chibumart (" +
			"chibuMartId bigint(20) unsigned not null auto_increment, " + 
			"firstName varchar(23), " +
			"lastName varchar(23), " +     
			"gender varchar(11), " +
			"emailAddress varchar(55), " +
			"phoneNumber varchar(23), " +
			"town varchar(55), " +
			"state varchar(55), " +
			"country varchar(55), " +
			"password varchar(333) not null, " +
			"regDate varchar(23) not null, " +
			"profilePicture varchar(55), " +
			"billingAddressId bigint(20) unsigned, " +
			"primary key(chibuMartId), " +
			"unique key(userName) " +
			") engine = InnoDB default charset = utf8", 
            
		"create table if not exists usertable (" + 
			"userTableId bigint(20) unsigned not null auto_increment, " +
			"chibuMartId bigint(20) unsigned not null, " + 
			"archiveTableName varchar(111) not null, " +
			"virtualAccountTable varchar(111) not null, " + 
			"primary key(userTableId), " +           
			"unique key(chibuMartId) " +          
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
	 
	_, executionError := connector.Exec("create database if not exists chibumart");

	if executionError != nil {
		utility.Exception(executionError);
	} else {
		utility.Println("Successfully created database.....");
	}
	
	connector.Close();
}


