#!/bin/bash

loginuser="$SUDO_USER"
if [ "$loginuser" = '' ];then
	loginuser="$USER"
fi
# shellcheck disable=SC2060
if [ "$(echo "$1" | tr [:upper:] [:lower:])"  = 'update-success' ]; then
	su "$loginuser" -c 'notify-send -i face-smile -t 6000 "OS updated successfully"'
elif [ "$(echo "$1" | tr [:upper:] [:lower:])" = 'update-failed' ];then
  su "$loginuser" -c 'notify-send -i face-sad -t 6000 "Failed to update the OS"'
fi


if [ "$(echo "$1" | tr [:upper:] [:lower:])"  = 'meeting reminder' ]; then
  su "$loginuser" -c 'notify-send -i face-smile -t 6000 "All the best for your meeting."'
fi