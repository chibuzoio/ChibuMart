package service

import (  
    "./control";      
    "github.com/jasonlvhit/gocron"; 
)

func checkGenericCartTable() { 
    control.CheckGenericCartTable();
}

func ExecuteCronJob() {
    gocron.Every(1).Minute().Do(checkGenericCartTable);
    
    <- gocron.Start();
}

/*
gocron.Every(1).Second().Do(myTask) // Every 1 Second
gocron.Every(1).Minute().Do(myTask) // Every Minute
gocron.Every(2).Hours().Do(myTask) // Every 2 Hours
gocron.Every(3).Days().Do(myTask) // Every 3 Days
gocron.Every(4).Weeks().Do(myTask) // Every 4 Weeks
*/


