package utility

import (
    "os";
    "fmt";
    "log";
    "time";
    "image";
    "strings";
    "image/png";
    "math/rand";
    "image/jpeg";
    "database/sql";
    "encoding/base64";
    _ "github.com/go-sql-driver/mysql";
)

type TableCompositeParts struct {     
	tableColumnName, generatedTableName string;      
	memberId int;     
}    

func GenerateSessionToken(email string) string {
	return email + fmt.Sprintf("%d", time.Now().Unix());
}

func GenerateInterfaceString(collectionArray []interface{}) string {
	return fmt.Sprintf("%v", collectionArray);     
}        
          
func SliceIntegerArray(composite []int, key int) []int {       
    return append(composite[ : key], composite[key + 1 : ]...);            
}        
        
func CaseInsensitiveSearch(searchComposite, searchKey string) bool {    
	return strings.Contains(strings.ToLower(searchComposite), strings.ToLower(searchKey));        
}     

func CompareInterfaceArray(firstArray, secondArray []interface{}) bool {   
    // If one is nil, the other must also be nil.
    if (firstArray == nil) != (secondArray == nil) { 
        return false; 
    }

    if len(firstArray) != len(secondArray) {
        return false;
    }

    for index := range firstArray {
        if firstArray[index] != secondArray[index] {
            return false;
        }
    }

    return true;
}
 
func GetRandomNumber(maxNumber int) int {
    randSource := rand.NewSource(time.Now().UnixNano());
    randObject := rand.New(randSource);
    return randObject.Intn(maxNumber);
}
  
func getValueKey(composite []string, searchKey string) int {
    var gottenKey int;

    for key, value := range composite {
        if searchKey == value {
            gottenKey = key;
            break;
        }
    }

    return gottenKey;
}

func ArraySplice(stringArray []string, element string) []string {
    var returnStringArray []string;
    searchKey := getValueKey(stringArray, element);

    for key, value := range stringArray {
        if searchKey != key {
            returnStringArray = append(returnStringArray, value);
        }
    }

    return returnStringArray;
}

func StoreBase64Png(imageString, imageName string) {
    reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageString));

    imageFile, formatString, error := image.Decode(reader);

    Exception(error);

    bounds := imageFile.Bounds();

    log.Println(bounds, formatString);

    storage, error := os.Create("chibumart/image/" + imageName);

    Exception(error);

    error = png.Encode(storage, imageFile);

    Exception(error);

    log.Println("Png file", imageName, "created...");
}

func StoreBase64Jpeg(imageString, imageName string) {
    reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageString));

    imageFile, formatString, error := image.Decode(reader);

    Exception(error);

    bounds := imageFile.Bounds();

    log.Println(bounds, formatString);

    storage, error := os.Create("chibumart/image/" + imageName);

    Exception(error);

    error = jpeg.Encode(storage, imageFile, &jpeg.Options{Quality: 100});

    Exception(error);

    log.Println("Jpeg file", imageName, "created...");
}

func StoreAverageImage(imageString, imageName string) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageString));

    imageFile, formatString, error := image.Decode(reader);

    Exception(error);

    bounds := imageFile.Bounds();

    log.Println(bounds, formatString);

	imageName = strings.Replace(imageName, ".jpg", "average.jpg", -1);

    storage, error := os.Create("chibumart/image/" + imageName);

    Exception(error);     

    error = jpeg.Encode(storage, imageFile, &jpeg.Options{Quality: 55});
    
    Exception(error); 
} 

func StoreThumbnailImage(imageString, imageName string) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageString));

    imageFile, formatString, error := image.Decode(reader);

    Exception(error);

    bounds := imageFile.Bounds();

    log.Println(bounds, formatString);

	imageName = strings.Replace(imageName, ".jpg", "thumbnail.jpg", -1);

    storage, error := os.Create("chibumart/image/" + imageName);

    Exception(error);     

    error = jpeg.Encode(storage, imageFile, &jpeg.Options{Quality: 25});
    
    Exception(error); 
} 

