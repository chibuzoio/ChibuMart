package main

import ( 
    "./service";     
	"github.com/gin-gonic/gin";   
)
      
func main() {
    gin.SetMode(gin.ReleaseMode);
	router := gin.Default();
	router.StaticFile("/", "chibu/index.html");
	router.GET("/install", service.Install);    
    router.GET("/fetchuserdata", service.FetchUserData);
	router.POST("/postuserdata", service.PostUserData);                           
    router.Run();                 
}                       


