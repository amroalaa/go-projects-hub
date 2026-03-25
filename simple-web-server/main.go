package main // تعريف أن هذا الملف هو نقطة انطلاق البرنامج

import (
	"fmt"      // استيراد حزمة الطباعة والتنسيق
	"net/http" // استيراد الحزمة الأساسية للتعامل مع بروتوكولات الويب (HTTP)
)

// دالة لمعالجة الطلبات التي تأتي للصفحة الرئيسية
func homeHandler(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w, "<h1>مرحباً بك في خادم جو!</h1>") // إرسال نص HTML كاستجابة لمتصفح المستخدم
	fmt.Fprintf(w, "<p>هذه الصفحة تعمل بواسطة لغة Go.</p>") // إضافة سطر آخر للاستجابة
} // تنتهي الدالة هنا

// دالة لمعالجة صفحة "عن الموقع"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>حول المشروع</h1>") // إرسال عنوان الصفحة للمتصفح
	fmt.Fprintf(w, "<p>لقد قمت ببناء هذا الخادم من الصفر!</p>") // إرسال محتوى الصفحة
} // تنتهي الدالة هنا

func main() { // الدالة الرئيسية التي تشغل الخادم
	http.HandleFunc("/", homeHandler)       // ربط الرابط الرئيسي (/) بالدالة homeHandler
	http.HandleFunc("/about", aboutHandler) // ربط الرابط (/about) بالدالة aboutHandler

	fmt.Println("الخادم يعمل الآن على: http://localhost:8080") // تنبيه المبرمج بأن الخادم جاهز
	
	// تشغيل الخادم على المنفذ 8080 والبدء في الاستماع للطلبات
	err := http.ListenAndServe(":8080", nil) 
	
	if err != nil { // التحقق مما إذا حدث خطأ أثناء تشغيل الخادم (مثل أن المنفذ محجوز)
		fmt.Println("فشل تشغيل الخادم:", err) // طباعة رسالة الخطأ في حال حدوثه
	}
}
