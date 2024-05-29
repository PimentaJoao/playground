term-setup:
	sudo apt install fonts-firacode
	# Installing preferences 
	cat ./gterminal.preferences | dconf load /org/gnome/terminal/

# Copies and modifies slightly the default zsh theme (robbyrussell)
custom-zsh-theme:
	# Makes a copy of the original theme
	cp ${ZSH}/themes/robbyrussell.zsh-theme $({ZSH_CUSTOM}/themes/robbyrussell.zsh-theme
	# Modifies copy, replaces 'c' with '~' 
	sed -i -e 's/}%c%{/}%~%{/g' ${ZSH_CUSTOM}/themes/robbyrussell.zsh-theme
	# Sources the theme
	exec zsh
