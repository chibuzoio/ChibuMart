package service 

import (
    "strings";
    "net/http";     
    
	"./control";
    "./control/model";
    "./control/utility";
	
	"github.com/gin-gonic/gin";    
    "github.com/gin-contrib/sessions";
)    


IncrementProduct
DecrementProduct
RemoveProduct

func AddNewProduct(context *gin.Context) {
    var addProductRequest model.AddProductRequest;
    var addProductResponse model.AddProductResponse;
    
    session := sessions.Default(context);
    
    sessionEmailAddress, _ := session.Get("emailAddress").(string);

    context.Bind(&addProductRequest);
    
    if addProductRequest.EmailAddress == sessionEmailAddress {
        StoreProductComposite
        StoreGenericImage
        StoreImageProperties
        
        addProductResponse.Success = false;
        addProductResponse.Message = "Product addition successful!";
    } else {
        addProductResponse.Success = false;
        addProductResponse.Message = "Product addition failed!";
    }
    
    context.JSON(http.StatusOK, addProductResponse);
}


