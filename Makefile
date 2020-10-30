update:
	docker pull prateek2408/juno

build:
	apt-get install -y libasound2-dev alsa-base pulseaudio
	vim +GoInstallBinaries +qall
	go build
start: 
	docker run -it -u ${UID}:1000 -e PULSE_SERVER=unix:${XDG_RUNTIME_DIR}/pulse/native -v ${XDG_RUNTIME_DIR}/pulse/native:${XDG_RUNTIME_DIR}/pulse/native -v ~/.config/pulse/cookie:/root/.config/pulse/cookie  --rm --device /dev/snd -v  ${PWD}:/opt/music-player prateek2408/juno
