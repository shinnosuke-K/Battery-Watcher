#!/bin/bash

# plist ファイルの準備
echo "<?xml version="1.0" encoding="UTF-8"?>" >>  battery-watch-daily.plist
echo "<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">" >>  battery-watch-daily.plist
echo "<plist version="1.0">" >>  battery-watch-daily.plist
echo "<dict>" >>  battery-watch-daily.plist
echo "<key>Label</key>" >>  battery-watch-daily.plist
echo "<string>batteryWattchDaily</string>" >>  battery-watch-daily.plist
echo "<key>ProgramArguments</key>" >>  battery-watch-daily.plist
echo "<array>" >>  battery-watch-daily.plist
echo "<string>/bin/bash</string>" >>  battery-watch-daily.plist
echo "<string>${PWD}/watch.sh</string>" >>  battery-watch-daily.plist
echo "</array>" >> battery-watch-daily.plist
echo "<key>StartCalendarInterval</key>" >> battery-watch-daily.plist
echo "<array>" >> battery-watch-daily.plist
echo "<dict>" >> battery-watch-daily.plist
echo "<key>Hour</key>" >> battery-watch-daily.plist
echo "<integer>0</integer>" >> battery-watch-daily.plist
echo "<key>Minute</key>" >> battery-watch-daily.plist
echo "<integer>0</integer>" >> battery-watch-daily.plist
echo "</dict>" >> battery-watch-daily.plist
echo "<dict>" >> battery-watch-daily.plist
echo "<key>Hour</key>" >> battery-watch-daily.plist
echo "<integer>6</integer>" >> battery-watch-daily.plist
echo "<key>Minute</key>" >> battery-watch-daily.plist
echo "<integer>0</integer>" >> battery-watch-daily.plist
echo "</dict>" >> battery-watch-daily.plist
echo "<dict>" >> battery-watch-daily.plist
echo "<key>Hour</key>" >> battery-watch-daily.plist
echo "<integer>12</integer>" >> battery-watch-daily.plist
echo "<key>Minute</key>" >> battery-watch-daily.plist
echo "<integer>0</integer>" >> battery-watch-daily.plist
echo "</dict>" >> battery-watch-daily.plist
echo "<dict>" >> battery-watch-daily.plist
echo "<key>Hour</key>" >> battery-watch-daily.plist
echo "<integer>18</integer>" >> battery-watch-daily.plist
echo "<key>Minute</key>" >> battery-watch-daily.plist
echo "<integer>0</integer>" >> battery-watch-daily.plist
echo "</dict>" >> battery-watch-daily.plist
echo "</array>" >> battery-watch-daily.plist
echo "</dict>" >> battery-watch-daily.plist
echo "</plist>" >> battery-watch-daily.plist

# 定期実行するスクリプトの作成
echo "#!/bin/bash" >> watch.sh
echo "cd ${PWD} && ./batwatch -path $1" >> watch.sh