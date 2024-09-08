-- Insert poll
INSERT INTO polls (id, poll_num, title, options, feedbacks) VALUES (
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb', 
    1, 
    'Yo‘riqnoma: Ta’kidlariga mos kelsa “ha” mos kelmasa “yo‘q” javobini belgilab qo‘ying.', 
    '[{"ball": 1, "variant": "ha"}, {"ball": 0, "variant": "yo‘q"}]'::JSONB,
    '[{"from": 0, "text": "0-10 ball topladingiz", "to": 10}]'::JSONB
);

INSERT INTO questions (id, num, content, poll_id) VALUES 
(
    'b8aeedae-be34-464d-a016-0ca6755cf904', 
    1,
    'Kuchli hayajonni sinash uchun sizda yangilikka intilish seziladimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf905', 
    2,
    'Sizni tushunadigan va qo‘llab quvvatlaydigan do‘st kerakligini his qilasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf906', 
    3,
    'Siz o‘zingizni tashvishi yuk odam deb hisoblaysizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf907', 
    4,
    'O‘ylagan narsangizdan qiyinchiliksiz voz kecha olasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf908', 
    5,
    'Bir ishni qilishdan oldin shoshmasdan, o‘ylab ish tuta olasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf909', 
    6,
    'Xatto sizga foydasi bo‘lmasa ham, vaʼdangizda turasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf90a', 
    7,
    'Ruxingiz tez tushib, tez ko‘tariladimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf90b', 
    8,
    'Ko‘p uylamay tez ish yurita olasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf90c', 
    9,
    'Jiddiy sabab bo‘lmasa ham o‘zingizni baxtsiz deb uylaysizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf90d', 
    10,
    'Birov bilan baxslashganingizda o‘zingizni ko‘p narsaga qodir ekanligingizga ishonasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf90e', 
    11,
    'Sizga yoqqan qarama-qarshi jinsli kishi bilan tanishsangiz, tortinasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf90f', 
    12,
    'Jaxlingiz chiqqanda o‘zingizni uyg‘otasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf910', 
    13,
    'Ko‘pincha o‘ylamasdan, sharoitga qarab ish tutasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf911', 
    14,
    'Gapirmasam bo‘lar edi, shu ishni qilmasam bo‘lar edi, degan xayol sizni tez-tez bezovta qiladimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf912', 
    15,
    'Sizga kitob o‘qishga nisbatan odamlar bilan uchrashish afzalmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf913', 
    16,
    'Hamma narsani o‘zingizga olasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf914', 
    17,
    'Do‘stlar davrasida bo‘lishni yoqtirasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf915', 
    18,
    'Birovlar bilishni istamagan sirlaringiz bormi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf916', 
    19,
    'Baʼzan g‘ayratingiz jo‘shib qadamingizdan o‘t chaqnaydi, baʼzan esa hamma ishdan xavfsalangiz pir bo‘lib loqayd bulishingiz haqiqatmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf917', 
    20,
    'Do‘stlaringiz davrasini eng yaxshi do‘stlaringiz bilan chegaralashni istaysizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf918', 
    21,
    'Siz ko‘p narsani orzu qilasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf919', 
    22,
    'Sizga baqirib gapirishsa, shunday javob qaytara olasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf91a', 
    23,
    'Baʼzan o‘zingizni biron narsada aybdorman deb hisoblaysizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf91b', 
    24,
    'Siz o‘z odatlaringizni yaxshi ekanligiga ishonasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf91c', 
    25,
    'O‘z hislaringizga erk berib, do‘stlar davrasida o‘zingizni betashvish deb hisoblaysizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf91d', 
    26,
    'Asabingiz tarang bo‘lgan vaqtlar ko‘p bo‘lgannmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf91e', 
    27,
    'Sizni xushchaqchaq va tiyrak odam deb hisoblashadimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf91f', 
    28,
    'Biron ishni bajarib bo‘lgandan so‘ng, shu ishni bundan ham yaxshiroq bajarishingiz mumkinligi haqida o‘ylasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf920', 
    29,
    'Katta davralarda o‘zingizni hotirjam his qilasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf921', 
    30,
    'Siz g‘iybat qilasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf922', 
    31,
    'Miyangizga har xil xayollar kelib, sizga uyqu bermagan paytlar bo‘lganmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf923', 
    32,
    'Siz biron narsani bilishni istasangiz, do‘stlaringizdan so‘rab surishtirishga nisbatan kitobdan qidirib topishni maʼqul ko‘rasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf924', 
    33,
    'Sizni tez-tez yuragingiz o‘ynab turadimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf925', 
    34,
    'Diqqat eʼtiborni bir joyga qo‘yadigan ish sizga yoqadimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf926', 
    35,
    'Sizni qaltiroq tutadimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf927', 
    36,
    'Siz doim haq gapirasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf928', 
    37,
    'Bir-birini kamsitadigan davralarda o‘zingizni hotirjam his qilasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf929', 
    38,
    'Siz serjaxlmisiz?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf92a', 
    39,
    'Tez harakat qiladigan ish sizga yoqadimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf92b', 
    40,
    'Hammasi yaxshilik bilan tugagan, lekin yomon oqibatlarga olib kelishi mumkin bo‘lgann voqealar sizni xayolingizni bezovta qiladimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf92c', 
    41,
    'Xarakatingiz sekin, chaqqon emasligi haqiqatmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf92d', 
    42,
    'Siz ishga yoki birov bilan uchrashuvga kechikkanmisiz?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf92e', 
    43,
    'Beʼmani tushlar ko‘rasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf92f', 
    44,
    'Siz gaplashishni sevganingiz uchun, yangi odamlarni ko‘rganda paytni boy bermasdan u bilan gaplashishga harakat qilasiz?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf930', 
    45,
    'Sizni biror og‘riq bezovta qiladimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf931', 
    46,
    'Do‘stlaringiz bilan uzoq uchrashmasangiz xafa bo‘lasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf932', 
    47,
    'O‘zingizni asabi yomon kishi deb hisoblaysizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf933', 
    48,
    'Sizning tanishlaringiz orasida o‘zingizga yoqmagan insonlar ham bormi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf934', 
    49,
    'Siz o‘zingizni o‘ziga ishongan inson deb hisoblaysizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf935', 
    50,
    'Sizni kamchiliklaringiz va ishingizni tanqid qilishsa, shaxsiyatingizga tegadimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf936', 
    51,
    'Ko‘pchilik bilan bajarilgan ishdan qoniqish qiyinmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf937', 
    52,
    'Sizni boshqalardan nima bilandir yomonman, degan fikr bezovta qiladimi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf938', 
    53,
    'Zerikarli davrani qiziqtira olasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf939', 
    54,
    'O‘zingiz tushunmagan narsalar haqida gapirgan vaqtlaringiz bo‘lganmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf93a', 
    55,
    'O‘z sog‘ligingiz haqida qayg‘urasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf93b', 
    56,
    'Birovlar ustidan hazil qilishni yoqtirasizmi?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
),
(
    'b8aeedae-be34-464d-a016-0ca6755cf93c', 
    57,
    'Uyqusizlikka duchormisiz?',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
