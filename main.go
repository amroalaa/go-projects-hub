package mmain

import (
	"html/template"
	"log"
	"net/http"
)

type project struct {
	ID    int
	Title string
	Skill string
}

func main() {
	projects := []project{
		{1, "الآلة الحاسبة", "المتغيرات و Switch"},
		{2, "تخمين الرقم", "الحلقات التكرارية (For)"},
		{3, "محول الوحدات", "العمليات الحسابية و Printf"},
		{4, "مولد الكلمات", "النصوص (Strings) و Slices"},
		{5, "قائمة المهام", "append و range"},
		{6, "محلل النصوص", "الخرائط (Maps)"},
		{7, "إدارة المكتبة", "الهياكل (Structs)"},
		{8, "خادم الويب", "حزمة net/http"},
		{9, "فحص المواقع", "التزامن (Goroutines)"},
		{10, "تطبيق الطقس", "JSON و APIs"},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, projects)
	})
	log.Println("الموقع يعمل على http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
