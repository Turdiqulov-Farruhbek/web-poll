-- Insert poll
INSERT INTO polls (id, poll_num, title, options) VALUES (
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb', 
    1, 
    'Yo‘riqnoma: Mulohazalar sizga mos kelishiga qarab, raqamlardan birini aylana bilan belgilab qo‘ying.', 
    '{"1": "Hech qachon", "2": "Odatda", "3": "Har doim"}'::JSONB
);


-- Insert questions to the poll
INSERT INTO questions (id, num, content, poll_id) VALUES (
    'b8aeedae-be34-464d-a016-0ca6755cf904', 
    1, 
    'Turli topshiriqlarni bajarishda muvaffaqiyat qozonganman, lekin, ularni bajarishdan avval buni qila olmayman degan qo‘rquv bo‘lar edi.',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
INSERT INTO questions (id, num, content, poll_id) VALUES (
    '5ee40db7-bfd8-48a7-b612-8bf4579ef902',
    2,
    'Men o‘zim mavjud qobiliyatdan ko‘ra ko‘proq qobiliyatli ekanligim haqida taassurot qoldira olaman',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
INSERT INTO questions (id, num, content, poll_id) VALUES (
    '86253556-126a-4353-819e-86e68554735c',
    3,
    'Iloji bo‘lsa baholanishdan qochaman va boshqalar meni baholashlaridan qo‘rqaman',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
INSERT INTO questions (id, num, content, poll_id) VALUES (
    '03072365-6553-4674-8500-519394844387',
    4,
    'Qachonki, odamlar meni erishgan narsam uchun  maqtashsa, kelajakda umidlarini oqlay olmayman, deb qo‘rqaman',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
INSERT INTO questions (id, num, content, poll_id) VALUES (
    '6013f474-103a-4655-922e-271148386343',
    5,
    'To‘g‘ri vaqtda kerakli joyda bo‘lganim va kerakli odamlarni bilganim uchun hozirgi mavqeyimga va mufavvaqiyatlarimga erishdim deb o‘ylayman',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
INSERT INTO questions (id, num, content, poll_id) VALUES (
    '88e6e231-5091-4954-a18a-587212425013',
    6,
    'Men uchun muhim odamlar ular o‘ylagandek qobiliyatli emasligimni bilib qolishlaridan qo‘rqaman',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
INSERT INTO questions (id, num, content, poll_id) VALUES (
    '05124628-7e5a-4512-8288-53915571721c',
    7,
    'Qo‘limdan kelganicha harakat qilgan paytlarimda ham undan ham ko‘proq harakat qilishim kerak edi deb eslab turaman',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
INSERT INTO questions (id, num, content, poll_id) VALUES (
    '56096997-447e-4286-8848-b99661193614',
    8,
    'Men kamdan-kam hollarda loyiha yoki topshiriqni ko‘nglimdagidek bajaraman.',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
INSERT INTO questions (id, num, content, poll_id) VALUES (
    '6170190e-2932-414d-8229-49533168256c',
    9,
    'Hayotimdagi muvaffaqiyatlarim kimningdir e’tiborsizligi yoki omad natijasida bo‘lgan deb hisoblayman.',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);
INSERT INTO questions (id, num, content, poll_id) VALUES (
    '15103056-798e-4c5e-8d50-114931787266',
    10,
    'Intellektim yoki yutuqlarim haqidagi maqtovlar yoki e’tiroflarni qabul qilish men uchun qiyin',
    'e3a2920e-8f23-4603-8086-3dfd8cffaffb'
);