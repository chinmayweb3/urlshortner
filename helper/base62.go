package helper

import (
	"fmt"

	"github.com/chinmayweb3/urlshortner/model"
	"github.com/deatil/go-encoding/base62"
)

// var counterNum int64 = 0

// func Base62Encoding(t time.Time) string {
// return string(base62.StdEncoding.Encode([]byte(t.Format("2006-01-02 15:04:05" + strconv.Itoa(rand.Intn(10000))))))[0:7]

func Base62Encoding() string {
	// counterNum++
	countNum := model.GetCountNumber()
	encode := string(base62.StdEncoding.Encode([]byte(fmt.Sprintf("%09d", countNum))))
	return encode
}
