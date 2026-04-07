#!/usr/bin/env python3
"""
Генератор README.md для репозитория с тренировками по алгоритмам.
Скрипт собирает данные из каталогов contest0..contest6 и less1..less3,
понижает уровень заголовков в URL.md на 1, добавляет ссылки на решения,
и формирует итоговый README.md с разделами "Контесты" и блоками уроков.
"""

import os
import re
import glob
from pathlib import Path

# Корневая директория репозитория (текущая)
ROOT = Path.cwd()

def slugify(text):
    """Преобразует текст в якорь для ссылки на заголовок."""
    # Убираем Markdown-разметку (ссылку, жирный, курсив)
    text = re.sub(r'\[([^\]]+)\]\([^)]+\)', r'\1', text)
    text = re.sub(r'[*_`]', '', text)
    # Приводим к нижнему регистру, заменяем пробелы и спецсимволы на дефисы
    slug = re.sub(r'[^\w\s-]', '', text).strip().lower()
    slug = re.sub(r'[-\s]+', '-', slug)
    return slug

def process_contest_file(file_path):
    """Обрабатывает файл contestN/URL.md:
    - Понижает уровень заголовков на 1
    - Для каждой задачи добавляет ссылку на решение contestN/номер/main.go
    Возвращает обработанное содержимое в виде строки.
    """
    with open(file_path, 'r', encoding='utf-8') as f:
        lines = f.readlines()

    contest_path = file_path.parent
    new_lines = []

    for line in lines:
        # Понижаем уровень заголовков (добавляем один #)
        if line.startswith('#'):
            stripped = line.lstrip()
            level = len(line) - len(stripped)
            hash_count = len(stripped) - len(stripped.lstrip('#'))
            if hash_count > 0:
                new_line = ' ' * level + '#' * (hash_count + 1) + stripped[hash_count:]
                new_lines.append(new_line)
                continue

        # Обработка нумерованного списка задач (1. Название)
        match_num = re.match(r'^(\s*)(\d+)\.\s+(.+)$', line)
        if match_num:
            prefix, number, title = match_num.groups()
            task_num = int(number)
            solution_path = contest_path / str(task_num) / 'main.go'
            if solution_path.exists() and '([решение]' not in line:
                link = f'([решение]({solution_path.as_posix()}))'
                new_line = f'{prefix}{number}. {title.rstrip()} {link}\n'
                new_lines.append(new_line)
            else:
                new_lines.append(line)
            continue

        # Обработка маркированного списка с буквой (- A. Название) – на всякий случай
        match_bullet = re.match(r'^(\s*)-?\s*([A-Za-z])\.\s+(.+)$', line)
        if match_bullet:
            prefix, letter, title = match_bullet.groups()
            letter_lower = letter.lower()
            solution_path = contest_path / letter_lower / 'main.go'
            if solution_path.exists() and '([решение]' not in line:
                link = f'([решение]({solution_path.as_posix()}))'
                new_line = f'{prefix}- {letter}. {title.rstrip()} {link}\n'
                new_lines.append(new_line)
            else:
                new_lines.append(line)
            continue

        new_lines.append(line)

    return ''.join(new_lines)


def process_lesson_file(file_path):
    """Обрабатывает файл lessN/URL.md:
    - Понижает уровень заголовков на 1
    - Для каждой задачи добавляет ссылку на решение lessN/буква/main.go
    Возвращает обработанное содержимое в виде строки.
    """
    with open(file_path, 'r', encoding='utf-8') as f:
        lines = f.readlines()

    lesson_path = file_path.parent
    new_lines = []

    for line in lines:
        # Понижаем уровень заголовков
        if line.startswith('#'):
            stripped = line.lstrip()
            level = len(line) - len(stripped)
            hash_count = len(stripped) - len(stripped.lstrip('#'))
            if hash_count > 0:
                new_line = ' ' * level + '#' * (hash_count + 1) + stripped[hash_count:]
                new_lines.append(new_line)
                continue

        # Обработка маркированного списка задач (- A. Название)
        match_bullet = re.match(r'^(\s*)-?\s*([A-Za-z])\.\s+(.+)$', line)
        if match_bullet:
            prefix, letter, title = match_bullet.groups()
            letter_lower = letter.lower()
            solution_path = lesson_path / letter_lower / 'main.go'
            if not solution_path.exists():
                solution_path = lesson_path / letter_lower / 'solution.txt'
            if solution_path.exists() and '([решение]' not in line:
                link = f'([решение]({solution_path.as_posix()}))'
                new_line = f'{prefix}- {letter}. {title.rstrip()} {link}\n'
                new_lines.append(new_line)
            else:
                new_lines.append(line)
            continue

        # Нумерованные списки – пропускаем (в уроках не используются)
        new_lines.append(line)

    return ''.join(new_lines)


