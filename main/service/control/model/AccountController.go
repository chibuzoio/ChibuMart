package model  

import (

)

type LoginData struct {
    ChibuMartId int `json:"chibuMartId"`;
    FirstName string `json:"firstName"`;
    LastName string `json:"lastName"`;
    Gender string `json:"gender"`;
    EmailAddress string `json:"emailAddress"`;
    PhoneNumber string `json:"phoneNumber"`;
    Town string `json:"town"`;
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
}     

type RegistrationRequest struct {	
    EmailAddress string `json:"emailAddress" binding:"required"`;   
    Password string `json:"password" binding:"required"`; 
}


