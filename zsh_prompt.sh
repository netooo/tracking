# load this file on your .zshrc, like below

tracking_current() {
  local current=$(tracking c)
  local duration=$(echo $current | grep Duration | cut -d ' ' -f 2)
  if [ ! -n "$duration" ]; then
      echo "NotTracking"
  else
      echo "$duration"
  fi
}