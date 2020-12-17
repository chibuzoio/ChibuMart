package model

import (  

)

type UserData struct {	
    LastName string `json:"lastName" binding:"required"`;   
    FirstName string `json:"firstName" binding:"required"`; 
}

type AllUserData struct {
    LastName string `json:"lastName" binding:"required"`;   
    UserDataId int `json:"userDataId" binding:"required"`;   
    FirstName string `json:"firstName" binding:"required"`; 
}
  
    
