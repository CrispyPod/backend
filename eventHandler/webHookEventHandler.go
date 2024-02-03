package eventhandler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"crispypod.com/crispypod-backend/db"
	"crispypod.com/crispypod-backend/dbModels"
	"github.com/google/uuid"
)

func TriggerHook(hook dbModels.Hook, data any) {

	j, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(hook.Method, hook.WebURL, bytes.NewReader(j))
	if err != nil {
		log.Fatal(err)
		return
	}
	var headerMap map[string]interface{}
	if err := json.Unmarshal([]byte(hook.Headers.String), &headerMap); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Triggerint hook %v", hook.ID.String())

	for k, v := range headerMap {
		req.Header.Add(k, fmt.Sprintf("%v", v))
	}

	dbHookLog := dbModels.HookLog{
		ID:         uuid.New(),
		HookID:     hook.ID,
		Status:     dbModels.HookLogStatusType_Started,
		CreateTime: time.Now(),
	}
	db.DB.Create(dbHookLog)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	dbHookLog.Duration = time.Since(dbHookLog.CreateTime)
	dbHookLog.Status = dbModels.HookLogStatusType_Finished
	if headerMarshaled, err := json.Marshal(resp.Header); err != nil {
		dbHookLog.ResponseHeader = sql.NullString{String: string(headerMarshaled), Valid: true}
	}

	if b, err := io.ReadAll(resp.Body); err != nil {
		dbHookLog.ResponseBody = sql.NullString{String: string(b), Valid: true}
	}

	db.DB.Save(dbHookLog)
}
