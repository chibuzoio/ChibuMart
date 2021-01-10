package model  

import (
    
)

type FetchProductData struct { 
    ProductId int `json:"productId"`;
    ProductName string `json:"productName"`;
    ProductCategory string `json:"productCategory"`;
    ProductQuantityRemaining int `json:"productQuantityRemaining"`;
    ProductQuantityRetailed int `json:"productQuantityRetailed"`;
    ProductQuantityTotal int `json:"productQuantityTotal"`;
    ProductPreviousPrice string `json:"productPreviousPrice"`;
    ProductCurrentPrice string `json:"productCurrentPrice"`;
    PlacementDate string `json:"placementDate"`;
    IncrementDate string `json:"incrementDate"`;
    RetailDate string `json:"retailDate"`;
    DescriptionId int `json:"descriptionId"`;
    NumberOfComments int `json:"numberOfComments"`;
    NumberOfLikes int `json:"numberOfLikes"`;
    AllReactionsTotal int `json:"allReactionsTotal"`;
    CommentTableName string `json:"commentTableName"`;
    LikeTableName string `json:"likeTableName"`;
    ProductLocation string `json:"productLocation"`; 
    ProductImageId int `json:"productImageId"`;
    ProductImageName string `json:"productImageName"`; 
    ProductImageWidth int `json:"productImageWidth"`; 
    ProductImageHeight int `json:"productImageHeight"`; 
}

type FetchProductResponse struct {
    Success bool `json:"success"`;
    Message string `json:"message"`;
    Data []FetchProductData `json:"data"`;
}
   

