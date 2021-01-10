package utility

import ( 
    "database/sql";
)

func CreateProductWishTable(connector *sql.DB, tableName string) {                                                       
	query := "create table if not exists " + tableName + " (" +
		"productWishId bigint(15) unsigned not null auto_increment, " +
		"productId bigint(20) unsigned not null, " +
		"wishDate varchar(23) not null, " +
		"primary key(productWishListId) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}                                                        

func CreateProductCartTable(connector *sql.DB, tableName string) {
	query := "create table if not exists " + tableName + " (" +
		"productCartId bigint(20) unsigned not null auto_increment, " +
		"productId bigint(20) unsigned not null, " +
		"productQuantity int(9) not null default 0, " +
		"primary key(productCartId), " +
		"unique key(productId) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
} 

func CreateAllWatchUserTable(connector *sql.DB, tableName string) {                                             
	query := "create table if not exists " + tableName + " (" +
		"allWatchUserTableId bigint(20) unsigned not null auto_increment, " +
		"computeBoneId bigint(20) unsigned not null, " +     
		"primary key(allWatchUserTableId), " +     
		"unique key(computeBoneId) " +   
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	ErrorStackTrace(error, "18 of TableFactoryManager");

	Println("Successfully Created " + tableName + " Table...");
}                                                                 

func CreateProfileWatcherTable(connector *sql.DB, tableName string) {
	query := "create table if not exists " + tableName + " (" +
		"watcherId bigint(20) unsigned not null auto_increment, " +
		"computeBoneId bigint(20) unsigned not null, " +      
		"allWatchUserTable varchar(111) not null, " +           
		"primary key(watcherId), " +                
		"unique key(allWatchUserTable), " +           
		"foreign key(computeBoneId) references " +           
		"computebone(computeBoneId) on update cascade on delete cascade" +          
		") engine = InnoDB default charset = utf8";           

	_, error := (*connector).Exec(query);

	ErrorStackTrace(error, "36 of TableFactoryManager");

	Println("Successfully Created " + tableName + " Table...");
}

func CreateInterimAdvertisementTable(connector *sql.DB, tableName string) {    
	query := "create table if not exists " + tableName + " (" +
		"actionCounterId bigint(20) unsigned not null auto_increment, " +
		"computeBoneId bigint(20) unsigned not null, " +
		"actionCounterType varchar(7) not null, " +
		"primary key(actionCounterId), " + 
		"key(computeBoneId), " +          
		"key(actionCounterType) " +       
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	ErrorStackTrace(error, "53 of TableFactoryManager");

	Println("Successfully Created " + tableName + " Table...");   
} 

func AlterMessengerTable(connector *sql.DB, tableName string) {
	query := "alter table " + tableName + " add column " + 
		"gotNewMessage tinyint(1) not null, add column changeIndexer " +
		"bigint(20) unsigned not null, add index gotNewMessage (gotNewMessage), " + 
		"add index changeIndexer (changeIndexer)";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully added gotNewMessage and changeIndexer columns to " + tableName + " table...");
}

func CreateProfilePictureTable(connector *sql.DB, tableName string) {
	query := "create table if not exists " + tableName + " (" +
		"profilePictureId bigint(20) unsigned not null auto_increment, " +
		"profilePictureImage varchar(55) not null, " +
		"commentTableName varchar(111) not null, " +
		"likeTableName varchar(111) not null, " +
		"uploadDate varchar(23) not null, " +
		"primary key(profilePictureId), " + 
		"unique key(profilePictureImage) " + 
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
} 

func CreateMessengerMessageTable(connector *sql.DB, tableName string) {
	query := "create table if not exists " + tableName + " (" +
		"chatMessageId bigint(20) unsigned not null auto_increment, " +
		"chatMessage text not null, " +
		"messageDate varchar(23) not null, " +
		"messageOrigin bigint(20) unsigned not null, " +
		"readStatus tinyint(1) not null, " +
		"seenStatus tinyint(1) not null, " +        
		"deleteMessage varchar(23) not null, " +   
		"primary key(chatMessageId) " + 
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}

func CreateMessengerCompositeTable(connector *sql.DB, tableName string) {
	query := "create table if not exists " + tableName + " (" +   
		"messengerId bigint(20) unsigned not null auto_increment, " +
		"messengerChatmateId bigint(20) unsigned not null, " +
		"messengerMessageTable varchar(111) not null, " +
		"requestStatus tinyint(1) not null, " +
		"unseenMessage tinyint(1) not null, " +
		"gotNewMessage tinyint(1) not null, " +
		"changeIndexer bigint(20) unsigned not null, " +
		"primary key(messengerId), " +
		"unique key(messengerChatmateId), " + 
		"key(unseenMessage), " + 
		"key(gotNewMessage), " + 
		"key(changeIndexer) " + 
		") engine = InnoDB default charset = utf8";
  
	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}

func CreateProductReceptionTable(connector *sql.DB, tableName string) {   
	query := "create table if not exists " + tableName + " (" +
		"productReceptionId bigint(20) unsigned not null auto_increment, " +
		"productDeliveryId bigint(20) unsigned not null, " +
		"productId bigint(20) unsigned not null, " +
		"productDeliveryTable varchar(111) not null, " +
		"receptionStatus varchar(23) not null, " +
		"primary key(productReceptionId), " +
		"key(productDeliveryTable), " +
		"key(productId) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}

func CreateProductDeliveryTable(connector *sql.DB, tableName string) { 
	query := "create table if not exists " + tableName + " (" +
		"productDeliveryId bigint(20) unsigned not null auto_increment, " +
		"productId bigint(20) unsigned not null, " +
		"computeBoneId bigint(20) unsigned not null, " +
		"productQuantity int(9), " +
		"deliveryStatus varchar(23) not null, " +
		"deliveryDate varchar(23) not null, " +
		"productReceptionTable varchar(111) not null, " +
		"primary key(productDeliveryId), " +
		"key(deliveryStatus), " +
		"key(computeBoneId) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}
       
func CreateSalesPerformanceTable(connector *sql.DB, tableName string) {
	query := "create table if not exists " + tableName + " (" + 
		"salesPerformanceId bigint(20) unsigned not null auto_increment, " +
		"theClientId varchar(55) not null, " +
		"productId bigint(20) unsigned not null, " +
		"computeBoneId bigint(20) unsigned not null, " +
		"productQuantity int(9), " +
		"rollbackTime bigint(20) unsigned not null, " +
		"stockroomWagonTable varchar(111) not null, " +
		"primary key(salesPerformanceId), " +
		"unique key(theClientId), " +
		"key(computeBoneId) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}

func CreateAllFollowedStockroomTable(connector *sql.DB, tableName string) {                                             
	query := "create table if not exists " + tableName + " (" +
		"allFollowedStockroomId bigint(20) unsigned not null auto_increment, " +
		"stockroomId bigint(20) unsigned not null, " +
		"primary key(allFollowedStockroomId), " +
		"unique key(stockroomId) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}                                                                 

func CreateStockroomSubscriberTable(connector *sql.DB, tableName string) {
	query := "create table if not exists " + tableName + " (" +
		"subscriberId bigint(20) unsigned not null auto_increment, " +
		"computeBoneId bigint(20) unsigned not null, " + 
		"allFollowedStockroomTable varchar(111) not null, " +
		"primary key(subscriberId), " +
		"unique key(allFollowedStockroomTable), " + 
		"foreign key(computeBoneId) references " + 
		"computebone(computeBoneId) on update cascade on delete cascade" + 
		") engine = InnoDB default charset = utf8";    

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}

func CreateAllTheWatchTable(connector *sql.DB, tableName string) {                                             
	query := "create table if not exists " + tableName + " (" +
		"allTheWatchTableNameId bigint(20) unsigned not null auto_increment, " +
		"billboardCategoryId bigint(20) unsigned not null, " +
		"specialty varchar(55) not null, " +
		"primary key(allTheWatchTableNameId), " +
		"unique key(billboardCategoryId), " +
		"key(specialty) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}                                                                 

func CreateGenericLikeTable(connector *sql.DB, tableName string) {                                             
	query := "create table if not exists " + tableName + " (" +
		"likeId bigint(20) unsigned not null auto_increment, " +
		"computeBoneId bigint(20) unsigned not null, " +
		"primary key(likeId), " +
		"unique key(computeBoneId) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}                                                                 

func CreateReplyTable(connector *sql.DB, tableName string) {                                                       
	query := "create table if not exists " + tableName + " (" +
		"replyId int(7) unsigned not null auto_increment, " +
		"theClientId varchar(55) not null, " +
		"computeBoneId bigint(20) unsigned not null, " +
		"reply text, " +
		"numberOfLikes int(7), " +
		"replyDate varchar(23) not null, " +
		"likeTableName varchar(111), " +
		"primary key(replyId), " +
		"unique key(theClientId) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}                                                        

func CreateCommentTable(connector *sql.DB, tableName string) {                                                       
	query := "create table if not exists " + tableName + " (" +
		"commentId int(7) unsigned not null auto_increment, " +
		"theClientId varchar(55) not null, " +
		"computeBoneId bigint(20) unsigned not null, " +
		"comment text, " +
		"numberOfLikes int(9), " +
		"numberOfReplies int(7), " +
		"commentDate varchar(23) not null, " +
		"likeTableName varchar(111), " +
		"replyTableName varchar(111), " +
		"primary key(commentId), " +
		"unique key(theClientId) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}                                                        

func CreateRequirementTable(connector *sql.DB, tableName string) {                                                       
	query := "create table if not exists " + tableName + " (" +
		"requirementId int(7) unsigned not null auto_increment, " +
		"title varchar(111) not null, " +
		"content varchar(777) not null, " +
		"primary key(requirementId), " +
		"key(title) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}                                                        

func CreateExhibitArchiveTable(connector *sql.DB, tableName string) {                                                       
	query := "create table if not exists " + tableName + " (" +
		"exhibitArchiveId bigint(15) unsigned not null auto_increment, " +
		"exhibitId bigint(20) unsigned not null, " +
		"specialty varchar(55) not null, " +
		"archiveDate varchar(23) not null, " +
		"primary key(exhibitArchiveId), " +
		"key(specialty) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}                                                        

func CreateNotificationTable(connector *sql.DB, tableName string) {
	query := "create table if not exists " + tableName + " (" +
		"notificationId bigint(20) unsigned not null auto_increment, " +
		"notificationOriginId bigint(20) unsigned not null, " +
		"notificationEffectorId bigint(20) unsigned not null, " +
		"notificationOriginTable varchar(111) not null, " +
		"genericNotification varchar(151) not null, " +
		"notificationDate varchar(23) not null, " +
		"notificationType varchar(111) not null, " +
		"seenStatusFront tinyint(1) not null, " +
		"seenStatusBack tinyint(1) not null, " +
		"readStatus tinyint(1) not null, " +
		"notificationRequirement mediumtext, " +
		"primary key(notificationId), " +
		"key(notificationOriginId), " +
		"key(notificationOriginTable) " +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}

func CreateWatchersTable(connector *sql.DB, tableName string) {
	query := "create table if not exists " + tableName + " (" +
		"watcherId bigint(20) unsigned not null auto_increment, " +
		"computeBoneId bigint(20) unsigned not null, " +
		"allTheWatchTableName varchar(111) not null, " +
		"primary key(watcherId), " +
		"unique key(allTheWatchTableName), " +
		"foreign key(computeBoneId) references " +
		"computebone(computeBoneId) on update cascade on delete cascade" +
		") engine = InnoDB default charset = utf8";

	_, error := (*connector).Exec(query);

	Exception(error);

	Println("Successfully Created " + tableName + " Table...");
}


