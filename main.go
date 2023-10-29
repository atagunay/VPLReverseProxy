package main

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func handleWebSocket(req *http.Request, clientConn *websocket.Conn) {
	req.URL.Scheme = "ws"
	conn, _, err := websocket.DefaultDialer.Dial(req.URL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal()
		}
	}(conn)

	for {
		// Read web-socket messages from the main server
		messageType, response, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		// Return web-socket message to client
		err = clientConn.WriteMessage(messageType, response)
		if err != nil {
			log.Fatal(err)
		}

		// Log web-socket response
		color.Cyan("[ws-response] %s \n", response)
	}
}

// Set request for Server
func prepareRequestForServer(req *http.Request, body io.ReadCloser) *http.Request {
	originServerURL, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}

	req.Host = originServerURL.Host
	req.URL.Host = originServerURL.Host
	req.URL.Scheme = originServerURL.Scheme
	req.RequestURI = ""
	req.Body = body

	return req
}

// In Go, once you read the HTTP request body using io.ReadAll or any other method that reads the entire request body,
// the body is effectively consumed. After you've read and processed the body,
// it is no longer available for further use in your code
func replicate(element io.ReadCloser) (io.ReadCloser, io.ReadCloser) {
	buf, _ := io.ReadAll(element)
	body1 := io.NopCloser(bytes.NewBuffer(buf))
	body2 := io.NopCloser(bytes.NewBuffer(buf))

	return body1, body2
}

// URL Define Jail Server URL
const URL = "http://192.168.71.2"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	var reverseProxy http.HandlerFunc = func(rw http.ResponseWriter, req *http.Request) {
		clientReq := req

		// Log request information
		color.Green("[%s] %s %s%s \n", time.Now().Format("2006-01-02 15:04:05"),
			req.Method, req.RemoteAddr, req.URL.Path)

		// Log headers
		for key, element := range req.Header {
			color.Magenta("%s: %s\n", key, element)
		}

		// Replicate bodies
		body1, body2 := replicate(req.Body)

		// Log body with body1
		color.Yellow("%s \n", body1)

		// Prepare request with body2
		prepareRequestForServer(req, body2)

		// Check request type
		if req.Header.Get("Upgrade") == "websocket" {
			// Open web-socket between client and this server
			conn, err := upgrader.Upgrade(rw, clientReq, nil)
			if err != nil {
				log.Fatal(err)
			}

			// Send web request to main server
			handleWebSocket(req, conn)
		} else {
			// save the response from the origin server
			originServerResponse, err := http.DefaultClient.Do(req)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprint(rw, err)
				color.Red("[ERROR!] => %s", err)
				return
			}

			// Replicate response body
			body1, body2 = replicate(originServerResponse.Body)

			// Log response body
			color.Cyan("%s \n", body1)

			// Return response to the client
			rw.WriteHeader(http.StatusOK)
			_, err = io.Copy(rw, body2)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Fatal(http.ListenAndServe(":80", reverseProxy))
}
