package model

import (  

)

type RegistrationResponse struct {      
    
}     

type UserData struct {	
    LastName string `json:"lastName" binding:"required"`;   
    FirstName string `json:"firstName" binding:"required"`; 
}

type Registration struct {	
    EmailAddress string `json:"emailAddress" binding:"required"`;   
    Password string `json:"password" binding:"required"`; 
}

type AllUserData struct {
    LastName string `json:"lastName" binding:"required"`;   
    UserDataId int `json:"userDataId" binding:"required"`;   
    FirstName string `json:"firstName" binding:"required"`; 
}

    
