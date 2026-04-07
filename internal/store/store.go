package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type DB struct{ db *sql.DB }

// Recipe is a single curated recipe. Ingredients and Instructions are
// stored as long-form strings (the dashboard treats Ingredients as one
// item per line). Times are integer minutes. Rating is 0-5 (0 = unrated).
type Recipe struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
	PrepTime     int    `json:"prep_time"`
	CookTime     int    `json:"cook_time"`
	Servings     int    `json:"servings"`
	Category     string `json:"category"`
	Tags         string `json:"tags"`
	Rating       int    `json:"rating"`
	CreatedAt    string `json:"created_at"`
}

func Open(d string) (*DB, error) {
	if err := os.MkdirAll(d, 0755); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", filepath.Join(d, "curator.db")+"?_journal_mode=WAL&_busy_timeout=5000")
	if err != nil {
		return nil, err
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS recipes(
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		ingredients TEXT DEFAULT '',
		instructions TEXT DEFAULT '',
		prep_time INTEGER DEFAULT 0,
		cook_time INTEGER DEFAULT 0,
		servings INTEGER DEFAULT 4,
		category TEXT DEFAULT '',
		tags TEXT DEFAULT '',
		rating INTEGER DEFAULT 0,
		created_at TEXT DEFAULT(datetime('now'))
	)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_recipes_category ON recipes(category)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_recipes_rating ON recipes(rating)`)
	db.Exec(`CREATE TABLE IF NOT EXISTS extras(
		resource TEXT NOT NULL,
		record_id TEXT NOT NULL,
		data TEXT NOT NULL DEFAULT '{}',
		PRIMARY KEY(resource, record_id)
	)`)
	return &DB{db: db}, nil
}

func (d *DB) Close() error { return d.db.Close() }

func genID() string { return fmt.Sprintf("%d", time.Now().UnixNano()) }
func now() string   { return time.Now().UTC().Format(time.RFC3339) }

func (d *DB) Create(e *Recipe) error {
	e.ID = genID()
	e.CreatedAt = now()
	if e.Servings == 0 {
		e.Servings = 4
	}
	_, err := d.db.Exec(
		`INSERT INTO recipes(id, title, ingredients, instructions, prep_time, cook_time, servings, category, tags, rating, created_at)
		 VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		e.ID, e.Title, e.Ingredients, e.Instructions, e.PrepTime, e.CookTime, e.Servings, e.Category, e.Tags, e.Rating, e.CreatedAt,
	)
	return err
}

func (d *DB) Get(id string) *Recipe {
	var e Recipe
	err := d.db.QueryRow(
		`SELECT id, title, ingredients, instructions, prep_time, cook_time, servings, category, tags, rating, created_at
		 FROM recipes WHERE id=?`,
		id,
	).Scan(&e.ID, &e.Title, &e.Ingredients, &e.Instructions, &e.PrepTime, &e.CookTime, &e.Servings, &e.Category, &e.Tags, &e.Rating, &e.CreatedAt)
	if err != nil {
		return nil
	}
	return &e
}

func (d *DB) List() []Recipe {
	rows, _ := d.db.Query(
		`SELECT id, title, ingredients, instructions, prep_time, cook_time, servings, category, tags, rating, created_at
		 FROM recipes ORDER BY title ASC`,
	)
	if rows == nil {
		return nil
	}
	defer rows.Close()
	var o []Recipe
	for rows.Next() {
		var e Recipe
		rows.Scan(&e.ID, &e.Title, &e.Ingredients, &e.Instructions, &e.PrepTime, &e.CookTime, &e.Servings, &e.Category, &e.Tags, &e.Rating, &e.CreatedAt)
		o = append(o, e)
	}
	return o
}

func (d *DB) Update(e *Recipe) error {
	_, err := d.db.Exec(
		`UPDATE recipes SET title=?, ingredients=?, instructions=?, prep_time=?, cook_time=?, servings=?, category=?, tags=?, rating=?
		 WHERE id=?`,
		e.Title, e.Ingredients, e.Instructions, e.PrepTime, e.CookTime, e.Servings, e.Category, e.Tags, e.Rating, e.ID,
	)
	return err
}

