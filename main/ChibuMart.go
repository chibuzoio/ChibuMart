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
    router.GET("/fetchuserdata", service.FetchUserData);
	router.POST("/postuserdata", service.PostUserData);                           
    router.Run();                 
}                       


