package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-curator/internal/server";"github.com/stockyard-dev/stockyard-curator/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="9700"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./curator-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("curator: %v",err)};defer db.Close();srv:=server.New(db)
fmt.Printf("\n  Curator — Self-hosted recipe and meal planner\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("curator: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