def extract_headings(content, level=2):
    """Возвращает список (текст заголовка, якорь) для всех заголовков указанного уровня."""
    headings = []
    pattern = re.compile(r'^#{' + str(level) + r'}\s+(.+)$', re.MULTILINE)
    for match in pattern.finditer(content):
        text = match.group(1).strip()
        # Убираем Markdown-ссылку, чтобы в оглавлении был читаемый текст
        clean_text = re.sub(r'\[([^\]]+)\]\([^)]+\)', r'\1', text)
        slug = slugify(clean_text)
        headings.append((clean_text, slug))
    return headings


def build_toc(headings):
    """Строит оглавление в виде Markdown-списка."""
    if not headings:
        return ''
    toc = ['## Оглавление', '']
    for text, slug in headings:
        toc.append(f'- [{text}](#{slug})')
    toc.append('')
    return '\n'.join(toc)


def main():
    # 1. Читаем readme_top.md, если есть
    top_path = ROOT / 'readme_top.md'
    top_content = ''
    if top_path.exists():
        with open(top_path, 'r', encoding='utf-8') as f:
            top_content = f.read()
    else:
        top_content = '# Тренировки по алгоритмам'

    # 2. Обрабатываем контесты
    contest_files = sorted(
        glob.glob('contest*/URL.md'),
        key=lambda path: int(os.path.basename(os.path.dirname(path)).replace('contest', ''))
    )
    contest_content = []
    for f in contest_files:
        path = Path(f)
        print(f'Processing {path}')
        contest_content.append(process_contest_file(path))

    # 3. Обрабатываем уроки
    lesson_files = sorted(glob.glob('less*/URL.md'))
    lesson_content = []
    for f in lesson_files:
        path = Path(f)
        print(f'Processing {path}')
        lesson_content.append(process_lesson_file(path))

    # 4. Собираем все заголовки второго уровня для оглавления
    all_headings = []

    # Заголовки из readme_top.md (если есть)
    if top_content:
        all_headings.extend(extract_headings(top_content, level=2))

    # Добавляем раздел "Контесты", если его ещё нет
    has_contest_section = any(text == 'Контесты' for text, _ in all_headings)
    if not has_contest_section:
        # Запомним, что добавим этот раздел в итоговый README
        contest_section_title = 'Контесты'
        contest_section_slug = slugify(contest_section_title)
        all_headings.append((contest_section_title, contest_section_slug))

    # Заголовки из уроков (после обработки они стали второго уровня)
    for lesson in lesson_content:
        all_headings.extend(extract_headings(lesson, level=2))

    # Строим оглавление
    toc = build_toc(all_headings)

    # 5. Формируем итоговый README
    readme_lines = []

    # Вставляем содержимое readme_top.md
    if top_content:
        readme_lines.append(top_content.rstrip('\n'))
        readme_lines.append('\n\n')

    # Вставляем оглавление
    if toc:
        readme_lines.append(toc)
        readme_lines.append('\n\n')

    # Добавляем раздел "Контесты", если его нет в top_content
    if not has_contest_section:
        readme_lines.append('## Контесты\n\n')

    # Добавляем обработанные контесты
    for content in contest_content:
        readme_lines.append(content)
        readme_lines.append('\n')
    readme_lines.append('[Оглавление](#оглавление)\n')

    # Добавляем обработанные уроки (без дополнительного заголовка)
    for content in lesson_content:
        readme_lines.append(content)
        readme_lines.append('\n')
        readme_lines.append('[Оглавление](#оглавление)\n')

    # Записываем README.md
    output_path = ROOT / 'README.md'
    with open(output_path, 'w', encoding='utf-8') as f:
        f.writelines(readme_lines)

    print(f'README.md успешно сгенерирован: {output_path}')


if __name__ == '__main__':
    main()
