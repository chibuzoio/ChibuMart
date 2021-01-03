package model

import (  

)

type ImageProperties struct {            
    ContentId int;        
    ImageName string;
    ImageWidth int;
    ImageHeight int;
}

type GenericImage struct {               
    ContentId int;
    ImageString string;
    ImageType string;
}

type UserTable struct {
    ChibuMartId int `json:"chibuMartId"`;
    StockroomCartTable string `json:"stockroomCartTable"`;
    ProductWishListTable string `json:"productWishListTable"`;
    NotificationTableName string `json:"notificationTableName"`;
    ProductReceptionTable string `json:"productReceptionTable"`;
} 

type UserData struct {	
    LastName string `json:"lastName" binding:"required"`;   
    FirstName string `json:"firstName" binding:"required"`; 
}

type AllUserData struct {
    LastName string `json:"lastName" binding:"required"`;   
    UserDataId int `json:"userDataId" binding:"required"`;   
    FirstName string `json:"firstName" binding:"required"`; 
}

    
