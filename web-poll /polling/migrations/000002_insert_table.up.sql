-- Insert poll
INSERT INTO polls (id, poll_num, title, subtitle, options) VALUES (
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb', 
    1, 
    'Temperament',
    'Yo‘riqnoma: Ta’kidlariga mos kelsa “ha” mos kelmasa “yo‘q” javobini belgilab qo‘ying.',
    '[{"ball": 1, "variant": "ha"}, {"ball": 0, "variant": "yo‘q"}]'::JSONB
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
    'Siz o‘zingizni tashvishi yo`q odam deb hisoblaysizmi?',
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
    'Jiddiy sabab bo‘lmasa ham o‘zingizni baxtsiz deb o`ylaysizmi?',
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
    'Jaxlingiz chiqqanda o‘zingizni o`yg‘otasizmi?',
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
    'Baʼzan g‘ayratingiz jo‘shib qadamingizdan o‘t chaqnaydi, baʼzan esa hamma ishdan xavfsalangiz pir bo‘lib loqayd bo`lishingiz haqiqatmi?',
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


-- surovnoma - 2 /////////////////////////////////////////////////////////////////////////////////////////////

-- Insert the main poll information
INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES (
    'ff133d3b-6c03-4857-9c91-77ab72d200aa',
    2,
    'Imposter sindromi',
    'Mulohazalar sizga mos kelishiga qarab, raqamlardan birini aylana bilan belgilab qo''ying',
    '[{"ball": 1, "variant": "Hech qachon"}, {"ball": 2, "variant": "Kamdan kam"}, {"ball": 3, "variant": "Ba''zan"}, {"ball": 4, "variant": "Tez-tez"}, {"ball": 5, "variant": "Mutloqo to''g''ri"}]'::jsonb,
    '[{"to": 40, "from": 1, "text": "Sizda imposter sindromi juda past darajada. Siz o''z qobiliyatlaringizga ishonasiz. Hayotdagi muvaffasiyatlaringizga o''zingizni loyiq ko''ra olasiz. Ajoyib! "}, {"to": 60, "from": 41, "text": "Sizda imposter sindorimi  o''rta darajada mavjud. Ba''zan o''z qobiliyatlaringizga ishonchsizlik paydo bo''ladi. Ba''zida hayotda yangi tajribalarga qo''l urishdan oldin o''z kuchingizga ishonch yetishmay qoladi."}, {"to": 80, "from": 61, "text": "Sizda imposter sindorimi tez-tez kuzatiladi. Hayotda muvaffaqiyatlarga erishgan bo''lsangiz ham doim o''z qobiliyatlaringizga ishonchsiz qaraysiz.   Erishgan yutuqlaringizga o''zingizni loyiq emasdek his qilasiz. Biror ishni boshlashdan oldin sizni stress va xavotir qurshab oladi. O''z kuchingizga va qobiliyatlaringizga ishoning"}, {"to": 150, "from": 81, "text": "Sizda imposter sindromi juda yuqori. Imposter sindromi - bu psixologik holat bo''lib, unda odamlar o''zlarining yutuqlariga shubha qiladilar va o''zlarining qobiliyatlari va muvaffaqiyatlariga qaramay, qobiliyatsiz sifatida fosh bo''lishdan qo''rqishadi. Imposter sindromini boshdan kechirayotgan odamlar ko''pincha o''zlarining yutuqlariga loyiq emasligini his qilishadi va boshqalar oxir-oqibat ular ko''rinadigan darajada qobiliyatli emasligini bilib olishlaridan xavotirlanishadi. "}]'::jsonb
);

-- Insert the questions associated with this poll
INSERT INTO questions (id, num, content, poll_id)
VALUES
('5accb331-d548-4f4c-8e6d-3e1bfa3f3343', 1, 'Turli topshiriqlarni bajarishda muvaffaqiyat qozonganman, lekin, ularni bajarishdan avval buni qila olmayman degan qo''rquv bo''lar edi', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('06da654b-11c1-4c96-b2de-fe411b4b6253', 2, 'Men o''zimda mavjud qobiliyatdan ko''ra ko''proq qobiliyatli ekanligim haqida taassurot qoldira olaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('92817d59-e477-4e8e-86b8-337d5480208f', 3, 'Iloji bo''lsa baholanishdan qochaman va boshqalar meni baholashlaridan qo''rqaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('ac71fc63-7c11-4730-adb1-d22e9a438505', 4, 'Qachonki, odamlar meni erishgan narsam uchun  maqtashsa, kelajakda umidlarini oqlay olmayman, deb qo''rqaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('90fc6407-e8bb-47f0-8c81-ee0dcf042743', 5, 'To''g''ri vaqtda kerakli joyda bo''lganim va kerakli odamlarni bilganim uchun hozirgi mavqeyimga va mufavvaqiyatlarimga erishdim deb o''ylayman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('435ac740-0cf8-40b8-947c-656ba178f197', 6, 'Men uchun muhim odamlar ular o''ylagandek qobiliyatli emasligimni bilib qolishlaridan qo''rqaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('617181b0-7267-4e33-884c-985651a7996d', 7, 'Qo''limdan kelganicha harakat qilgan paytlarimda ham undan ham ko''proq harakat qilishim kerak edi deb eslab turaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('e876cdcb-7b86-4b5a-a921-a745c199cffe', 8, 'Men kamdan-kam hollarda loyiha yoki topshiriqni ko''nglimdagidek bajaraman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('8bcfc7f8-e232-448e-ad67-6fba48f5cc1c', 9, 'Hayotimdagi muvaffaqiyatlarim kimningdir e''tiborsizligi yoki omad natijasida bo''lgan deb hisoblayman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('734ae9b9-6a67-4627-b0f6-0dfdc58d95e3', 10, 'Intellektim yoki yutuqlarim haqidagi maqtovlar yoki e''tiroflarni qabul qilish men uchun qiyin', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('aa551337-33ee-4a98-a8c0-6d1036ccceb5', 11, 'Ba''zida muvaffaqiyatlarim qandaydir omad tufayli bo''lgan deb hisoblayman Ba''zida hozirgi yutuqlarimdan hafsalam pir bo''ladi va men ko''proq narsaga erishgan bo''lardim deb o''ylayman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('37cd686a-c26b-4a5b-87e6-55bf7cf69489', 12, 'Ba''zida menda bilim yoki qobiliyat yetishmasligini boshqalar bilib olishidan qo''rqaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('2085202c-4986-454c-b58a-02966689b550', 13, 'Odatda harakat qilgan narsamni yaxshi bajara olgan bo''lsam ham  ko''pincha yangi topshiriq yoki ishni bajara olmay qolishimdan qo''rqaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('afa81279-b755-409a-b071-0cb3abcc893e', 14, 'Men biror narsada muvaffaqiyatga erishganimda va yutuqlarim uchun eʼtirof etilganimda,  bu muvaffaqiyatni davom ettirishimga shubha qilaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('468bf927-2ce1-4cb7-acb0-4ebad2529d8b', 15, 'Agar men erishgan narsam uchun katta maqtov va eʼtirofga sazovor boʻlsam, qilgan ishim bu darajada e''tirofga loyiq deb hisoblamayman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('f3b77ce7-a4fa-49f3-8b74-681af7e1da46', 16, 'Men ko''pincha o''z qobiliyatimni atrofimdagilar bilan solishtiraman va ular mendan ko''ra aqlliroq bo''lishi mumkin deb o''ylayman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('add7e8ba-6f1c-45c8-a54d-b47d42094e79', 17, 'Garchi atrofimdagilar yaxshi uddalashimga ishonsalar ham men ko''pincha loyiha yoki imtihonda muvaffaqiyat qozona olmasligimdan tashvishlanaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('7bcffb74-2095-40d9-9366-34b3838cdcf4', 18, 'Agar men lavozimga ko''tarilmoqchi bo''lsam yoki biron bir e''tirofga sazovor bo''lmoqchi bo''lsam, buni amalga oshirilgunga qadar boshqalarga aytishga ikkilanaman', 'ff133d3b-6c03-4857-9c91-77ab72d200aa'),
('056e4dfc-475e-4411-95ba-75ed4429ecfc', 19, 'Agar men muvaffaqiyatga erishgan vaziyatlarimda "eng yaxshi" yoki hech bo''lmaganda "juda maxsus" deb e''tirof etilmasam, o''zimni yomon his qilaman va kayfiyatim tushadi', 'ff133d3b-6c03-4857-9c91-77ab72d200aa');


-- surovnoma - 3 /////////////////////////////////////////////////////////////////////////////////////////////

-- Asosiy so'rov ma'lumotlarini kiritish
INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES (
    '3e8be13b-056b-4dbf-942a-40e7619a5024',
    3,
    'Baholanishdan qo''rquv',
    'Mulohazalar sizga mos kelishiga qarab, raqamlardan birini belgilab qo''ying',
    '[{"ball": 1, "variant": "Hech qachon"}, {"ball": 2, "variant": "Kamdan kam"}, {"ball": 3, "variant": "Ba''zan"}, {"ball": 4, "variant": "To''g''ri"}, {"ball": 5, "variant": "Aynan shunday"}]'::jsonb,
    '[{"to": 40, "from": 1, "text": "Sizda boshqalar tomonidan hukm qilinish, tanqid qilinish va salbiy baholanishdan qo''rquv juda past darajada, Ajoyib!  Qolgan testlarni ham ishlashda davom eting."}, {"to": 47, "from": 41, "text": "Sizda salbiy baholanishdan qo''qruv mavjud bo''lib, boshqalar siz haqingizda qanday o''ylashlari, tanqid qilishlari va sizni salbiy baholashlari xavotirga soladi. Insonlar sizni qanday baholashsa ham o''zingizni boringizcha qabul qiling. Qolgan testlarni ishlashda davom eting."}, {"to": 100, "from": 48, "text": "Sizda salbiy baholanishdan qo''rquv darajasi yuqori ekanligi anqilandi. Boshqalar sizni muhokama qilishi, tanqid qilishlari va siz haqingizda salbiy baholashlari sizda kuchli xavotir hamda stressni yuzaga keltiradi. Bunday qo''rquv sizda ijtimoiy xavotir va o''ziga nisbatan ishonchsizlikni paydo qilishi mumkin. O''z qobiliyatlaringiz va yutuqli jihatlaringiz haqida o''ylab ko''ring. Qolgan testlarni ham ishlashda davom eting."}]'::jsonb
);

-- So'rovga tegishli savollarni kiritish
INSERT INTO questions (id, num, content, poll_id)
VALUES
('1b0ebfc9-b95e-4582-bff8-cff9fb6c58cc', 1, 'Xatto bu hech qanday farq qilmasligini bilsam ham, boshqalar men haqimda nima deb o''ylashlari haqida xavotirlanaman', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('96908f3e-727f-4171-9958-d5cace33cffb', 2, 'Odamlar men haqimda noxush taassurot  shakllantirishini bilsam ham tashvishlanmayman', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('a5d526de-1076-4c57-9658-7f0f6018e4f1', 3, 'Ko''pincha boshqalar mening kamchiliklarimni ko''rishidan qo''rqaman', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('6a026938-d80d-4fff-a2c6-b36af478896e', 4, 'Men kimgadir qanday taassurot qoldirayotganim haqida kamdan-kam tashvishlanaman', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('35daacfd-fc4f-4bf4-a4d1-2fcffcb1803d', 5, 'Boshqalar meni ma''qullamaydi deb qo''rqaman', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('b0b8db72-f560-4e11-b6c1-b1c496fe01c2', 6, 'Odamlar mendan ayb topishlaridan qo''rqaman', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('3db3718d-5130-4f66-875e-b7045e05c8b9', 7, 'Boshqalarning men haqimdagi fikrlari meni bezovta qilmaydi', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('3ccc2c0d-1593-4a26-875f-7a0a88216d4f', 8, 'Biror kishi bilan gaplashayotganda, ular men haqimda nima deb o''ylashlari mumkinligi haqida tashvishlanaman', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('ab92af30-c728-4bc9-91f7-6856f751db08', 9, 'Men odatda qanday taassurot qoldirishimdan xavotirlanaman', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('d18d3f27-7f4b-462c-a955-098642a841ae', 10, 'Agar kimdir meni muhokama qilayotganini bilsam, bu menga biroz ta''sir qiladi', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('e4a3318a-2bf6-428d-b5e1-e7d83d34553f', 11, 'Ba''zan men boshqalarning men haqimda o''ylashlari bilan juda qiziqaman deb o''ylayman', '3e8be13b-056b-4dbf-942a-40e7619a5024'),
('7ac2f472-d437-4f34-a85d-5cae4fee55bf', 12, 'Men ko''pincha noto''g''ri narsalarni aytaman yoki qilaman deb xavotirlanaman', '3e8be13b-056b-4dbf-942a-40e7619a5024');





-- surovnoma - 4 /////////////////////////////////////////////////////////////////////////////////////////////

-- Asosiy so'rov ma'lumotlarini kiritish
INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES (
    '4590ee7c-0870-497f-ae93-b1a2bb7ef58e',
    4,
    'Stress',
    'Mulohazalar sizga mos kelishiga qarab, raqamlardan birini belgilab qo''ying',
    '[{"ball": 1, "variant": "Hech qachon"}, {"ball": 2, "variant": "Mutloqo kamdan kam"}, {"ball": 3, "variant": "Juda kam"}, {"ball": 4, "variant": "Kam"}, {"ball": 5, "variant": "Ba''zan"}, {"ball": 6, "variant": "Tez-tez"}, {"ball": 7, "variant": "Juda tez-tez"}, {"ball": 8, "variant": "Doimo"}]'::jsonb,
    '[{"to": 99, "from": 1, "text": "Sizda stress past darajada. Psixologik jihatdan zo''riqish mavjud emas. Hayotingizdagi qiyinchiliklar yoki talablarni yenga olish ko''nikmasi shakllangan. Bunday vaziyatlar sizni stressga tushirmaydi. Qolgan testlarni ham ishlashda davom eting 😊"}, {"to": 125, "from": 100, "text": "Sizda stressning o''rta darajada ekanligi aniqlandi. Hayotingizda uchraydigan qiyinchiliklar sizni tashvishga soladi va buning natijasida sizda hissiy zo''riqish yuzaga keladi. O''zingizga dam bering. Tinchlantiruvchi mashg''ulotlar bilan shug''ullaning. Keyingi testlarni ham ishlashda davom eting 😊"}, {"to": 200, "from": 126, "text": "Sizda stress juda yuqori darajada ekanligi aniqlandi. Sizda quhiy va hissiy zo''riqish holati mavjud. Ish yoki hayotingizdagi muammolar sabab sizda stress yuqori darajada bo''lishi mumkin. Muammolarni hal etish yo''llarini o''ylab ko''ring. Ruhiy farovonlikni ta''minlaydigan mashg''ulot bilan shug''ullaning. Keyingi testlarni ham bajarishda davom eting 😊"}]'::jsonb
);

-- So'rovga tegishli savollarni kiritish
INSERT INTO questions (id, num, content, poll_id)
VALUES
('36c93e7e-00d7-4942-b85f-88cf6b0c92fe', 1, 'Men zoʻriqqan va hayajonlanganman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('a4ef0906-18d5-4e67-afe4-9f094807a2d9', 2, 'Mening boʻgʻzimda nimadir tiqilgandek  (ogʻzim qurib qolganini his qilaman)', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('4b15d375-ec13-4eba-a0ef-40c49ffd262d', 3, 'Men ish bilan oʻralashib qolganmanki, vaqt yetishmaydi', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('bf64410a-d01e-438f-bfc6-7d7a9f7cd912', 4, 'Men ovqatni chala chaynab yutaman yoki ovqatlanganimni unutib qoʻyaman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('e0230b3e-4560-4a47-a888-94f2eaf6f594', 5, 'Gʻoyalarim haqida qayta-qayta oʻylayveraman; oʻz rejalarimni oʻzgartiraman; fikrlarim doimo takrorlanaveradi', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('18cb490f-94b9-4704-bfa3-3f49b194679d', 6, 'Men oʻzimni oʻzim tushunmayman, oʻzimni yolgʻiz va ajralib qolgandek his qilaman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('eb1557ca-3de6-461b-a6d9-312681ba7e6b', 7, 'Men jismoniy darmonsizlikdan azob chekaman; boshim ogʻriydi, boʻynimdagi muskullar zoʻriqqandek, belimda ogʻriq bor, oshqozonim qisadi', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('a0acf835-bd16-4015-9d7d-3e2e50866514', 8, 'Men fikrlarimga oʻralashib qolganligim sababli undan azoblanaman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('f72683d0-09e2-47e9-aa17-4bf6c795bc85', 9, 'Menga tasodifan issiq va sovuq  taʼsir koʻrsatayotgandek boʻladi', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('65bb90d6-b09d-431e-b797-41282b548041', 10, 'Men bajarilishi lozim boʻlgan ish va uchrashuvlarni unutib qoʻyaman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('a2f8737a-c6d1-47d8-a34c-c00bcdc69d4e', 11, 'Men tez yigʻlab yuboraman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('91ac4936-f372-4f4d-9aaf-bec9b4f56b1f', 12, 'Men oʻzimni toliqqandek his etayapman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('c6b58422-a3fc-44c3-bc08-57696a4b7f2d', 13, 'Men tishlarimni mahkam  siqib olganman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('b4d080e0-3454-4187-8985-6b612af314b2', 14, 'Men xotirjam emasman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('19679a92-b933-4c16-b6e4-5cbb0ee42466', 15, 'Menga nafas olish qiyin yoki menda kutilmaganda havo yetishmay qoladi', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('3ac96874-09be-4c53-a98c-d7f6754aade9', 16, 'Mening ovqat hazm qilishim va ichaklarim bilan bogʻliq muammolarim bor (ogʻriq, sanchish, buzilish yoki ich qotishi )', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('8f971af9-54cf-44a7-9f16-53bc04f42127', 17, 'Men hayajondaman, behalovatman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('903b1923-ff9b-4afd-819e-beb0bd52486d', 18, 'Men tez choʻchiyman, shovqun va shitirlashlardan seskanib ketaman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('5c6762a6-230b-45fc-a821-1eefbadae5ad', 19, 'Menga uxlab qolishim uchun yarim soatdan ortiq vaqt kerak boʻladi', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('d541f392-207c-4dc7-9cfe-95f00260cdf0', 20, 'Men fikrlarim bilan chalkashib qolaman, mening fikrlarim chalkash; diqqatimni bir joyga toʻplay olmayman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('3f58618b-37ea-4002-80f9-08aca844d0a8', 21, 'Men toliqqandek, koʻzlarim osti salqigan va shishgandek koʻrinishdaman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('903c4cc7-2fb7-4a9a-918b-d9291e6acec1', 22, 'Yelkamdan yuk bosgandek his qilayapman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('b3881ea4-aa3c-44a0-8738-6e9f10e259c0', 23, 'Men bezovtaman.Shu bois bir joyda tura olmayman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('7846f0ea-6174-448c-9a04-440891319241', 24, 'Menga xatti-harakatlarim, imo-ishoralarim va emotsiyamni nazorat qilish qiyin', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e'),
('ce1a98e6-5243-46ae-88a5-7f616e5a8407', 25, 'Men zoʻriqishdaman', '4590ee7c-0870-497f-ae93-b1a2bb7ef58e');


-- surovnoma - 5 /////////////////////////////////////////////////////////////////////////////////////////////

INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES
    (
        'd4a3b6fa-98ad-4442-91b1-83a95edf6167',
        5,
        'Perfektsionizm',
        'Takidlar sizga mos kelsa “ha” mos kelmasa “yo‘q” javobini belgilab qo‘ying',
        '[{"ball": 1, "variant": "Ha"}, {"ball": 0, "variant": "Yo`q"}]'::jsonb,
        '[{"to": 4, "from": 1, "text": "Siz perfektsionist emassiz. Lekin, ayrim bir topshiriqlar va vazifalarni bajarishda e‘tiborliroq bo‘ling. Keyingi testlarni ham bajarishda davom eting 😊"}, {"to": 10, "from": 5, "text": "Siz o‘rta darajada perfektsionistsiz. Ayrim vazifalarni bajarishda mukammallikka ko‘p e‘tibor qaratasiz. Bu esa sizni vazifalarni bajarishda kechikishingizga olib kelishi mumkin. O‘z psixik holatlaringiz haqida bilish uchun qolgan testlarni ham bajarishda davom eting😊"}, {"to": 16, "from": 11, "text": "Siz yuqori darajada perfeksionistsiz. Topshiriqlar va vazifalarni bajarish sizda doim ortiqcha kuch sarflashni talab etadi. O‘zingizga qo‘ygan talabingiz yuqoriligi uchun bajarishingiz kerak bo‘lgan ishlarni ideal darajada qilishga urinasiz. Bu esa sizda turli emotsional qiyinchiliklarni keltirib chiqaradi. O‘z psixik holatingiz haqida ko‘proq bilish uchun qolgan testlarni ham bajarishda davom eting😊"}]'::jsonb
    );


INSERT INTO questions (id, num, content, poll_id)
VALUES
    ('899bccbf-e632-43cc-97d6-057eb5ea5d52', 1, 'Men ko‘pincha bajarganimdan ko‘ra yaxshiroq qilishim kerak edi deb hisoblayman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('95184a00-e7fa-4bb6-9741-d03a67c58765', 2, 'Agar ularni mukammal bajarish uchun vaqtim bo‘lmasa, men ularni kechiktirishga moyilman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('d8ef0320-920c-4212-9cd0-7a6c5196873b', 3, 'Muhim loyiha ustida ishlayotganda muvaffaqiyatsizlikka uchrashdan qo‘rqaman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('4effff02-8349-4be8-a161-6be0116d99e1', 4, 'Men boshqalarni eng yaxshi fazilatlarim yoki yutuqlarim bilan hayratda qoldirishga intilaman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('2a893283-e7df-4992-94fc-8db691d8b521', 5, 'Agar biror xatoni takrorlasam o‘zimni past baholayman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('cb7477ca-d52f-4c6e-8212-af09613dff53', 6, 'Men har doim his-tuyg‘ularimni nazorat qilishga intilaman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('82130529-7ae4-45ee-b100-eb1fb800568d', 7, 'Ishlar rejalashtirganimdek bo‘lmasa, bundan kayfiyatim tushadi', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('33c31118-a928-49f9-9bfd-46738cef9313', 8, 'Men ko‘pincha boshqa odamlarning bajargan ish sifatidan ko‘nglim to‘lmaydi', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('6a94a5ef-4d2b-4d5e-b2be-5e6d7d0c655a', 9, 'Mening talablarim juda yuqori deb o‘ylamayman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('b34e35f9-3893-4f6a-8741-ecce6d975b7f', 10, 'Agar muvaffaqiyatsizlikka uchrasam odamlar meni past baholashlaridan qo‘rqaman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('721064ef-8964-4cd9-9fa3-392d69f1285f', 11, 'Men doimo o‘zimni rivojlantirishga harakat qilaman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('b665e576-5b46-4511-8139-b3c94d4f765f', 12, 'Agar men bajargan ishim o‘rta darajada baholansa kayfiyatim tushadi', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('3e6f7040-15c0-46da-9154-0f915e1b33b6', 13, 'Mening uyim va ofisim har doim toza va tartibli bo‘lishi kerak', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('7b4fc386-c163-4b66-980c-2e0f664aea83', 14, 'Mendan ko‘ra aqilli, jozibali, muvaffaqiyatli bo‘lgan odamlar oldida o‘zimni past deb his qilaman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167'),
    ('5286483f-caf5-4797-b516-de25bfd1ee71', 15, 'Odamlar orasida bo‘lganimda o‘zimni imkon qadar eng yaxshi darajada tutishga harakat qilaman', 'd4a3b6fa-98ad-4442-91b1-83a95edf6167');



-- surovnoma - 6 /////////////////////////////////////////////////////////////////////////////////////////////

INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES
    (
        'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7',
        6,
        'Ish va hayot balansi',
        'Hayotingiz va ish muvozanati haqida baho bering',
        '[{"ball": 1, "variant": "To''g''ri"}, {"ball": 0, "variant": "Noto''g''ri"}]'::jsonb,
        '[{"to": 2, "from": 1, "text": "Sizda ish va hayot o''rtasidagi balans me''yorida. Ishdagi mas''uliyatlaringizni shaxsiy hayotingizga aralashtirib yubormagansiz. Bu juda yaxshi. Psixologik holatingiz haqida ko''proq ma''lumotga ega bo''lish uchun testlarni bajarishda davom eting😊"}, {"to": 5, "from": 3, "text": "Sizda ish va shaxsiy hayot o''rtasida muvozanat buzilgan. Ishdagi mas''uliyatlaringizni shaxsiy hayotingizga aralashtirib yubormang. Ishdagi topshiriqlar uchun chegarani belgilab olishingiz kerak. Shaxsiy hayotingizda qilishni rejalashtirgan maqsadlaringiz haqida qayta o''ylab chiqishingiz kerak. Psixologik holatingiz haqida ko''proq ma''lumotga ega bo''lish uchun testlarni bajarishda davom eting 😊"}, {"to": 16, "from": 6, "text": "Sizning shaxsiy hayotingiz va ishingiz o''rtasida muvozanat jiddiy darajada buzilgan. Ish shaxsiy hayotingiz o''rini egallab olgan. Ishdagi topshiriqlar chegarasini jiddiy belgilab olish kerak. Ish bilan bog''liq bo''lgan muammolardan uzoqlashib, shaxsiy hayotingizdagi asosiy maqsadlarni yodga olishingiz psixologik salomatligingiz yaxshilanishiga yordam beradi. Psixologik holatingiz haqida ko''proq ma''lumotga ega bo''lish uchun testlarni bajarishda davom eting😊"}]'::jsonb
    );


INSERT INTO questions (id, num, content, poll_id)
VALUES
    ('318514c4-fc52-44bd-bec3-915c2e8173d0', 1, 'Ish bilan bogʻliq boʻlgan loyihalarga koʻpdan koʻp vaqt sarflayman', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('cc6759fb-84a6-4cff-bcf0-8d47d768165e', 2, 'Menda koʻpincha oʻzim, oilam va doʻstlarim uchun vaqtim yoʻq bo‘ladi', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('ba544a95-dbf8-4122-a6b3-70f29fcb7c66', 3, 'Nima qilsam ham, koʻpincha har bir kunning har bir daqiqasi doimo biror narsaga rejalashtirilganga oʻxshaydi', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('49453abe-fad4-4b94-a6da-a4fce523fbec', 4, 'Ba''zan oʻzimni kimligimni va nima uchun bu ishni/karyerani tanlaganimni tushunmay qolaman', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('eadba422-d970-4c70-bc6f-1e99e85ede34', 5, 'Oxirgi marta oʻzim uchun qachon dam olishga vaqt topa olganimni eslay olmayman', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('599d78d7-123a-445d-8b95-7474456af59f', 6, 'Koʻp hollarda stressga duch kelaman', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('5f031418-cd20-462f-b1a5-28ae4c778888', 7, 'Oxirgi marta qachon dam olish va shaxsiy ishlarim uchun ajratgan vaqtimni ham eslay olmayman', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('190bf5dc-1bf6-453a-ab40-63a9a2857942', 8, 'Eski loyihani tugatib, yangi loyiha boshlashdan oldin nafas olishga ham vaqtim bo‘lmay qoladi', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('067cda6f-e919-481d-9d10-26a1b4908393', 9, 'Faqat zavq uchun oʻqigan kitobni oxirgi marta qachon oʻqiganimni va tugatganimni eslay olmayman', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('0bc50967-8f63-4c66-994f-a2c2ac1b30e5', 10, 'Qaniydi tashqarida qiziqarli koʻngil ochar mashgʻulotlar va sevimli hobbilarim bilan shugʻullanish uchun koʻproq vaqtim bo‘lsaydi, afsuski unday emas', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('b1f29345-4be7-4630-b17a-ed8e2eba04ea', 11, 'Haftaning boshlanishidanoq charchoq his qilaman', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('30477523-d10b-44fb-aa03-e03868df113e', 12, 'Men oxirgi marta qachon kinoga borganimni yoki muzeyga tashrif buyurganimni yoki boshqa madaniy tadbirda qatnashganimni eslay olmayman', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('1ad07e62-2a3e-4610-bb1c-5b7590a7cc9f', 13, 'Ishda vazifalarimni bajarishim shart, chunki koʻp odamlar (bolalar, hamkasblar, ota-onalar) mening koʻmagimga muhtoj', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('f94247f1-b5a7-4664-a925-d8ec5f6e5408', 14, 'Men ish bilan bogʻliq vaqt bosimi va mas''uliyatlarim tufayli oilamdagi ko‘plab muhim voqealarni o‘tkazib yubordim', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7'),
    ('7f25098d-4e4d-4429-a67b-45071747d30b', 15, 'Men deyarli har doim oʻzim bilan uyga ishimdagi yumushlarni olib kelaman', 'e0a68ebb-0f1d-4e7e-84d3-6fd060059ab7');



-- surovnoma - 7 /////////////////////////////////////////////////////////////////////////////////////////////

INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES
    (
        '5c62eebf-a35a-487a-af91-ef40a3723db6',
        7,
        'Yolg`izlik',
        'Mulohazalar sizga mos kelishiga qarab javoblardan birini belgilang',
        '[{"ball": 1, "variant": "Hech qachon"}, {"ball": 2, "variant": "Kamdan kam"}, {"ball": 3, "variant": "Ba''zan"}, {"ball": 4, "variant": "Tez-tez"}]'::jsonb,
        '[{"to": 44, "from": 1, "text": "Siz o''zingizni yolg''iz his etmaysiz. Sizda shaxslararo munosabatlar yaxshi darajada. Maroqli vaqt o''tkazish uchun yetarlicha do''stlarga egasiz. Bu juda yaxshi😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 59, "from": 45, "text": "Siz vaqti vaqti bilan o''zingizni yolg''iz his etasiz. Shaxslararo munosabatlarni yaxshi tashkillasangiz ham qalbingizga yaqin insonni topishingiz qiyin bo''ladi. Ba''zan dildan suhbat qurish uchun hamroh yetishmaydi. Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 90, "from": 60, "text": "Sizda yolg''izlik darajasi kuchli. O''zingizni ko''p vaziyatlarda yolg''iz his etasiz. Hayotingizdagi qiyinchiliklarni yengib o''tish paytida dardlashadigan insoningiz yo''q. Yolg''izlik hissini yo''qotish uchun qiziqishlaringiz va dunyoqarashingiz mos keladigan do''stlar orttirishdan qo''rqmang. Ba''zan ijtimoiy yordam so''rash mustahkam do''stlikka ham sabab bo''lishi mumkin😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}]'::jsonb
    );


INSERT INTO questions (id, num, content, poll_id)
VALUES
    ('ac51ec3c-bec2-4128-97d2-6ca8ecdb5804', 8, 'Qanchalik sizning qiziqishlaringiz va g‘oyalaringiz atrofingizdagilar tomonidan qiziqish uyg‘otmagandek tuyiladi?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('55b54c0a-9cea-4b94-b633-7a095a73a73d', 9, 'Qanchalik o‘zingizni do‘stona va ochiq his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('0814a572-bf92-4e4b-b4e6-503e75323e94', 10, 'O‘zingizni qanchalik odamlarga yaqin his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('4ed7c99b-41e6-43a7-8b13-d60ea878e060', 11, 'Qanchalik o‘zingizni chetda qolgandek his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('27acea1b-b356-4729-9744-5ae3a4920c62', 12, 'Qanchalik boshqalar bilan munosabatlaringiz ma’nosizdek tuyiladi?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('87ec428a-c68b-4a46-9902-4bf9da1c6603', 13, 'Qanchalik boshqalar sizni yaqindan bilmaydigandek his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('36bc8904-fd26-45e4-ad29-97c012bda253', 14, 'O‘zingizni qanchalik boshqalardan ajralgan his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('c5b79a36-86d4-4883-803f-1274e9a0fc91', 15, 'O‘zingiz xohlagan payt vaqt o‘tkazish uchun sherik topa olasizmi?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('cf34a8f5-b8d7-41bc-91cd-13fb44ebb525', 16, 'Qanchalik sizni haqiqatdan ham tushuna oladigan odamlar borligini his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('eaf4bae5-dd42-40f4-84d6-71fd3b5b100a', 17, 'Qanchalik uyalasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('00a896bb-9427-4d48-9d96-558527c830d5', 18, 'Qanchalik atrofingizdagi odamlar siz bilan birga emasliklarini his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('8a3a9da1-740a-4e0f-b070-f6503296a4a8', 19, 'Qanchalik gaplashishingiz uchun odamlar borligini his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('09c8bfd3-85f7-47a1-b4f8-a124316751db', 20, 'Qanchalik siz dildan suhbatlashadigan odamlar borligini his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('676faf22-74a6-4069-ac3b-a2441519af28', 1, 'Qanchalik o‘zingizni atrofdagi odamlar bilan bir to‘lqindadek his qilasiz', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('6f42614c-ad16-45e0-a736-c6ae2d8c6521', 2, 'Sizga do‘stlik yetishmayotgandek tuyiladimi?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('b304574b-6f5c-4f75-bde9-8b2834074da8', 3, 'Ko‘nglingizni ochish uchun odam yo‘qligini his qilasizmi?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('b0a61b3c-74b1-426c-8810-05048872e941', 4, 'Qanchalik o‘zingizni yolg‘iz his qilasiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('5fe21268-2d28-40b9-9fdd-1e7f6721c970', 5, 'Qanchalik o‘zingizni do‘stlar guruhining bir qismi deb hisoblaysiz?', '5c62eebf-a35a-487a-af91-ef40a3723db6'),
    ('e208dbba-f24a-43ce-8a90-0440cfd1f673', 6, 'Atrofingizdagi odamlar bilan o‘zingizda umumiylik bor deb o‘ylaysizmi?', '5c62eebf-a35a-487a-af91-ef40a3723db6');



-- surovnoma - 8 /////////////////////////////////////////////////////////////////////////////////////////////

INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES
    (
        '3cbe1c37-91cb-4780-9d52-2ad451869c94',
        8,
        'Emotsional kuyish',
        'Mulohazalar sizga mos kelishiga qarab, birini belgilab qo‘ying',
        '[{"ball": 1, "variant": "Hech qachon"}, {"ball": 2, "variant": "Kamdan kam"}, {"ball": 3, "variant": "Ba''zan"}, {"ball": 4, "variant": "Tez-tez"}, {"ball": 5, "variant": "Juda tez-tez"}]'::jsonb,
        '[{"to": 18, "from": 1, "text": "Sizda emotsional kuyish mavjud emas. Sizda emotsional zo''riqish holati juda kam kuzatiladi. Juda ajoyib😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq bilib oling"}, {"to": 32, "from": 19, "text": "Sizda emotsional kuyish past darajada. Ba''zan emotsional jihatdan toliqish va charchoq kuzatiladi. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq bilib oling"}, {"to": 49, "from": 33, "text": "Sizda emotsional kuyish xavfi mavjud. Ish va hayotiy muammolar sizda ruhiy jihatdan charchoqni yuzaga keltiradi. Emotsional jihatdan farovonlikka erishish uchun o''zingizga e''tiborli bo''ling. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq bilib oling"}, {"to": 59, "from": 50, "text": "Sizda emotsional kuyish xavfi katta. Ish va hayotingizdagi muammolar natijasida tez-tez emotsional jihatdan kuchli charchoq yuzaga keladi. Bu esa ish va shaxslararo munosabatlarni to''g''ri tashkillashtirishga halaqit beradi. Ruhiy holatingizni yaxshilash uchun hotirjamlik beradigan mashg''ulotlar bilan shug''ullaning. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq bilib oling"}, {"to": 80, "from": 60, "text": "Siz emotsional jihatdan kuyish yuqori darajada. Sizda kuchli ruhiy zo''riqish holati mavjud. Ish va hayotingizdagi mas''uliyatlardan ozroq chekinib, o''z ruhiy salomatligingiz haqida qayg''urishingiz kerak. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq bilib oling"}]'::jsonb
    );


INSERT INTO questions (id, num, content, poll_id)
VALUES
    ('14ab90ea-a224-4fc5-bf70-2f5e9dddd5a0', 1, 'Men toliqqan va jismoniy yoki emotsional jihatdan energiyam tugaganday his qilaman', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('0717dd4b-f1f7-4be9-9550-531872f2711d', 2, 'Menda ishimga nisbatan salbiy fikrlar bor', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('b8def531-98e6-4c73-b7e0-eafe8aec7455', 3, 'Men odamlarga ular loyiq bo‘lganidan ko‘ra kamroq hamdardman va sovuq munosbatda bo‘laman', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('a2bec953-89a9-4818-a21d-b57d760818ff', 4, 'Kichkina muammolar yoki hamkasblarim tufayli tezda bezovtalanman', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('a11a6e08-8285-42a1-a7e3-c5bad2e9f071', 5, 'Men hamkasblarim tomonidan noto‘g‘ri tushunilgan yoki qadrlanmagan his qilaman', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('14cb3b4c-015e-4de9-9b69-51626807be95', 6, 'Men bilan gaplashadigan hech kim yo‘q', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('2b5d34c3-55ee-4402-a4ac-f8f9e56add6a', 7, 'Men erishishim kerak bo‘lganidan ko‘ra kamroq narsaga erishyapman', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('8ff0c1ab-e602-48e4-8fd0-fe25f359183d', 8, 'Muvaffaqiyatga erishish uchun yoqimsiz bosim ostida ekanligimni his qilaman', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('8182e049-6629-48d4-8e39-4236a21c292b', 9, 'Men o‘z ishimdan xohlagan narsamni ololmayotganimni his qilyapman', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('28566508-92bb-47bd-bdcb-7addbed01f9e', 10, 'Men o‘zimni noto‘g‘ri tashkilot yoki kasbdaman deb hisoblayman', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('491696d2-863c-4b26-b528-fa8929372e64', 11, 'Ish faoliyatimning ayrim holartlarida umidsizlikka tushaman', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('60e27964-177c-4bab-9617-182f7c35676c', 12, 'Ishdagi byurokratiya va tashkiliy masalalar samarali ish faoliyatimga putur yetkazadi', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('4bda909d-a4d0-4fe3-b5b2-3893d9a4b23d', 13, 'Qobiliyatimdan ortiqroq ish qilishimga to‘g‘ri kealdi', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('1410b5a3-d27f-49d0-b664-029e6641541e', 14, 'Yaxshi sifatli ishni bajarish uchun muhim bo‘lgan ko‘p narsalarni qilishga vaqtim yetmaydi', '3cbe1c37-91cb-4780-9d52-2ad451869c94'),
    ('a29062ba-a5c9-4678-821b-26e7c53453fd', 15, 'Ishlarimni reja qilish uchun o‘zim xohlagancha vaqt ajrata olmayman', '3cbe1c37-91cb-4780-9d52-2ad451869c94');


-- surovnoma - 9 /////////////////////////////////////////////////////////////////////////////////////////////

-- Polls jadvaliga ma'lumot qo'shish
INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES
    (
        '274ad998-e1bc-471a-8d97-196fb9fea8f7',
        9,
        'Kasbiy yangilik',
        'Mulohazalar sizga mos kelishiga qarab, javoblardan birini belgilab qo‘ying',
        '[{"ball": 1, "variant": "Hech qachon"}, {"ball": 2, "variant": "Kamdan kam"}, {"ball": 3, "variant": "Ba''zan"}, {"ball": 4, "variant": "Tez-tez"}, {"ball": 5, "variant": "Har doim"}]'::jsonb,
        '[{"to": 30, "from": 1, "text": "Sizning kasbingizda yangi bilimlarni o''rganish zarurati juda past. Mavjud bilimlar bilan kasbiy faoliyatingizni davom ettirish yetarli bo''ladi 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling."}, {"to": 44, "from": 31, "text": "Siz ishlayotgan sohada yangi bilimlarni o''zlashtirish zarurati o''rta darajada. Ba''zan yangi bilimlarni o''rganish shart bo''lib qoladi😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling."}, {"to": 79, "from": 45, "text": "Siz ishlayotgan sohada yangi bilimlarni o''zlashtirish juda muhim. Doimiy tajriba va bilimlarni oshirib borish kasbiy o''sishingizni ta''minlaydi😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling."}]'::jsonb
    );

-- Questions jadvaliga ma'lumot qo'shish
INSERT INTO questions (id, num, content, poll_id)
VALUES
    ('f9362fae-3b2b-40a2-84d5-7cbb3f5c8fd2', 5, 'Sohangizda yangiliklarga qanchalik tez duch kelasiz?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('a5fb50d2-33e7-4547-a4bc-1fa071300a97', 6, 'Kasbingizda siz umuman tushunmaydigan ma’lumotlarga qanchalik tez duch kelasiz?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('398e5ceb-ff5c-4154-b874-f172d4b196f9', 7, 'Qanchalik o‘zingizni yangi texnologiyalar qurshovida qolgandek his qilasiz?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('ebcd67ea-7a68-46ab-9d2c-6e9da9332654', 8, 'Sohangizdagi yangi texnalogiyalarni va bilimlarni o‘zlashtirishda o‘zingizga ishonch darajangizni qay darajada baholaysiz?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('9647736c-993b-4af9-ab39-005a2cb2bcb2', 9, 'Sohangizdagi yangiliklar sizni qanchalik charchatadi?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('dd5c0598-278a-4a9d-b440-b09055df9648', 10, 'Hamkabslaringiz orasida sohangizdagi yangiliklarni tushunmaydigan odamlarga qanchalik tez duch kelasiz?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('1bc8687c-6292-4850-a4c4-bfe77932ee32', 11, 'Kasbiy faoliyatingiz doirasida qo‘shimcha kurslardan qanchalik tez olib turasiz?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('b8fd0708-0fca-478d-bcbd-cd8a798a16d0', 12, 'Siz faoliyat olib borayotgan ishxona tomonidan sohangizda bo‘layotgan texnologik o‘zgarishlar bilan hamnafas bo‘lish uchun yordam yoki manbalar qanchalik tez taqdim etiladi?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('6dd61723-6f98-4e38-9c7e-1b076c85830d', 13, 'Sohangizdagi o‘zgarishlardan qanchalik stressga tushasiz?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('1eca1de9-9168-4261-a6e9-b292aed46d06', 14, '“Yana nimadir o‘rganishim kerakmi?” degan fikrga qanchalik tez duch kelasiz?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('f83cfdda-d10e-4ece-abe8-86b6d5f8450a', 15, 'Yangiliklar sababli sohangizni tark etib ketayotgan hamkasblaringiz ham uchrab turadimi?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('9e8bf729-3054-4f29-ae60-028babe88bac', 1, 'Sohangizdagi uchraydigan texnologik o‘zgarishlarni qay darajada baholaysiz?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('7f4c5b9b-d9c1-4502-9292-85090d3c3c0e', 2, 'Texnologik taraqqiyot tufayli mahoratingiz eskirib borayotgandek tuyiladimi?', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('8bd5d810-e4f6-4ee1-af6b-8a3962962255', 3, 'Quyidagi fikrga rozilik darajangizni baholang: "Doimiy o‘rganish va yangi texnologiyalarga moslashish dasturiy ta''minotni ishlab chiqishda muvaffaqiyatga erishish uchun muhim"', '274ad998-e1bc-471a-8d97-196fb9fea8f7'),
    ('ae9b0dcf-9335-4210-91ec-b23ce695c25a', 4, 'Sohangizda yangi bilimlarni o‘rganish qanchalik zarur?', '274ad998-e1bc-471a-8d97-196fb9fea8f7');



-- surovnoma - 10 /////////////////////////////////////////////////////////////////////////////////////////////

-- Polls jadvaliga ma'lumot qo'shish
INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES
    (
        '48e56a48-f31d-46fc-a157-7e60e03a0e68',
        10,
        'Raqobat',
        'Mulohazalarni o''qing va o''zingizga mos javobni belgilang',
        '[{"ball": 1, "variant": "Hech qachon"}, {"ball": 2, "variant": "Kamdan kam"}, {"ball": 3, "variant": "Ba''zan"}, {"ball": 4, "variant": "Tez-tez"}, {"ball": 5, "variant": "Har doim"}]'::jsonb,
        '[{"to": 17, "from": 1, "text": "Siz faoliyat olib borayotgan sohada ish o''rni uchun raqobat juda past. Hozirgi kunda egallab turgan lavozimingizni saqlab qolish uchun raqobat oqimiga duch kelmaysiz 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 27, "from": 18, "text": "Siz faoliyat olib borayotgan ish o''rnida raqobat mavjud. Hozirgi paytda ish joyingizni saqlab qolish uchun mahoratlaringizni rivojolantirishingiz zarur.😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 60, "from": 28, "text": "Siz faoliyat olib borayotgan sohada ish o''rni uchun raqobat juda yuqori. Shu bois mazkur ish o''rningizni saqlab qolish uchun qo''shimcha mahoratlaringizni rivojlantirishingiz talab etiladi😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}]'::jsonb
    );

-- Questions jadvaliga ma'lumot qo'shish
INSERT INTO questions (id, num, content, poll_id)
VALUES
    ('d57bc1e5-1906-43e3-bc84-fe9cbabc98db', 1, 'Faoliyat olib borayotgan sohangizdagi arzimas kamchiliklarga yo‘l qo‘yish ishingizdan ayrilishingizga sabab bo‘ladimi?', '48e56a48-f31d-46fc-a157-7e60e03a0e68'),
    ('8bc4f0da-b9a8-4b79-93a1-b4f8b5a486e5', 2, 'Ish joyingizda hamkasblaringizdan ko‘ra yaxshiroq natija ko‘rsatish muhimmi?', '48e56a48-f31d-46fc-a157-7e60e03a0e68'),
    ('f270affa-12f0-404f-86d4-129f52c85583', 3, 'Siz ishlayotgan lavozimda qolish uchun yangi bilimlarni o‘zlashtirish qanchalik muhim?', '48e56a48-f31d-46fc-a157-7e60e03a0e68'),
    ('4e66bcd5-71fd-4706-8610-206a7728f552', 4, 'Ishxonangizda yangi ish joyi uchun talabgorlar ko‘plab uchrab turadimi?', '48e56a48-f31d-46fc-a157-7e60e03a0e68'),
    ('9a04f674-2cc4-46e1-9ac2-228e4fd01b18', 5, 'Talabgorlar ko‘pligidan ish joyingizdan ayrilib qolish qo‘rquvi bo‘ladimi?', '48e56a48-f31d-46fc-a157-7e60e03a0e68'),
    ('53fc3e38-d6a8-4e6f-a5c1-58d609c32860', 6, 'Ishdagi raqobat muhitidan charchab stressga tushgan paytlaringiz bo‘ladimi?', '48e56a48-f31d-46fc-a157-7e60e03a0e68'),
    ('386d135d-11eb-42f5-bd40-fcb5b9949711', 7, 'Ishdagi loyiha yoki boshqa topshiriqlarni boshqalardan ko‘ra yaxshiroq bajarishim kerak degan fikrlar sizni qiynab turadimi?', '48e56a48-f31d-46fc-a157-7e60e03a0e68'),
    ('6ccac93f-6936-43c7-8f56-0a2b530440c0', 8, 'Sohangizda ishga kirish uchun ko‘plab talabgorlar bilan raqobatlashishga to‘g‘ri keladimi?', '48e56a48-f31d-46fc-a157-7e60e03a0e68'),
    ('c4ccd3f4-0253-4fdc-82c6-39c961249aaf', 9, 'Raqobatlashadigan talabgorlar mavjudligi sabab ko‘plab mahoratlarga ega bo‘lishga intilasizmi?', '48e56a48-f31d-46fc-a157-7e60e03a0e68'),
    ('73a9be15-e6bd-4980-838d-cdfe7da35b9e', 10, 'Malakali talabgorlar ko‘pligi sababli ishga rad javobini oladigan dasturchilarga duch kelib turasizmi?', '48e56a48-f31d-46fc-a157-7e60e03a0e68');



-- surovnoma - 11 /////////////////////////////////////////////////////////////////////////////////////////////

-- Polls jadvaliga ma'lumot qo'shish
INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES
    (
        'b3e2b18a-1997-4157-bec2-f05bff14373b',
        11,
        'Iroda kuchi',
        'Mulohazalarni o''qing va o''zingizga mos javobni belgilang',
        '[{"ball": 2, "variant": "Ha"}, {"ball": 1, "variant": "Bo''lib turadi"}, {"ball": 0, "variant": "Yo''q"}]'::jsonb,
        '[{"to": 12, "from": 1, "text": "Sizning irodangiz kuchsiz. Sizni osongina ishontirish mumkin. Ammo, o‘zining «ojiz joylarini» bilish insonni yanada kuchliroq qiladi. O‘zingizning irodaviy  boshqarishingizni takomillashtirib, o‘zingiz ustingizda ishlang. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 21, "from": 13, "text": "Sizda iroda kuchi o‘rtacha. Siz har xil holatlarda turlicha harakat qilasiz, ayrim hollarda yon berishni, ayrim hollarda esa qa’tiyatlilik va sabr-bardoshlilikni namoyon qilasiz. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 40, "from": 22, "text": "Sizda iroda kuchi yuqori. Siz irodali odamsiz. Lekin, esingizda bo‘lsin siz yolg‘iz emassiz, yon berishlar, kelishishlar va o‘zaro munosabatlar  ham muloqot  davaishda muvaffaqiyatlarga olib keladi. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}]'::jsonb
    );

-- Questions jadvaliga ma'lumot qo'shish
INSERT INTO questions (id, num, content, poll_id)
VALUES
    ('fdca16b6-8413-446e-8d66-1e151770e67c', 1, 'Agarda vaqt va sharoit undan kechish va keyinchalik yana unga qaytish imkoniyatini bersa ham sizga qiziq bo‘lmagan, lekin boshlagan ishingizni tugata olasizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('ce98e18b-1de3-44e9-8993-a3dae91128ec', 2, 'Siz, sizga yoqimsiz narsani qilish kerak bo‘lganda, hech qanday qiyinchiliksiz ichki qarama-qarshiliklarni yenga olganmisiz?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('74d58caa-0a72-4da5-91b9-8d0a5aaefb84', 3, 'Siz qachondir ishda yoki hayotingizda konflikt holatga tushganingizda, bu holatga maksimal obyektivlikday qarab, o‘zingizni qo‘lga ola bilasizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('85ae4ae1-dd41-4f2c-847f-265bb79c7c52', 4, 'Agarda sizdan parhez qilish talab qilingan bo‘lsa, siz barcha mazali narsalarni yeyish ehtiyojingizni yenga olasizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('fa1cbb08-4e1a-4867-97dd-9ce95986d871', 5, 'Siz ertalab kecha rejalashtirilgandan ertaroq uyg‘onishga kuch topa bilasizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('f8f574d2-a0ff-474e-a709-069bd363c49a', 6, 'Siz guvohlik berish uchun voqe’a sodir bo‘lgan joyda qolasizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('456d2b4a-9f44-483d-b91d-28f21220a414', 7, 'Siz xatlarga tez javob bera olasizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('ab993a4e-0125-469d-b196-e1056d4dd6a9', 8, 'Agar sizda samolyotda uchish yoki tish shifokori xonasiga borish oldidan qo‘rquv paydo bo‘lsa, siz alohida qiyinchiliklarsiz bu hissiyotni yenga olasizmi va oxirgi daqiqada o‘zingizning fikringizni o‘zgartirmaysizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('64d688a6-decc-4572-a50b-d58d3b8ba334', 9, 'Siz shifokor yozib bergan juda yoqimsiz dorini qabul qilasizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('48f56ef4-f931-4c91-855e-5829cfc19306', 10, 'Siz qiziqqonlik bilan bergan va’dangizda turasizmi? Uni bajarish sizga biroz qiyinchiliklar tug‘dirsa ham?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('06bc1bd8-f3a0-4e2a-ba5c-28ae6243af81', 11, 'Agarda bu zarur bo‘lsa, siz hech qanday ikkilanmasdan notanish shaharga safarga jo‘naysizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('83955862-47f4-49b5-91eb-48279d0febaf', 12, 'Siz kun tartibiga qat’iy rioya qilasizmi: vaqtida turish, ovqatlanish, mashg‘ul bo‘lish, tozalash va boshqa narsalar.', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('32c8c20d-9930-4ac6-ac5e-d500e56a7bce', 13, 'Siz kutubxonadan qarzdorlarga yomon munosabatda bo‘lasizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('e0b66088-9f07-456e-8a80-2be3d85b01b3', 14, 'Juda qiziqarli teleko‘rsatuv sizga tezda va muhim bajariladigan ishni qoldirishga majbur qiladimi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b'),
    ('2b086e56-62ff-44af-b1e3-382a1f621abf', 15, 'Sizni qarama-qarshi tomon so‘zlari qanchalik xafa qiladigan bo‘lmasin tortishuvni to‘xtatib va jim bo‘la olasizmi?', 'b3e2b18a-1997-4157-bec2-f05bff14373b');



-- surovnoma - 12 /////////////////////////////////////////////////////////////////////////////////////////////

-- Polls jadvaliga ma'lumot qo'shish
INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES
    (
        '9b0e731a-97d9-420f-8fc6-c889b84b6a6d',
        12,
        'Xavotirlanish',
        'Mulohazalarni o''qing va o''zingizga mos javobni belgilang',
        '[{"ball": 0, "variant": "Hech qachon"}, {"ball": 1, "variant": "Juda kam"}, {"ball": 2, "variant": "Ba''zan"}, {"ball": 3, "variant": "Har doim"}]'::jsonb,
        '[{"to": 21, "from": 1, "text": "Sizda xavotirlanish darajasi juda past. Sizni psixologik jihatdan sarosimaga soladigan vaziyatlarga nisbatan mustahkam psixologik himoya shakllangan. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 35, "from": 22, "text": "Sizda xavotirlanish mavjud. Psixologik jihatdan sizni sarosimaga soaldigan vaziyatlar mavjud va ularga nisbatan psixologik himoya shakllantirishingiz zarur. Xavotirga solayotgan vaziyatlarni tahlil qilib ko''rish psixologik holatingizni yaxshilashga yordam beradi. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 90, "from": 36, "text": "Sizda xavotirlanish darajasi yuqori. Hayotingizda psixologik jihatdan sarosimaga soluvchi va stressga tushuruvchi vaziyatlar mavjud. Ularni tahlil qilib ko''rish psixologik holatingizni yaxshilanishiga yordam beradi. Har bir muammoning yechimi mavjud. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}]'::jsonb
    );

-- Savollar jadvaliga ma'lumot qo'shish
INSERT INTO questions (id, num, content, poll_id)
VALUES
    ('8960f964-b688-48a7-80b7-b174ce9ce668', 1, 'Tanangizda uvishish bo‘ladimi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('0517dd83-4ba8-4474-b706-38df38104dad', 2, 'Qizib ketasizmi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('4671ed9f-60b8-4453-b047-46baf2fef5e8', 3, 'Oyoqlaringizda qaltirash holati kuzatiladimi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('bab306a9-e20a-4355-8a9c-c603e02cc991', 4, 'Dam olish siz uchun qiyinmi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('0dce3857-c3fb-467c-bed6-b10631d717e8', 5, 'Yomon narsa sodir bo‘ladi, deb qo‘rquvga tushasizmi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('8baa3fd6-a27c-4427-a78e-5ebd42135801', 6, 'Boshingiz aylanadimi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('dc563cbc-bfcc-4f06-8e14-1e8e694e44da', 7, 'Yuragingiz tez urib ketadimi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('4fd07355-d1f3-4c79-9e52-6813dd495381', 8, 'Beqarormisiz?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('bef83041-51d7-4abd-8c58-101267d75116', 9, 'Sizda qo‘rquv bormi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('232e1503-8bae-4ba6-b44a-cfcb6e7847a3', 10, 'Asabiymisiz?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('8d396d45-a64e-4e8d-b109-70e53d6755bc', 11, 'Bo‘g‘ilgandek his qilasizmi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('f7428d8a-ba55-4cf7-ac7d-1c310fa46e6d', 12, 'Qo‘llaringiz titraydimi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('73878d9b-213c-4f94-a1c1-e763ffed38fe', 13, 'Muvozanatni yo‘qotib titrab ketasizmi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('6fb4008d-64c3-4e71-bddd-56642caa21dc', 14, 'Nazoratni yo‘qotishdan qo‘rqasizmi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('1cd83c64-dbc1-4bb2-b795-42ca37cc2007', 15, 'Nafas olishga qiynalasizmi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('73fa406f-4cf9-4255-b626-ecc07c3e9264', 16, 'O‘limdan qo‘rquv bo‘ladimi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('a22e4054-99d1-4a78-b617-42fadfe0059a', 17, 'Daxshatga tushasizmi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('fb3e2506-7de0-422a-b172-a351c72e54f9', 18, 'Sizda ovqat hazm qilish bilan bog‘liq muammo bormi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('cb421d8b-df94-4d49-b566-1a7e2a78e22f', 19, 'Tanangizda zaiflik his qilasizmi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('32df3b74-8e79-481f-a569-b9f2375bda54', 20, 'Yuzingiz qizarib ketadimi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d'),
    ('9a234f41-3bec-46ca-bdaa-ef00acaf7f02', 21, 'Tanangizda issiq yoki sovuq ter chiqib ketadimi?', '9b0e731a-97d9-420f-8fc6-c889b84b6a6d');



-- surovnoma - 13 /////////////////////////////////////////////////////////////////////////////////////////////

INSERT INTO polls (id, poll_num, title, subtitle, options, feedbacks)
VALUES
    (
        'd4e66eea-5860-44ee-9227-d079a6ce8fbc',
        13,
        'Linkedin',
        'Mulohazalarni o''qing va o''zingizga mos javobni belgilang',
        '[{"ball": 2, "variant": "Ha"}, {"ball": 0, "variant": "Yo''q"}, {"ball": 1, "variant": "Ba''zan"}]'::jsonb,
        '[{"to": 6, "from": 1, "text": "Sizga Linkedin professional platformasining salbiy ta''sir ko''rsatmaydi. O''z qobiliyatlaringiz va o''zingizga bo''lgan bahoyingiz yaxshi shakllangan. Ajoyib 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 13, "from": 7, "text": "Sizga Linkedin professional platformasi biroz salbiy ta''sir ko''rsatmoqda. O''z qobiliyatlaringizni kashf eting va o''zingizni boshqalar bilan taqqoslamang. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}, {"to": 60, "from": 14, "text": "Qobiliyatlaringizga ishonishingizda va o''zingizni baholashingizda Linkedin professional platformasi salbiy ta''sir ko''rsatmoqda. Har bir inson individual ekanligini inobatga olgan holda shaxsiy yutuqlaringizga va qobiliyatlaringizga diqqat qarating. 😊 Testlarni bajarishda davom eting va psixologik holatingiz haqida ko''proq ma''lumotga ega bo''ling"}]'::jsonb
    );


INSERT INTO questions (id, num, content, poll_id)
VALUES
    ('96bd7ae1-383d-4cb5-aea5-0499237f68a0', 1, 'Linkedin ilovasida sizdan ko‘ra muvaffaqiyatliroq bo‘lgan dasturchilarni kuzatish sizda sarosima paydo qiladimi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc'),
    ('d51f207c-b53d-4bd8-9eb3-86df23d09827', 2, 'Linkedinda dasturchilarning yutuqlarini ko‘rganingizda ular mendan ko‘ra mahoratliroq deb o‘zingizni taqqoslaysizmi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc'),
    ('803b665e-7f22-48b9-9751-f90daa548880', 3, 'Linkedin kabi professional ilovalarda boshqalarning loyiha va ishlarini ko‘rganingizda men ham ko‘proq harakat qilishim kerak degan his paydo bo‘ladimi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc'),
    ('487f4cad-414a-49ba-bbb6-0567a430ac33', 4, 'Linkedinda postlar orasida sohangizda qandaydir yangilik yaratgan dasturchini ko‘rganingizda, sizda bu uchun qobiliyatlaringiz yetarli emasdek tuyiladimi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc'),
    ('d55741fe-8adf-4b78-9459-7ef4ac4738ee', 5, 'Linkedinda HR lar tomonidan dasturchilarga qo‘yilgan talablar sizni dovdiratib qo‘yadimi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc'),
    ('346c53ad-2272-469c-b931-46bbffb71093', 6, 'Linkedinda dasturchilar portfoliosini kuzatish sizning o‘zingizga bo‘lgan ishonchingizni tushiradimi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc'),
    ('1630f594-840e-4488-92a5-7462008299b6', 7, 'Linkedindagi sohangizga oid postlarni kuzatish sizni ko‘proq harakat qilishim kerak deb asabiylashtiradimi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc'),
    ('48ff2572-11bd-4efb-a937-aa0948f313aa', 8, 'Linkedindagi muvaffaqiyatli dasturchilarni kuzatib, men orqada qolib ketyapman degan fikr keladimi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc'),
    ('1f484ab5-219f-4e7e-9aaf-1cffb8e9fc03', 9, 'Linkedin ilovasidagi muvaffaqiyatli insonlar darajasiga yetish qiyin deb hisoblaysizmi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc'),
    ('62f984d1-bf1d-4944-bff6-d0cbbe5e52c9', 10, 'Linkedindagi professionallarning ko‘plab yutuq va bajarilgan ishlari sizga uddalab bo‘lmaydigandek tuyiladimi?', 'd4e66eea-5860-44ee-9227-d079a6ce8fbc');
