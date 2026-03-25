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
	baseRepoURL := "https://github.com/amroalaa"
	projects := []project{
		{1, "الآلة الحاسبة", "المتغيرات و Switch", baseRepoURL + "calculator"},
		{2, "تخمين الرقم", "الحلقات التكرارية", baseRepoURL + "guessing-game"},
		{3, "محول الوحدات", "العمليات الحسابية", baseRepoURL + "unit-converter"},
		{4, "مولد الكلمات", "النصوص و Slices", baseRepoURL + "password-generator"},
		{5, "قائمة المهام", "append و range", baseRepoURL + "todo-list"},
		{6, "محلل النصوص", "الخرائط (Maps)", baseRepoURL + "word-counter"},
		{7, "إدارة المكتبة", "الهياكل (Structs)", baseRepoURL + "library-system"},
		{8, "خادم الويب", "حزمة net/http", baseRepoURL + "simple-web-server"},
		{9, "فحص المواقع", "التزامن (Goroutines)", baseRepoURL + "site-checker"},
		{10, "تطبيق الطقس", "JSON و APIs", baseRepoURL + "weather-cli"},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, projects)
	})
	log.Println("الموقع يعمل على http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
