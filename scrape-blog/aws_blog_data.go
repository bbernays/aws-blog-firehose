package main

import (
	"time"
)

type blogsData struct {
	Metadata struct {
		Count     int `json:"count"`
		TotalHits int `json:"totalHits"`
	} `json:"metadata"`
	FieldTypes struct {
		FeaturedImageURL string `json:"featuredImageUrl"`
		PostExcerpt      string `json:"postExcerpt"`
		CreatedDate      string `json:"createdDate"`
		DisplayDate      string `json:"displayDate"`
		ModifiedDate     string `json:"modifiedDate"`
		Link             string `json:"link"`
		Contributors     string `json:"contributors"`
		Title            string `json:"title"`
		ContentType      string `json:"contentType"`
		Slug             string `json:"slug"`
	} `json:"fieldTypes"`
	Items []struct {
		Tags []struct {
			TagNamespaceID string `json:"tagNamespaceId"`
			CreatedBy      string `json:"createdBy"`
			Name           string `json:"name"`
			DateUpdated    string `json:"dateUpdated"`
			Locale         string `json:"locale"`
			LastUpdatedBy  string `json:"lastUpdatedBy"`
			DateCreated    string `json:"dateCreated"`
			Description    string `json:"description"`
			ID             string `json:"id"`
		} `json:"tags"`
		Item struct {
			CreatedBy        string  `json:"createdBy"`
			Locale           string  `json:"locale"`
			Author           string  `json:"author"`
			DateUpdated      string  `json:"dateUpdated"`
			Score            float64 `json:"score"`
			Name             string  `json:"name"`
			NumImpressions   int     `json:"numImpressions"`
			DateCreated      string  `json:"dateCreated"`
			AdditionalFields struct {
				FeaturedImageURL string    `json:"featuredImageUrl"`
				PostExcerpt      string    `json:"postExcerpt"`
				CreatedDate      time.Time `json:"createdDate"`
				DisplayDate      string    `json:"displayDate"`
				Link             string    `json:"link"`
				ModifiedDate     time.Time `json:"modifiedDate"`
				Contributors     string    `json:"contributors"`
				Title            string    `json:"title"`
				ContentType      string    `json:"contentType"`
				Slug             string    `json:"slug"`
			} `json:"additionalFields"`
			ID            string `json:"id"`
			DirectoryID   string `json:"directoryId"`
			LastUpdatedBy string `json:"lastUpdatedBy"`
		} `json:"item"`
	} `json:"items"`
}
