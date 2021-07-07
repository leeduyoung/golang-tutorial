package main

import (
	"net/http"
)

// func main() {
// 	fmt.Println("start server")

// 	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
// 		res.Write([]byte("Hello, world!")) // 웹 브라우저에 응답
// 	}) // 경로에 접속했을때 생성할 함수 설정

// 	http.ListenAndServe(":80", nil)
// }

func main() {
	// HTTP 요청 멀티플렉서 인스턴스 생성
	// HTTP 요청 멀티플렉서는 URL 경로를 통시에 여러개를 처리한다고 하여 멀티플렉서라고 부른다.
	// http.HandleFunc 함수도 내부적으로는 http.NewServeMux로 DefaultServeMux 인스턴스를 생성하여 사용하므로 완전히 같은 기능이다.)
	mux := http.NewServeMux()

	s := "Hello, world!"

	// 함수의 매개변수 req에는 웹 브라우저가 접속했을때의 메서드, 쿠키, 헤더 등의 정보가 들어있고
	// res로 웹 브라우저에 응답 할 수 있다.
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// HTML로 웹 페이지 작성
		html := `
		<html>
		<head>
			<title>Hello</title>
			<script type="text/javascript" src="/assets/hello.js"></script>
			<link href="/assets/hello.css" rel="stylesheet" />
		</head>
		<body>
			<span class="hello">` + s + `</span>
		</body>
		</html>
		`

		res.Header().Set("Content-Type", "text/html") // HTML 헤더 설정
		res.Write([]byte(html))                       // 웹 브라우저에 응답
	})

	mux.Handle( // /assets/ 경로에 접근했을 때
		// 파일 서버를 동작시킴
		"/assets/",
		http.StripPrefix( // 파일 서버를 실행할 때 assets
			// 디렉터리를 지정했으므로 URL 경로에서 /assets/ 삭제
			"/assets/",
			http.FileServer(http.Dir("assets")), // 웹 서버에서 assets
			// 디렉터리 아래의 파일을 표시
		),
	)

	// 80번 포트에 웹 서버를 실행하고, mux를 이용하여 HTTP 요청을 처리
	http.ListenAndServe(":80", mux)
}
