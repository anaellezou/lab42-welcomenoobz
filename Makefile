compile:
	go build

drop:
	rm -rf lab42.db

clean:
	rm lab42-welcomenoobz

fclean: clean drop

re: fclean compile