package control

import (
	"fmt"; 
	"time";
	"./model";
	"math/big";
	"./utility";
)

func PlaceProductOrder(emailAddress string) bool {
    var cartTableIdArray []int;
    var placeProductOrderArray []model.PlaceProductOrder;
    
    connector := utility.GetConnection();
	currentUnixTime := fmt.Sprintf("%d", time.Now().Unix()); 
     
    defer connector.Close();
                
    chibuMartId := GetChibuMartId(emailAddress);

    query := "select cartTableId, productId, productQuantity " +  
        "from chibumartcart where chibuMartId = ?";
    
    resultSet, error := connector.Prepare(query);
    
    utility.Exception(error);
    
    rows, error := resultSet.Query(chibuMartId);
    
    utility.Exception(error);
    
    for rows.Next() {
        var placeProductOrder model.PlaceProductOrder;
        
        error = rows.Scan(&placeProductOrder.CartTableId, 
                          &placeProductOrder.ProductId, 
                          &placeProductOrder.ProductQuantity);
        
        utility.Exception(error);
        
        cartTableIdArray = append(cartTableIdArray, placeProductOrder.CartTableId);
        placeProductOrderArray = append(placeProductOrderArray, placeProductOrder);
    }

    resultSet.Close();
    rows.Close();
            
    if len(cartTableIdArray) > 0 {
        query = "insert into productdeliverytable (productDeliveryId, productId, " + 
            "chibuMartId, productQuantity, deliveryStatus, deliveryDate) values (?, ?, ?, ?, ?, ?)";
        
        stmt, error := connector.Prepare(query);
        
        utility.Exception(error);
        
        // PENDING, DELIVERING, DELIVERED AND CANCELED     
        for _, value := range placeProductOrderArray {
            _, error = stmt.Exec(0, value.ProductId, chibuMartId, 
                                 value.ProductQuantity, "PENDING", currentUnixTime);
        
            utility.Exception(error);
        }
        
        stmt.Close();
        
        utility.DeleteCartProductArray(cartTableIdArray);
    }
       
    connector.Close();
    
    return true;
}

func AddCartProduct(cartProductRequest model.CartProductRequest) bool {
    connector := utility.GetConnection();
     
    defer connector.Close();
                
    currentUnixTime := time.Now().Unix();
    timeout := currentUnixTime + 1800;
    chibuMartId := GetChibuMartId(cartProductRequest.EmailAddress);
    
    query := "insert into chibumartcart (cartTableId, chibuMartId, " + 
        "productId, productQuantity, timeout) values (?, ?, ?, ?, ?)";
    
    stmt, error := connector.Prepare(query);
    
    utility.Exception(error);
    
    _, error = stmt.Exec(0, chibuMartId, cartProductRequest.ProductId, 
                         cartProductRequest.ProductQuantity, timeout);
    
    utility.Exception(error);
    
    stmt.Close();
    connector.Close();
    
    return true;
}

func AddWishedProduct(addWishedProduct model.AddWishedProduct) bool {
	timeNow := time.Now(); 
	connector := utility.GetConnection();
	currentUnixTime := fmt.Sprintf("%d", timeNow.Unix()); 
    
    defer connector.Close();
 
    query := "insert into " + addWishedProduct.ProductWishTable + 
        " (productWishId, productId, wishDate) values (?, ?, ?)";
    
    stmt, error := connector.Prepare(query);
    
	if utility.EvaluateTable(addWishedProduct.ProductWishTable, error) {
		utility.CreateProductWishTable(connector, addWishedProduct.ProductWishTable);
		
		stmt, error = connector.Prepare(query);

		utility.Exception(error);
	}

    _, error = stmt.Exec(0, addWishedProduct.ProductId, currentUnixTime);
    
    utility.Exception(error);
    
    stmt.Close();
    connector.Close();
    
    return true;
}

func GetProductWishTable(chibuMartId int) string {
    connector := utility.GetConnection();
    
    defer connector.Close();
    
    var productWishTable string;
    query := "select productWishTable from usertable where chibuMartId = ?";
    
    resultSet, error := connector.Prepare(query);
    
    utility.Exception(error);
    
    rows, error := resultSet.Query(chibuMartId);
    
    utility.Exception(error);
    
    for rows.Next() {
        error = rows.Scan(&productWishTable);
        
        utility.Exception(error);
    }
    
    resultSet.Close();
    rows.Close();
    connector.Close();
    
    return productWishTable;
}

func GetProductCartTable(chibuMartId int) string {
    connector := utility.GetConnection();
    
    defer connector.Close();
    
    var productCartTable string;
    query := "select productCartTable from usertable where chibuMartId = ?";
    
    resultSet, error := connector.Prepare(query);
    
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


