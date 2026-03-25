package main // تعريف أن هذا الملف هو نقطة انطلاق البرنامج

import (
	"encoding/json" // حزمة لتحويل بيانات JSON إلى هياكل Go والعكس
	"fmt"           // حزمة للطباعة والقراءة
	"net/http"      // حزمة لإرسال طلبات الإنترنت
)

// تعريف هيكل يطابق شكل البيانات القادمة من الإنترنت (JSON)
type WeatherResponse struct {
	Name string `json:"name"` // اسم المدينة
	Main struct {
		Temp float64 `json:"temp"` // درجة الحرارة
	} `json:"main"` // هيكل فرعي يحتوي على البيانات الأساسية
}

func main() { // الدالة الرئيسية للبرنامج
	city := "Cairo"                                                    // تحديد المدينة المراد جلب طقسها
	apiKey := "YOUR_API_KEY"                                           // مكان وضع مفتاح الـ API الخاص بك
	url := fmt.Sprintf("https://api.openweathermap.org", city, apiKey) // بناء رابط الطلب

	fmt.Printf("جاري جلب حالة الطقس لمدينة %s...\n", city) // رسالة للمستخدم أثناء الانتظار

	resp, err := http.Get(url) // إرسال طلب GET للموقع لجلب البيانات
	if err != nil {            // التحقق من وجود خطأ في الاتصال بالإنترنت
		fmt.Println("خطأ في الاتصال:", err) // طباعة رسالة الخطأ
		return                              // إنهاء البرنامج في حال وجود خطأ
	}
	defer resp.Body.Close() // إغلاق جسم الاستجابة بعد الانتهاء لتوفير موارد الذاكرة

	var data WeatherResponse                                         // إنشاء متغير من نوع الهيكل الذي صممناه لتخزين البيانات فيه
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil { // تفكيك بيانات JSON ووضعها داخل المتغير data
		fmt.Println("خطأ في معالجة البيانات:", err) // طباعة خطأ في حال فشل تحويل الـ JSON
		return                                      // إنهاء البرنامج
	}

	// طباعة النتائج النهائية للمستخدم بشكل مرتب
	fmt.Println("---------------------------")               // خط فاصل للتنظيم
	fmt.Printf("المدينة: %s\n", data.Name)                   // طباعة اسم المدينة المستلم من الـ API
	fmt.Printf("درجة الحرارة: %.1f مئوية\n", data.Main.Temp) // طباعة درجة الحرارة بتقريب رقم عشري واحد
	fmt.Println("---------------------------")               // خط فاصل ختامي
}