func (d *DB) Delete(id string) error {
	_, err := d.db.Exec(`DELETE FROM recipes WHERE id=?`, id)
	return err
}

func (d *DB) Count() int {
	var n int
	d.db.QueryRow(`SELECT COUNT(*) FROM recipes`).Scan(&n)
	return n
}

func (d *DB) Search(q string, filters map[string]string) []Recipe {
	where := "1=1"
	args := []any{}
	if q != "" {
		where += " AND (title LIKE ? OR ingredients LIKE ? OR tags LIKE ?)"
		s := "%" + q + "%"
		args = append(args, s, s, s)
	}
	if v, ok := filters["category"]; ok && v != "" {
		where += " AND category=?"
		args = append(args, v)
	}
	if v, ok := filters["min_rating"]; ok && v != "" {
		where += " AND rating >= ?"
		args = append(args, v)
	}
	rows, _ := d.db.Query(
		`SELECT id, title, ingredients, instructions, prep_time, cook_time, servings, category, tags, rating, created_at
		 FROM recipes WHERE `+where+`
		 ORDER BY title ASC`,
		args...,
	)
	if rows == nil {
		return nil
	}
	defer rows.Close()
	var o []Recipe
	for rows.Next() {
		var e Recipe
		rows.Scan(&e.ID, &e.Title, &e.Ingredients, &e.Instructions, &e.PrepTime, &e.CookTime, &e.Servings, &e.Category, &e.Tags, &e.Rating, &e.CreatedAt)
		o = append(o, e)
	}
	return o
}

// Stats returns aggregate metrics: total recipes, distinct categories,
// average rating (across rated recipes), favorites count (rating>=4),
// total cook time across the collection, and breakdowns by category.
func (d *DB) Stats() map[string]any {
	m := map[string]any{
		"total":       d.Count(),
		"categories":  0,
		"avg_rating":  0.0,
		"favorites":   0,
		"by_category": map[string]int{},
	}

	var cats int
	d.db.QueryRow(`SELECT COUNT(DISTINCT category) FROM recipes WHERE category != ''`).Scan(&cats)
	m["categories"] = cats

	var avg float64
	d.db.QueryRow(`SELECT COALESCE(AVG(CASE WHEN rating > 0 THEN rating END), 0) FROM recipes`).Scan(&avg)
	m["avg_rating"] = avg

	var fav int
	d.db.QueryRow(`SELECT COUNT(*) FROM recipes WHERE rating >= 4`).Scan(&fav)
	m["favorites"] = fav

	if rows, _ := d.db.Query(`SELECT category, COUNT(*) FROM recipes WHERE category != '' GROUP BY category`); rows != nil {
		defer rows.Close()
		by := map[string]int{}
		for rows.Next() {
			var s string
			var c int
			rows.Scan(&s, &c)
			by[s] = c
		}
		m["by_category"] = by
	}

	return m
}

// ─── Extras: generic key-value storage for personalization custom fields ───

func (d *DB) GetExtras(resource, recordID string) string {
	var data string
	err := d.db.QueryRow(
		`SELECT data FROM extras WHERE resource=? AND record_id=?`,
		resource, recordID,
	).Scan(&data)
	if err != nil || data == "" {
		return "{}"
	}
	return data
}

func (d *DB) SetExtras(resource, recordID, data string) error {
	if data == "" {
		data = "{}"
	}
	_, err := d.db.Exec(
		`INSERT INTO extras(resource, record_id, data) VALUES(?, ?, ?)
		 ON CONFLICT(resource, record_id) DO UPDATE SET data=excluded.data`,
		resource, recordID, data,
	)
	return err
}

func (d *DB) DeleteExtras(resource, recordID string) error {
	_, err := d.db.Exec(
		`DELETE FROM extras WHERE resource=? AND record_id=?`,
		resource, recordID,
	)
	return err
}

func (d *DB) AllExtras(resource string) map[string]string {
	out := make(map[string]string)
	rows, _ := d.db.Query(
		`SELECT record_id, data FROM extras WHERE resource=?`,
		resource,
	)
	if rows == nil {
		return out
	}
	defer rows.Close()
	for rows.Next() {
		var id, data string
		rows.Scan(&id, &data)
		out[id] = data
	}
	return out
}
