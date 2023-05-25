
create_db:
	@docker start mongo 2> /dev/null || docker run --name mongo\
		-d -p 27017:27017 --rm mongo:5

remove_db:
	@docker stop mongo 2> /dev/null || true

create_redis:
	@docker start redis 2> /dev/null || docker run --name redis \
		-d -p 6379:6379 --rm redis:7

remove_redis:
	@docker stop redis 2> /dev/null || true

clean: remove_db

.PHONY: create_db remove_db clean
