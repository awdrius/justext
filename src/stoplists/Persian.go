package gojustext

var PersianStoplist = map[string]bool {"و": true,
"در": true,
"به": true,
"از": true,
"که": true,
"این": true,
"را": true,
"با": true,
"است.": true,
"آن": true,
"است": true,
"سال": true,
"یک": true,
"برای": true,
"او": true,
"بر": true,
"سیارک": true,
"خود": true,
"شد.": true,
"تا": true,
"یا": true,
"شده": true,
"نیز": true,
"نام": true,
"ایران": true,
"قرار": true,
"بود": true,
"شهر": true,
"یکی": true,
"پس": true,
"دو": true,
"می": true,
"کشف": true,
"وی": true,
"بود.": true,
"هم": true,
"می‌شود.": true,
"استان": true,
"استفاده": true,
"بخش": true,
"شد": true,
"دیگر": true,
"عنوان": true,
"هر": true,
"شهرستان": true,
"کرد.": true,
"دارد.": true,
"توسط": true,
"برابر": true,
"وجود": true,
"کار": true,
"آنها": true,
"هزار": true,
"اما": true,
"کرد": true,
"کشور": true,
"می‌شود": true,
"مورد": true,
"نفر": true,
"صورت": true,
"روی": true,
"شده‌است": true,
"دارد": true,
"زمان": true,
"همچنین": true,
"دست": true,
"زبان": true,
"شده‌است.": true,
"دارای": true,
"دانشگاه": true,
"جمعیت": true,
"بوده": true,
"واقع": true,
"بعد": true,
"می‌باشد.": true,
"بین": true,
"دهستان": true,
"مرکز": true,
"روز": true,
"بسیار": true,
"کرده": true,
"مردم": true,
"بیشتر": true,
"کتاب": true,
"روستا": true,
"پیوند": true,
"نظر": true,
"پیش": true,
"شرکت": true,
"سه": true,
"داشته": true,
"قدر": true,
"آب": true,
"ولی": true,
"انجام": true,
"گروه": true,
"مطلق": true,
"فیلم": true,
"میان": true,
"تاریخ": true,
"بزرگ": true,
"بیرون.": true,
"آمریکا": true,
"زندگی": true,
"مانند": true,
"تهران": true,
"جنگ": true,
"اولین": true,
"می‌کند.": true,
"میلادی": true,
"چند": true,
"منطقه": true,
"برخی": true,
"بسیاری": true,
"راه": true,
"مرکزی": true,
"آغاز": true,
"داده": true,
"دوره": true,
"حال": true,
"اصلی": true,
"اول": true,
"حدود": true,
"می‌توان": true,
"بار": true,
"همین": true,
"بازی": true,
"تشکیل": true,
"توابع": true,
"تنها": true,
"همراه": true,
"محمد": true,
"جنوب": true,
"می‌کند": true,
"گرفته": true,
"جهان": true,
"ملی": true,
"دوران": true,
"کردن": true,
"سرشماری": true,
"ایجاد": true,
"تولید": true,
"شدن": true,
"دولت": true,
"سازمان": true,
"باید": true,
"شمال": true,
"است،": true,
"جمعیت.": true,
"بوده‌است.": true,
"اسلامی": true,
"اثر": true,
"جمله": true,
"شامل": true,
"هستند.": true,
"شرقی": true,
"زیر": true,
"منابع.": true,
"ساخته": true,
"ایران.": true,
"دوم": true,
"شاه": true,
"اگر": true,
"روستایی": true,
"شکل": true,
"باشد.": true,
"موسیقی": true,
"اساس": true,
"هستند": true,
"آثار": true,
"نوع": true,
"زیادی": true,
"نشان": true,
"دلیل": true,
"طول": true,
"ماه": true,
"بیش": true,
"ایرانی": true,
"مختلف": true,
"شود.": true,
"سر": true,
"همان": true,
"ای": true,
"چون": true,
"توجه": true,
"غربی": true,
"بن": true,
"نقش": true,
"نخستین": true,
"داشت": true,
"تمام": true,
"جمهوری": true,
"پایان": true,
"طور": true,
"همه": true,
"فارسی": true,
"گفته": true,
"می‌کنند.": true,
"کند.": true,
"زمین": true,
"بود،": true,
"سطح": true,
"علی": true,
"تیم": true,
"می‌باشد": true,
"جهانی": true,
"قبل": true,
"هجری": true,
"جهت": true,
"آن‌ها": true,
"اینکه": true,
"حزب": true,
"سیاسی": true,
"قابل": true,
"باعث": true,
"سپس": true,
"باشد": true,
"چهار": true,
"تحت": true,
"تغییر": true,
"دیگری": true,
"یعنی": true,
"تعداد": true,
"وارد": true,
"منتشر": true,
"داشت.": true,
"آذربایجان": true,
"ادامه": true,
"نسبت": true,
"بی": true,
"بودند": true,
"اشاره": true,
"۱۳۸۵": true,
"سوی": true,
"ساخت": true,
"غیر": true,
"داد": true,
"خواهد": true,
"معروف": true,
"بیست": true,
"قرن": true,
"انقلاب": true,
"مجموعه": true,
"داد.": true,
"فعالیت": true,
"حکومت": true,
"می‌تواند": true,
"جدید": true,
"خط": true,
"اطلاعات": true,
"جای": true,
"متر": true,
"سمت": true,
"تبدیل": true,
"محل": true,
"شروع": true,
"معمولاً": true,
"زمانی": true,
"نه": true,
"آمار": true,
"انتخاب": true,
"کنار": true,
"برنامه": true,
"ممکن": true,
"بدون": true,
"رشته": true,
"تاریخی": true,
"افزایش": true,
"سیستم": true,
"آنجا": true,
"می‌شوند.": true,
"گرفت.": true,
"طی": true,
"مجلس": true,
"افراد": true,
"دارند.": true,
"باز": true,
"مرگ": true,
"براساس": true,
"اعلام": true,
"دهه": true,
"حتی": true,
"علوم": true,
"طریق": true,
"مدت": true,
"مربوط": true,
"پیدا": true,
"روستای": true,
"بنا": true,
"مناطق": true,
"هیچ": true,
"ما": true,
"چه": true,
"دنیا": true,
"شناخته": true,
"دریافت": true,
"کشورهای": true,
"کوچک": true,
"آنان": true,
"غرب": true,
"جریان": true,
"دارند": true,
"ایالات": true,
"کمک": true,
"حضور": true,
"روش": true,
"نوشته": true,
"هنگام": true,
"عمل": true,
"چاپ": true,
"کوه": true,
"سید": true,
"شدند.": true,
"چنین": true,
"سپتامبر": true,
"درجه": true,
"بالا": true,
"۲": true,
"کیلومتر": true,
"فرانسه": true,
"موجود": true,
"کند": true,
"رنگ": true,
"رئیس": true,
"قانون": true,
"قدرت": true,
"سی": true,
"جنوبی": true,
"روستاهای": true,
"فوتبال": true,
"مهم": true,
"دور": true,
"نزدیک": true,
"حرکت": true,
"داستان": true,
"متحده": true,
"فرهنگ": true,
"کردند.": true,
"من": true,
"نمود.": true,
"معنی": true,
"نیروهای": true,
"کم": true,
"شهرهای": true,
"نیروی": true,
"سازی": true,
"۳": true,
"قسمت": true,
"طرف": true,
"آلمان": true,
"خارج": true,
"علت": true,
"بودن": true,
"آخرین": true,
"مسجد": true,
"مواد": true,
"می‌کنند": true,
"سایر": true,
"تقویم": true,
"خان": true,
"زنان": true,
"پسر": true,
"دیده": true,
"شود": true,
"خانواده": true,
"گرفت": true,
"خانه": true,
"علمی": true,
"آزاد": true,
"شمار": true,
"بودند.": true,
"آلبوم": true,
"پنج": true,
"سوم": true,
"عضو": true,
"طبیعی": true,
"تر": true,
"میلیون": true,
"کاهش": true,
"انسان": true,
"اسلام": true,
"فقط": true,
"نتیجه": true,
"طراحی": true,
"آموزش": true,
"شبکه": true,
"ایالت": true,
"بهترین": true,
"(به": true,
"گذشته": true,
"اکنون": true,
"گردید.": true,
"عمومی": true,
"نظامی": true,
"بین‌المللی": true,
"اجتماعی": true,
"انگلیسی": true,
"دومین": true,
"شورای": true,
"رسمی": true,
"نیاز": true,
"حالت": true,
"۵": true,
"ارائه": true,
"می‌دهد.": true,
"کننده": true,
"پدر": true,
"شمالی": true,
"سال‌های": true,
"۴": true,
"خانوار)": true,
"خاطر": true,
"کامل": true,
"قلعه": true,
"سرعت": true,
"بعضی": true,
"کنند.": true,
"بازار": true,
"پی": true,
"ریاست": true,
"ایران،": true,
"صد": true,
"اعضای": true,
"اروپا": true,
"ویژه": true,
"حاضر": true,
"می‌دهد": true,
"(در": true,
"یافت": true,
"شیخ": true,
"حمله": true,
"جایزه": true,
"نمایش": true,
"کردند": true,
"شرق": true,
"ترتیب": true,
"زن": true,
"بیماری": true,
"قمری": true,
"شد،": true,
"یافت.": true,
"فرهنگی": true,
"کیلومتری": true,
"جزیره": true,
"درصد": true,
"داخل": true,
"ابتدا": true,
"علاوه": true,
"آمریکایی": true,
"دکتر": true,
"اکتبر": true,
"انواع": true,
"بررسی": true,
"کنترل": true,
"آمده": true,
"وزارت": true,
"خورشیدی": true,
"عربی": true,
"حقوق": true,
"عراق": true,
"بخشی": true,
"۱۰": true,
"زمینه": true,
"البته": true,
"امروزه": true,
"انتخابات": true,
"توسعه": true,
"منبع.": true,
"مشهور": true,
"بدن": true,
"رو": true,
"مازندران": true,
"تقسیم": true,
"جان": true,
"منابع": true,
"کل": true,
"فرزند": true,
"سبک": true,
"ارتباط": true,
"حسین": true,
"اقتصادی": true,
"درباره": true,
"میدان": true,
"هنوز": true,
"هنر": true,
}