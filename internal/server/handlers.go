package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-curator/internal/store")
func(s *Server)handleListRecipes(w http.ResponseWriter,r *http.Request){q:=r.URL.Query().Get("q");list,_:=s.db.ListRecipes(q);if list==nil{list=[]store.Recipe{}};writeJSON(w,200,list)}
func(s *Server)handleCreateRecipe(w http.ResponseWriter,r *http.Request){var rec store.Recipe;json.NewDecoder(r.Body).Decode(&rec);if rec.Title==""{writeError(w,400,"title required");return};s.db.CreateRecipe(&rec);writeJSON(w,201,rec)}
func(s *Server)handleDeleteRecipe(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.DeleteRecipe(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handlePlan(w http.ResponseWriter,r *http.Request){var m store.MealPlan;json.NewDecoder(r.Body).Decode(&m);if m.Week==""||m.Day==""{writeError(w,400,"week and day required");return};if m.Meal==""{m.Meal="dinner"};s.db.PlanMeal(&m);writeJSON(w,201,m)}
func(s *Server)handleGetPlan(w http.ResponseWriter,r *http.Request){week:=r.URL.Query().Get("week");if week==""{writeError(w,400,"week required");return};list,_:=s.db.GetPlan(week);if list==nil{list=[]store.MealPlan{}};writeJSON(w,200,list)}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
