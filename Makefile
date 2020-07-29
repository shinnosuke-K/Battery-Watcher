PATH = ${PWD}

init:
	go get -u ./...
	go build -o batwatch main.go
	touch battery-watch-daily.plist
	touch watch.sh
	touch cap.csv
	chmod 666 cap.csv

	/bin/bash ./init.sh ${PATH}

	sudo /bin/cp ./battery-watch-daily.plist /Library/LaunchDaemons/battery-watch-daily.plist
	sudo /bin/launchctl load /Library/LaunchDaemons/battery-watch-daily.plist

unload:
	sudo /bin/launchctl unload /Library/LaunchDaemons/battery-watch-daily.plist

reload:
	sudo /bin/cp ./battery-watch-daily.plist /Library/LaunchDaemons/battery-watch-daily.plist
	sudo /bin/launchctl load /Library/LaunchDaemons/battery-watch-daily.plist
