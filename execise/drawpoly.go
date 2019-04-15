package exe

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/*var mu sync.Mutex
var count int*/
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/poly", poly)
	log.Fatal(http.ListenAndServe("127.0.0.1:8084", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	mu.Unlock()
}

func poly(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	myWidth, err := strconv.ParseFloat(q["width"][0], 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	myHeight, err := strconv.ParseFloat(q["height"][0], 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; strokewidth:0.7' "+"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner2(i+1, j, myWidth, myHeight)
			bx, by := corner2(i, j, myWidth, myHeight)
			cx, cy := corner2(i, j+1, myWidth, myHeight)
			dx, dy := corner2(i+1, j+1, myWidth, myHeight)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("<==================================================================>")
	fmt.Printf("%f - width : %f - height", myWidth, myHeight)
	fmt.Println("<==================================================================>")
	fmt.Println("</svg>")
}

func corner2(i, j int, mywidth, myheight float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - .5)
	y := xyrange * (float64(j)/cells - .5)
	z := f(x, y)

	sx := mywidth/2 + (x-y)*cos30*xyscale
	sy := myheight/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
