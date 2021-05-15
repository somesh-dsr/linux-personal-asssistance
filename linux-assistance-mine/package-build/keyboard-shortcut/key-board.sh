#!/bin/sh

existingkShortcut=$(gsettings get org.gnome.settings-daemon.plugins.media-keys custom-keybindings)

existingkShortcut=$(echo "$existingkShortcut" | sed 's/]//')
existingkShortcut="$existingkShortcut,'/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/linux-personal-assistance/']"
existingkShortcut=$(echo "$existingkShortcut" | sed "s/\[\,/[/")
echo "Creating key board shortcut to interact with your linux assistance."
gsettings set org.gnome.settings-daemon.plugins.media-keys custom-keybindings "$existingkShortcut"
gsettings set org.gnome.settings-daemon.plugins.media-keys.custom-keybinding:/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/linux-personal-assistance/ name 'linux-personal-assistance'
gsettings set org.gnome.settings-daemon.plugins.media-keys.custom-keybinding:/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/linux-personal-assistance/ command "/opt/linux-personal-assistance/linux-assistance 'generic'"
gsettings set org.gnome.settings-daemon.plugins.media-keys.custom-keybinding:/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/linux-personal-assistance/ binding '<Alt>d'
echo "Successfully created key board shortcut. Now you can interact with the application by pressing <Alt> +  d "