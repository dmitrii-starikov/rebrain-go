#!/bin/bash
set -euo pipefail

log_action() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

if [ -z "$1" ]; then
    echo "Usage: $0 <directory>"
    exit 1
fi

BASE_DIR="$1"

mkdir -p "$BASE_DIR"
cd "$BASE_DIR" || exit 1

log_action "Created base directory: $BASE_DIR"

echo "New file content" > new.txt
log_action "Created new.txt with content"

sleep 2

echo "Appended content" >> new.txt
log_action "Appended content to new.txt"

sleep 2

echo "New2 file content" > new2.txt
log_action "Created new2.txt with content"

sleep 2

mkdir -p changed
log_action "Created directory: changed"

sleep 2

echo "Changed content" > new.txt
log_action "Changed content of new.txt"

sleep 2

mv new.txt changed/
log_action "Moved new.txt to changed/"

sleep 2

rm -f new2.txt
log_action "Deleted new2.txt"

sleep 2

mkdir -p texts/one texts/two texts/three
log_action "Created texts/one, texts/two, texts/three"

sleep 2

for dir in one two three; do
    for file in one two three; do
        echo "Content of $file in $dir" > "texts/$dir/$file"
    done
done
log_action "Created files one, two, three in each subdirectory"

sleep 2

ln -s texts/one symlink_one
log_action "Created symlink: symlink_one -> texts/one"

sleep 2

echo "Last file content" > last.txt
log_action "Created last.txt in root"

sleep 2

for file in one two three; do
    mkdir -p -m 755 "texts/two/dir_$file"
    cp last.txt "texts/two/dir_$file/last_$file"
done
log_action "Copied last.txt to texts/two/dir_{one,two,three} as last_*"

sleep 2

mv last.txt texts/three/
log_action "Moved last.txt to texts/three/"

sleep 2

chmod 777 texts/one
chmod 766 texts/two
chmod 755 texts/three
log_action "Set permissions: texts/one 777, texts/two 666, texts/three 555"

log_action "All actions completed!"