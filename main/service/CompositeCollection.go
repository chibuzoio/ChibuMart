package service

import (  
    "net/http";     
    
	"./control";
    "./control/model";
    "./control/utility";
	
	"github.com/gin-gonic/gin";     
)

func FetchProducts(context *gin.Context) { 
    var fetchProductResponse model.FetchProductResponse;
 
    context.Bind(&fetchProductRequest);
 
    fetchProductResponse.Data = control.FetchProducts();
    fetchProductResponse.Message = "Products fetched successfully!";
    fetchProductResponse.Success = true;
     
    context.JSON(http.StatusOK, fetchProductResponse);
}


