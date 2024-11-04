import openpyxl

def count_triangles(filename):
    count = 0
    workbook = openpyxl.load_workbook(filename)
    sheet = workbook.active  # Выбираем активный лист

    for row in sheet.iter_rows(values_only=True):
        a, b, c = row
        if a + b > c and a + c > b and b + c > a:
            count += 1

    return count

def is_triangle(text_file):  
    count = 0
    with open(text_file, 'r', encoding='utf-8') as file:
        for line in file:
            a, b, c = map(int, line.split())
            if a + b > c or a + c > b or b + c > a:
                count += 1
    return count

print(count_triangles("F:\\Downloads\\Треугольники.xlsx"))