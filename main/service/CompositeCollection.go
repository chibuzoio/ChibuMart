package service

import ( 
    "fmt";
    "net/http";     
    
	"./control";
    "./control/model";
    "./control/utility";
	
	"github.com/gin-gonic/gin";    
    "github.com/gin-contrib/sessions";
)

func FetchProducts(context *gin.Context) {
    var fetchProductRequest model.FetchProductRequest;
    var fetchProductResponse model.FetchProductResponse;
  
    session := sessions.Default(context);
    
    sessionEmailAddress, _ := session.Get("emailAddress").(string);

    context.Bind(&fetchProductRequest);

    if fetchProductRequest.EmailAddress == sessionEmailAddress {
        fetchProductResponse.Data = control.FetchProducts();
        fetchProductResponse.Message = "Products fetched successfully!";
        fetchProductResponse.Success = true;
    } else {
        fetchProductResponse.Success = false;
        fetchProductResponse.Message = "User not signed in!";
    }
    
    utility.Println("Response gotten here is this " + fmt.Sprintf("%v", fetchProductResponse));
    
    context.JSON(http.StatusOK, fetchProductResponse);
}


