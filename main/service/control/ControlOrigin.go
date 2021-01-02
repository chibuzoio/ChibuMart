package control

import (
	"os";
	"fmt";
	"time";
	"strings";
	"./model";
	"math/big";
	"./utility";
)

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


