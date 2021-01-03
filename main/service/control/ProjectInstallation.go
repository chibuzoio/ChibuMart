package control

import (      
	"./utility";
    "database/sql";
    _ "github.com/go-sql-driver/mysql";
)
 
func CreateProjectTables() {
	tableCollection := []string{         
		"create table if not exists productimages (" +
			"productImageId bigint(20) unsigned not null auto_increment, " +   
			"productId bigint(20) unsigned not null, " +
			"productImageName varchar(55), " +
			"commentTableName varchar(111), " +
			"likeTableName varchar(111), " +
			"numberOfComments int(7), " +
			"numberOfLikes int(7), " +
			"primary key(productImageId) " + 
			") engine = InnoDB default charset = utf8", 

		"create table if not exists chibumartimages (" +
			"image varchar(55) not null, " +
			"width int(7) unsigned not null, " +
			"height int(7) unsigned not null, " +
			"primary key(image) " +
			") engine = InnoDB default charset = utf8",                   

		"create table if not exists productcollection (" +
			"productId bigint(20) unsigned not null auto_increment, " +      
			"productName varchar(111) not null, " +
			"productCategory varchar(111) not null, " +
			"productQuantityRemaining bigint(15) unsigned, " +
			"productQuantityRetailed bigint(15) unsigned, " + // for top selling products, different from trending products              
			"productQuantityTotal bigint(15) unsigned not null, " +
			"productPreviousPrice decimal(33, 11) not null, " +
			"productCurrentPrice decimal(33, 11) not null, " +
			"placementDate varchar(23) not null, " +
			"incrementDate varchar(23), " +
			"retailDate varchar(23), " +
			"descriptionId bigint(20) unsigned, " +
			"numberOfComments int(7), " +
			"numberOfLikes int(7), " +
			"allReactionsTotal int(9), " + // for trending products, different from top selling products                         
			"commentTableName varchar(111), " +
			"likeTableName varchar(111), " +
			"productLocation varchar(111), " +
			"primary key(productId), " +    
			"key(descriptionId) " + 
			") engine = InnoDB default charset = utf8",                    
      
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
			"passwordTimestamp varchar(23) not null, " + 
			"primary key(chibuMartId), " +
			"unique key(emailAddress) " +
			") engine = InnoDB default charset = utf8", 
            
		"create table if not exists usertable (" + 
			"userTableId bigint(20) unsigned not null auto_increment, " +
			"chibuMartId bigint(20) unsigned not null, " +          
			"notificationTableName varchar(111) not null, " +
			"productReceptionTable varchar(111) not null, " +
			"productWishListTable varchar(111) not null, " +
			"stockroomCartTable varchar(111) not null, " +
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


