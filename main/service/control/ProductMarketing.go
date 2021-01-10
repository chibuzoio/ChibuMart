package control

import (
	"fmt"; 
	"time";
	"./model";
	"math/big";
	"./utility";
)

func AddWishedProduct(addWishedProduct model.AddWishedProduct) bool {
    connector := utility.GetConnection();
    
    defer connector.Close();
    
    query := "insert into ";
}

func GetProductCartTable(chibuMartId int) string {
    connector := utility.GetConnection();
    
    defer connector.Close();
    
    var productCartTable string;
    query := "select productCartTable from usertable where chibuMartId = ?";
    
    resultSet, error := connector.Prepare();
    
    utility.Exception(error);
    
    rows, error := resultSet.Query(chibuMartId);
    
    utility.Exception(error);
    
    for rows.Next() {
        error = rows.Scan(&productCartTable);
        
        utility.Exception(error);
    }
    
    resultSet.Close();
    rows.Close();
    connector.Close();
    
    return productCartTable;
}

func StoreProductImage(imageProperties model.ImageProperties) {
    connector := utility.GetConnection();
    
    defer connector.Close();
    
    query := "insert into productimages (productImageId, productId, productImageName, " +  
        "commentTableName, likeTableName, numberOfComments, numberOfLikes) values (?, ?, ?, ?, ?, ?, ?)";
            
    stmt, error := connector.Prepare(query);
    
    utility.Exception(error);
    
    _, error = stmt.Exec(0, imageProperties.ContentId, imageProperties.ImageName, "", "", 0, 0);
    
    utility.Exception(error);
    
    stmt.Close();
    connector.Close();
}

func StoreProductComposite(addProductRequest model.AddProductRequest) int {
	timeNow := time.Now(); 
	connector := utility.GetConnection();
	currentUnixTime := fmt.Sprintf("%d", timeNow.Unix()); 
    
	defer connector.Close();

	var tenRat, hundredRat, productPreviousPriceRat, productCurrentPriceRat big.Rat;

	tenRat.SetFloat64(10);
	hundredRat.SetFloat64(100);
	productCurrentPriceRat.SetString(addProductRequest.ProductPrice);

	productPreviousPriceRat.Mul(&productCurrentPriceRat, &tenRat);
	productPreviousPriceRat.Quo(&productPreviousPriceRat, &hundredRat);
	productPreviousPriceRat.Add(&productPreviousPriceRat, &productCurrentPriceRat);

    // for big.Rat arithmetics, you do not use the conventional arithmetic signs 
    // (+, -, * or /) for calculations, but the Mul, Quo, Add and more arithmetic 
    // functions provided by the math/big library   
	/*productPreviousPrice := productCurrentPrice + ((productCurrentPrice * 10) / 100);*/

    productCurrentPrice := productCurrentPriceRat.FloatString(15);
	productPreviousPrice := productPreviousPriceRat.FloatString(15);

    query := "insert into productcollection (productId, productName, productCategory, " +
        "productQuantityRemaining, productQuantityRetailed, productQuantityTotal, productPreviousPrice, " + 
        "productCurrentPrice, placementDate, incrementDate, retailDate, descriptionId, numberOfComments, " + 
        "numberOfLikes, allReactionsTotal, commentTableName, likeTableName, productLocation) values " +  
        "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)";   

    transaction, error := connector.Begin();
    
    utility.Exception(error);
 
    stmt, error := transaction.Prepare(query);
    
    utility.Exception(error);
                      
    _, error = stmt.Exec(0, addProductRequest.ProductName, addProductRequest.ProductCategory, 
        addProductRequest.ProductQuantity, 0, addProductRequest.ProductQuantity, productPreviousPrice, 
        productCurrentPrice, currentUnixTime, currentUnixTime, "", 0, 0, 0, 0, "", "", "");
    
    utility.Exception(error);
    
    stmt.Close();

    var productId int;
    query = "select productId from productcollection order by productId desc limit 1";
    
    rows, error := transaction.Query(query);
    
    utility.Exception(error);
    
    for rows.Next() {
        error = rows.Scan(&productId);
        
        utility.Exception(error);
    }
    
    rows.Close();
    
    utility.Exception(transaction.Commit());
    
    connector.Close();
    
    return productId;
}


