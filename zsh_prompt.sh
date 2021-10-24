# load this file on your .zshrc, like below

tracking_current() {
  local current=$(tracking c)
  local current_time=$(echo $current | grep Duration | cut -d ' ' -f 2)
  if [ ! -n "$currrent_time" ]; then
      echo "NotTracking"
  else
      echo "[$current_time]"
  fi
}