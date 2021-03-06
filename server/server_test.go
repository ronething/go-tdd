// author: ashing
// time: 2020/4/21 12:51 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestPlayerServer(t *testing.T) {
//	t.Run("returns Pepper's score", func(t *testing.T) {
//		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
//		response := httptest.NewRecorder()
//
//		PlayerServer(response, request)
//
//		got := response.Body.String()
//		want := "10"
//
//		if got != want {
//			t.Errorf("got '%s', want '%s'", got, want)
//		}
//	})
//}

func TestGETPlayers(t *testing.T) {

	//server := &PlayerServer{}

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}


	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper() // 标记为辅助函数
	if got != want {
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}
