import os
import re
import sys


def merge_md_block(main, branch):
    new_md = []
    main_h2 = dict()
    branch_h2 = dict()
    main_h2_block, main_h2_block_title = [], None
    branch_h2_block, branch_h2_block_title = [], None
    # generate main h2 block
    for line in main:
        if len(line) > 3 and line[:3] == "## ":
            if main_h2_block_title is not None:
                main_h2[main_h2_block_title] = main_h2_block
            main_h2_block = [line]
            main_h2_block_title = line
        else:
            main_h2_block.append(line)
    if main_h2_block_title is not None:
        main_h2[main_h2_block_title] = main_h2_block
    # generate branch h2 block
    for line in branch:
        if len(line) > 3 and line[:3] == "## ":
            if branch_h2_block_title is not None:
                branch_h2[branch_h2_block_title] = branch_h2_block
            branch_h2_block = [line]
            branch_h2_block_title = line
        else:
            branch_h2_block.append(line)
    if branch_h2_block_title is not None:
        branch_h2[branch_h2_block_title] = branch_h2_block
    # merge h2 block
    for key, value in main_h2.items():
        if branch_h2.get(key) is not None:
            tmp = merge_h2_block(value, branch_h2[key])
            for line in tmp:
                new_md.append(line)
        else:
            for line in value:
                new_md.append(line)
    for key, value in branch_h2.items():
        if main_h2.get(key) is None:
            for line in value:
                new_md.append(line)
    return new_md


def merge_h2_block(main, branch):
    new_md = [main[0]]
    main_h3 = dict()
    branch_h3 = dict()
    main_h3_block, main_h3_block_title = [], None
    branch_h3_block, branch_h3_block_title = [], None
    # generate main h3 block
    for line in main[1:]:
        if len(line) > 4 and line[:4] == "### ":
            if main_h3_block_title is not None:
                main_h3[main_h3_block_title] = main_h3_block
            main_h3_block = [line]
            main_h3_block_title = line
        else:
            main_h3_block.append(line)
    if main_h3_block_title is not None:
        main_h3[main_h3_block_title] = main_h3_block
    # generate branch h3 block
    for line in branch[1:]:
        if len(line) > 4 and line[:4] == "### ":
            if branch_h3_block_title is not None:
                branch_h3[branch_h3_block_title] = branch_h3_block
            branch_h3_block = [line]
            branch_h3_block_title = line
        else:
            branch_h3_block.append(line)
    if branch_h3_block_title is not None:
        branch_h3[branch_h3_block_title] = branch_h3_block
    # merge h3 block
    for key, value in main_h3.items():
        if branch_h3.get(key) is not None:
            tmp = merge_h3_block(value, branch_h3[key])
            for line in tmp:
                new_md.append(line)
        else:
            for line in value:
                new_md.append(line)
    for key, value in branch_h3.items():
        if main_h3.get(key) is None:
            for line in value:
                new_md.append(line)
    return new_md


def merge_h3_block(main, branch):
    new_ul = []
    new_text = [main[0]]
    for line in main[1:]:
        if len(line) > 2 and line[:2] == "* ":
            new_ul.append(line)
        else:
            new_text.append(line)
    for line in branch[1:]:
        if len(line) > 2 and line[:2] == "* ":
            new_ul.append(line)
        else:
            new_text.append(line)
    return new_text


def change_url(base_url, line):
    r = re.compile("^\* \[.*\]\((.*)\).*$")
    res = r.match(line)
    if res is not None:
        relative_url = res.group(1)
        line = line.replace(relative_url, os.path.join("../../" + base_url, relative_url))
    return line


def load_md(path, is_change_url=False):
    md = None
    dir = os.path.dirname(path)
    with open(path, encoding='utf-8') as file:
        md = file.readlines()
    for i in range(len(md)-1, -1, -1):
        if md[i] == "\n":
            del(md[i])
        else:
            if is_change_url:
                md[i] = change_url(dir, md[i])
    return md


def merge_mds(paths):
    main_md = ''
    for path in paths:
        branch_md = load_md(path, True)
        main_md = merge_md_block(main_md, branch_md)
    return main_md


def main():
    # the path of base.md
    main_dir = sys.argv[1]
    # the dir of branchs' parent dir
    branch_dir = sys.argv[2]
    # ths path of README.md to save
    save_dir = sys.argv[3]
    # merge branch mds
    branchs = os.listdir(branch_dir)
    branch_path = []
    for i in branchs:
        path = "/".join([branch_dir, i, "README.md"])
        if os.path.exists(path):
            branch_path.append(path)
    branch_md = merge_mds(branch_path)
    # load main md
    main_md = load_md(main_dir)
    # merge main-md & branch-mds
    main_md = [*main_md, *branch_md]
    # save md
    with open(save_dir, encoding='utf-8', mode='w') as file:
        file.write("".join(main_md))


if __name__ == "__main__":
    main()