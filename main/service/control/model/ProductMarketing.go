package model  

import (
    
)

type CartProductResponse struct {
    Success bool `json:"success"`;
    Message string `json:"message"`;
}

type CartProductRequest struct {
    ProductId int `json:"productId" binding:"required"`;   
    EmailAddress string `json:"emailAddress" binding:"required"`;   
}

type AddWishedProduct struct {
    ProductId int;
    ProductWishTable string;
}

type WishProductResponse struct {
    Success bool `json:"success"`;
    Message string `json:"message"`;
}

type WishProductRequest struct {
    ProductId int `json:"productId" binding:"required"`;   
    EmailAddress string `json:"emailAddress" binding:"required"`;   
}

type AddProductResponse struct {
    Success bool `json:"success"`;
    Message string `json:"message"`;
}

type AddProductRequest struct { 
    ProductPrice string `json:"productPrice" binding:"required"`;    
    ProductCategory string `json:"productCategory" binding:"required"`;    
    ProductImage string `json:"productImage" binding:"required"`;    
    ProductName string `json:"productName" binding:"required"`;   
    ImageWidth int `json:"imageWidth" binding:"required"`; 
    ImageHeight int `json:"imageHeight" binding:"required"`;  
    ProductQuantity int `json:"productQuantity" binding:"required"`; 
    EmailAddress string `json:"emailAddress" binding:"required"`;   
}


