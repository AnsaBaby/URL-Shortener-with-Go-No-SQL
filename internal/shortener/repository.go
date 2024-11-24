package shortener

import (
	"errors"
	"fmt"

	"github.com/gocql/gocql"
)

type URLRepository struct {
	session *gocql.Session
}

func NewCassandraSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1") // Cassandra address
	cluster.Keyspace = "urlshortener"
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func NewURLRepository(session *gocql.Session) *URLRepository {
	return &URLRepository{session: session}
}

func (r *URLRepository) Save(originalURL, shortURL string) error {
	err := r.session.Query(`INSERT INTO urls (short_url, original_url) VALUES (?, ?)`,
		shortURL, originalURL).Exec()
	if err != nil {
		return fmt.Errorf("failed to save URL: %v", err)
	}
	return nil
}

func (r *URLRepository) Get(shortURL string) (string, error) {
	var originalURL string
	err := r.session.Query(`SELECT original_url FROM urls WHERE short_url = ?`, shortURL).Scan(&originalURL)
	if err != nil {
		if err == gocql.ErrNotFound {
			return "", errors.New("URL not found")
		}
		return "", fmt.Errorf("failed to retrieve URL: %v", err)
	}
	return originalURL, nil
}