func GenerateMediumImage(mediumImageName string) {
	_, error := os.Stat("/chibumart/image/" + mediumImageName);
	
	if error != nil {
		imageName := strings.Replace(mediumImageName, "average.jpg", ".jpg", -1);
		
		imageFile, error := os.Open("chibumart/image/" + imageName);
		
		Exception(error);
		
		defer imageFile.Close();  
		
		imageString, formatString, error := image.Decode(imageFile);
		
		Exception(error);
		
		bounds := imageString.Bounds();
	
	    log.Println(bounds, formatString);
	
		storage, error := os.Create("chibumart/image/" + mediumImageName);
	
	    Exception(error);     
	
	    error = jpeg.Encode(storage, imageString, &jpeg.Options{Quality: 55});
	    
	    Exception(error); 
	} 
} 

func CompressExistingImage(thumbnailName string) {
	_, error := os.Stat("/chibumart/image/" + thumbnailName);
	
	if error != nil {
		imageName := strings.Replace(thumbnailName, "thumbnail.jpg", ".jpg", -1);
		
		imageFile, error := os.Open("chibumart/image/" + imageName);
		
		Exception(error);
		
		defer imageFile.Close();  
		
		imageString, formatString, error := image.Decode(imageFile);
		
		Exception(error);
		
		bounds := imageString.Bounds();
	
	    log.Println(bounds, formatString);
	
		storage, error := os.Create("chibumart/image/" + thumbnailName);
	
	    Exception(error);     
	
	    error = jpeg.Encode(storage, imageString, &jpeg.Options{Quality: 25});
	    
	    Exception(error); 
	} 
} 

func CheckTableError(exception error) bool {
    returnValue := false;

    if exception != nil {
        Exception(exception);

        localException := fmt.Sprintf("%s", exception);

        if (strings.Contains(strings.ToLower(localException), strings.ToLower("changeIndexer")) || 
            strings.Contains(strings.ToLower(localException), strings.ToLower("gotNewMessage"))) && 
            strings.Contains(strings.ToLower(localException), strings.ToLower("Unknown column")) {
            returnValue = true;
        }
    }

    return returnValue;
}

func EvaluateTable(databaseTable string, exception error) bool {
    returnValue := false;

    if exception != nil {
        Exception(exception);

        localException := fmt.Sprintf("%s", exception);

        if strings.Contains(localException, databaseTable) && 
            strings.Contains(localException, "doesn't exist") {
            returnValue = true;
        }
    }

    return returnValue;
}
          
func ErrorStackTrace(exception error, errorLocation string) {
    if exception != nil {
        Println("The Error Below Is From Line " + errorLocation);
        Exception(exception);        
    }
}

func Println(statement string) {
    _, error := os.Stat("chibumart/chibumart.txt");

    if error != nil {   
        logFile, error := os.OpenFile("chibumart/chibumart.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644);

        log.Println("File doesn't exist; " + fmt.Sprintf("%v", error));

        log.SetOutput(logFile);
        log.Println(statement);

        logFile.Close();
    } else {
        logFile, error := os.OpenFile("chibumart/chibumart.txt", os.O_WRONLY|os.O_APPEND, 0644);

        log.Println("File does exist; " + fmt.Sprintf("%v", error));

        log.SetOutput(logFile);
        log.Println(statement);

        logFile.Close();
    }
}

func Exception(exception error) {
    _, error := os.Stat("chibumart/chibumart.txt");
 
    if error != nil {   
        logFile, error := os.OpenFile("chibumart/chibumart.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644);

        log.Println("File doesn't exist; " + fmt.Sprintf("%v", error));

        log.SetOutput(logFile);

        if exception != nil {
            log.Println(exception);
        }

        logFile.Close();
    } else {
        logFile, error := os.OpenFile("chibumart/chibumart.txt", os.O_WRONLY|os.O_APPEND, 0644);

        log.Println("File does exist; " + fmt.Sprintf("%v", error));

        log.SetOutput(logFile);

        if exception != nil {
            log.Println(exception);
        }

        logFile.Close();
    }
}

func SearchIntegerArray(composite []int, searchKey int) bool {
    for _, value := range composite {
        if searchKey == value {
            return true;
        }
    }

    return false;
}

func SearchStringArray(composite []string, searchKey string) bool {
    for _, value := range composite {
        if searchKey == value {
            return true;
        }
    }

    return false;
}

func GetConnection() *sql.DB { 
	connector, error := sql.Open("mysql", "computebone:2352C/C++solu+++;@/chibumart"); 

/*    
    connector.SetMaxOpenConns(777);
    connector.SetMaxIdleConns(222);  
*/    

    Exception(error);

	log.Println("Connection to chibumart database acquired!!!!!!");
	
	return connector;
}


