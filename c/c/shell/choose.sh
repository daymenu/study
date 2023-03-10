
function choose {
    local default="$1"
    local prompt="$2"
    local choice_yes="$3"
    local choice_no="$4"
    local answer

    read -p "$prompt" answer
    [ -z "$answer" ] && answer="$default"

    case "$answer" in 
        [yY1] ) eval "$choice_yes"
            ;;
        [nN0] ) eval "$choice_no"
            ;;
        * ) printf "%b" "Unexpected answer '$answer'!" >&2 ;;
    esac
}

choose "y" "Do you want to play a game?" "ls" 'printf "%b" See you later Professor Falkin. \n' >&2
