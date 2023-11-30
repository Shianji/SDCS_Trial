package main

import (
	pb "SDCS/json_cache"
	"context"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc"
)

var address [4]string
var client [2]pb.CacheClient
var conn [2]*grpc.ClientConn
var cache Cache
var hashselect int //这个变量用来控制将键值对根据其hash值存储到指定服务器

// Cache 结构体用于存储缓存数据
type Cache struct {
	data map[string]string
	mu   sync.RWMutex // 用于保护并发访问的互斥锁
	pb.UnimplementedCacheServer
}

// Post实现gRPC的Post方法，用于向缓存中写入数据
func (c *Cache) Post(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[req.Key] = req.Value
	return &pb.PostResponse{}, nil
}

// Get实现gRPC的Get方法，用于从缓存中获取数据
func (c *Cache) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if _, ok := c.data[req.Key]; ok {
		return &pb.GetResponse{Value: c.data[req.Key]}, nil
	} else {
		return &pb.GetResponse{Value: ""}, nil
	}
}

// Delete实现gRPC的Delete方法，用于从缓存中删除数据
func (c *Cache) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[req.Key]; ok {
		delete(c.data, req.Key)
		return &pb.DeleteResponse{Num: 1}, nil
	} else {
		return &pb.DeleteResponse{Num: 0}, nil
	}
}

func setAddress() {
	if os.Args[1] == "1" {
		hashselect = 1
		address[0] = "127.0.0.1:9527"
		address[1] = "127.0.0.1:9530"

		address[2] = "127.0.0.1:9531"
		address[3] = "127.0.0.1:9532"
	} else if os.Args[1] == "2" {
		hashselect = 2
		address[0] = "127.0.0.1:9528"
		address[1] = "127.0.0.1:9531"

		address[2] = "127.0.0.1:9530"
		address[3] = "127.0.0.1:9532"
	} else if os.Args[1] == "3" {
		hashselect = 3
		address[0] = "127.0.0.1:9529"
		address[1] = "127.0.0.1:9532"

		address[2] = "127.0.0.1:9530"
		address[3] = "127.0.0.1:9531"
	} else {
		fmt.Println("only 3 cacheserver.")
	}
}

func httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 处理 HTTP Get请求
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 2 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		rkey := parts[1]
		getReq := pb.GetRequest{Key: rkey}
		var getResp pb.GetResponse
		if value, ok := cache.data[rkey]; ok {
			getResp.Value = value
		} else {
			response := GetCache(client[0], &getReq)
			if response.Value != "" {
				getResp.Value = response.Value
			} else {
				response = GetCache(client[1], &getReq)
				if response.Value != "" {
					getResp.Value = response.Value
				} else {
					w.WriteHeader(http.StatusNotFound)
					return
				}
			}
		}
		jsonResponse := map[string]string{rkey: getResp.Value}
		jsonData, err := json.Marshal(jsonResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	case http.MethodPost:
		// 处理 HTTP Post请求
		var postReq pb.PostRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var tempMap map[string]interface{}
		json.Unmarshal(body, &tempMap)
		for key, value := range tempMap {
			postReq.Key = key
			postReq.Value = fmt.Sprintf("%v", value)
			break
		}
		defer r.Body.Close()
		h := fnv.New32a()
		h.Write([]byte(postReq.Key))
		hashValue := h.Sum32()
		hashnum := int(hashValue % 3)
		if hashnum == 0 {
			if hashselect == 1 {
				cache.data[postReq.Key] = postReq.Value
			} else {
				PostCache(client[0], &postReq)
			}
		} else if hashnum == 1 {
			if hashselect == 1 {
				PostCache(client[0], &postReq)
			} else if hashselect == 2 {
				cache.data[postReq.Key] = postReq.Value
			} else {
				PostCache(client[1], &postReq)
			}
		} else if hashnum == 2 {
			if hashselect == 3 {
				cache.data[postReq.Key] = postReq.Value
			} else {
				PostCache(client[1], &postReq)
			}
		}

		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		// 处理 HTTP Delete请求
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 2 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		var num int32
		rkey := parts[1]
		deleteReq := pb.DeleteRequest{Key: rkey}
		if _, ok := cache.data[rkey]; ok {
			delete(cache.data, rkey)
			num = 1
		} else {
			num = DeleteCache(client[0], &deleteReq)
			if num == 0 {
				num = DeleteCache(client[1], &deleteReq)
			}
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, strconv.Itoa(int(num)))

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// 启动HTTP服务器
func startHttpServer() {
	http.HandleFunc("/", httpRequestHandler)
	fmt.Println("Listening http at", address[0])
	err := http.ListenAndServe(address[0], nil)
	if err != nil {
		fmt.Println("Listten failed:", err)
	}
}

// 启动gRPC服务器
func startRpcServer() {
	fmt.Println("Listening rpc on", address[1])
	listen, err := net.Listen("tcp", address[1])
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cache = Cache{data: make(map[string]string)}
	pb.RegisterCacheServer(grpcServer, &cache)
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve:%v\n", err)
		return
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please specify server index(1-3).")
		return
	}

	setAddress()
	go startHttpServer()
	go startRpcServer()
	startClient()

	select {}
}
