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
    
	router.StaticFile("/", "chibu/index.html");
	router.GET("/install", service.Install);    
    router.POST("/login", service.LoginUser);
	router.GET("/logout", service.LogoutUser);    
    router.POST("/register", service.RegisterUser);
	router.POST("/postuserdata", service.PostUserData);                           
    router.GET("/fetchuserdata", service.FetchUserData);
    router.Run();                 
}                       


