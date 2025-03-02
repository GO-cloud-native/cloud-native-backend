package services

import (
	"cloud/src/config"
	"cloud/src/models"
	"log"

	"github.com/gocql/gocql"
)

// GetMediaByID retrieves a media record by ID
func GetMediaByID(id string) (*models.Media, error) {
	var media models.Media
	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		log.Println("Invalid UUID format:", err)
		return nil, err
	}

	query := "SELECT id, user_id, file_name, file_type, file_url, uploaded_at FROM media WHERE id = ? LIMIT 1"
	err = config.DBSession.Query(query, uuid).Scan(&media.ID, &media.UserID, &media.FileName, &media.FileType, &media.FileURL, &media.UploadedAt)
	if err != nil {
		log.Println("Error fetching media by ID:", err)
		return nil, err
	}

	return &media, nil
}
