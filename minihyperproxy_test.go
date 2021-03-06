package minihyperproxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestGenerale(t *testing.T) {
	os.Setenv("PROXY_SERVER", "7053")

	m := NewMinihyperProxy()

	m.startHopperServer("prova", "")
	m.startHopperServer("prova2", "")
	m.AddHop("prova", &url.URL{Host: "www.google.com", Scheme: "http"}, &url.URL{Host: "localhost:7055", Scheme: "http"})
	m.ReceiveHop("prova2", &url.URL{Host: "www.google.com", Scheme: "http"}, &url.URL{Host: "localhost:7055", Scheme: "http"})
	//target, err := url.Parse("https://google.com/")

	//m.startProxyServer("prova2")
	//m.addProxyRedirect("prova2", &url.URL{Path: "/google", Scheme: "http"}, target)

	time.Sleep(5 * time.Second)

	resp, err := http.Get("http://localhost:7054/www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body))
	data, _ := json.Marshal(m.GetServersInfo())
	fmt.Println(string(data))
	m.stopServer("prova")
	m.stopServer("prova2")
}
