build:
	go build -o co2mini main.go

install:
	cp -f co2mini /usr/local/sbin/co2mini
	cp -f co2mini.service /etc/systemd/system/co2mini.service
	sudo systemctl enable co2mini	
	sudo systemctl start co2mini

clean:
	rm -f co2mini
