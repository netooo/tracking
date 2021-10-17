# Install fzf and load this file on your .zshrc, like below

function insert-in-buffer () {
    if [ -n "$1" ]; then
        local new_left=""
        if [ -n "$LBUFFER" ]; then
            new_left="${new_left}${LBUFFER} "
        fi
        if [ -n "$2" ]; then
            new_left="${new_left}${2} "
        fi
        new_left="${new_left}$1"
        BUFFER=${new_left}${RBUFFER}
        CURSOR=${#new_left}
    fi
}

# tracking find task
function fzf-find-task () {
    local SELECTED_TASK="$(tracking list | fzf | head -n1 | cut -d ' ' -f 1)"
    insert-in-buffer "${SELECTED_TASK}" "-t"
}
zle -N fzf-find-task
bindkey "^P" fzf-find-task # Assign to your favorite key bind
