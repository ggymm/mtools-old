package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/goccy/go-json"

	"mtools-pro/backend"
)

var HandlerMap = make(map[string]Handler)

type Handler struct {
	svr    any
	method string
	params reflect.Type
}

type Req struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

func returnErr(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err))
}

func returnSuccess(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(data)
	if err != nil {
		returnErr(w, "response marshal error")
		return
	}
	_, _ = w.Write(resp)
}

func runDevServer() {
	m := http.NewServeMux()
	m.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// 设置允许跨域访问
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("origin"))
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		var (
			err  error
			body []byte

			req    Req
			method reflect.Value
			params []reflect.Value
		)

		body, err = io.ReadAll(r.Body)
		if err != nil {
			returnErr(w, "read body error")
			return
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			returnErr(w, "unmarshal body error")
			return
		}

		h, ok := HandlerMap[req.Method]
		if !ok {
			returnErr(w, "func error, not found")
			return
		}

		// 获取方法
		method = reflect.ValueOf(h.svr).MethodByName(h.method)
		if method == (reflect.Value{}) {
			returnErr(w, "func method error, not found")
			return
		}

		// 获取方法参数
		if h.params != nil {
			param := reflect.New(h.params).Elem()
			err = json.Unmarshal([]byte(req.Params), param.Addr().Interface())
			if err != nil {
				returnErr(w, "func params error, unmarshal error")
				return
			}
			params = append(params, param)
		}

		// 调用方法
		result := method.Call(params)
		if len(result) == 0 {
			returnSuccess(w, nil)
			return
		}

		// 返回结果
		returnSuccess(w, result[0].Interface())
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        m,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	slog.Error("start server", s.ListenAndServe())
}

var wg = sync.WaitGroup{}

func main() {
	app := backend.NewApp()
	app.InitLog()
	err := app.InitStorage()
	if err != nil {
		slog.Error("init storage", err)
		return
	}
	app.ServicesStartup(context.Background())

	services := app.Services()
	for _, svr := range services {
		t := reflect.TypeOf(svr)
		for i := 0; i < t.NumMethod(); i++ {
			sName := t.Elem().Name()
			mName := t.Method(i).Name
			mType := t.Method(i).Type
			key := fmt.Sprintf("%s.%s", sName, mName)
			if _, ok := HandlerMap[key]; ok {
				continue
			}

			numParams := mType.NumIn()
			if numParams <= 1 {
				HandlerMap[key] = Handler{
					svr:    svr,
					method: mName,
				}
			} else {
				HandlerMap[key] = Handler{
					svr:    svr,
					method: mName,
					params: mType.In(1),
				}
			}
		}
	}

	wg.Add(1)
	go runDevServer()
	wg.Wait()
}
