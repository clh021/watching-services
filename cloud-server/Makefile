.PHONY: init run


init:
	cp -f database/.ht.sqlite.example database/.ht.sqlite
	cd plugins && \
		rm -rf wp-api-swaggerui && \
		git clone --depth=1 https://github.com/jonyextenz/wp-api-swaggerui.git && \
		cd wp-api-swaggerui && \
		npm install && \
		npm run build
	@echo "[OK] App binary was created!"

run:
	@docker-compose up
