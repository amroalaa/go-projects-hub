package main

import (
	"html/template"
	"log"
	"net/http"
)

type project struct {
	ID    int
	Title string
	Skill string
	link  string
}

func main() {
	projects := []project{
		{1, "الآلة الحاسبة", "المتغيرات و Switch", "https://github.com/amroalaa"},
		{2, "تخمين الرقم", "الحلقات التكرارية (For)", "https://github.com/amroalaa"},
		{3, "محول الوحدات", "العمليات الحسابية و Printf", "https://github.com/amroalaa"},
		{4, "مولد الكلمات", "النصوص (Strings) و Slices", "https://github.com/amroalaa"},
		{5, "قائمة المهام", "append و range", "https://github.com/amroalaa"},
		{6, "محلل النصوص", "الخرائط (Maps)", "https://github.com/amroalaa"},
		{7, "إدارة المكتبة", "الهياكل (Structs)", "https://github.com/amroalaa"},
		{8, "خادم الويب", "حزمة net/http", "https://github.com/amroalaa"},
		{9, "فحص المواقع", "التزامن (Goroutines)", "https://github.com/amroalaa"},
		{10, "تطبيق الطقس", "JSON و APIs", "https://github.com/amroalaa"},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, projects)
	})
	log.Println("الموقع يعمل على http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
