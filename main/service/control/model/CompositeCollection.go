package model  

import (
    
)

type FetchProductData struct { 
    ProductId int `json:"productId"`;
    ProductName String `json:"productName"`;
    ProductCategory String `json:"productCategory"`;
    ProductQuantityRemaining int `json:"productQuantityRemaining"`;
    ProductQuantityRetailed int `json:"productQuantityRetailed"`;
    ProductQuantityTotal int `json:"productQuantityTotal"`;
    ProductPreviousPrice String `json:"productPreviousPrice"`;
    ProductCurrentPrice String `json:"productCurrentPrice"`;
    PlacementDate String `json:"placementDate"`;
    IncrementDate String `json:"incrementDate"`;
    RetailDate String `json:"retailDate"`;
    DescriptionId int `json:"descriptionId"`;
    NumberOfComments int `json:"numberOfComments"`;
    NumberOfLikes int `json:"numberOfLikes"`;
    AllReactionsTotal int `json:"allReactionsTotal"`;
    CommentTableName String `json:"commentTableName"`;
    LikeTableName String `json:"likeTableName"`;
    ProductLocation String `json:"productLocation"`; 
    ProductImageId int `json:"productImageId"`;
    ProductImageName String `json:"productImageName"`; 
    ProductImageWidth int `json:"productImageWidth"`; 
    ProductImageHeight int `json:"productImageHeight"`; 
}

type FetchProductRequest struct {
    Success bool `json:"success"`;
    Message string `json:"message"`;
    Data []FetchProductData `json:"data"`;
}

type FetchProductResponse struct {
    EmailAddress string `json:"emailAddress" binding:"required"`;   
}


