package control

import ( 
	"./model"; 
	"./utility";
)

func FetchProducts() []model.FetchProductData {
    connector := utility.GetConnection();
    
    var fetchProductData model.FetchProductData;
    var fetchProductDataArray []model.FetchProductData;
                       
    defer connector.Close();
    
    query := "select productId, productName, productCategory, productQuantityRemaining, " + 
        "productQuantityRetailed, productQuantityTotal, productPreviousPrice, productCurrentPrice, " + 
        "placementDate, incrementDate, retailDate, descriptionId, numberOfComments, numberOfLikes, " + 
        "allReactionsTotal, commentTableName, likeTableName, productLocation from productcollection";
        
    rows, error := connector.Query(query);
    
    utility.Exception(error);
    
    for rows.Next() {
        error = rows.Scan(&fetchProductData.ProductId, &fetchProductData.ProductName, 
            &fetchProductData.ProductCategory, &fetchProductData.ProductQuantityRemaining, 
            &fetchProductData.ProductQuantityRetailed, &fetchProductData.ProductQuantityTotal, 
            &fetchProductData.ProductPreviousPrice, &fetchProductData.ProductCurrentPrice, &fetchProductData.PlacementDate, 
            &fetchProductData.IncrementDate, &fetchProductData.RetailDate, &fetchProductData.DescriptionId, 
            &fetchProductData.NumberOfComments, &fetchProductData.NumberOfLikes, &fetchProductData.AllReactionsTotal, 
            &fetchProductData.CommentTableName, &fetchProductData.LikeTableName, &fetchProductData.ProductLocation);
        
        utility.Exception(error);
              
        query = "select productImageId, productImageName from productimages where productId = ?";
        
        firstResultSet, error := connector.Prepare(query);
        
        utility.Exception(error);
        
        firstRows, error := firstResultSet.Query(fetchProductData.ProductId);
        
        utility.Exception(error);
        
        for firstRows.Next() {
            error = firstRows.Scan(&fetchProductData.ProductImageId, &fetchProductData.ProductImageName);
            
            utility.Exception(error);
  
            query = "select width, height from chibumartimages where image = ?";
            
            secondResultSet, error := connector.Prepare(query);
            
            utility.Exception(error);
            
            secondRows, error := secondResultSet.Query(fetchProductData.ProductImageName);
            
            utility.Exception(error);
            
            for secondRows.Next() {
                error = secondRows.Scan(&fetchProductData.ProductImageWidth, &fetchProductData.ProductImageHeight);
                
                utility.Exception(error);                       
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


