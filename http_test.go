package got

import (
	"fmt"
	"net/url"
	"testing"
)

func TestHttp(t *testing.T) {
	j, e := HttpDefault.Get("http://www.baidu.com/")
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Println(j.Status)
}

func TestHttpGetJson(t *testing.T) {
	j, e := HttpGetJson("http://www.baidu.com/")
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Println(j)
}

func TestHttpPostJson(t *testing.T) {
	form := url.Values{}
	form.Set("a", "b")
	j, e := HttpPostJson("http://www.baidu.com/", form)
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Println(j)
}
