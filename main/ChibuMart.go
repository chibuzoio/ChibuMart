package main

import ( 
    "./service";     
    
    "github.com/google/uuid";
	"github.com/gin-gonic/gin";   
	"github.com/gin-contrib/sessions";
	"github.com/gin-contrib/sessions/cookie";
)
      
func main() {
    gin.SetMode(gin.ReleaseMode);
    
	router := gin.Default();
	store := cookie.NewStore([]byte(uuid.New().String()));
	
    router.Use(sessions.Sessions("ChibuMart", store));
    
	router.Static("/js", "./chibumart/js");
	router.GET("/install", service.Install);    
	router.Static("/css", "./chibumart/css");
    router.POST("/login", service.LoginUser);
	router.GET("/logout", service.LogoutUser);    
	router.Static("/image", "./chibumart/image");
    router.POST("/register", service.RegisterUser);
	router.StaticFile("/", "./chibumart/index.html");
	router.POST("/postuserdata", service.PostUserData);                           
    router.GET("/fetchproducts", service.FetchProducts);
    router.GET("/fetchuserdata", service.FetchUserData);
    router.POST("/addnewproduct", service.AddNewProduct);
    router.POST("/wishproduct", service.AddWishedProduct);
    router.Run();                 
}                       


