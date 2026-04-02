package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Recipe struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Ingredients string `json:"ingredients"`
	Instructions string `json:"instructions"`
	PrepTime int `json:"prep_time"`
	CookTime int `json:"cook_time"`
	Servings int `json:"servings"`
	Category string `json:"category"`
	Tags string `json:"tags"`
	Rating int `json:"rating"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"curator.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS recipes(id TEXT PRIMARY KEY,title TEXT NOT NULL,ingredients TEXT DEFAULT '[]',instructions TEXT DEFAULT '',prep_time INTEGER DEFAULT 0,cook_time INTEGER DEFAULT 0,servings INTEGER DEFAULT 4,category TEXT DEFAULT '',tags TEXT DEFAULT '',rating INTEGER DEFAULT 0,created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Recipe)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO recipes(id,title,ingredients,instructions,prep_time,cook_time,servings,category,tags,rating,created_at)VALUES(?,?,?,?,?,?,?,?,?,?,?)`,e.ID,e.Title,e.Ingredients,e.Instructions,e.PrepTime,e.CookTime,e.Servings,e.Category,e.Tags,e.Rating,e.CreatedAt);return err}
func(d *DB)Get(id string)*Recipe{var e Recipe;if d.db.QueryRow(`SELECT id,title,ingredients,instructions,prep_time,cook_time,servings,category,tags,rating,created_at FROM recipes WHERE id=?`,id).Scan(&e.ID,&e.Title,&e.Ingredients,&e.Instructions,&e.PrepTime,&e.CookTime,&e.Servings,&e.Category,&e.Tags,&e.Rating,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Recipe{rows,_:=d.db.Query(`SELECT id,title,ingredients,instructions,prep_time,cook_time,servings,category,tags,rating,created_at FROM recipes ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Recipe;for rows.Next(){var e Recipe;rows.Scan(&e.ID,&e.Title,&e.Ingredients,&e.Instructions,&e.PrepTime,&e.CookTime,&e.Servings,&e.Category,&e.Tags,&e.Rating,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Update(e *Recipe)error{_,err:=d.db.Exec(`UPDATE recipes SET title=?,ingredients=?,instructions=?,prep_time=?,cook_time=?,servings=?,category=?,tags=?,rating=? WHERE id=?`,e.Title,e.Ingredients,e.Instructions,e.PrepTime,e.CookTime,e.Servings,e.Category,e.Tags,e.Rating,e.ID);return err}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM recipes WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM recipes`).Scan(&n);return n}

func(d *DB)Search(q string, filters map[string]string)[]Recipe{
    where:="1=1"
    args:=[]any{}
    if q!=""{
        where+=" AND (title LIKE ?)"
        args=append(args,"%"+q+"%");
    }
    if v,ok:=filters["category"];ok&&v!=""{where+=" AND category=?";args=append(args,v)}
    rows,_:=d.db.Query(`SELECT id,title,ingredients,instructions,prep_time,cook_time,servings,category,tags,rating,created_at FROM recipes WHERE `+where+` ORDER BY created_at DESC`,args...)
    if rows==nil{return nil};defer rows.Close()
    var o []Recipe;for rows.Next(){var e Recipe;rows.Scan(&e.ID,&e.Title,&e.Ingredients,&e.Instructions,&e.PrepTime,&e.CookTime,&e.Servings,&e.Category,&e.Tags,&e.Rating,&e.CreatedAt);o=append(o,e)};return o
}

func(d *DB)Stats()map[string]any{
    m:=map[string]any{"total":d.Count()}
    return m
}
