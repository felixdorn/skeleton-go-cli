#!/bin/bash
# 'return' when run as "source <script>" or ". <script>", 'exit' otherwise
[[ "$0" != "${BASH_SOURCE[0]}" ]] && safe_exit="return" || safe_exit="exit"

script_name=$(basename "$0")

ask_question() {
    # ask_question <question> <default>
    local ANSWER
    read -r -p "$1 ($2): " ANSWER
    echo "${ANSWER:-$2}"
}

confirm() {
    # confirm <question> (default = N)
    local ANSWER
    read -r -p "$1 (y/N): " -n 1 ANSWER
    echo " "
    [[ "$ANSWER" =~ ^[Yy]$ ]]
}

repo_url=$(ask_question "Repository URL")
binary_name=$(ask_question "Binary name" "myapp")
git_email=$(git config user.email)
git_name=$(git config user.name)
author_name=$(ask_question "Author name" "$git_name")
author_email=$(ask_question "Author email" "$git_email")

if ! confirm "Modify files?"; then
    $safe_exit 1
fi

mv "cmd/name" "cmd/$binary_name"
mv "core/handler/cli/name.go" "core/handler/cli/$binary_name.go"

grep -Erli ":bin|owner/repo|:email" --exclude-dir=bin ./* ./.github/* | grep -v "$script_name" \
| while read -r file ; do
    echo "adapting $file"
        temp_file="$file.temp"
        < "$file" \
          sed "s#:bin#$binary_name#g" \
          sed "s#name.#$binary_name#g" \
          sed "s#name :=#$binary_name#g" \
        | sed "s#owner/repository#$repo_url#g" \
        | sed "s#:author#$author_name#g" \
        | sed "s#:email#$author_email#g" \
        > "$temp_file"
        mv "$temp_file" "$file"
done

rm README.md

if confirm "Remove stub files?"; then
  rm core/handler/cli/cmd/testdata/output/*.golden
  rm core/handle/cli/cmd/inspire*.go
fi

rm -rf .git
git init
git remote add origin "git@github.com:$repo_url.git"

if confirm 'Let this script delete itself?'; then
    sleep 1 && rm -- "$0"
fi