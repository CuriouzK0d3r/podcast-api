package models

import (
	"encoding/json"
	"fmt"
	"podcast-api/database"
)

type Podcast struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Host   string `json:"host"`
	Link   string `json:"link"`
}

func (p *Podcast) Save() error {
	db := database.GetDB()
	query := `INSERT INTO podcasts (title, host, link) VALUES (?, ?, ?)`
	result, err := db.Exec(query, p.Title, p.Host, p.Link)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	p.ID = int(id)
	return nil
}

func FetchAllPodcasts() ([]Podcast, error) {
	db := database.GetDB()
	query := `SELECT id, title, host, link FROM podcasts`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var podcasts []Podcast
	for rows.Next() {
		var p Podcast
		if err := rows.Scan(&p.ID, &p.Title, &p.Host, &p.Link); err != nil {
			return nil, err
		}
		podcasts = append(podcasts, p)
	}
	return podcasts, nil
}

func (p *Podcast) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

func FromJSON(data []byte) (Podcast, error) {
	var p Podcast
	err := json.Unmarshal(data, &p)
	return p, err
}

func DeletePodcast(id int) error {
	db := database.GetDB()
	query := `DELETE FROM podcasts WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("podcast with id %d not found", id)
	}

	return nil
}