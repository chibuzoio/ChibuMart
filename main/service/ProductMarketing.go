package service 

import (   
    "net/http";     
    
	"./control";
    "./control/model";   
	
	"github.com/gin-gonic/gin";    
    "github.com/gin-contrib/sessions";
)    
               
func PlaceProductOrder(context *gin.Context) {
    var placeOrderRequest model.PlaceOrderRequest;
    var placeOrderResponse model.PlaceOrderResponse;
    
    session := sessions.Default(context);
    
    sessionEmailAddress, _ := session.Get("emailAddress").(string);

    context.Bind(&placeOrderRequest);
    
    if placeOrderRequest.EmailAddress == sessionEmailAddress {
        committed := control.PlaceProductOrder(placeOrderRequest.EmailAddress);
        
        if committed {
            placeOrderResponse.Success = true;
            placeOrderResponse.DeliveryStatus = "PENDING";
            placeOrderResponse.Message = "Product order placement successful!";
        } else {
            placeOrderResponse.Success = false;
            placeOrderResponse.DeliveryStatus = "";
            placeOrderResponse.Message = "Product order placement failed!";
        }
    } else {
        placeOrderResponse.Success = false;
        placeOrderResponse.DeliveryStatus = "";
        placeOrderResponse.Message = "Product order placement failed!";
    }
    
    context.JSON(http.StatusOK, placeOrderResponse);
}

func AddCartProduct(context *gin.Context) {
    var cartProductRequest model.CartProductRequest;
    var cartProductResponse model.CartProductResponse;
    
    session := sessions.Default(context);
    
    sessionEmailAddress, _ := session.Get("emailAddress").(string);

    context.Bind(&cartProductRequest);
        
    if cartProductRequest.EmailAddress == sessionEmailAddress {  
        committed := control.AddCartProduct(cartProductRequest);
        
        if committed {
            cartProductResponse.Success = true;
            cartProductResponse.Message = "Product added successful!";
        } else {
            cartProductResponse.Success = false;
            cartProductResponse.Message = "Product addition failed!";
        }
    } else {
        cartProductResponse.Success = false;
        cartProductResponse.Message = "Product addition failed!";
    }
    
    context.JSON(http.StatusOK, cartProductResponse);
}

func AddWishedProduct(context *gin.Context) {
    var addWishedProduct model.AddWishedProduct;
    var wishProductRequest model.WishProductRequest;
    var wishProductResponse model.WishProductResponse;
    
    session := sessions.Default(context);
    
    sessionEmailAddress, _ := session.Get("emailAddress").(string);

    context.Bind(&wishProductRequest);
        
    if wishProductRequest.EmailAddress == sessionEmailAddress {
        chibuMartId := control.GetChibuMartId(wishProductRequest.EmailAddress);
        productWishTable := control.GetProductWishTable(chibuMartId);
                
        addWishedProduct.ProductWishTable = productWishTable;
        addWishedProduct.ProductId = wishProductRequest.ProductId;
                    
        committed := control.AddWishedProduct(addWishedProduct);
        
        if committed {
            wishProductResponse.Success = true;
            wishProductResponse.Message = "Product wish successful!";
        } else {
            wishProductResponse.Success = false;
            wishProductResponse.Message = "Product wish failed!";   
        }
    } else {
        wishProductResponse.Success = false;
        wishProductResponse.Message = "User not signed in!";
    }
    
    context.JSON(http.StatusOK, wishProductResponse);
}

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
            imageProperties.ContentId = productId;
            imageProperties.ImageName = productImageName;
            imageProperties.ImageWidth = addProductRequest.ImageWidth;
            imageProperties.ImageHeight = addProductRequest.ImageHeight;
           
            control.StoreProductImage(imageProperties);
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


