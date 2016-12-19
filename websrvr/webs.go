package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//import _ "github.com/go-sql-driver/mysql"

// var db *sql.DB
// var err error

func main() {

	// db, err = sql.Open("mysql", "root:1234567@tcp(127.0.0.1:3306)/test_1")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // sql.DB should be long lived "defer" closes it once this function ends
	// defer db.Close()

	// // Test the connection to the database
	// err = db.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// rows, _ := db.Query("SELECT * FROM teacher")
	// //checkErr(err)

	// for rows.Next() {
	// 	var NAME string
	// 	var DOB string
	// 	var Add string
	// 	var Pin int
	// 	err = rows.Scan(&NAME, &DOB, &Add, &Pin)
	// 	//checkErr(err)
	// 	fmt.Println(NAME)
	// 	fmt.Println(DOB)
	// 	fmt.Println(Add)
	// 	fmt.Println(Pin)
	// }

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:9999", nil))
}
func comma(s string) string {
	var res string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		res = s[dot:]
		s = s[:dot]
	}
	n := len(s)
	if n <= 3 {
		return s
	}

	// if dot := strings.LastIndex(s, "."); dot >= 0 {
	// 	if dot > 3 {
	// 		n = dot
	// 	} else {
	// 		return s
	// 	}
	// }

	return comma(s[:n-3]) + "," + s[n-3:] + res
}

func comma2(s string) string {
	var buf [10][]byte
	var temp bytes.Buffer

	var res string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		res = s[dot:]
		s = s[:dot]
	}

	count := len(s) / 3

	buf[count] = []byte(res)

	for i := len(s); i > 2; {
		count--
		buf[count] = []byte(s[i-3:])
		i -= 3
	}

	return string( bytes.Join([][]byte(buf), []byte(","))

	// n := len(s)
	// if n <= 3 {
	// 	return s
	// }

	// if dot := strings.LastIndex(s, "."); dot >= 0 {
	// 	if dot > 3 {
	// 		n = dot
	// 	} else {
	// 		return s
	// 	}
	// }

	//return comma(s[:n-3]) + "," + s[n-3:] + res
}

// intsToString is like fmt.Sprintf(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	//var ww io.Writer

	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%d", v))
		//fmt.Fprintf(ww, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s\n\n", intsToString([]int{1, 2, 3}))

	tstr := comma2(r.URL.String()[1:])

	// fstr := tstr[:6]
	// lstr := tstr[6:]

	// ttstr := tstr[9:]

	//fmt.Println(tstr, fstr, lstr, ttstr)
	fmt.Fprintf(w, "\n%s\n\n", tstr)
	fmt.Fprintf(w, "%v %s %s\n", r.Method, r.URL.String(), r.Proto)
	fmt.Fprintf(w, "Hostname: %s\n", string(r.Host))

	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v[0])
	}
	username1 := string(r.URL.RequestURI())
	username2 := string(r.URL.String())

	fmt.Fprintf(w, "\nusername1=%v\n", username1)
	fmt.Fprintf(w, "\nusername2=%v\n", username2)

	username := username1[1:]
	//username := r.FormValue("username")

	fmt.Fprintf(w, "\nusername=%s\n", username)

	o := 0666
	fmt.Fprintf(w, "%d %o %#o\n", o, o, o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Fprintf(w, "%d %x %#x %#X\n", x, x, x, x)

	// qr := fmt.Sprintf("SELECT * FROM teacher where name like '%s'", username)

	// fmt.Printf("query=%s\n", qr)
	// rows, err := db.Query(qr)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// for rows.Next() {
	// 	var NAME string
	// 	var DOB string
	// 	var Add string
	// 	var Pin int
	// 	err = rows.Scan(&NAME, &DOB, &Add, &Pin)
	// 	fmt.Fprintf(w, "%s:", NAME)
	// 	fmt.Fprintf(w, "%s:", DOB)
	// 	fmt.Fprintf(w, "%s:", Add)
	// 	fmt.Fprintf(w, "%d:", Pin)
	// }

	//r.Header.Write(w)
}
