package model   

import (
    
)

type ImageProperties struct {            
    ImageName string `json:"imageName"`;
    ImageWidth int `json:"imageWidth"`;
    ImageHeight int `json:"imageHeight"`;
}

type GenericImage struct {               
    ContentId int `json:"contentId"`;
    ImageString string `json:"ImageString"`;
    ImageType string `json:"ImageType"`;
}


