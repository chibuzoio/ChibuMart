package model  

import (

)

type LogoutResponse struct {            
    Success bool `json:"success"`;
    Message string `json:"message"`;
}

type LoginData struct {
    ChibuMartId int `json:"chibuMartId"`;
    FirstName string `json:"firstName"`;
    LastName string `json:"lastName"`;
    Gender string `json:"gender"`;
    EmailAddress string `json:"emailAddress"`;
    PhoneNumber string `json:"phoneNumber"`;
    Town string `json:"town"`;
    State string `json:"state"`;
    Country string `json:"country"`;
    RegDate string `json:"regDate"`;
    ProfilePicture string `json:"profilePicture"`;
    BillingAddressId int `json:"billingAddressId"`;
}

type LoginRequest struct {
    EmailAddress string `json:"emailAddress" binding:"required"`;   
    Password string `json:"password" binding:"required"`; 
}

type LoginResponse struct {
    Success bool `json:"success"`;
    Message string `json:"message"`;
    Data LoginData `json:"data"`;
}

type RegistrationResponse struct {      
    Success bool `json:"success"`;
    Message string `json:"message"`;
    Data LoginData `json:"data"`;
}     

type RegistrationRequest struct {	
    EmailAddress string `json:"emailAddress" binding:"required"`;   
    Password string `json:"password" binding:"required"`; 
}


