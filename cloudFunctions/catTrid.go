package cloudFunctions

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

func catTrid(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		CatTrid string `json:"cat_trid"`
		PtJwt string `json:"pt_jwt"`
	}
	var ptJwtValue string
	var catTridValue string
	ptJwt, err := r.Cookie("pt_jwt")
	if err != nil {
		ptJwtValue = ""
	} else {
		ptJwtValue = ptJwt.Value
	}
	catTrid, err := r.Cookie("cat_trid")
	if err != nil {
		now := strconv.FormatInt(time.Now().UnixNano(), 10)
		catTridValue = uuid.New().String() + now[:10] + "." + now[10:17]
	} else {
		catTridValue = catTrid.Value
	}
	http.SetCookie(w, &http.Cookie{
		Name:       "cat_trid",
		Value:      catTridValue,
		Path:       "/",
		Domain:     ".breaktime.com.tw",
		Expires:    time.Now().AddDate(2, 0, 0),
		RawExpires: "",
		MaxAge:     60 * 60 * 24 * 365 * 2,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	})
	res := Response{CatTrid: catTridValue, PtJwt: ptJwtValue}
	jsRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(jsRes)
	fmt.Println("write", write)
	w.WriteHeader(200)
}
