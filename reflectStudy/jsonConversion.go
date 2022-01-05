package reflectStudy

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name string     `json:"title"`
	Year int	    `json:"year"`
	Price int	    `json:"price""`
	Actors []string `json:"actors"`
}

func JsonConversion() {
	movie := Movie{"movie_name", 1998, 100, []string{"a", "b"}}
	jsonStr, err  := json.Marshal(movie)
	if err != nil{
		fmt.Println("json marshal failed.", err)
		return
	}

	//fmt.Println(jsonStr)
	//这才是正确的打印json字符串方式
	fmt.Printf("%s\n", jsonStr)

	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		return
	}
	fmt.Println(myMovie)

}
