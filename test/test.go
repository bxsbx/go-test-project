package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type One struct {
	Id    string `json:"ids"`
	Name  string `json:"names"`
	Count int
	Tiem  time.Time
	Two   Two
	three Three
	Uid   []Two
	Pp    map[string]Two
	Hh    []One
}

type One2 struct {
	Id2 string `json:"Id"`
	//Name2  string `json:"Name"`
	//Count2 int    `json:"Count"`
	Id    string
	Name  string
	Count json.Number
	Two   Two
	three Three
	Tiem  time.Time
	Uid   []Two
	Pp    map[string]Two
}

type Two struct {
	Uid   []string
	Uname int
	Three Three
	//One   []One
}

type Three struct {
	Nu string
	DD int64
	//one []One
}

func main() {
	var ones []One

	for i := 0; i < 10; i++ {
		ones = append(ones, One{
			Id:    strconv.Itoa(rand.Intn(30)),
			Name:  strconv.Itoa(rand.Intn(89)),
			Count: 67,
			Tiem:  time.Now(),
			three: Three{
				"233",
				34354,
				//nil,
			},
		})
	}
	var two = Two{
		Uid:   []string{"75656"},
		Uname: 3434,
		Three: Three{
			"222",
			4554,
			//nil,
		},
	}
	ones[0].Uid = []Two{two}
	ones[0].Two = Two{
		Uid:   []string{"vfef"},
		Uname: 4545,
		Three: Three{
			"233",
			3435477,
			//nil,
		},
		//One: []One{ones[0]},
	}
	ones[0].Pp = make(map[string]Two)
	ones[0].Pp["11222"] = two
	ones[0].Pp["87878"] = two
	//ones[0].Hh = make([]One, 10)
	//ones[0].Hh[0] = ones[0]
	//fmt.Println(ones[0].Hh)
	//ones[0].Two.three.one = []One{ones[0]}

	//-----------1
	//toMap := util.ListObjToMap(ones, func(one One) string {
	//	return one.Id
	//})
	//for k, one := range toMap {
	//	str := fmt.Sprintf("%s->%s-%s_%d", k, one.Id, one.Name, one.Count)
	//	fmt.Println(str)
	//}

	//--------------2
	//obj := util.ListObjToListObj(ones, func(one One) Two {
	//	return Two{
	//		one.Id,
	//		one.Count,
	//	}
	//})
	//fmt.Println(obj)

	//------------------3
	//list := []int{1, 1, 6, 8, 0, 1, 3, 6, 3, 2, 4, 2, 2, 1}
	//
	//list = util.RemoveRepeatFromList(list)
	//fmt.Println(list)
	//
	//fromList := util.RemoveRepeatFromListObj(ones, func(one One) string {
	//	return one.Id
	//})
	//fmt.Println(fromList)

	//-----------4
	//ones[0].two = Two{
	//	Uid:   []string{"vfef"},
	//	Uname: 4545,
	//	three: Three{
	//		"233",
	//		34354,
	//	},
	//}
	//structToMap, err := util.StructToMap(Three{
	//	"vvd",
	//	3434,
	//})
	//ones[0].Count = 23
	//structToMap, err := util.StructToMap(ones[0])
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for k, v := range structToMap {
	//
	//	fmt.Println(k, v)
	//}
	//fmt.Println(structToMap["Count"])
	//sprintf := fmt.Sprintf("%v", structToMap["two"])
	//fmt.Println(sprintf)

	//--------------5
	//fmt.Println(reflect.TypeOf(structToMap["two"]).Kind())
	//if reflect.ValueOf(structToMap["two"]).Kind() == reflect.Map {
	//	reflect.ValueOf(structToMap["two"]).
	//}
	//
	//m, ok := structToMap["two"].(map[string]interface{})
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
	//fmt.Println(m["Uid"], ok)
	//
	////aa := One2(ones[0])
	//fmt.Println(ones[0])
	//err := util.Obj1ToObj2WithJson(ones[0], &aa)
	//fmt.Println(err)
	//fmt.Println(aa)
	//
	//jsonString := `{"Count": 123456565767}`
	//var bb One
	////util.JsonToStruct(jsonString, &bb)
	//jsoMap := map[string]interface{}{"ids": "cssc", "Count": 123456565767}
	//util.ObjToObjByJson(jsoMap, &bb)
	//fmt.Println(bb)
	//
	//err = json.Unmarshal([]byte(jsonString), &bb)
	//fmt.Println(err)
	//fmt.Println(bb)

	//var one2 One2
	//fmt.Printf("%p\n", &one2)
	//fmt.Printf("%p\n", ones[0].Two)
	//
	////fmt.Println(ones[0].Two)
	////fmt.Println(ones[0])
	//fmt.Printf("%p\n", ones[0].Two.three.one)
	////fmt.Printf("%p\n", ones[0].Uid)
	//fmt.Println("_________")
	//util.ObjToObjByReflect(ones[0], &one2)
	////fmt.Printf("%p", &one2)
	////fmt.Printf("%p\n", one2.Two)
	////ones[0].Two.Uname = 2323555
	////ones[0].Two.Uid[0] = "77777"
	////fmt.Println(ones[0].Two)
	////fmt.Println(one2.Two)
	////fmt.Println(one2)
	//fmt.Printf("%p\n", one2.Two.three.one)
	////fmt.Printf("%p\n", one2.Uid)
	//
	////var ki = []string{"vferf", "csaasc"}
	////var kk []string
	////util.ObjToObjByReflect(ones[0], &kk)
	////fmt.Println(kk)

	//var one2 One2
	//fmt.Printf("%p\n", ones[0].Uid)
	//fmt.Println(ones[0])
	//util.ObjToObjByReflect1(&ones[0], &one2)
	//fmt.Printf("%p\n", one2.Uid)
	//fmt.Println(one2)

	//a := []string{"2323"}
	//b := []string{"6767"}
	//util.CopyStructShallow(&a, &b)
	//fmt.Println(b)

	//var one2 One2
	//fmt.Printf("%p\n", &ones[0].Tiem)
	//fmt.Println(ones[0])
	//util.ObjToObjByReflect(&ones[0], &one2, true)
	////util.ObjToObjByJson(ones[0], &one2)
	//fmt.Printf("%p\n", &one2.Tiem)
	//fmt.Println(one2)

	//list := util.CopyList(ones[0].Two.One)
	//fmt.Println(list)
	//fmt.Printf("%p\n", &ones[0].Two.One[0])
	//fmt.Printf("%p\n", &list[0])
	//
	//var oo = [12]string{"cksock"}
	////var oo = make([]string, 10)
	//fmt.Println(reflect.Array == reflect.TypeOf(oo).Kind())

	//fmt.Println(ones[0])
	//fmt.Printf("%p\n", ones[0].Pp)
	//obj := util.CopyObj(&ones[0], true)
	//fmt.Println(obj)
	//fmt.Printf("%p\n", obj.Pp)
	//
	//s := []string{"233"}
	//fmt.Printf("%p\n", s)
	//copyObj := util.CopyObj(&s, false)
	//fmt.Println(copyObj)
	//fmt.Printf("%p\n", copyObj)
	var flagMap = make(map[int]bool, 0)
	flagMap[10] = true
	fmt.Println(len(flagMap))

}
