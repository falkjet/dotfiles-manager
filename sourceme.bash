_repotool_source_dir="$(realpath "$(dirname "${BASH_SOURCE[0]}")")"

_repotool_build() {
    (cd "$_repotool_source_dir" && go build .)
}

repotool() {
    _repotool_build
    "$_repotool_source_dir"/repotool "$@"
}

eval "$(repotool completion bash)"
