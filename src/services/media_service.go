package services

import (
	"fmt"
	"log"
	"time"

	"cloud/src/config"
	"cloud/src/models"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

// SaveMedia stores media file metadata in Cassandra
func SaveMedia(userID, fileName, fileType, fileURL string) (*models.Media, error) {
	// Ensure DB session is initialized
	if config.Session == nil {
		log.Println("Cassandra session is not initialized")
		return nil, fmt.Errorf("database connection error")
	}

	// Create a new media record
	media := &models.Media{
		ID:         gocql.UUID(uuid.New()),
		UserID:     gocql.UUID(uuid.New()),
		FileName:   fileName,
		FileType:   fileType,
		FileURL:    fileURL,
		UploadedAt: time.Now(),
	}

	// Insert query
	query := `INSERT INTO media (id, user_id, file_name, file_type, file_url, uploaded_at) 
	          VALUES (?, ?, ?, ?, ?, ?)`
	err := config.Session.Query(query, media.ID, media.UserID, media.FileName, media.FileType, media.FileURL, media.UploadedAt).Exec()
	if err != nil {
		log.Println("Error inserting media:", err)
		return nil, err
	}

	fmt.Println("Media saved:", media.FileURL)
	return media, nil
}

func GetAllMedia() ([]models.Media, error) {
	if config.Session == nil {
		log.Println("Cassandra session is not initialized")
		return nil, fmt.Errorf("database connection error")
	}

	var mediaList []models.Media
	query := "SELECT id, user_id, file_name, file_type, file_url, uploaded_at FROM media"

	iter := config.Session.Query(query).Iter()
	defer iter.Close()

	var media models.Media
	for iter.Scan(&media.ID, &media.UserID, &media.FileName, &media.FileType, &media.FileURL, &media.UploadedAt) {
		mediaList = append(mediaList, media)
	}

	if err := iter.Close(); err != nil {
		log.Println("Error fetching media:", err)
		return nil, err
	}

	return mediaList, nil
}
