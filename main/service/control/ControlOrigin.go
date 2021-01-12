package control

import ( 
	"fmt";
	"time"; 
    "strconv";
	"./model";     
	"./utility";    
)

func CheckGenericCartTable() {
	connector := utility.GetConnection(); 

	defer connector.Close();

    var dueCartTableIdArray []int;
    
    currentUnixTime := time.Now().Unix();
    
    query := "select cartTableId, timeout from chibumartcart";
    
    rows, error := connector.Query(query);
    
    utility.Exception(error);
    
    for rows.Next() {
        var timeout string;   
        var cartTableId int;
        
        error = rows.Scan(&cartTableId, &timeout);
        
        utility.Exception(error);
        
        integerTimeout, error := strconv.ParseInt(timeout, 10, 64);
        
        utility.Exception(error);
        
        timeDifference := currentUnixTime - integerTimeout;
        
        if timeDifference > 0 {
            dueCartTableIdArray = append(dueCartTableIdArray, cartTableId);
        }
    }
    
    rows.Close();
           
    query = "delete from chibumartcart where cartTableId = ?";
    
    stmt, error := connector.Prepare(query);
    
    utility.Exception(error);
    
    for _, value := range dueCartTableIdArray {
        _, error = stmt.Exec(dueCartTableIdArray[value]);
        
        utility.Exception(error);
    }
    
    stmt.Close();
    connector.Close();
}

func GetChibuMartId(emailAddress string) int {
	connector := utility.GetConnection(); 

	defer connector.Close();

    var chibuMartId int;
    query := "select chibuMartId from chibumart where emailAddress = ?";
    
    resultSet, error := connector.Prepare(query);
    
    utility.Exception(error);
    
    rows, error := resultSet.Query(emailAddress);
    
    utility.Exception(error);
    
    for rows.Next() {
        error = rows.Scan(&chibuMartId);
        
        utility.Exception(error);
    }
    
    resultSet.Close();
    rows.Close();
    connector.Close();
    
    return chibuMartId;
}

func StoreImageProperties(imageProperties model.ImageProperties) bool {
    var committed bool;
	connector := utility.GetConnection(); 

	defer connector.Close();

	query := "insert into chibumartimages (image, width, height) values (?, ?, ?)";

	if imageProperties.ImageName != "" || imageProperties.ImageHeight > 0 || imageProperties.ImageWidth > 0 {  
		stmt, error := connector.Prepare(query);

		utility.Exception(error);

		_, error = stmt.Exec(imageProperties.ImageName, imageProperties.ImageWidth, imageProperties.ImageHeight);

		utility.Exception(error);

		stmt.Close();

        committed = true;
	} else {
        committed = false;
    }

	connector.Close();

    return committed;
}

func StoreGenericImage(genericImage model.GenericImage) string {
	timeNow := time.Now(); 
	contentIdPart := fmt.Sprintf("%05d", genericImage.ContentId);
	currentUnixTime := fmt.Sprintf("%d", timeNow.Unix()); 
	contentIdPart = contentIdPart[len(contentIdPart) - 5 : len(contentIdPart)];     
    gottenImageName := genericImage.ImageType + currentUnixTime + contentIdPart + ".jpg";

	utility.StoreBase64Jpeg(genericImage.ImageString, gottenImageName);
	utility.StoreAverageImage(genericImage.ImageString, gottenImageName);  
	utility.StoreThumbnailImage(genericImage.ImageString, gottenImageName);

	return gottenImageName;
}


