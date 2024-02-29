#!/usr/bin/python
# -*- coding: UTF-8 -*-
import os
import sys

def load_label_mapping(mapping_file):
    label_mapping = {}
    with open(mapping_file, 'r', encoding='utf-8') as f:
        for line in f:
            label, word = line.strip().split(': ')
            label_mapping[label] = word
    return label_mapping

def restore_words(root_dir, file_extensions, label_mapping):
    for root, dirs, files in os.walk(root_dir):
        for file in files:
            if file.endswith(file_extensions):
                file_path = os.path.join(root, file)
                with open(file_path, 'r', encoding='utf-8') as f:
                    content = f.read()

                # Восстанавливаем слова из меток
                for label, word in label_mapping.items():
                    content = content.replace(label, word)

                with open(file_path, 'w', encoding='utf-8') as f:
                    f.write(content)

root_dir='.'
exts = ('.html', '.go', '.js', '.css', '.sql')

lang = sys.argv[1]

# Восстановление label_mapping из файла
label_mapping = load_label_mapping(f'label_mapping_{lang}.txt')

# Восстановление слов из меток
restore_words(root_dir, exts, label_mapping)
