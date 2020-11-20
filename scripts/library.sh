. ./scripts/environment.sh

pause() {
  read -n1
}

shopt -s expand_aliases

command -v bat >/dev/null 2>&1 && {
  export BAT_THEME="Monokai Extended Origin"
  alias cat=bat
}

set -v
