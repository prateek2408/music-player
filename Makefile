build:
	apt-get install -y libasound2-dev alsa-base pulseaudio
	go build
start: 
	docker run -it -e PULSE_SERVER=unix:${XDG_RUNTIME_DIR}/pulse/native -v ${XDG_RUNTIME_DIR}/pulse/native:${XDG_RUNTIME_DIR}/pulse/native -v ~/.config/pulse/cookie:/root/.config/pulse/cookie  --rm --device /dev/snd -v  ${PWD}:/opt/music-player prateek2408/juno
