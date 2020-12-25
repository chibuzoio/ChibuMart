package model  

import (

)

type RegistrationResponse struct {      
    Success bool `json:"success"`;
    Message string `json:"message"`;
}     

type RegistrationRequest struct {	
    EmailAddress string `json:"emailAddress" binding:"required"`;   
    Password string `json:"password" binding:"required"`; 
}


