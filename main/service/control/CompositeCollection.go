package control

import ( 
	"./model"; 
	"./utility";
)

func FetchProducts() []model.FetchProductData {
    connector := utility.GetConnection();
    
    var fetchProductData model.FetchProductData;
    var fetchProductDataArray []model.FetchProductData;
    
    var productImageWidth, productImageHeight int;
    var descriptionId, numberOfComments, numberOfLikes, allReactionsTotal, productImageId int;
    var productId, productQuantityRemaining, productQuantityRetailed, productQuantityTotal int; 
    var productName, productCategory, productPreviousPrice, productCurrentPrice, placementDate string;     
    var incrementDate, retailDate, commentTableName, likeTableName, productLocation, productImageName string;
    
    defer connector.Close();
    
    query := "select productId, productName, productCategory, productQuantityRemaining, " + 
        "productQuantityRetailed, productQuantityTotal, productPreviousPrice, productCurrentPrice, " + 
        "placementDate, incrementDate, retailDate, descriptionId, numberOfComments, numberOfLikes, " + 
        "allReactionsTotal, commentTableName, likeTableName, productLocation from productcollection";
        
    rows, error := connector.Query(query);
    
    utility.Exception(error);
    
    for rows.Next() {
        error = rows.Scan(&productId, &productName, &productCategory, &productQuantityRemaining, 
            &productQuantityRetailed, &productQuantityTotal, &productPreviousPrice, &productCurrentPrice, 
            &placementDate, &incrementDate, &retailDate, &descriptionId, &numberOfComments, &numberOfLikes,  
            &allReactionsTotal, &commentTableName, &likeTableName, &productLocation);
        
        utility.Exception(error);
        
        fetchProductData.ProductId = productId;
        fetchProductData.ProductName = productName;
        fetchProductData.ProductCategory = productCategory;
        fetchProductData.ProductQuantityRemaining = productQuantityRemaining;
        fetchProductData.ProductQuantityRetailed = productQuantityRetailed;
        fetchProductData.ProductQuantityTotal = productQuantityTotal;
        fetchProductData.ProductPreviousPrice = productPreviousPrice;
        fetchProductData.ProductCurrentPrice = productCurrentPrice;
        fetchProductData.PlacementDate = placementDate;
        fetchProductData.IncrementDate = incrementDate;
        fetchProductData.RetailDate = retailDate;
        fetchProductData.DescriptionId = descriptionId;
        fetchProductData.NumberOfComments = numberOfComments;
        fetchProductData.NumberOfLikes = numberOfLikes;
        fetchProductData.AllReactionsTotal = allReactionsTotal;
        fetchProductData.CommentTableName = commentTableName;
        fetchProductData.LikeTableName = likeTableName;
        fetchProductData.ProductLocation = productLocation;    
             
        query = "select productImageId, productImageName from productimages where productId = ?";
        
        firstResultSet, error := connector.Prepare(query);
        
        utility.Exception(error);
        
        firstRows, error := firstResultSet.Query(productId);
        
        utility.Exception(error);
        
        for firstRows.Next() {
            error = firstRows.Scan(&productImageId, &productImageName);
            
            utility.Exception(error);

            fetchProductData.ProductImageId = productImageId;
            fetchProductData.ProductImageName = productImageName;
  
            query = "select width, height from chibumartimages where image = ?";
            
            secondResultSet, error := connector.Prepare(query);
            
            utility.Exception(error);
            
            secondRows, error := secondResultSet.Query(productImageName);
            
            utility.Exception(error);
            
            for secondRows.Next() {
                error = secondRows.Scan(&productImageWidth, &productImageHeight);
                
                utility.Exception(error);
                
                fetchProductData.ProductImageWidth = productImageWidth;
                fetchProductData.ProductImageHeight = productImageHeight;
            }
            
            secondResultSet.Close();
            secondRows.Close();
            
            break;
        }
        
        firstResultSet.Close();
        firstRows.Close();
        
        fetchProductDataArray = append(fetchProductDataArray, fetchProductData);
    }
    
    rows.Close();
    connector.Close();
    
    return fetchProductDataArray;
}


