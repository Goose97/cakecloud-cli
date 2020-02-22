install:
	 sudo install ./cakecloud /usr/local/bin
	 mkdir -p ~/.oh-my-zsh/plugins/cakecloud
	 cp -r ./cakecloud_autocomplete/. ~/.oh-my-zsh/plugins/cakecloud
	 source ~/.bash_profile