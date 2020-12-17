package service   

import ( 
    "net/http";     
    
	"./control";      
	
	"github.com/gin-gonic/gin";   
)
           
func Install(context *gin.Context) { 
	control.CreateProjectDatabase();
    control.CreateProjectTables();   
    context.JSON(http.StatusOK, gin.H{"message" : "Database and tables created successfully!"});
}
    
    
    
