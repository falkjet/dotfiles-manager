_dotfiles_manager_source_dir="$(realpath "$(dirname $0)")"

_dotfiles-manager_build() {
    (cd "$_dotfiles_manager_source_dir" && go build .)
}

dotfiles-manager() {
    _dotfiles-manager_build
    (HOME="$PWD/home" "$_dotfiles_manager_source_dir/dotfiles-manager" "$@")
}

eval "$(dotfiles-manager completion zsh)"

compdef _dotfiles-manager dotfiles-manager
