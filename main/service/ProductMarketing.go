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
        productId := control.StoreProductComposite(addProductRequest);
        
        if productId > 0 {
            var genericImage model.GenericImage;
            genericImage.ContentId = productId;  
            genericImage.ImageString = addProductRequest.ProductImage;
            genericImage.ImageType = "PRODUCT";
            
            productImageName := control.StoreGenericImage(genericImage);
            
            var imageProperties model.ImageProperties;
            imageProperties.ImageName = productImageName;
            imageProperties.ImageWidth = addProductRequest.ImageWidth;
            imageProperties.ImageHeight = addProductRequest.ImageHeight;
            
            control.StoreImageProperties(imageProperties);
            
            addProductResponse.Success = true;
            addProductResponse.Message = "Product addition successful!";
        } else {
            addProductResponse.Success = false;
            addProductResponse.Message = "Product addition failed!";
        }
    } else {
        addProductResponse.Success = false;
        addProductResponse.Message = "User not signed in!";
    }
    
    context.JSON(http.StatusOK, addProductResponse);
}


