// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b

package main

import "fmt"

func main() {
	var str1 = []string{"Java", "Python", "Javascript", "C ++", "C#"}
	var str2 = []string{"Swift", "Java", "Kotlin", "Python"}
	fmt.Println(Intersection(str1, str2))
}

func Intersection(str1, str2 []string) (inter []string) {
	//return []string{} // TODO: replace this
	m := make(map[string]bool)
	for _, item := range str1 {
		m[item] = true
}

for _, item := range str2 {
		if _, ok := m[item]; ok {
				inter = append(inter, item)
		}
}
return
}

func RemoveDuplicates(elements []string) (nodups []string) {
	//return []string{} // TODO: replace this
	keys := make(map[string]bool)
    list := []string{}

	for _, entry := range string {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list

}
